package models

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=db user=your_username password=your_password dbname=your_db sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("Failed to connect to database")
    }

    DB = db
}
