# Kobraa Chess

Realtime fullstack chess web application built with Golang.

## Tech Stack

### Backend
- Golang
- Gin
- WebSocket
- GORM
- PostgreSQL

### Frontend
- SSR (Server-Side Rendering)
- HTMX
- TailwindCSS

---

# Features

## Current / Planned Features
- Real-time multiplayer chess
- Friend room matchmaking
- Random matchmaking
- Local multiplayer (1 device)
- Authentication system
- Match history
- ELO rating system
- Leaderboard
- Reconnect handling
- Chess move validation
- Check / checkmate detection

---

# Project Structure

```txt
cmd/
internal/
web/
migrations/
```

## Main Architecture

```txt
internal/
├── config/
├── handlers/
├── middlewares/
├── services/
├── repositories/
├── websocket/
├── domain/
├── models/
├── routes/
└── utils/
```

---

# Architecture Overview

This project uses:
- Fullstack monolithic architecture
- Server-side rendering (SSR)
- Domain-driven structure
- WebSocket realtime communication

---

# Important Principles

## Domain Layer

The `domain/` directory contains:
- chess rules
- move validation
- game engine
- board state
- matchmaking logic

This layer must remain independent from:
- Gin
- Database
- WebSocket implementation
- GORM

---

## Services

Services handle:
- business flow
- orchestration
- communication between layers

---

## Repositories

Repositories are responsible only for:
- database access
- persistence logic

---

## Handlers

Handlers should remain thin.

Handlers only:
- parse requests
- call services
- return responses

---

# Development Setup

## Requirements

- Go 1.25+
- PostgreSQL
- Air

---

# Installation

## Clone Repository

```bash
git clone https://github.com/SulthanFarras/kobraa.chess.git
```

## Enter Project

```bash
cd kobraa.chess
```

## Install Dependencies

```bash
go mod tidy
```

## Install Air

```bash
go install github.com/air-verse/air@latest
```

---

# Environment Variables

Create `.env`

```env
APP_ENV=development
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=kobraa_chess

JWT_SECRET=supersecret
```

---

# Run Application

## Development Mode

```bash
air
```

## Standard Run

```bash
go run ./cmd/app
```

---

# Database Migration

Migration files are stored inside:

```txt
migrations/
```

---

# Development Phases

## Phase 1
- Chess board
- Piece movement
- Turn system

## Phase 2
- Move validation
- Check/checkmate
- Local multiplayer

## Phase 3
- WebSocket realtime
- Friend rooms
- Matchmaking

## Phase 4
- Authentication
- Match history
- ELO system
- Leaderboard

---

# Code Style

## Important Rules

- Keep handlers thin
- Keep business logic inside services/domain
- Do not place chess logic inside handlers
- Do not place database queries inside handlers
- Prefer composition over inheritance
- Avoid overengineering

---

# Realtime Architecture

```txt
Client
  ↓
WebSocket
  ↓
Game Service
  ↓
Chess Domain Engine
  ↓
Broadcast State
```

---

# Deployment Goals

- Single binary deployment
- Docker-ready
- Production-ready architecture
- Low memory usage
- Fast realtime communication

---

# Future Plans

- Spectator mode
- Match replay
- Tournament system
- AI opponent
- Mobile support