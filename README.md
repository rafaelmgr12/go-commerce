# **go-commerce**

A flexible solution for processing online orders in an e-commerce system, built with Go.

## **Overview**

The **`go-commerce`** system is designed to process online orders by applying a series of rules based on the characteristics of the order. The solution aims to replace a previous implementation that was hard to maintain due to numerous if/else conditions. The new design is modular, making it easy to add or remove rules as business requirements evolve.

## **Features**

- **Entities**:
    - **`Product`**: Represents a product with attributes like category and value.
    - **`Order`**: Represents an order with attributes like product, payment, and labels.
    - **`Payment`**: Represents a payment method and its value.
- **Rules**:
    - **`FreeShippingRule`**: Adds a **`free-shipping`** label for products with a value greater than 1000.
    - **`FragileProductRule`**: Adds a **`fragile`** label for products in the **`home-appliance`** category.
    - **`ChildProductRule`**: Adds a **`gift`** label for products in the **`child`** category.
    - **`BoletoDiscountRule`**: Provides a 10% discount for payments made with "Boleto".

## **Getting Started**

1. **Clone the Repository**:
    
    ```bash
    git clone https://github.com/rafaelmgr12/go-commerce.git
    
    ```
    
2. **Navigate to the Directory**:
    
    ```bash
    cd go-commerce
    
    ```
    
3. **Install Dependencies**:
    
    ```go
    go mod download
    
    ```
    
4. **Run Tests**:
    
    ```bash
    go test ./...
    
    ```
    

## **Usage**

To process an order and apply the relevant rules:

```go
import (
    "github.com/rafaelmgr12/go-commerce/internal/domain/entity"
    "github.com/rafaelmgr12/go-commerce/internal/usecase"
    "github.com/rafaelmgr12/go-commerce/internal/usecase/rules"
)

// Create a new product, payment, and order
product := entity.NewProduct("home-appliance", 1100)
payment := entity.NewPayment("Boleto", 1100)
order := entity.NewOrder(product, payment, nil)

// Set up the chain of rules
r1 := &rules.FreeShippingRule{}
r2 := &rules.FragileProductRule{}
r3 := &rules.ChildProductRule{}
r4 := &rules.BoletoDiscountRule{}

r1.SetNext(r2)
r2.SetNext(r3)
r3.SetNext(r4)

// Process the order
usecase.ProcessOrder(order, r1)

```

## **Contributing**

Contributions are welcome! Please read the **[CONTRIBUTING.md](https://chat.openai.com/CONTRIBUTING.md)** for details on how to contribute.

## **License**

This project is licensed under the MIT License. See the **[LICENSE.md](https://chat.openai.com/LICENSE.md)** file for details.
