package controllers

import (
	"context"
	"go-block-api/auth"
	"go-block-api/config"
	"go-block-api/evm"
	"go-block-api/model"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/spruceid/siwe-go"
)

var ctx = context.Background()

type signinParams struct {
	Message   string `json:"message" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

// instead of storing the nonce in db for an inexistant user we just put it in some redis that expires
func Nonce(c *gin.Context) {
	app := c.MustGet("app").(*config.App)
	address := c.Param("address")

	if !evm.IsValidAddress(address) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid address format"})
		return
	}

	nonce := siwe.GenerateNonce()

	err := app.Rdb.Set(ctx, address, nonce, 1*time.Minute).Err()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"nonce": nonce,
	})
}

func Signin(c *gin.Context) {
	app := c.MustGet("app").(*config.App)

	var signinP signinParams
	if err := c.ShouldBindJSON(&signinP); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// parse message to siwe
	siweMessage, err := siwe.ParseMessage(signinP.Message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get the nonce in cache for address
	addr := siweMessage.GetAddress().String()
	nonce, err := app.Rdb.Get(ctx, addr).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	// domain will be cors restricted its fine to just use the one from the message
	domain := siweMessage.GetDomain()
	// verify signature
	publicKey, err := siweMessage.Verify(signinP.Signature, &domain, &nonce, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addr = crypto.PubkeyToAddress(*publicKey).Hex()

	token, err := auth.GenerateJWT(addr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	// if user exist we return it
	res := app.Db.Where("address = ?", addr).First(&user)
	if res.RowsAffected == 1 {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
			"jwt":  token,
		})
		return
	}

	// if user not exist we create it
	user = model.User{Address: addr}
	res = app.Db.Create(&user)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"jwt":  token,
	})
}
