package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-auth-product/database"
	"go-auth-product/helpers"
	"go-auth-product/models"
	"net/http"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	err := db.Debug().Create(&User).Error
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.FullName,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email or password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email or password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email, User.Role)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
