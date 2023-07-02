package models

import (
	"gorm.io/gorm"
	"network.com/network/config"
)

type Company struct {
	gorm.Model
	companyName      string
	address          string
	phoneNumber      uint64
	totalConnections uint32
}

func InitCompanyMigration() {
	config.GetDbConnection().AutoMigrate(&Company{})
}
