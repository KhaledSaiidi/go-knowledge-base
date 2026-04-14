package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	// ctx := context.WithValue(context.Background(), "user-id", 123)
	userID := 123
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Result:", val)
	fmt.Printf("Time taken: %s\n", time.Since(start))

}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	// val := ctx.Value("user-id") // => 123
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel()
	responseChannel := make(chan Response)
	go func() {
		val, err := fetchThirdPartyStuffWithCanBeSlow(userID)
		responseChannel <- Response{value: val, err: err}
	}()
	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from 3rd party took too long: %w", ctx.Err())
		case resp := <-responseChannel:
			return resp.value, resp.err
		}
	}
}

func fetchThirdPartyStuffWithCanBeSlow(userID int) (int, error) {
	time.Sleep(time.Millisecond * 500)

	return userID, nil
}
