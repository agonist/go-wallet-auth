package evm

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	client *ethclient.Client
}

func New(rpcUrl string) *Client {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	c := &Client{client: client}

	return c
}

func (c *Client) GetBalance(address string) (*big.Int, error) {
	balance, err := c.client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (c *Client) GetGasPrice() (*big.Int, error) {
	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	return gasPrice, nil
}
