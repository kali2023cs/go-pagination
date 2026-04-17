# Go Pagination Project

This project demonstrates different pagination strategies and search techniques in Go using Gin and GORM with PostgreSQL.

## Features
- **Offset Pagination**: `/products/offset?page=1&limit=10`
- **Cursor Pagination**: `/products/cursor?cursor=10&limit=10`
- **Dynamic Filtering**: `/products/search?category=Electronics&min_price=100`
- **Full-Text Search**: `/products/search?q=laptop`
- **Seeding**: Automatically seeds 1000 products on first run.

## Performance Tuning (Offset vs Cursor)
Pagination performance degrades in SQL as the offset increases because the database still has to scan all preceding rows.

### Benchmark Comparison (Deep Pagination at Index 990)
- **Offset**: `SELECT * FROM products ORDER BY id LIMIT 10 OFFSET 990`
- **Cursor**: `SELECT * FROM products WHERE id > 990 ORDER BY id LIMIT 10`

Cursor pagination is consistently faster and more stable because it uses the index directly without scanning.

## How to Run
1. Ensure PostgreSQL is running and you have a database named `go_pagination_db`.
2. Update `.env` with your credentials.
3. Run `go run main.go`.
4. The server will start on `http://localhost:8080`.
