package main

import (
	"fmt"
	"github.com/mert-yigittop/cxp-api-starter/config"
	"github.com/mert-yigittop/cxp-api-starter/pkg/database"
)

func main() {
	config.LoadConfig()

	db, err := database.ConnectToPostgresql()
	if err != nil {
		fmt.Println("Database connection failed")
		return
	}
	fmt.Println("Database connection established")

}
