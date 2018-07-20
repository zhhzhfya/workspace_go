package users

import (
	"gopkg.in/gin-gonic/gin.v1"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func CreateUser(c *gin.Context) {
	c.String(200, "Success")
}

func FetchAllUsers(c *gin.Context)  {

}

func FetchSingleUser(c *gin.Context)  {

}

func UpdateUser(c *gin.Context)  {

}

func DeleteUser(c *gin.Context)  {

}

//func startPage(c *gin.Context) {
//	var person Person
//	if c.ShouldBindQuery(&person) == nil {
//		log.Println("====== Only Bind By Query String ======")
//		log.Println(person.Name)
//		log.Println(person.Address)
//	}
//	c.String(200, "Success")
//}
