# Distributed Job Queue & Workflow Engine

A distributed asynchronous job processing system built using Go, Redis, and PostgreSQL.

## Features

- Asynchronous job processing
- Worker pool architecture
- Redis-backed queue
- PostgreSQL integration
- Dockerized local setup

## Run Locally

```bash
docker compose up -d
make run-api
make run-worker
```
