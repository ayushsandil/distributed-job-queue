package workers

import (
	"fmt"
	"time"

	"distributed-job-queue/internal/queue"
)

func StartWorker(queueName string) {

	for {

		jobID, err := queue.PopJob(queueName)

		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Printf(
			"[%s] Processing Job: %s\n",
			queueName,
			jobID,
		)

		time.Sleep(3 * time.Second)

		fmt.Printf(
			"[%s] Completed Job: %s\n",
			queueName,
			jobID,
		)
	}
}