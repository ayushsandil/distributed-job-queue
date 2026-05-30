package main

import (
	"sync"

	"distributed-job-queue/internal/queue"
	"distributed-job-queue/internal/workers"
)

func main() {

	queue.InitRedis()

	queues := []string{
		"email_queue",
		"reporting_queue",
		"etl_queue",
		"analytics_queue",
		"notification_queue",
	}

	var wg sync.WaitGroup

	for _, q := range queues {

		wg.Add(1)

		go func(queueName string) {
			defer wg.Done()
			workers.StartWorker(queueName)
		}(q)
	}

	wg.Wait()
}