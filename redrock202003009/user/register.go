package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"redrock20200309/sql"
)

type User struct {
	gorm.Model
	Num      int    //`gorm:"AUTO_INCREMENT;unique;not null"`
	Username string //`gorm:"-"`  //`json:"username" bind"required"`
	Password string //`gorm:"-"`  //`json:"password" bind"required"`
	Add      bool
}

func Register(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	//var user User
	//err := c.ShouldBindJSON(&user);
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	db := sql.Link()
	defer db.Close()
	db.AutoMigrate(&User{})
	db.Create(&User{Num: 0, Username: username, Password: password, Add: false})
	c.JSON(200, gin.H{"status": http.StatusOK, "message": "注册成功"})
}

