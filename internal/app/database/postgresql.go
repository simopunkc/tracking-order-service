package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewConnDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DryRun:      false,
		PrepareStmt: false,
	})
	if err != nil {
		log.Fatalln("can't connect to database")
	}
	return db
}

func NewDatabaseRead() *gorm.DB {
	dsn := "host=" + os.Getenv("DBR_HOST") + " user=" + os.Getenv("DBR_USER") + " password=" + os.Getenv("DBR_PASS") + " dbname=" + os.Getenv("DBR_NAME") + " port=" + os.Getenv("DBR_PORT") + " sslmode=disable TimeZone=Asia/Jakarta"
	return NewConnDatabase(dsn)
}

func NewDatabaseWrite() *gorm.DB {
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Asia/Jakarta"
	return NewConnDatabase(dsn)
}
