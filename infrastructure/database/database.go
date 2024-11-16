package database

import (
	"clean-reservations/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	cnf := config.GetConfig()

	database := cnf.MySQLDatabase

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cnf.MySQLUser,
		cnf.MySQLPassword,
		cnf.MySQLHost,
		cnf.MySQLPort,
		database,
	)
	fmt.Println(dsn)

	db, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func MigrateExecution() error {
	err := db.AutoMigrate()
	if err != nil {
		return fmt.Errorf("migration execution failed: %w", err)
	}
	fmt.Println("migration was successfully performed")
	return nil
}
