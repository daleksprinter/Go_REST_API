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

 		defer db.Close()
 		db.LogMode(true)
}