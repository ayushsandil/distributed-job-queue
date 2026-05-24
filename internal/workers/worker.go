package workers

import (
    "fmt"
    "time"

    "distributed-job-queue/internal/queue"
)

func StartWorker() {
    for {
        jobID, err := queue.PopJob()

        if err != nil {
            time.Sleep(2 * time.Second)
            continue
        }

        fmt.Println("Processing Job:", jobID)

        time.Sleep(3 * time.Second)

        fmt.Println("Completed Job:", jobID)
    }
}
