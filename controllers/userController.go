package controllers

import (
	"gojwtproject/config"
	"gojwtproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid Input"})
		return
	}

	if result := config.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Failed to create User in the database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message" : "User Created Succesfully"})
}