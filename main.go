package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"encoding/json"
)

func gormConnect() *gorm.DB{
	DBMS := "mysql"
	USER := "root"
	PASSWORD := "password"
	PROTOCOL := "tcp(0.0.0.0:3306)"
	DBNAME := "gin"
	CONNECT := USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
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

type Ranking struct{
	gorm.Model
	Title string `gorm:"size:255"`
	UserID uint
}

func getUserByEmail(email string) (User, error) {
	db := gormConnect()
	var u User
	if err := db.Where("email = ?", email).First(&u).Error; err != nil {
        return u, err
    }

    return u, nil
}

func main() {

	db := gormConnect()
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&User{})
	db.AutoMigrate(Ranking{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	defer db.Close()
	db.LogMode(true)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
			c.String(200, "Hello,Gin!")
	})

	r.GET("api/user", func(c *gin.Context){
		var users []User
		if err := db.Find(&users).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		c.JSON(200, users)
	})
			
	r.POST("/api/user", func(c *gin.Context){
		c.Request.ParseForm()
		user := User{
			Name: c.PostForm("username"), 
			Password: c.PostForm("password"), 
			Email: c.PostForm("email"),
		}
		db.Create(&user)
	})

	r.GET("api/ranking", func(c *gin.Context){
		var rankings []Ranking
		if err := db.Find(&rankings).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		c.JSON(200, rankings)
	})

	r.POST("api/ranking", func(c *gin.Context){

		c.Request.ParseForm()
		email := c.PostForm("email")
		user, err := getUserByEmail(email)
		if err != nil {
			c.AbortWithStatus(400)
		}
		ranking := Ranking{
			Title: c.PostForm("title"),
			UserID: user.ID,
		}

		db.Create(&ranking)
	})

	

	r.Run()


}