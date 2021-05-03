package main

import (
	"user/bootstrap"
	"user/utils"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	// Connect to databse
	utils.ConnectDatabase()
	//Load env
	godotenv.Load()
	fx.New(bootstrap.Module).Run()

}
