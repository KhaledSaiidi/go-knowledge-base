package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// controlling timeouts
// cancelling go routines
// passing metadata across your Go application

func main() {
	ctx := context.Background()
	exampleTimeout(ctx)
	exampleWityhValues(ctx)

	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func exampleTimeout(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 3)
		close(done)
	}()
	select {
	case <-done:
		fmt.Println("Called the API")
	case <-ctxWithTimeout.Done():
		fmt.Println("API call took too long", ctxWithTimeout.Err())
		// logic
	}

}

func exampleWityhValues(ctx context.Context) {
	type key int
	const UserKey key = 0

	ctxWithValue := context.WithValue(ctx, UserKey, "123")
	if userID, ok := ctxWithValue.Value(UserKey).(string); ok {
		fmt.Println("User ID:", userID)
	} else {
		fmt.Println("User ID not found in context")
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "Hello, World!")
	case <-ctx.Done():
		http.Error(w, "Request timed out", http.StatusGatewayTimeout)
	}
}
