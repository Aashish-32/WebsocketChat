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
├── .vscode/                # Visual Studio Code configuration files
├── cmd/
│   └── web/                # Main application entry point
│       └── main.go         # Starts the web server
├── html/                   # HTML templates for the frontend
│   └── index.html          # Main chat interface
├── internal/
│   └── handlers/           # Application logic and WebSocket handlers
│       ├── client.go       # Manages individual WebSocket clients
│       ├── hub.go          # Manages client connections and broadcasting
│       └── serveWs.go      # Handles WebSocket requests
├── .gitignore              # Specifies files to ignore in Git
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
└── README.md               # Project documentation
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


