package queue

type Producer struct {
	queue *Queue
}

func QueueProducer(queue *Queue) *Producer {
	return &Producer{queue: queue}
}