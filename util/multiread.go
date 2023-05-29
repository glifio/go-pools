package util

import (
	"sync"
)

type TaskFunc func() (interface{}, error)

// @dev executes a list of tasks in parallel and returns the results.
func Multiread(tasks []TaskFunc) ([]interface{}, error) {
	var wg sync.WaitGroup
	results := make([]interface{}, len(tasks))
	errChan := make(chan error, 1)

	for i, task := range tasks {
		wg.Add(1)
		go func(i int, task TaskFunc) {
			defer wg.Done()
			result, err := task()
			if err != nil {
				// Stop on first error.
				select {
				case errChan <- err:
				default:
				}
				return
			}
			results[i] = result
		}(i, task)
	}

	wg.Wait()
	close(errChan)

	// If there was an error, return it.
	if err, ok := <-errChan; ok {
		return nil, err
	}

	return results, nil
}
