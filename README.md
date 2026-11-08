# Smart Pot Irrigation — Backend Service

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Docker Compose](https://img.shields.io/badge/Docker-Compose-2496ED?style=flat&logo=docker)](https://www.docker.com)
[![MQTT](https://img.shields.io/badge/MQTT-Mosquitto_2.0-660066?style=flat&logo=eclipsemosquitto)](https://mosquitto.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18_Alpine-4169E1?style=flat&logo=postgresql)](https://www.postgresql.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

Backend telemetry ingestion service and REST API foundation for the **Smart Pot Irrigation System**, built with **Go**, **Eclipse Mosquitto (MQTT)**, and **PostgreSQL**.

---

## 🏗️ Project Architecture

```text
  [ ESP32-S3 IoT Node ]
           │
           │ MQTT Pub (smart_irrigation/+/telemetry)
           ▼
┌────────────────────────────────────────────────────────┐
│             SMART POT IRRIGATION BACKEND               │
│                                                        │
│  [ Mosquitto Broker ] ──> [ Go Ingestion Worker ]      │
│         :1883                      │                   │
│                                    ▼                   │
│  [ Go REST Engine ]   <──> [ PostgreSQL 18 ]           │
│         :8080                      :5432               │
└────────────────────────────────────────────────────────┘
           │
           │ REST API (JSON)
           ▼
  [ Web App / Mobile Clients ]
```

---

## 🚀 Quick Start

### 1. Local Development
```bash
# Run backend server directly
go run ./cmd/server
```
Server will be available at `http://localhost:8080`.

### 2. Infrastructure via Docker Compose
```bash
# Start Mosquitto Broker & PostgreSQL database
docker compose up -d mosquitto postgres
```

---

## 📡 Base Endpoints

| Method | Endpoint | Description | Status |
|---|---|---|---|
| `GET` | `/healthz` | System health check and stage verification | ✅ Initialized |
| `GET` | `/api/v1/telemetry/latest` | Real-time sensor readings | ⏳ Next Phase |
| `GET` | `/api/v1/telemetry/history` | Historical time-series query | ⏳ Next Phase |
| `GET` | `/api/v1/export/csv` | Research dataset CSV export | ⏳ Next Phase |

---

## 📂 Directory Structure

```text
smart-pot-irrigation-backend/
├── cmd/
│   └── server/main.go       # Server entrypoint & graceful shutdown
├── internal/
│   ├── api/router.go        # HTTP ServeMux router & CORS
│   ├── config/config.go     # Environment variable loader
│   ├── models/telemetry.go  # Telemetry data contracts (ICD standard)
│   └── database/            # Database layer
├── mosquitto/
│   └── config/              # Eclipse Mosquitto configuration
├── scripts/
│   └── init.sql             # PostgreSQL DDL schema & seed
├── Dockerfile               # Production container build
├── docker-compose.yml       # Stack orchestration
├── go.mod
└── README.md
```

---

## 📄 License
MIT License. Developed as part of the Smart Pot Precision Irrigation & TinyML Research.
