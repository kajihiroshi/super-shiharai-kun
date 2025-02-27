
# **スーパー支払い君.com (Super Shiharai Kun)**

## **Overview**
This project is a REST API for a fictional web service called **スーパー支払い君.com** (Super Shiharai Kun). The service allows users to register invoices with future payment due dates. On the due date, the system automatically processes the payment via bank transfer, even if the user's account balance is zero, effectively delaying cash outflow by up to one month.

The API is built using **Golang** and follows best practices for code organization, separation of concerns, and security.

---

## **Features**
- **Create Invoices**: Users can create new invoice records.
- **List Invoices**: Users can retrieve a list of invoices within a specified date range.
- **Automatic Fee Calculation**: The system calculates the total invoice amount, including a 4% fee and 10% consumption tax on the fee.
- **Secure Configuration**: Database credentials and other sensitive information are managed using environment variables.

---

## **Technologies Used**
- **Golang**: Primary programming language.
- **Gin**: Web framework for building the REST API.
- **MySQL**: Relational database for storing invoice, user, and company data.
- **Godotenv**: Library for loading environment variables from a `.env` file.
- **Git**: Version control system.

---

## **Project Structure**
```
super-payment-kun/
├── .env                     # Environment variables (not committed to Git)
├── .gitignore               # Specifies files to ignore in Git
├── go.mod                   # Go module file
├── go.sum                   # Go checksum file
├── README.md                # Project documentation
├── /cmd
│   └── /api
│       └── main.go          # Entry point for the API server
├── /internal
│   ├── /models              # Core business models
│   │   ├── invoice.go
│   │   ├── user.go
│   │   └── company.go
│   ├── /db                  # Database interaction layer
│   │   ├── invoice_db.go
│   │   ├── user_db.go
│   │   └── company_db.go
│   └── /service             # Business logic layer
│       ├── invoice_service.go
│       ├── user_service.go
└──     └── company_service.go
```

---

## **Setup Instructions**

### **Prerequisites**
1. **Go**: Install Go from [https://golang.org/dl/](https://golang.org/dl/).
2. **MySQL**: Install MySQL from [https://dev.mysql.com/downloads/](https://dev.mysql.com/downloads/).
3. **Git**: Install Git from [https://git-scm.com/](https://git-scm.com/).

### **Steps**
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/kajihiroshi/super-shiharai-kun.git
   cd super-shiharai-kun
   ```

2. **Set Up the Database**:
   - Create a MySQL database named `super_shiharai_kun`.
   - Run the necessary SQL scripts to create tables (if applicable).

3. **Configure Environment Variables**:
   - Create a `.env` file in the root directory:
     ```env
     DB_USER=your_db_user
     DB_PASSWORD=your_db_password
     DB_HOST=127.0.0.1
     DB_PORT=3306
     DB_NAME=super_payment_kun
     ```
   - Replace the placeholders with your actual database credentials.

4. **Install Dependencies**:
   ```bash
   go mod download
   ```

5. **Run the Application**:
   ```bash
   go run cmd/api/main.go
   ```

6. **Access the API**:
   - The API will be available at `http://localhost:8080`.

---

## **API Documentation**

### **Base URL**
```
http://localhost:8080/api
```

### **Endpoints**

#### **1. Create a New Invoice**
- **URL**: `/invoices`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "company_id": 1,
    "client_id": 1,
    "issue_date": "2023-10-01T00:00:00Z",
    "payment_amount": 10000,
    "due_date": "2023-11-01T00:00:00Z"
  }
  ```
- **Response**:
  ```json
  {
    "id": 1,
    "company_id": 1,
    "client_id": 1,
    "issue_date": "2023-10-01T00:00:00Z",
    "payment_amount": 10000,
    "fee": 400,
    "fee_rate": 0.04,
    "tax": 40,
    "tax_rate": 0.10,
    "total_amount": 10440,
    "due_date": "2023-11-01T00:00:00Z",
    "status": "未処理"
  }
  ```

#### **2. List Invoices Within a Date Range**
- **URL**: `/invoices`
- **Method**: `GET`
- **Query Parameters**:
  - `start`: Start date (e.g., `2024-10-01`).
  - `end`: End date (e.g., `2025-02-26`).
- **Response**:
  ```json
  [
    {
      "id": 1,
      "company_id": 1,
      "client_id": 1,
      "issue_date": "2023-10-01T00:00:00Z",
      "payment_amount": 10000,
      "fee": 400,
      "fee_rate": 0.04,
      "tax": 40,
      "tax_rate": 0.10,
      "total_amount": 10440,
      "due_date": "2023-11-01T00:00:00Z",
      "status": "未処理"
    }
  ]
  ```

---

## **Testing**
To run unit tests:
```bash
go test ./...
```

---