package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

var DB *sql.DB

// ConnectDatabase membuat koneksi ke MySQL/MariaDB Laragon.
func ConnectDatabase() {

	// Membaca file .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env tidak ditemukan, menggunakan environment variable")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Default configuration
	if host == "" {
		host = "127.0.0.1"
	}

	if port == "" {
		port = "3306"
	}

	if user == "" {
		user = "root"
	}

	// DSN MySQL
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbName,
	)

	fmt.Println("===================================")
	fmt.Println("DATABASE CONFIGURATION")
	fmt.Println("===================================")
	fmt.Println("HOST :", host)
	fmt.Println("PORT :", port)
	fmt.Println("USER :", user)
	fmt.Println("DB   :", dbName)
	fmt.Println("===================================")

	// Membuka koneksi
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("❌ Database Open Error:", err)
	}

	// Mengecek koneksi
	if err := db.Ping(); err != nil {
		log.Fatal("❌ Database Connection Error:", err)
	}

	DB = db

	fmt.Println("✅ Database Connected Successfully")
}
