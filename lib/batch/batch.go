package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var ids = make(chan int64)
	var i int64
	var wg sync.WaitGroup
	var mx sync.Mutex
	for i = 0; i < pool; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for id := range ids {
				user := getOne(id)

				mx.Lock()
				res = append(res, user)
				mx.Unlock()
			}
		}()
	}

	for i = 0; i < n; i++ {
		ids <- i
	}

	close(ids)
	wg.Wait()
	return
}
