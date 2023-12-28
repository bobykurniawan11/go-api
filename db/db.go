package db

import (
	"os/exec"

	"github.com/bobykurniawan11/starter-go/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func Init(env string) *gorm.DB {
	config := config.GetConfig()
	dsn := config.GetString("database.url")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_db = db

	//run cli

	// migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank" -verbose up
	exec.Command("migrate", "-path", "db/migration", "-database", config.GetString("fulldatabase.url"), "-verbose", "up")

	return db
}

func GetDB() *gorm.DB {
	return _db
}
