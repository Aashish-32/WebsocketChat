# 💬 WebSocket Chat Application

A real-time chat application built with Go (Golang) using WebSockets and the Gin framework. This project demonstrates how to create a simple chat room where multiple users can communicate with each other in real-time.

---

## 🚀 Features

- Real-time communication using WebSockets
- REST API built with Gin framework
- Simple and clean user interface
- Dynamic user join/leave notifications
- Efficient message broadcasting to all connected clients

---

## 🛠 Tech Stack

- **Backend:** Go (Golang), Gin, WebSockets
- **Frontend:** HTML, CSS, JavaScript
- **WebSocket Library:** Gorilla WebSocket
- **Others:** Go modules for dependency management

---

## 📁 Project Structure

```
.
├── main.go             # Entry point of the application
├── hub.go              # Manages clients and message broadcasting
├── client.go           # Handles individual WebSocket clients
├── templates/
│   └── index.html      # Frontend UI
└── static/
    └── main.js         # Client-side WebSocket handling
```



---

## 🚴‍♂️ Getting Started

### Prerequisites

- Go 1.16 or higher installed

### Installation & Running

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Aashish-32/WebsocketChat.git
   cd WebsocketChat


