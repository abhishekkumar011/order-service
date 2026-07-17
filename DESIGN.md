# Design Document

## Architecture

The system consists of two independent applications.

1. A Node.js event generator continuously creates randomized food order events.
2. A Go backend service receives these events over HTTP, processes them, stores the latest order state in MongoDB, and exposes an API to retrieve the current state of all orders.

```
Generator
    │
    ▼
  HTTP
    │
    ▼
Go Backend
    │
    ▼
 MongoDB
```

---

## Data Model

Each order contains:

* orderId
* customerId
* restaurantId
* items
* status
* lastUpdated

MongoDB automatically generates `_id`, while the application generates a separate `orderId` to keep the business identifier independent of the database implementation.

---

## Event Flow

### order.create

* Backend generates a unique orderId.
* Creates a new order.
* Stores it in MongoDB.
* Returns the generated orderId.

### order.update.status

* Finds an existing order.
* Updates only the status.
* Updates the lastUpdated timestamp.

### order.update.items

* Finds an existing order.
* Replaces the existing items with the newly received list.
* Updates the lastUpdated timestamp.

---

## Design Decisions

### Database

MongoDB was chosen because the order structure contains nested arrays (items), making it a natural fit for document storage.

### Transport

HTTP was selected because it is simple to implement, easy to test using Postman.

### Order Identifier

A separate `orderId` is generated instead of using MongoDB's `_id`. This keeps the application independent of any specific database and allows the event format to remain unchanged if another database is used.

---

## Handling Edge Cases

### Out-of-order events

Update events referencing a non-existing order are ignored because the corresponding order has not yet been created.

### Duplicate events

Duplicate detection has not been implemented. Each incoming event is processed independently.

### Throughput

The service processes events continuously over HTTP. For higher throughput, the transport layer could be replaced by a message broker such as Kafka or RabbitMQ.

### Latest State Guarantee

MongoDB always stores the latest version of each order. The `GET /orders` endpoint retrieves the current state directly from the database.

---
