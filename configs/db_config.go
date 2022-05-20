package config

import (
	"ecommerce/models"
	"ecommerce/utils"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {

	environment := utils.Getenv("ENVIRONMENT", "development")
	if environment == "production" {
		username := os.Getenv("DATABASE_USERNAME")
		password := os.Getenv("DATABASE_PASSWORD")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")
		database := os.Getenv("DATABASE_NAME")
		// production
		dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=require"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		db.AutoMigrate(
			&models.Users{},
			&models.UsersProfile{},
			&models.Stores{},
			&models.StoresProfile{},
			&models.ProductCategory{},
			&models.Products{},
			&models.Carts{},
		)

		return db
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		database := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err.Error())
		}

		db.AutoMigrate(
			&models.Users{},
			&models.UsersProfile{},
			&models.Stores{},
			&models.StoresProfile{},
			&models.ProductCategory{},
			&models.Products{},
			&models.Carts{},
		)

		return db
	}

}
