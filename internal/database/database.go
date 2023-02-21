package database

import (
	"os"

	"github.com/danielcesario/finman/internal/transactions"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

func OpenDb() (db *gorm.DB, err error) {
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_DATABASE := os.Getenv("DB_DATABASE")

	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		DryRun:                                   false,
	})
}

func UpdateSchema(db *gorm.DB) {
	db.AutoMigrate(
		&transactions.Role{},
		&transactions.User{},
	)

	roleAdmin := &transactions.Role{Role: "ADMIN"}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&roleAdmin)

	roleUser := &transactions.Role{Role: "USER"}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&roleUser)
}
