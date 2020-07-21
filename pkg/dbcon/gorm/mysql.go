package dbcon

import (
	"fmt"
	"strconv"
	"time"

	// load database driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB :nodoc:
var DB *gorm.DB

// DBInit :nodoc:
func DBInit(prop MSQLProp) (*gorm.DB, error) {

	mysqlCon := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		prop.User,
		prop.Pass,
		prop.Host,
		strconv.Itoa(prop.Port),
		prop.Name,
		prop.Charset,
	)
	var err error
	DB, err = gorm.Open("mysql", mysqlCon)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("Failed connected to database %s\n", mysqlCon)
		return DB, err
	}
	DB.DB().SetConnMaxLifetime(time.Duration(prop.MaxLifeTimeMS) * time.Millisecond)
	DB.DB().SetMaxIdleConns(prop.MaxIdle)
	DB.DB().SetMaxOpenConns(prop.MaxOpen)
	fmt.Printf("Connection created.\n(%s)\n", mysqlCon)
	return DB, err
}

// GetConnection :nodoc:
func GetConnection() *gorm.DB {
	if DB == nil {
		fmt.Println("Creating Connection..")
		prop := MSQLProp{}
		DB, _ = DBInit(prop)
	} else {
		fmt.Println("Get Connection..")
	}
	return DB
}
