# Overview
This project focuses on creating and deploying a Golang backend to support an Angular web application. The backend will enable the management of a product catalog, with capabilities for creating, updating, listing, and searching products. It will also include features for sorting and paginating the product list.

# Features
**1. Product Management** \
**Create Product:** Add a new product with details such as name, type, price and  description. Validation added for all the fields, (*asterik is not added in UI) \
**Edit Product:** Modify the details of an existing product. \
**2. Product Listing** \
**Grid View:** List all products in a grid layout, If there is no records, then displaying "No record found". \
**Pagination:** Use pagination to manage large sets of products, ensuring only a limited number of products are displayed per page. \
**3. Sorting** \
**Sorting by Criteria:** Sort products by various criteria such as price (ascending/descending), product type, or any other field. \
**4. Search** \
**Search Functionality:** Allow users to search for products by name, displaying only the matched results.

# Project Structure
The project will consist of the following components: \
**Backend:** 
  - Golang APIs to handle product-related requests (CRUD operations, search, pagination, and sorting). 
  -	PostgreSQL database for product data storage. 
  -	Docker configuration for the backend application. \
**Frontend (Angular):** 
  -	Angular application to display products in a grid, implement search functionality, and provide sorting/paging features.

# Installation and Setup
**Prerequisites**
Before you start, make sure you have the following installed:

-  	Go (Golang) v1.23
-  	Docker  // Not able to use it in my laptop to test it
-  	PostgreSQL 
-  	Angular CLI v19.0.2 (for frontend development)
-  	Git 

# Backend Setup
Clone the repository:
```
Bash
git clone <repository-url>
cd <project-directory>
```
Build and run the Go application:
```
bash
go build -o app
./app
```
The backend will start running on a specified port (e.g., localhost:8000).

# Set up the PostgreSQL database:
Configure the postgresql like below
```
DB_HOST: localhost.
DB_PORT: Port of the PostgreSQL database (default: 5432).
DB_USER: postgres
DB_PASSWORD: yokini22
```

# Frontend Setup (Angular)
Navigate to the frontend directory:
```
bash
cd frontend
```
Install dependencies:
```
bash
npm install
```

Run the Angular application:
```
bash
ng serve
```
The frontend will be available at http://localhost:4200.

# API Endpoints
The backend will expose the following API endpoints:

- POST /api/products: Create a new product.
- GET /api/products: List all products with pagination and sorting options.
- GET /api/products/{id}: Get details of a single product.
- PUT /api/products/{id}: Update an existing product.
