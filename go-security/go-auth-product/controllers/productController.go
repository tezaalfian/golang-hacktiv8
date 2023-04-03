package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-auth-product/database"
	"go-auth-product/helpers"
	"go-auth-product/models"
	"net/http"
	"strconv"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	err := db.Debug().Create(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Product)
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{Title: Product.Title, Description: Product.Description}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Product)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	Product.ID = uint(productId)

	err := db.Debug().Delete(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}

func GetProduct(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	Product.ID = uint(productId)

	err := db.Debug().Where("id = ?", productId).Take(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Product)
}

func GetAllProducts(c *gin.Context) {
	db := database.GetDB()
	Products := []models.Product{}

	err := db.Debug().Find(&Products).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Products)
}
