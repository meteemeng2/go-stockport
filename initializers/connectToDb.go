package initializers

import (
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	// github.com/denisenkom/go-mssqldb
	dsn := os.Getenv("STOCKPORT_DSN")
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true, // skip the snake_casing of names
		},
	})
	if err != nil {
		panic("Failed to connect to db")
	}
}
