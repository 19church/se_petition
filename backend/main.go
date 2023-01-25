package main

import (
	"github.com/gin-gonic/gin"

	"github.com/19church/se_petition/controller"
	"github.com/19church/se_petition/entity"
)

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	/////////////////////////////////////////////////////////////

	//Student
	r.POST("/CreateStudent", controller.CreateStudent)
	r.GET("/GetStudent/:id", controller.GetStudent)
	r.GET("/ListStudent", controller.ListStudent)
	r.DELETE("/DeleteAdmin/:id", controller.DeleteStudent)
	r.PATCH("/UpdateStudent", controller.UpdateStudent)

	//PetitionType
	r.POST("/CreatePetitionType", controller.CreatePetitionType)
	r.GET("/GetPetitionType/:id", controller.GetPetitionType)
	r.GET("/ListPetitionType", controller.ListPetitionType)
	r.DELETE("/DeletePetitionType/:id", controller.DeletePetitionType)
	r.PATCH("/UpdatePetitionType", controller.UpdatePetitionType)

	//PetitionPeriod
	r.POST("/CreatePetitionPeriod", controller.CreatePetitionPeriod)
	r.GET("/GetPetitionPeriod/:id", controller.GetPetitionPeriod)
	r.GET("/ListPetitionPeriod", controller.ListPetitionPeriod)
	r.DELETE("/DeletePetitionPeriod/:id", controller.DeletePetitionPeriod)
	r.PATCH("/UpdatePetitionType", controller.UpdatePetitionType)

	//TODO....abcde

	/////asdfasdfj;lj;j///

	//Petition
	r.POST("/CreatePetition", controller.CreatePetition)
	r.GET("/GetPetition/:id", controller.GetPetition)
	r.GET("/ListPetition", controller.ListPetition)
	r.DELETE("/DeletePetition", controller.DeletePetition)
	r.PATCH("/UpdatePetition", controller.UpdatePetition)

	r.Run()
}