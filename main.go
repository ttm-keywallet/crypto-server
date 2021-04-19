package main

import (
	"fmt"
	"github.com/ttmbank/crypto-server/app"
)

func main() {
	srv := new(app.App)
	srv.Initialize()
	srv.Run(":8020")

	fmt.Println("Starting crypto server")
}
