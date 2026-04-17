package main

import (
	"go-pagination/config"
	"go-pagination/models"
	"testing"
)

func BenchmarkOffsetPagination(b *testing.B) {
	config.ConnectDatabase()
	var products []models.Product
	limit := 10
	offset := 990 // Simulate deep pagination

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config.DB.Limit(limit).Offset(offset).Order("id asc").Find(&products)
	}
}

func BenchmarkCursorPagination(b *testing.B) {
	config.ConnectDatabase()
	var products []models.Product
	limit := 10
	cursor := 990 // Direct access via index

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config.DB.Limit(limit).Where("id > ?", cursor).Order("id asc").Find(&products)
	}
}
