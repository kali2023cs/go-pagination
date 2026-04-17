package controllers

import (
	"go-pagination/config"
	"go-pagination/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Pagination Response Structure
type PaginationResponse struct {
	Total      int64            `json:"total"`
	Page       int              `json:"page,omitempty"`
	Limit      int              `json:"limit"`
	NextCursor string           `json:"next_cursor,omitempty"`
	Data       []models.Product `json:"data"`
}

// 1. Offset Pagination: Classic LIMIT/OFFSET
func GetOffsetPagination(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 { page = 1 }
	offset := (page - 1) * limit

	var products []models.Product
	var total int64

	config.DB.Model(&models.Product{}).Count(&total)
	config.DB.Limit(limit).Offset(offset).Order("id asc").Find(&products)

	c.JSON(http.StatusOK, PaginationResponse{
		Total: total,
		Page:  page,
		Limit: limit,
		Data:  products,
	})
}

// 2. Cursor Pagination: More performant for scaling
func GetCursorPagination(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	cursor, _ := strconv.Atoi(c.Query("cursor")) // Last seen ID

	var products []models.Product

	query := config.DB.Limit(limit).Order("id asc")
	if cursor > 0 {
		query = query.Where("id > ?", cursor)
	}

	query.Find(&products)

	var nextCursor string
	if len(products) > 0 {
		nextCursor = strconv.Itoa(int(products[len(products)-1].ID))
	}

	c.JSON(http.StatusOK, PaginationResponse{
		Limit:      limit,
		NextCursor: nextCursor,
		Data:       products,
	})
}

// 3. Dynamic Search & Filtering
func GetDynamicSearch(c *gin.Context) {
	category := c.Query("category")
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max_price")
	searchQuery := c.Query("q")

	query := config.DB.Model(&models.Product{})

	// Dynamic Filtering
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if minPrice != "" {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPrice != "" {
		query = query.Where("price <= ?", maxPrice)
	}

	// Full-Text Search Basics (using LIKE for simplicity, or tsvector for advanced)
	if searchQuery != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}

	var products []models.Product
	var total int64

	query.Count(&total)
	query.Limit(10).Find(&products)

	c.JSON(http.StatusOK, PaginationResponse{
		Total: total,
		Limit: 10,
		Data:  products,
	})
}
