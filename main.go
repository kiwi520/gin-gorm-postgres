package main

import (
	"gin-gorm-postgres/database"
	"gin-gorm-postgres/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	database.Start()
	model.Migrate()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/create/class", func(c *gin.Context) {

		className := c.DefaultPostForm("class", "") // 此方法可以设置默认值
		if err := database.DB.Create(&model.Class{ClassName: className}).Error; err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.POST("/create/teacher", func(c *gin.Context) {

		teacherName := c.DefaultPostForm("teacher", "") // 此方法可以设置默认值
		if err := database.DB.Create(&model.Teacher{TeacherName: teacherName}).Error; err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.POST("/create/student", func(c *gin.Context) {

		stuName := c.DefaultPostForm("name", "")        // 此方法可以设置默认值
		class := c.DefaultPostForm("class", "1")        // 此方法可以设置默认值
		teacherName := c.DefaultPostForm("teacher", "") // 此方法可以设置默认值

		var teachers []model.Teacher
		//database.DB.Table("teachers").Where("teacher_name = ?",teacherName).Find(&teachers)
		database.DB.Model(&model.Teacher{}).Where("teacher_name = ?", teacherName).Find(&teachers)

		classId, _ := strconv.Atoi(class)
		if err := database.DB.Create(&model.Student{StudentName: stuName, ClassID: uint(classId), Teachers: teachers}).Error; err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.POST("/update/teacher", func(c *gin.Context) {

		teacherName := c.DefaultPostForm("teacher", "") // 此方法可以设置默认值
		teachName := c.DefaultPostForm("teach", "")     // 此方法可以设置默认值

		database.DB.Model(&model.Teacher{}).Where("teacher_name = ?", teacherName).Update("teach_name", teachName)

		c.JSON(200, gin.H{
			"message": "ok",
		})
	})


   //获取老师任课信息
	r.GET("/teacher", func(context *gin.Context) {
		teacherName := context.DefaultQuery("name", "Guest") // 此方法可以设置默认值
		var data []model.Teacher
		database.DB.Select([]string{"TeacherName","TeachName","CreatedAt","ID","UpdatedAt"}).Where("teacher_name = ?", teacherName).First(&data)
		context.JSON(200, gin.H{
			"code":200,
			"message": "ok",
			"data":data,
		})
	})


   //获取老师所属学生
	r.GET("/teacherStuList", func(context *gin.Context) {
		teacherName := context.DefaultQuery("name", "Guest") // 此方法可以设置默认值


		//关联查询（多对多）
		var data model.Teacher
		database.DB.Model(model.Teacher{}).Where("teacher_name = ?",teacherName).First(&data)
		database.DB.Model(&data).Association("Students").Find(&data.Students)

		context.JSON(200, gin.H{
			"code":200,
			"message": "ok",
			"data":data,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
