# 🚀 Go Pagination API: Advanced Performance & Scalability

Welcome to the **Go Pagination API** project! This repository is designed to teach you how to handle large datasets efficiently using **Go**, **Gin**, and **PostgreSQL**. You will learn the critical differences between pagination strategies, how to build dynamic search filters, and how to optimize for millisecond-level response times even with millions of records.

---

## 🏗️ Integrated Technologies & Concepts

-   **Offset Pagination**: The traditional `LIMIT` and `OFFSET` approach for UI-based page numbers.
-   **Cursor Pagination**: A modern, scalable approach using "tokens" or "markers" for infinite scroll.
-   **Dynamic Query Building**: Constructing SQL queries on-the-fly based on optional user input.
-   **Full-Text Search (FTS)**: Moving beyond simple `LIKE` queries for intelligent text matching.
-   **GIN Indexing**: Specialized PostgreSQL indexes for ultra-fast search performance.
-   **Performance Benchmarking**: Comparing algorithms at high depths (e.g., page 1000+).

---

## 🔄 The Full Flow: Step-by-Step Logic

### 1. Database Configuration (`config/database.go`)
-   **Connection**: Establishes a robust connection to PostgreSQL using GORM.
-   **Environment**: Uses `.env` for secure credential management, keeping sensitive data out of the source code.

### 2. Scalable Seeding (`models/product.go`)
-   **Mass Data**: Automatically generates **1000 items** across diverse categories (Electronics, Fashion, etc.).
-   **Real-world Entropy**: Randomizes prices and timestamps to ensure your pagination tests mirror real-world data entropy.

### 3. Pagination Logic (`controllers/product_controller.go`)
-   **Offset Method**: Demonstrates how to calculate page offsets. Highlights why this strategy slows down as users go deeper into the results.
-   **Cursor Method**: Implements "Keyset Pagination" using the last seen ID. This ensures the database can "jump" directly to the next set of results without scanning previous ones.

### 4. Search & Filtering (`controllers/product_controller.go`)
-   **Flexible API**: Combines category filters, price ranges, and keyword search into a single unified endpoint.
-   **Clean Querying**: Shows how to conditionally add `WHERE` clauses to a GORM query object without cluttering the code.

---

## 🎓 Go Backend Interview Questions & Answers

### 🟢 Level 1: Pagination Basics
1.  **What is Offset-based pagination?**
    *   *Answer:* It uses `LIMIT X OFFSET Y` to skip `Y` rows and return the next `X`. It's the simplest method but suffers from performance degradation as the offset grows.
2.  **Why is Cursor pagination preferred for Mobile/Infinite Scroll?**
    *   *Answer:* Because it prevents "data shifting." If a new item is added while a user is scrolling, Offset pagination might show the same item twice. Cursor pagination stays pinned to the last record seen.

### 🟡 Level 2: Performance & Scalability
3.  **Why does `OFFSET 1000000 LIMIT 10` kill database performance?**
    *   *Answer:* The database must internally fetch and count **1,000,010** rows from the disk, only to discard the first million. This results in high I/O and CPU usage.
4.  **How do you implement a stable Cursor with multiple sort-keys?**
    *   *Answer:* By combining your sort-key (e.g., `CreatedAt`) with a unique ID (e.g., `(created_at, id) < (cursor_time, cursor_id)`). This ensures no records are skipped if two items have the exact same timestamp.

### 🔴 Level 3: Advanced Search
5.  **What is the difference between `LIKE %search%` and Full-Text Search?**
    *   *Answer:* `LIKE %...%` usually forces a "Full Table Scan," which is extremely slow on large tables. Full-Text Search uses **Inverted Indexes (GIN)** to map words directly to their row locations, similar to the index at the back of a textbook.
6.  **What is a GIN Index in PostgreSQL?**
    *   *Answer:* **Generalized Inverted Index**. It's specially designed to handle composite data types and is the gold standard for high-performance text search and JSON filtering in Postgres.

---

## 🛠️ Project Checklist
- [x] Initialize Go module and Gin/GORM dependencies.
- [x] Configure PostgreSQL with optimized connection pooling.
- [x] Implement Offset Pagination with `Total Count` metadata.
- [x] Implement Cursor Pagination using ID-based keysets.
- [x] Build Dynamic Search filters for Category and Price ranges.
- [x] Apply Basic Full-Text Search pattern.
- [x] Create automated seeder for 1000+ records.
- [x] Implement Benchmark tests to compare Offset vs Cursor performance.
- [x] Document the architecture in a professional Postman collection.
