package core

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func ExampleGo(n int) {
	work := func(id int) {
		fmt.Printf("worker %d starting\n", id)
		time.Sleep(time.Second)
		fmt.Printf("worker %d ending\n", id)
	}
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			work(i)
		}(i)
	}
	wg.Wait()
}

func MaybeError(n int) (string, error) {
	if n >= 0 {
		return fmt.Sprintf("asdf %d", n), nil
	} else {
		return "", errors.New("positive only")
	}
}
