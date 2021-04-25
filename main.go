package main

import (
	"fmt"

	"github.com/ttmbank/crypto-server/app"
)

func main() {
	srv := new(app.App)
	srv.Initialize()
	fmt.Println("Starting crypto server")
	srv.Run(":8020")
}
