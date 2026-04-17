package models

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"index" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"index" json:"price"`
	Category    string    `gorm:"index" json:"category"`
	CreatedAt   time.Time `gorm:"index" json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func SeedProducts(db *gorm.DB) {
	var count int64
	db.Model(&Product{}).Count(&count)
	if count > 0 {
		fmt.Println("Products already seeded, skipping...")
		return
	}

	categories := []string{"Electronics", "Fashion", "Home", "Books", "Toys"}
	names := []string{"Pro", "Ultra", "Lite", "Grand", "Smart", "Mini", "Super", "Elite"}
	items := []string{"Laptop", "Phone", "Watch", "Camera", "Headphones", "Speaker", "Monitor", "Tablet"}

	fmt.Println("Seeding 1000 products...")

	var products []Product
	for i := 1; i <= 1000; i++ {
		cat := categories[rand.Intn(len(categories))]
		prodName := fmt.Sprintf("%s %s %s %d", names[rand.Intn(len(names))], cat, items[rand.Intn(len(items))], i)
		
		products = append(products, Product{
			Name:        prodName,
			Description: fmt.Sprintf("High quality %s for everyday use. Model number %d.", prodName, i),
			Price:       float64(rand.Intn(1000)) + rand.Float64(),
			Category:    cat,
			CreatedAt:   time.Now().Add(time.Duration(-rand.Intn(10000)) * time.Hour), // Randomized creation time for cursor demo
		})

		// Batch insert every 100 items
		if len(products) == 100 {
			db.Create(&products)
			products = []Product{}
		}
	}
	fmt.Println("Seeding complete.")
}
