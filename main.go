package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB{
	DBMS := "mysql"
	USER := "root"
	PASSWORD := "password"
	PROTOCOL := "tcp(0.0.0.0:3306)"
	DBNAME := "gin"
	CONNECT := USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
        panic(err.Error())
    }
    fmt.Println("db connected: ", &db)
    return db	
}

type User struct {
    gorm.Model
    Name string `gorm:"size:255"`
    Password string `gorm:"size:255"`
    Email string `gorm:"size:255"`
}

func main() {
		if false {
			r := gin.Default()
			r.GET("/", func(c *gin.Context) {
					c.String(200, "Hello,World!")
			})
			
			r.GET("/hoge", func(c *gin.Context){
				c.String(200, "hello hogehoge")
			})

			r.Run()
		}
		fmt.Println("hoge")
		db := gormConnect()
		db.Set("gorm:table_options", "ENGINE=InnoDB")
		db.AutoMigrate(&User{})
 		defer db.Close()
 		db.LogMode(true)
}