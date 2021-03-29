//package main
//
//import (
//	"fmt"
//	"log"
//	"strconv"
//	"database/sql"
//
//	"github.com/gin-gonic/gin"
//	"github.com/go-gorp/gorp"
//	"golang.org/x/crypto/bcrypt"
//
//	_ "github.com/go-sql-driver/mysql"
//)
//
//var dbmap = initDb()
//
//func checkErr(err error, msg string) {
//	if err != nil {
//		log.Fatalln(msg, err)
//	}
//}
//
//func initDb() *gorp.DbMap {
//	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/test")
//	checkErr(err, "sql.Open failed")
//	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
//	err = dbmap.CreateTablesIfNotExists()
//	checkErr(err, "Create tables failed")
//	return dbmap
//}
//
//func HashPassword(password string) string {
//	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
//	return string(bytes)
//}
//
//func CheckPasswordHash(password, hash string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//	return err == nil
//}
//
//type User struct {
//	Id        int64  `db:"ID" json:"id"`
//	Username  string `db:"Username" json:"username"`
//	Password  string `db:"Password" json:"password"`
//	Firstname string `db:"Firstname" json:"firstname"`
//	Lastname  string `db:"Lastname" json:"lastname"`
//}
//
//func GetUserDetail(c *gin.Context) {
//	id := c.Params.ByName("id")
//	var user User
//	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)
//	if err == nil {
//		userId, _ := strconv.ParseInt(id, 0, 64)
//		content := &User{
//			Id:        userId,
//			Username:  user.Username,
//			Password:  user.Password,
//			Firstname: user.Firstname,
//			Lastname:  user.Lastname,
//		}
//		c.JSON(200, content)
//	} else {
//		c.JSON(404, gin.H{"error": "user not found"})
//	}
//}
//
//func RegisterUser(c *gin.Context) {
//	uname := c.PostForm("username")
//	password := HashPassword(c.PostForm("password"))
//	first := c.PostForm("firstname")
//	last := c.PostForm("lastname")
//
//	var user User
//	err := dbmap.Insert(&user, "INSERT into user username=? password=? firstname=? lastname=?", uname, password, first, last)
//	if err == nil {
//		c.JSON(201, err)
//	} else {
//		c.JSON(204, gin.H{"error": "cannot create data, invalid params"})
//	}
//}
//
////func UpdateUser(c *gin.Context) {
////	id := c.Param("id")
////	uname := c.PostForm("username")
////	first := c.PostForm("firstname")
////	last := c.PostForm("lastname")
////
////	var user User
////	//err := dbmap.Update(&user, "Update user set username=? firstname=? lastname=?", uname, first, last, "where id=?", id)
////	result := dbmap.SelectOne(&user, "Select * from user where id=?", id)
////	result.Update()
////	if result == nil {
////		c.JSON(201, result)
////	} else {
////		c.JSON(204, gin.H{"error": "cannot create data, invalid params"})
////	}
////}
//
//func SplitString(c *gin.Context) {
//	var value string = c.Param("value")
//	for i:= 0; i < len(value); i++ {
//		fmt.Println(value[i])
//	}
//}
//
//func main() {
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//	r.GET("/search/:id", GetUserDetail)
//	r.POST("/register", RegisterUser)
//	r.GET("/split/:value", SplitString)
//	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//}

package main

import (
	"fmt"
	"test-kulina/config"
	"test-kulina/server"
)

func ulala() {
	fmt.Println("Uwuw")
}
func main() {
	ulala()
	config.Init()
	server.Start()
}