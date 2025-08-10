# byfood-app
## Book Management + URL Cleaner App
Technical test for Software Engineer at ByFood

A Dockerized application for storing books data + URL cleaner. It includes:
- Next.js app for frontend
- Go app for handling books data and URL cleaner
- PostgreSQL for database

## Getting Started
### 0. Install Dependencies & Tools
- **[Docker & Docker Compose](https://docs.docker.com/compose/)**
### 1. Clone The Repo

```bash
git clone https://github.com/rahadianir/byfood.git
cd byfood
```
### 2. Build & Run Services
Make sure you have docker compose installed and run this command to build and run the services.

```bash
docker compose up --build -d
```
### 3. Browse the app

DASHBOARD: http://localhost

SWAGGER: http://localhost:8080/swagger/index.html 

* Next.js serves dashboard on port 80
* Backend API served on port 8080
* Postgresql accessible on port 5432

### 4. Clean up
```bash
docker compose down --volumes
```

## Overview
### Next.js App :
* Serves dashboard to manage books data on port 80
* Use dynamic routing to show book data details
* Use modal dialog to enhance user experience
* Client side validation for most inputs

### Go App :
Serves API to interact with books data and URL cleaner

#### GET /books
Get all books data from database

**Request Example:**
```bash
curl --request GET --url http://localhost:8080/books 
```
**Response Example:**
```json
{
    "message": "books fetched",
    "data": [
        {
            "id": 1,
            "title": "To Kill a Mockingbird",
            "author": "Harper Lee",
            "publish_year": 1960,
            "created_at": "2025-08-10T15:30:46.064356Z",
            "updated_at": "2025-08-10T15:30:46.064356Z"
        },
        {
            "id": 2,
            "title": "1984",
            "author": "George Orwell",
            "publish_year": 1949,
            "created_at": "2025-08-10T15:30:46.064356Z",
            "updated_at": "2025-08-10T15:30:46.064356Z"
        },
        {
            "id": 3,
            "title": "Pride and Prejudice",
            "author": "Jane Austen",
            "publish_year": 1813,
            "created_at": "2025-08-10T15:30:46.064356Z",
            "updated_at": "2025-08-10T15:30:46.064356Z"
        },
        {
            "id": 4,
            "title": "The Great Gatsby",
            "author": "F. Scott Fitzgerald",
            "publish_year": 1925,
            "created_at": "2025-08-10T15:30:46.064356Z",
            "updated_at": "2025-08-10T15:30:46.064356Z"
        },
        {
            "id": 5,
            "title": "Moby-Dick",
            "author": "Herman Melville",
            "publish_year": 1851,
            "created_at": "2025-08-10T15:30:46.064356Z",
            "updated_at": "2025-08-10T15:30:46.064356Z"
        },
        {
            "id": 6,
            "title": "War and Peace",
            "author": "Leo Tolstoy",
            "publish_year": 1869,
            "created_at": "2025-08-10T15:30:46.064356Z",
            "updated_at": "2025-08-10T15:30:46.064356Z"
        },
        {
            "id": 7,
            "title": "The Catcher in the Rye",
            "author": "J.D. Salinger",
            "publish_year": 1951,
            "created_at": "2025-08-10T15:30:46.064356Z",
            "updated_at": "2025-08-10T15:30:46.064356Z"
        },
        {
            "id": 8,
            "title": "The Hobbit",
            "author": "J.R.R. Tolkien",
            "publish_year": 1937,
            "created_at": "2025-08-10T15:30:46.064356Z",
            "updated_at": "2025-08-10T15:30:46.064356Z"
        },
        {
            "id": 9,
            "title": "Fahrenheit 451",
            "author": "Ray Bradbury",
            "publish_year": 1953,
            "created_at": "2025-08-10T15:30:46.064356Z",
            "updated_at": "2025-08-10T15:30:46.064356Z"
        },
        {
            "id": 10,
            "title": "The Lord of the Rings",
            "author": "J.R.R. Tolkien",
            "publish_year": 1954,
            "created_at": "2025-08-10T15:30:46.064356Z",
            "updated_at": "2025-08-10T15:30:46.064356Z"
        }
    ]
}
```
#### GET /books/{id}
Get book data by ID from database

**Request Example:**
```bash
curl --request GET --url http://localhost:8080/books/2 
```
**Response Example:**
```json
{
    "message": "book data fetched",
    "data": {
        "id": 2,
        "title": "1984",
        "author": "George Orwell",
        "publish_year": 1949,
        "created_at": "2025-08-10T16:24:56.481163Z",
        "updated_at": "2025-08-10T16:24:56.481163Z"
    }
}
```
#### POST /books
Store book data to database

**Request Example:**
```bash
curl --request POST \
  --url http://localhost:8080/books \
  --header 'Content-Type: application/json' \
  --data '{
	"title": "judul",
	"author": "penulisr",
	"publish_year": 2002
}'
```
**Response Example:**
```json
{
	"message": "book data stored",
	"data": {
		"id": 11,
		"title": "judul",
		"author": "penulisr",
		"publish_year": 2002,
		"created_at": "2025-08-10T16:26:11.633963Z",
		"updated_at": "2025-08-10T16:26:11.633963Z"
	}
}
```
#### PUT /books/{id}
Update book data to database

**Request Example:**
```bash
curl --request PUT \
  --url http://localhost:8080/books/11 \
  --header 'Content-Type: application/json' \
  --data '{
	"title": "judul",
	"author": "ganti-author",
	"publish_year": 2002
}'
```
**Response Example:**
```json
{
	"message": "book data updated",
	"data": {
		"id": 11,
		"title": "judul",
		"author": "ganti-author",
		"publish_year": 2002,
		"updated_at": "2025-08-10T16:26:11.633963Z"
	}
}
```
#### GET /books/{id}
Delete book data by ID from database

**Request Example:**
```bash
curl --request DELETE --url http://localhost:8080/books/3 
```
**Response Example:**
```json
{
	"message": "book data deleted"
}
```
#### POST /url/cleanup
Clean up url by the given operation. Operations that can be done are `"canonical"`, `"redirection"`, and `"all"` that combines both
**Request Example:**
```bash
curl --request POST \
  --url http://localhost:8080/url/cleanup \
  --header 'Content-Type: application/json' \
  --data '{
	"url": "https://BYFOOD.com/food-EXPeriences?query=abc/",
	"operation": "all"
}'
```
**Response Example:**
```json
{
	"processed_url": "https://www.byfood.com/food-experiences"
}
```

### PostgreSQL :
| id  | title  | author  | publish_year  | created_at  | updated_at  | deleted_at  |
|---|---|---|---|---|---|---|
| 1  | One Piece  | Eiichiro Oda  | 1997  |  2025-08-09 15:57:49.056 | 2025-08-09 15:57:49.056  | null  |
| 2  | Naruto  | Masashi Kishimoto  | 1997  | 2025-08-09 15:57:49.056  | 2025-08-09 15:57:49.056  | null  |
* Stores books data in library.books table
* Initializes schema on first launch from migration/init/init.sql so further migration can be stored in migration directory

### Network separation :
* api-network: frontend and backend containers
* repository-network: backend and PostgreSQL containers

## Repository Structure
```
.
├── backend
│   ├── docs
│   └── internal
├── frontend
│   ├── app
│   ├── components
│   ├── context
│   ├── node_modules
│   ├── public
│   └── styles
└── migration
    └── init
```

