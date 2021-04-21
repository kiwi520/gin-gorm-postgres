package main

import (
	"gin-gorm-postgres/database"
	"gin-gorm-postgres/model"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Start()
    model.Migrate()


	//database.DB.
	i:= model.IDCard{Num: 123456}
	s:= model.Student{
		StudentName: "kiwi",
		IDCard: i,
	}


	t:= model.Teacher{
		TeacherName: "腾跃文",
		Students: []model.Student{s},
	}

	c:= model.Class{
		ClassName: "实验一班",
		Student: []model.Student{s},
	}
	if err := database.DB.Create(&c).Error; err != nil {
		panic(err)
	}

	if err := database.DB.Create(&s).Error; err != nil {
		panic(err)
	}
	if err := database.DB.Create(&t).Error; err != nil {
		panic(err)
	}


	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}