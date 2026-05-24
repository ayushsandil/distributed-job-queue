package main

import (
    "distributed-job-queue/internal/db"
    "distributed-job-queue/internal/handlers"
    "distributed-job-queue/internal/queue"

    "github.com/gin-gonic/gin"
)

func main() {
    db.InitDB()
    queue.InitRedis()

    router := gin.Default()
    router.POST("/jobs", handlers.CreateJob)
    router.Run(":8080")
}
