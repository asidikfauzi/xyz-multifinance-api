package database

import (
	"asidikfauzi/xyz-multifinance-api/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitDatabase() *gorm.DB {
	dbConfig := config.DBConfigFromEnv()
	dbName := dbConfig.DBName
	dbUser := dbConfig.User
	dbPass := dbConfig.Password
	dbHost := dbConfig.Host
	dbPort := dbConfig.Port
	dbCharset := dbConfig.Charset
	dbParseTime := dbConfig.ParseTime
	dbLoc := dbConfig.Locale

	initDSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, dbCharset, dbParseTime, dbLoc,
	)

	db, err := gorm.Open(mysql.Open(initDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	var exists int
	checkDBQuery := fmt.Sprintf("SELECT COUNT(*) FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = '%s'", dbName)
	db.Raw(checkDBQuery).Scan(&exists)

	if exists == 0 {
		createDBQuery := fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", dbName)
		if err := db.Exec(createDBQuery).Error; err != nil {
			log.Fatalf("Failed to create database: %v", err)
		}
		fmt.Println("Database created successfully:", dbName)
	}

	sqlDB, _ := db.DB()
	err = sqlDB.Close()
	if err != nil {
		log.Printf("Warning: Failed to close database connection: %v", err)
	}

	finalDSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, dbCharset, dbParseTime, dbLoc,
	)

	db, err = gorm.Open(mysql.Open(finalDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to newly created database: %v", err)
	}

	sqlDB, err = db.DB()
	if err != nil {
		log.Fatalf("Failed to access database connection pool: %v", err)
	}

	sqlDB.SetConnMaxIdleTime(10 * time.Second)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	fmt.Println("Database connection established successfully!")
	return db
}
