package database

import (
	_ "github.com/lib/pq"
	"github.com/spf13/cast"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetGormConn function
func GetGormConn(host, user, dbName, password string, port int) (*gorm.DB, error) {
	// return gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
	// 	host, port, user, dbName, password,
	// ))

	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=" + user + " password=" + password + " dbname=" + dbName + " port=" + cast.ToString(port) + " sslmode=disable",
		PreferSimpleProtocol: false,
	}), &gorm.Config{})
}
