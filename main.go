package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bobykurniawan11/starter-go/config"
	"github.com/bobykurniawan11/starter-go/db"
	"github.com/bobykurniawan11/starter-go/server"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func main() {

	environment := flag.String("e", "development", "")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()
	config.Init(*environment)

	db.Init(*environment)

	server.Init()

}
