package goblockapi

import (
	"os"

	"github.com/agonist/goblockapi/evm"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Rpc *evm.Client
	Rdb *redis.Client
	Db  *gorm.DB
}

func Init() *App {
	loadEnv()
	redis := setupRedis()
	db := setupDb()
	client := evm.New(os.Getenv("RPC_URL"))

	app := &App{
		Rpc: client,
		Rdb: redis,
		Db:  db,
	}
	return app
}

func setupRedis() *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return redis
}

func setupDb() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the db")
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed ti run migrations")
	}

	return db
}

func loadEnv() {
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")

	if "test" != env {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load()
}
