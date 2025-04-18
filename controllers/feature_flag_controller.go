package controllers

import (
	"finbox_assignment/database"
	"finbox_assignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeatureFlagController struct{}

type CreateFeatureFlagRequest struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	DefaultValue bool   `json:"defaultValue"`
}

func (fc *FeatureFlagController) Create(c *gin.Context) {
	var req CreateFeatureFlagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	featureFlag := models.FeatureFlag{
		Name:         req.Name,
		Description:  req.Description,
		DefaultValue: req.DefaultValue,
	}

	result := database.GetDB().Create(&featureFlag)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create feature flag"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":           featureFlag.ID,
		"name":         featureFlag.Name,
		"description":  featureFlag.Description,
		"defaultValue": featureFlag.DefaultValue,
	})
}
