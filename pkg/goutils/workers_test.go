package goutils

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestWorkersCancelOthers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < 10; i++ {
		go func(id int) {
			select {
			case <-ctx.Done():
				fmt.Printf("Worker %d stopped\n", id)
				return
			case <-time.After(time.Duration(rand.Intn(5)) * time.Second):
				fmt.Printf("Worker %d finished\n", id)
				cancel()
				return
			}
		}(i)
	}
	time.Sleep(8 * time.Second)
}

type Task struct {
	ID     int
	Result any
	Error  error
}

func TestWorkersFirstFinish(t *testing.T) {
	const numWorkers = 100
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	results := make(chan Task)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			duration := time.Duration(rand.Intn(5)) * time.Second

			select {
			case <-ctx.Done():
				return

			case <-time.After(duration):
				select {
				case results <- Task{ID: id, Result: rand.Intn(100)}:
				case <-ctx.Done():
					return
				}
			}
		}(i)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	loop := true
	for loop {
		select {
		case task := <-results:
			if task.Error != nil {
				continue
			}
			cancel()
			loop = false
			break

		case <-done:
			loop = false
			break
		}
	}

	<-done
}
