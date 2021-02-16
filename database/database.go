package database

import (
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Driver         string
	Host           string
	Username       string
	Password       string
	Port           int
	Database       string
	SqliteInMemory bool
}

type Database struct {
	*gorm.DB
}

func New(config *DatabaseConfig) (*Database, error) {
	var db *gorm.DB
	var err error
	switch strings.ToLower(config.Driver) {
	case "mysql":
		dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.Database + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=UTC"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		break
	case "postgresql", "postgres":
		dsn := "user=" + config.Username + " password=" + config.Password + " dbname=" + config.Database + " host=" + config.Host + " port=" + strconv.Itoa(config.Port) + " TimeZone=UTC"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		break
	case "sqlserver", "mssql":
		dsn := "sqlserver://" + config.Username + ":" + config.Password + "@" + config.Host + ":" + strconv.Itoa(config.Port) + "?database=" + config.Database
		db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		break
	case "sqlite", "sqlite3":
		dsn := ""

		if config.SqliteInMemory {
			log.Println("Setting SQLite in Memory as Database")
			dsn = "file::memory:?cache=shared"
		} else {
			_, b, _, _ := runtime.Caller(0)
			rootdir := filepath.Join(filepath.Dir(b), "../")
			db_loc := filepath.Join(rootdir, "wota.db")
			dsn = db_loc
		}

		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		break
	}
	return &Database{db}, err
}
