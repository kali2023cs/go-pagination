go mod init go-pagination
go mod tidy

# Core Dependencies
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/joho/godotenv

# Start Server
go run main.go

# Performance Testing (Benchmarks)
# Use this to compare Offset vs Cursor pagination performance
go test -bench=. -benchmem
