package main

import (
    "distributed-job-queue/internal/queue"
    "distributed-job-queue/internal/workers"
)

func main() {
    queue.InitRedis()
    workers.StartWorker()
}
