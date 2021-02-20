package config

import (
	"fmt"
	"log"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// NewDBConnection returns initialized sqlx.DB
func NewDBConnection() (*sqlx.DB, error) {
	user := GetEnvWithDefault("DB_USER", "root")
	password := GetEnvWithDefault("DB_PASSWORD", "")
	host := GetEnvWithDefault("DB_HOST", "localhost")
	port := GetEnvWithDefault("DB_PORT", "3306")
	name := GetEnvWithDefault("DB_NAME", "app")
	timezone := url.QueryEscape("Asia/Tokyo")
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&time_zone=%s&loc=Local", user, password, name, host, port, timezone)
	db, err := sqlx.Open("mysql", datasource)
	if err != nil {
		return nil, fmt.Errorf("faild to connect DB -> %w", err)
	}
	log.Println("connected to DB")
	return db, nil
}
