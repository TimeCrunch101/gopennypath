# 💰 GoPennyPath

GoPennyPath is a modern personal finance tracker that helps users log transactions, monitor balances, schedule recurring bills, and tag spending with customizable categories.

Built with:
- 🧠 **Go** (Fiber + GORM) backend
- 🎨 **Vue 3** (Composition API + Pinia) frontend
- 💾 **MySQL** for persistent storage

---

## 📦 Features (MVP)

- ✅ **Transaction Logging** — Add, edit, and delete individual transactions.
- 📊 **Balance Tracking** — View your current balance updated by transactions.
- 🔁 **Recurring Bills** — Automatically logs bill transactions on scheduled dates.
- 🏷️ **Categories** — Categorize spending with default and user-created tags.

---

## 🚀 Project Structure

```
gopennypath/
├── backend/         # Go API with GORM models and Fiber handlers
├── frontend/        # Vue 3 SPA with Composition API
├── db/              # SQL schema and seed data
├── docker-compose.yml
└── .env.example     # Sample environment variables
```

---

## 🛠 Tech Stack

| Layer      | Stack                             |
|------------|------------------------------------|
| Backend    | Go, Fiber, GORM, MySQL             |
| Frontend   | Vue 3, Composition API, Pinia      |
| Database   | MySQL                              |
| DevOps     | Docker, docker-compose             |

---

## 🧪 Setup Instructions

### 🔧 Prerequisites

- Go 1.20+
- Node.js 18+
- Docker + docker-compose
- MySQL 8.x (if not using Docker)

### ⚙️ Environment

Copy and configure `.env` file:

```bash
cp .env.example .env
```

### 🐳 Docker Setup (Recommended)

```bash
docker-compose up --build
```

- Frontend: [http://localhost:5173](http://localhost:5173)
- Backend API: [http://localhost:3000](http://localhost:3000)

### 🧱 Manual Setup (Alternative)

**Backend**

```bash
cd backend
go run cmd/main.go
```

**Frontend**

```bash
cd frontend
npm install
npm run dev
```

---

## 📂 Database

- Schema and seed files: `db/schema.sql`, `db/seed.sql`
- GORM auto-migrates models on startup

---

## 📌 Roadmap

- [ ] User authentication & sessions
- [ ] Multi-account support
- [ ] Reports / analytics dashboard
- [ ] Export data to CSV
- [ ] Dark mode 🌙

---

## 📝 License

MIT License © 2025 Aaron Allen & GoPennyPath Contributors
