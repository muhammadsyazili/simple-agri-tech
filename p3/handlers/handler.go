package handlers

import (
	"coding-test/p3/config"
	"coding-test/p3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetTopSpendingCountryUsers(c *gin.Context) {
	var result struct {
		Country    string
		TotalSpend int64
	}

	// Query to find top spending country
	err := config.DB.Table("users").
		Select("users.country, SUM(spendings.total_buy) as total_spend").
		Joins("JOIN spendings ON users.id = spendings.user_id").
		Group("users.country").
		Order("total_spend DESC").
		Limit(1).
		Scan(&result).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.Country == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "No data found"})
		return
	}

	// Get users from that country
	var users []models.User
	if err := config.DB.Where("country = ?", result.Country).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

type CreateUserInput struct {
	Country        string `json:"country" binding:"required"`
	CreditCardType string `json:"credit_card_type" binding:"required"`
	CreditCard     string `json:"credit_card" binding:"required"`
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Country:          input.Country,
		CreditCardType:   input.CreditCardType,
		CreditCardNumber: input.CreditCard,
		FirstName:        input.FirstName,
		LastName:         input.LastName,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
