package db

import "github.com/go-sql-driver/mysql"

var cfg = mysql.Config{
	User:   "root",
	Passwd: "123",
	Net:    "tcp",
	Addr:   "db:3306",
	DBName: "scraper",
}
