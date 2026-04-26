# Go Cheat Sheet: Basics

This file groups the syntax-heavy building blocks from the repository. Open it when you need fast recall for values, collections, structs, and config construction.

## Table of Contents

- [Rune](#rune)
- [Pointers](#pointers)
- [Slices](#slices)
- [Maps](#maps)
- [Structs](#structs)
- [Efficient Struct / Functional Options](#efficient-struct--functional-options)

## Rune

### Mental model

- Double quotes create a `string`.
- Single quotes create a `rune` (`int32` code point).
- `len(string)` counts bytes, not visible characters.
- `for _, r := range s` walks Unicode code points.

### Syntax patterns

```go
a := "S"       // string
b := 'S'       // rune
c := string(b) // rune -> string

fmt.Println(a)
fmt.Println(b)
fmt.Println(c)

for _, r := range "hello ❗" {
	fmt.Print(string(r) + " ")
}
fmt.Println()

fmt.Println(len("Hello !")) // 7 bytes
fmt.Println(len("hello ❗")) // 9 bytes; ❗ takes 3 bytes in UTF-8
```

### Common usage

```go
log := fmt.Sprintf("a: %T, b: %T, c: %T", a, b, c)
for _, r := range log {
	fmt.Print(string(r) + " ")
}
```

### Gotchas

- `'S'` is not a `string`; print or convert it deliberately.
- `len("hello ❗")` is byte length, not rune count.

### Quick memory trigger

Double quotes store text, single quotes store one code point.

## Pointers

### Mental model

- `&x` gives the address of `x`.
- `*p` dereferences a pointer, so you can read or write the pointed value.
- Structs use value semantics by default; pointer receivers share and mutate the same instance.
- Return `*T` when caller and callee should work on the same struct value.

### Syntax patterns

```go
i := 42
p := &i

fmt.Println(p)  // address of i
fmt.Println(*p) // value at that address

*p = 21
fmt.Println(i) // i changed through the pointer

func squareAdd(p *int) {
	*p *= *p
}

type person struct {
	name string
	age  int
}

func initPerson() *person {
	p := person{"Alice", 30}
	return &p
}
```

### Common usage

```go
type User struct {
	username string
	email    string
	file     []byte
}

func (u User) Email() string {
	return u.email // read-only
}

func (u *User) SetEmail(email string) {
	u.email = email // mutate original
}

func (u *User) toString() string {
	return fmt.Sprintf("%s <%s>", u.username, u.email)
}
```

### Gotchas

- If a method must change state, a value receiver changes only its copy.
- Pointer receivers are usually better when the struct is large or shared.
- Slices and maps already reference shared underlying data; structs do not.

### Quick memory trigger

Use pointers when the caller should see the change.

## Slices

### Mental model

- A slice is a lightweight view over an underlying array.
- `append` can reuse the same backing array.
- Variadic functions accept many values, and `slice...` expands a slice into arguments.
- A named slice type can have methods and satisfy interfaces like `sort.Interface`.

### Syntax patterns

```go
s1 := []string{"a", "b", "c"}
s2 := []string{"d", "e", "f"}

fmt.Println(slices.Equal(s1, s2))
fmt.Println(slices.Contains(s2, "d"))

s3 := append(s1, "d")
s4 := append(s1, s2...)

s5 := make([]string, len(s4), 6)
copy(s5, s4)

func addUsers(users ...string) {
	fmt.Println(users)
}
```

### Common usage

```go
type MySlice []int

func (m MySlice) Remove(index int) ([]int, error) {
	if index < 0 || index >= len(m) {
		return m, fmt.Errorf("index out of range: %d", index)
	}
	return append(m[:index], m[index+1:]...), nil
}

type SortNumbers []int

func (n SortNumbers) Len() int           { return len(n) }
func (n SortNumbers) Less(i, j int) bool { return n[i] < n[j] }
func (n SortNumbers) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

type byDec struct{ SortNumbers }

func (n byDec) Less(i, j int) bool {
	return n.SortNumbers[i] > n.SortNumbers[j]
}

nums := SortNumbers{5, 2, 9, 1}
sort.Sort(nums)
sort.Sort(byDec{nums})
```

### Gotchas

- `append(m[:i], m[i+1:]...)` can mutate the original slice view because backing storage may be reused.
- Always bounds-check before removing by index.
- Use `slice...` when passing a slice into a variadic function.

### Quick memory trigger

Slices are views, so `append` may affect shared backing storage.

## Maps

### Mental model

- Map lookup returns a value plus an `ok` flag when you need existence checks.
- A named map type can have helper methods.
- Write with `m[key] = value`.
- Remove with `delete(m, key)`.

### Syntax patterns

```go
type MP map[string]int

func (m MP) Get(key string) (int, bool) {
	value, ok := m[key]
	return value, ok
}

func (m MP) Set(key string, value int) {
	m[key] = value
}

func (m MP) Delete(key string) {
	delete(m, key)
}

myMap := MP{"five": 5, "two": 2}
value, ok := myMap.Get("two")
myMap.Set("three", 3)
myMap.Delete("five")
```

### Common usage

```go
func (m MP) String() string {
	result := "{"
	for key, value := range m {
		result += fmt.Sprintf("%s: %d, ", key, value)
	}
	if len(result) > 1 {
		result = result[:len(result)-2] // trim trailing comma
	}
	return result + "}"
}
```

### Gotchas

- Without `ok`, a missing key and a stored zero value look the same.
- Manual string building needs trailing separator cleanup.

### Quick memory trigger

Read with `ok`, write with assignment, remove with `delete`.

## Structs

### Mental model

- Structs are copied by value unless you pass pointers.
- Value receivers are fine for read-only behavior.
- Pointer receivers are required when you want to update the original value.
- JSON tags describe field names, but the field still has to be exported for `json.Marshal`.

### Syntax patterns

```go
type Book struct {
	title   string    `json:"title"`
	author  string    `json:"author"`
	pages   int       `json:"pages"`
	isSaved bool      `json:"isSaved"`
	savedAt time.Time `json:"savedAt"`
}

func (b Book) getBookInfo() string {
	return fmt.Sprintf("%s (%d pages)", b.title, b.pages)
}

func (b Book) saveBookBurned() {
	b.isSaved = true   // changes only the copy
	b.savedAt = time.Now()
}

func (b *Book) saveBook() {
	b.isSaved = true
	b.savedAt = time.Now()
}
```

### Common usage

```go
book := Book{
	title:   "The Great Gatsby",
	pages:   180,
	isSaved: false,
}

book.saveBookBurned() // value receiver: original stays unchanged
fmt.Println(book.getBookInfo())

book.saveBook() // pointer receiver: original is updated
fmt.Println(book.getBookInfo())

jsonData, _ := json.Marshal(book) // lowercase fields will not be included
fmt.Println(string(jsonData))
```

### Gotchas

- A write on a value receiver only changes the copied struct.
- `time.Time{}` is the zero-value "not set yet" state.
- JSON tags on lowercase fields do not make them marshalable.

### Quick memory trigger

Read by value, write by pointer.

## Efficient Struct / Functional Options

### Mental model

- Keep defaults in one place.
- Use small option functions to override only what changes.
- Apply options before building the final struct.
- Embedding config makes the chosen settings easy to access.

### Syntax patterns

```go
type OptFunc func(*Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 100,
		id:      "default",
		tls:     false,
	}
}

func withTLS(opts *Opts) {
	opts.tls = true
}

func withMaxConn(maxConn int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = maxConn
	}
}

func withID(id string) OptFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}
```

### Common usage

```go
type Server struct {
	Opts
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{Opts: o}
}

s := newServer(withTLS, withMaxConn(200), withID("pg-cluster-1"))
fmt.Printf("Server: %+v\n", s)
```

### Gotchas

- Always start from defaults before custom options.
- Option functions should change config, not perform side effects.
- If you bypass the constructor, you also bypass the default setup.

### Quick memory trigger

Defaults first, then stack `with...` overrides.
