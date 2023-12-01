package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(){
	var err error

	POSTGRES_DB_HOST := os.Getenv("POSTGRES_DB_HOST")
	POSTGRES_DB_PORT := os.Getenv("POSTGRES_DB_PORT")
	POSTGRES_DB_USER := os.Getenv("POSTGRES_DB_USER")
	POSTGRES_DB_PASSWORD := os.Getenv("POSTGRES_DB_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")

	dsn := "host="+POSTGRES_DB_HOST+ " user="+POSTGRES_DB_USER+ " password="+POSTGRES_DB_PASSWORD+ " dbname="+POSTGRES_DB+ " port="+POSTGRES_DB_PORT+ " sslmode=disable"
	
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if(err != nil){
		log.Fatal("Failed to connect with database")
	}
}