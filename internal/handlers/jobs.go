package handlers

import (
    "net/http"

    "distributed-job-queue/internal/models"
    "distributed-job-queue/internal/queue"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

func CreateJob(c *gin.Context) {
    var job models.Job

    if err := c.BindJSON(&job); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    job.ID = uuid.New().String()
    job.Status = "PENDING"

    err := queue.PushJob(job.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, job)
}
