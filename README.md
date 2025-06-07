# 🔐 LockBox
A secure, minimal password manager built with React, TypeScript, Go, and PostgreSQL. LockBox allows users to store, update, and delete login credentials with an intuitive UI that's safely encrypted with AES-256.

## 🚀 Features
🧑‍💻 Frontend (React + TypeScript):
Responsive and clean UI for vault management, built with reusable components and modern best practices.

⚙️ Backend (Go, net/http):
RESTful API built with Go for handling secure CRUD operations, encryption, and validation.

🛢️ Database (PostgreSQL):
Relational database schema storing encrypted credentials with timestamps.

🔐 AES-256 Encryption:
Passwords are encrypted before storage using Go’s crypto/aes library for secure storage at rest.

## 🧹 Future Improvements
- Add user authentication / multi-user support

- Password masking toggle

- Session timeout + lock screen

- Clipboard auto-clear after copying passwords

- Deployment via Docker or Fly.io
