# Project Summary: WS-System

## 📦 Requirements

- [Taskfile](https://taskfile.dev/installation/)
- Docker
- Docker Compose

## ⚙️ Available Commands

### 🔹 `task test`
Runs unit and integration tests.

### 🔹 `task build-and-up`
Starts the system.  
Access the web interface at: [http://localhost:8080](http://localhost:8080)

---

## 🌐 Web Application Overview

### 🔐 Authentication
- **Username:** `Marcos`  
- **Password:** `Marcos`

### 🖥️ User Interface Layout
- The screen is divided into **two halves**:
  - **Left Side:** Displays the state of the `Ingestor` service.
  - **Right Side:** Displays the state of the `Publisher` service.
- **Sending a Message:**  
  Type your message into the `"Mensaje"` input field and press `Enter`. The message will be transmitted from the Ingestor to all connected Publisher clients via Redis.

---

## 🧩 Service Architecture

### 🔸 Ingestor
- Handles validation and WebSocket connections.
- On receiving a message, it publishes the message via **Redis Pub/Sub** to a shared channel.

### 🔸 Publisher
- Subscribes to the Redis channel.
- Broadcasts any received messages to all currently connected WebSocket clients.

---

## 🛠️ Architecture & Design

- Layered architecture with a clear separation between:
  - **Services** (core logic)
  - **Handlers** (routing and I/O)
- Shared modules and logic between services to ensure consistency and reuse.
- Code is structured for clarity and testability.

---

## 💡 Technologies Used

- **[Fiber](https://gofiber.io/)** – Web framework for HTTP and WebSocket routing.
- **JWT (JSON Web Tokens)** – Used for secure authentication.
- **Redis Pub/Sub** – Used to relay messages between services.
- **Zerolog** – Lightweight structured logging for Go.

---

## ✅ Testing Strategy

- **Unit Tests:** Validate individual components and business logic.
- **Integration Tests:** Validate end-to-end functionality including Redis, JWT, and WebSocket connections.

---
