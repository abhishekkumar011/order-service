# Food Delivery Order Processing Service

## Overview

This project is a backend service built in **Go** that processes a continuous stream of food order events and maintains the latest state of each order in MongoDB.

A separate **Node.js Event Generator** continuously generates randomized order events and sends them to the backend using HTTP.

---

## Architecture

```
Node.js Event Generator
        │
        │ HTTP POST
        ▼
Go Backend Service (Gin)
        │
        ▼
     MongoDB
        │
        ▼
   GET /orders
```

---

## Technologies Used

### Backend

* Go
* Gin
* MongoDB

### Generator

* Node.js
* Axios

### Database

* MongoDB

### Transport

* HTTP

---

## Project Structure

```
Food-delivery-order/

cmd/
    server/

internal/
    config/
    database/
    handlers/
    models/
    routes/
    service/

generator/
    api.js
    data.js
    eventGenerator.js
    index.js
    .env

.env
README.md
API.md
DESIGN.md
```

---

## Prerequisites

* Go 1.24+
* Node.js
* MongoDB

---

## Backend Setup

Clone the repository.

Create a `.env` file in the project root.

Example:

```
MONGO_URI=YOUR_MONGO_URI
DATABASE_NAME=food_delivery
COLLECTION_NAME=orders
PORT=8000
```

Install Go dependencies.

```
go mod tidy
```

Run the backend.

```
go run cmd/server/main.go
```

---

## Generator Setup

Navigate to the generator folder.

```
cd generator
```

Install dependencies.

```
npm install
```

Create a `.env` file.

```
API_URL=http://localhost:8000/events
INTERVAL=1000
```

Run the generator.

```
npm start
```

---

## Features

* Continuously processes food order events
* Creates new orders
* Updates order status
* Updates order items
* Stores latest order state in MongoDB
* Returns all latest orders using the List Orders API

---

## Notes

Docker support has not been implemented in this submission. The application can be run locally by starting MongoDB, the Go backend, and the Node.js generator.
