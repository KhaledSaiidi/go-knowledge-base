# Go Cheat Sheet: Patterns + Runtime

This file groups behavior-oriented examples from the repository. Open it when you need fast recall for interfaces, cleanup flow, helper patterns, goroutines, request context, and `net/http`.

## Table of Contents

- [Interfaces](#interfaces)
- [Defer](#defer)
- [Small Helper Patterns](#small-helper-patterns)
- [Concurrency](#concurrency)
- [Context Package](#context-package)
- [Basic Web Server (`net/http`)](#basic-web-server-nethttp)

## Interfaces

### Mental model

- Interfaces describe behavior, not fields.
- Types satisfy interfaces implicitly by implementing the required methods.
- A slice of interfaces lets different concrete types share one call path.
- Small interfaces are useful at dependency boundaries like notifiers and databases.

### Syntax patterns

```go
type Player interface {
	KickBall()
}

type FootballPlayer struct {
	stamina int
	power   int
}

func (f FootballPlayer) KickBall() {
	fmt.Println("Kick the ball", f.stamina+f.power)
}

type CR7 struct {
	stamina int
	power   int
	SUI     int
}

func (f CR7) KickBall() {
	fmt.Println("CR7 Kick the ball", f.stamina+f.power*f.SUI)
}

team := []Player{
	FootballPlayer{stamina: 5, power: 7},
	CR7{stamina: 10, power: 10, SUI: 10},
}

for _, player := range team {
	player.KickBall()
}
```

### Common usage

```go
type AccountNotifier interface {
	NotifyAccountCreated(context.Context, Account) error
}

type AccountHandler struct {
	AccountNotifier AccountNotifier
}

func (h *AccountHandler) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.AccountNotifier.NotifyAccountCreated(r.Context(), account); err != nil {
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
}
```

### Gotchas

- If the method set does not match, the type does not satisfy the interface.
- Pointer receiver methods matter when you assign a concrete type to an interface.
- Keep the interface small, like the `dbcontract` example.

### Quick memory trigger

Define behavior once, swap concrete implementations behind it.

## Defer

### Mental model

- `defer` schedules cleanup when the current function returns.
- Put cleanup right after the resource is acquired successfully.
- A deferred closure can measure total function time.
- `recover` only works inside a deferred function.

### Syntax patterns

```go
func readFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}

func processData(data []int) {
	start := time.Now()
	defer func() {
		fmt.Println("Data processing completed in", time.Since(start))
	}()

	for _, d := range data {
		fmt.Println("Processing data:", d)
	}
}
```

### Common usage

```go
func safeOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("Something went wrong!")
}
```

### Gotchas

- Do not `defer` before checking the open/create error.
- `recover` outside a deferred function does nothing.

### Quick memory trigger

Acquire first, then defer cleanup immediately.

## Small Helper Patterns

### Mental model

- Small helper functions can flatten repetitive error handling.
- A generic `Must` works well in tiny programs and one-shot flows.
- Combine `Must` with `defer` to keep resource pipelines compact.
- Fail-fast helpers belong at program edges, not inside reusable libraries.

### Syntax patterns

```go
func Must[T any](x T, err error) T {
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	return x
}

func CheckErr(err error) {
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}
```

### Common usage

```go
src := "./template.txt"
dst := "./template_copy.txt"

r := Must(os.Open(src))
defer r.Close()

w := Must(os.Create(dst))
defer func() {
	CheckErr(w.Close()) // close once, but still surface the error
}()

Must(io.Copy(w, r))
```

### Gotchas

- `Must` exits the process, so do not use it where callers should recover.
- Avoid closing the same file twice; pair one resource with one close path.

### Quick memory trigger

`Must` is for fast fail, not graceful recovery.

## Concurrency

### Mental model

- Goroutines run work concurrently.
- Channels move data between goroutines and can be buffered or unbuffered.
- `select` lets you wait on multiple events like work or timeout.
- Shared mutable state still needs `sync.Mutex` protection.

### Syntax patterns

```go
type Order struct {
	ID     int
	Status string
	mutex  sync.Mutex
}

func processOrders(
	orderChan <-chan *Order,
	processedOrderChan chan<- *Order,
	wg *sync.WaitGroup,
) {
	defer func() {
		wg.Done()
		close(processedOrderChan)
	}()

	for order := range orderChan {
		order.Status = "Processed"
		processedOrderChan <- order
	}
}
```

### Common usage

```go
wg := sync.WaitGroup{}

orderChan := make(chan *Order, 20)         // buffered
processedOrderChan := make(chan *Order, 20)

wg.Add(1)
go func() {
	defer wg.Done()
	defer close(orderChan)

	for _, order := range generateOrders(20) {
		orderChan <- order
	}
}()

wg.Add(1)
go processOrders(orderChan, processedOrderChan, &wg)

wg.Add(1)
go func() {
	defer wg.Done()
	for {
		select {
		case order, ok := <-processedOrderChan:
			if !ok {
				return
			}
			fmt.Printf("Processed order %d with status: %s\n", order.ID, order.Status)
		case <-time.After(10 * time.Second):
			fmt.Println("No processed orders received")
			return
		}
	}
}()
```

### Gotchas

- Call `wg.Add(...)` before the goroutine starts.
- Close a channel from the sender side, and only once.
- Any shared counter or shared struct update needs locking to avoid races.

### Quick memory trigger

Goroutines do work, channels coordinate, mutexes protect shared writes.

## Context Package

### Mental model

- `context.Context` carries cancellation, deadlines, and request-scoped metadata.
- Derive child contexts with `WithTimeout`, then `defer cancel()`.
- `select` on `ctx.Done()` when work might block.
- Use a typed key for `WithValue`.

### Syntax patterns

```go
ctx := context.Background()

ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()

done := make(chan struct{})
go func() {
	time.Sleep(3 * time.Second)
	close(done)
}()

select {
case <-done:
	fmt.Println("Called the API")
case <-ctxWithTimeout.Done():
	fmt.Println("API call took too long", ctxWithTimeout.Err())
}
```

### Common usage

```go
func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	responseChannel := make(chan Response)
	go func() {
		val, err := fetchThirdPartyStuffWithCanBeSlow(userID)
		responseChannel <- Response{value: val, err: err}
	}()

	select {
	case <-ctx.Done():
		return 0, fmt.Errorf("fetching data took too long: %w", ctx.Err())
	case resp := <-responseChannel:
		return resp.value, resp.err
	}
}

type key int

const UserKey key = 0

ctxWithValue := context.WithValue(context.Background(), UserKey, "123")
userID, _ := ctxWithValue.Value(UserKey).(string)

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
```

### Gotchas

- If the slow function does not accept the context, the caller stops waiting but the goroutine may still keep running.
- `WithValue` is for request metadata, not optional function arguments.
- Raw string keys are easy to collide; a custom key type is safer.

### Quick memory trigger

Create child context, defer cancel, watch `Done()`.

## Basic Web Server (`net/http`)

### Mental model

- `http.NewServeMux()` maps routes to handlers.
- Handlers decode input, validate, mutate/read state, then write a response.
- `PathValue("id")` pulls route parameters from patterns like `GET /users/{id}`.
- Shared in-memory state needs `sync.RWMutex`.

### Syntax patterns

```go
type User struct {
	Name string `json:"name"`
}

var userCache = make(map[int]User)
var cacheMutex sync.RWMutex
var nextUserID = 1

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("GET /users/{id}", getUsers)
	mux.HandleFunc("GET /allusers", getAllUsers)
	mux.HandleFunc("PUT /users/{id}", updateUser)
	mux.HandleFunc("DELETE /users/{id}", deleteUsers)

	http.ListenAndServe(":8080", mux)
}
```

### Common usage

```go
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	id := nextUserID
	userCache[id] = user
	nextUserID++
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	cacheMutex.RLock()
	user, exists := userCache[id]
	cacheMutex.RUnlock()

	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
```

### Gotchas

- `len(userCache)+1` is not a safe ID strategy when deletes happen; use a separate counter.
- A map header copy still points to the same underlying map, so do not unlock and then keep reading it unsafely.
- Use `RLock` for reads and `Lock` for writes on shared cache state.

### Quick memory trigger

Decode, validate, lock, mutate/read, respond.
