package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Users struct {
	Id       int    `json:"id" gorm:"column:id"`
	FullName string `json:"fullName" gorm:"column:fullName"`
	Email    string `json:"email" gorm:"column:email"`
	Age      int    `json:"age" gorm:"column:age"`
}

func (Users) TableName() string {
	return "users"
}

type UserCreate struct {
	Id       int    `json:"id" gorm:"column:id"`
	FullName string `json:"fullName" gorm:"column:fullName"`
	Email    string `json:"email" gorm:"column:email"`
	Age      int    `json:"age" gorm:"column:age"`
}

func (UserCreate) TableName() string {
	return Users{}.TableName()
}

type UserUpdate struct {
	FullName *string `json:"fullName" gorm:"column:fullName"`
	Email    *string `json:"email" gorm:"column:email"`
	Age      *int    `json:"age" gorm:"column:age"`
}

func (UserUpdate) TableName() string {
	return Users{}.TableName()
}

func (res *UserCreate) Validate() error {
	res.Id = 0
	res.FullName = strings.TrimSpace(res.FullName)
	res.Email = strings.TrimSpace(res.Email)

	if len(res.FullName) == 0 {
		return errors.New("full name can't be blank")
	}

	if len(res.Email) == 0 {
		return errors.New("email can't be blank")
	}
	if res.Age == 0 {
		return errors.New("age can't be blank")
	}

	return nil
}

func createUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := data.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data.Id})
	}
}

func getListUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		type Paging struct {
			Page  int   `json:"page" form:"page"`
			Limit int   `json:"limit" form:"limit"`
			Total int64 `json:"total" form:"total"`
		}

		var paging Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 10
		}

		offset := (paging.Page - 1) * paging.Limit

		var data []Users

		if err := db.Table(Users{}.TableName()).
			Count(&paging.Total).
			Offset(offset).
			Limit(paging.Limit).
			Order("id desc").
			Find(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"paging": paging, "data": data})
	}
}

func getUserById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data Users

		id, err := strconv.Atoi(c.Param("user_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func updateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data UserUpdate

		id, err := strconv.Atoi(c.Param("user_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func deleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Table(Users{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func main() {
	dns := os.Getenv("MYSQL")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalf("Can not connect DB MYSQL ...")
	}

	log.Println("Connect to ", db)

	router := gin.Default()
	router.SetTrustedProxies([]string{"12.4.27.15"})
	v1 := router.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", createUser(db))
			users.GET("", getListUser(db))
			users.GET("/:user_id", getUserById(db))
			users.PUT("/:user_id", updateUser(db))
			users.DELETE("/:user_id", deleteUser(db))
		}
	}
	router.Run(":3010")
}
