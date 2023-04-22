package config

import (
	"fmt"
	"go-final-project-mygram/models"

	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	// db conf
	Username string
	Password string
	Port     string
	Address  string
	Database string
	// db conn
	DB *gorm.DB
}

type GormDB struct {
	*Gorm
}

var (
	GORM *GormDB
)

func InitGorm() error {
	GORM = new(GormDB)

	GORM.Gorm = &Gorm{
		Username: os.Getenv("DB_USER"),
		Port:     os.Getenv("DB_PORT"),
		Address:  os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
	}

	err := GORM.Gorm.OpenConnection()
	if err != nil {
		return err
	}
	return nil
}

func (g *Gorm) OpenConnection() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", g.Address, g.Port, g.Username, g.Database)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	g.DB = dbConn

	err = g.DB.Debug().AutoMigrate(models.Users{}, models.SocialMedias{}, models.Photos{}, models.Comments{})
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected to database using gorm")
	return nil
}
