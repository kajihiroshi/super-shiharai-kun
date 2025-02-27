package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"super-shiharai-kun/internal/db"
	"super-shiharai-kun/internal/models"
	"super-shiharai-kun/internal/service"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construct the database connection string
	dbConnectionString := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	// Initialize database
	dbConn, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Test the database connection
	if err := dbConn.Ping(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Initialize layers
	invoiceDB := db.NewInvoiceDB(dbConn)
	invoiceService := service.NewInvoiceService(invoiceDB)

	// Setup router
	r := gin.Default()

	// API endpoints
	r.POST("/api/invoices", func(c *gin.Context) {
		var invoice models.Invoice
		if err := c.ShouldBindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := invoiceService.CreateInvoice(&invoice); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, invoice)
	})

	r.GET("/api/invoices", func(c *gin.Context) {
		start := c.Query("start")
		end := c.Query("end")

		startTime, _ := time.Parse("2006-01-02", start)
		endTime, _ := time.Parse("2006-01-02", end)

		invoices, err := invoiceService.GetInvoicesByPeriod(startTime, endTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, invoices)
	})

	// Start server
	r.Run(":8080")
}
