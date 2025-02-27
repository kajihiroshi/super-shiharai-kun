package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"super-shiharai-kun/internal/db"
	"super-shiharai-kun/internal/models"
	"super-shiharai-kun/internal/service"

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
	clientDB := db.NewClientDB(dbConn)
	clientService := service.NewClientService(clientDB)

	clientBankAccountDB := db.NewClientBankAccountDB(dbConn)
	clientBankAccountService := service.NewClientBankAccountService(clientBankAccountDB)

	// Setup router
	r := gin.Default()

	// API endpoints for Clients
	r.POST("/api/clients", func(c *gin.Context) {
		var client models.Client
		if err := c.ShouldBindJSON(&client); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := clientService.CreateClient(&client); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, client)
	})

	r.GET("/api/clients", func(c *gin.Context) {
		companyIDStr := c.Query("company_id")
		companyID, err := strconv.Atoi(companyIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company_id"})
			return
		}

		clients, err := clientService.GetClientsByCompanyID(companyID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, clients)
	})

	// API endpoints for Client Bank Accounts
	r.POST("/api/client-bank-accounts", func(c *gin.Context) {
		var account models.ClientBankAccount
		if err := c.ShouldBindJSON(&account); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := clientBankAccountService.CreateAccount(&account); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, account)
	})

	r.GET("/api/client-bank-accounts", func(c *gin.Context) {
		clientIDStr := c.Query("client_id")
		clientID, err := strconv.Atoi(clientIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client_id"})
			return
		}

		accounts, err := clientBankAccountService.GetAccountsByClientID(clientID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, accounts)
	})

	// Start server
	r.Run(":8080")
}
