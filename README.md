# Tasker

Tasker is a premium, modern, and offline-first task management application. Built with a robust Go backend and a cutting-edge Svelte 5 frontend, it offers a seamless and high-performance experience for organizing your daily goals.

![Tasker Logo](client/static/images/logo.svg)

## ✨ Features

- **Premium Aesthetics**: High-fidelity glassmorphic UI powered by **Tailwind CSS v4**.
- **Offline-First (PWA)**: Fully functional Progressive Web App with service worker caching for offline use.
- **Svelte 5 Runes**: Built using the latest Svelte 5 reactivity system for extreme performance and type safety.
- **Robust Backend**: Lightweight and secure Go API using SQLite for reliable data persistence.
- **Interactive Dashboards**: Real-time task duration tracking (Days, Hours, Minutes).
- **Secure Authentication**: Bcrypt-hashed password security with persistent session management.

## 🚀 Tech Stack

### Frontend
- **Framework**: [SvelteKit](https://kit.svelte.dev/) (Svelte 5 Runes)
- **Styling**: [Tailwind CSS v4](https://tailwindcss.com/)
- **State Management**: Svelte Stores + Runes
- **PWA**: [vite-plugin-pwa](https://vite-pwa-org.netlify.app/)
- **Icons**: Heroicons (Inline SVG)

### Backend
- **Language**: [Go](https://go.dev/) (Standard Library)
- **Database**: SQLite3
- **Authentication**: JWT-style UUID tokens & Bcrypt
- **Routing**: Gorilla Mux

## 🛠️ Installation

### Prerequisites
- [Go 1.21+](https://go.dev/dl/)
- [Node.js 20+](https://nodejs.org/)
- [GCC](https://gcc.gnu.org/) (for SQLite support)

### Backend Setup
1. Navigate to the `api` directory:
   ```bash
   cd api
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the server:
   ```bash
   go run main.go
   ```
   *The server will start on `localhost:8080` (default) and initialize `db.sqlite`.*

### Frontend Setup
1. Navigate to the `client` directory:
   ```bash
   cd client
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Run the development server:
   ```bash
   npm run dev
   ```
4. Build for production:
   ```bash
   npm run build
   ```

## 📱 PWA Support
Tasker is a full Progressive Web App. You can install it on your mobile device or desktop via the browser's "Add to Home Screen" or "Install" option. It features:
- Offline access to your task list.
- Custom app icon and theme colors.
- Background service worker for resource caching.

## 📜 License
This project is for demonstration purposes. Feel free to use and modify!
