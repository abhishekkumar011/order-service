# API Documentation

## Base URL

```
http://localhost:8000
```

---

# POST /events

Processes an incoming order event.

## order.create

### Request

```json
{
    "type":"order.create",
    "customerId":"C101",
    "restaurantId":"R201",
    "items":[
        {
            "itemId":"Pizza",
            "qty":2
        }
    ]
}
```

### Response

```json
{
    "message":"Event processed successfully",
    "orderId":"ORD-xxxxxxxx"
}
```

---

## order.update.status

### Request

```json
{
    "type":"order.update.status",
    "orderId":"ORD-xxxxxxxx",
    "status":"Preparing"
}
```

### Response

```json
{
    "message":"Event processed successfully"
}
```

---

## order.update.items

### Request

```json
{
    "type":"order.update.items",
    "orderId":"ORD-xxxxxxxx",
    "items":[
        {
            "itemId":"Burger",
            "qty":1
        },
        {
            "itemId":"Coke",
            "qty":2
        }
    ]
}
```

### Response

```json
{
    "message":"Event processed successfully"
}
```

---

# GET /orders

Returns the latest state of every order.

### Response

```json
[
    {
        "orderId":"ORD-xxxxxxxx",
        "customerId":"C101",
        "restaurantId":"R201",
        "status":"Preparing",
        "items":[
            {
                "itemId":"Pizza",
                "qty":2
            }
        ],
        "lastUpdated":"2026-07-17T12:00:00Z"
    }
]
```
