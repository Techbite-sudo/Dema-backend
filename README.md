# **PhotoVault Backend**

Welcome to the PhotoVault backend repository! This repository contains the backend implementation using Go, GraphQL, GORM, and PostgreSQL.

## **Getting Started**
-------------------

### **Clone the Repository**

```bash
git clone <repository-url>
cd backend
```

### **Install Dependencies**

```bash
go mod tidy
```

### **Set Up PostgreSQL Database**

- Install PostgreSQL on your machine if you haven't already.
- Create a new PostgreSQL database for PhotoVault.

### **Configure Environment Variables**

- Create a `.env` file in the root directory.
- Add your PostgreSQL connection string and other environment variables as needed.

Example `.env` file:

```bash
DATABASE_URL=postgres://user:password@localhost:5432/photovault
```

### **Run the Application**

```bash
go run server.go
```

The application should now be running at `http://localhost:8080`. You can access the GraphQL playground to interact with the API.

**Documentation**
-----------------

For detailed documentation on the GraphQL API and available endpoints, refer to the GraphQL schema and documentation provided in the `schema.graphql` file.

**Issues and Contributions**
----------------------------

If you encounter any issues or have suggestions for improvements, please open an issue or submit a pull request. Your contributions are greatly appreciated!
