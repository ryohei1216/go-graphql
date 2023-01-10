package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("root:%s@tcp(mysql_c)/%s", "root", "test")
	// db, err := sql.Open("mysql", connectionInfo)
	// dsn := fmt.Sprintf("%s:%s@%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_PATH"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("データベース接続完了")
	return db, err
}

