# go

> source: [https://go.dev/tour/](https://go.dev/tour/)

## basic types

```go
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32, represents a Unicode code point
float32 float64
complex64 complex128
```

## zero values

```
0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.
```

## constants

constants can be character, string, boolean, or numeric values. 

```go
const Pi = 3.14
const (
	Big = 1 << 100
	Small = Big >> 99
)
```

## switch

```go
switch os := runtime.GOOS; os {
case "darwin":
	fmt.Println("OS X.")
case "linux":
	fmt.Println("Linux.")
default:
	// freebsd, openbsd,
	// plan9, windows...
	fmt.Printf("%s.\n", os)
}
```

## defer

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

## pointers

The type *T is a pointer to a T value. Its zero value is nil. 

```go
i, j := 42, 2701
p := &i  // point to i
fmt.Println(*p) // read i through the pointer
*p = 21 // set i through the pointer
fmt.Println(i)  // see the new value of i
p = &j // point to j
*p = *p / 37 // divide j through the pointer
fmt.Println(j) // see the new value of j
```

## struct 

```go
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{} 	   // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)
```

## arrays

```go
var a [2]string
a[0] = "Hello"
a[1] = "World"
fmt.Println(a[0], a[1])
fmt.Println(a)
```

## slices

 This selects a half-open range which includes the first element, but excludes the last one. 

```go
primes := [6]int{2, 3, 5, 7, 11, 13}

var s []int = primes[1:4] // 1 to 3 from primes
fmt.Println(s)
```

The length of a slice is the number of elements it contains.  
The capacity of a slice is the number of elements in the underlying array,

```go
len(s)
cap(s)
```

The zero value of a slice is nil. 

### make

The make function allocates a zeroed array and returns a slice that refers to that array: 

```go
a := make([]int, 5)  // len(a)=5
```

To specify a capacity, pass a third argument to make:

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

### append

the slice grows as needed

```go
s = append(s, 2, 3, 4)
```

## range

 The range form of the for loop iterates over a slice or map. 

```go
for i, v := range pow {
```
>  If you only want the index, you can omit the second variable: for i := range pow

## map

 A map maps keys to values  
The zero value of a map is nil. A nil map has no keys, nor can keys be added

```go
m = make(map[string]Vertex)
```

insert: 

```go
m[key] = elem
```

retrieve:

```go
elem = m[key]
```

delete:

```go
delete(m, key)
``
test key present:

```go
elem, ok := m[key]
	// If key is in m, ok is true. If not, ok is false
```

## functions

```go
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

## methods

```go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

Methods with pointer receivers can modify the value to which the receiver points 

```go
func (v *Vertex) Scale(f float64) {
```

## interface

An interface type is defined as a set of method signatures
A value of interface type can hold any value that implements those methods

```go
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}
```

 Under the hood, interface values can be thought of as a tuple of a value and a concrete type: (value, type)

 ### empty interface

 Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}

```go
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

### type assertion

 A type assertion provides access to an interface value's underlying concrete value: 
 
```go
t, ok := i.(T)
f, ok := i.(float64)
```

type switch:

```go
switch v := i.(type) {
case T:
// here v has type T
case S:
// here v has type S
default:
// no match; here v has the same type as i
}
```

## errors

```go
type error interface {
    Error() string
}
```

## goroutine

 A goroutine is a lightweight thread managed by the Go runtime

```go
go f(x, y, z)
```

## channels

Channels are a typed conduit through which you can send and receive values with the channel operator, <-

```go
ch := make(chan int)
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v
```

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables

```go
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}
```

### buffered channels

Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel

```go
ch := make(chan int, 100)
```

you can close a channel with: 

```go
close(c)
```

For example a loop for i := range c receives values from the channel repeatedly until it is closed

### select 

The select statement lets a goroutine wait on multiple communication operations

```go
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

### sync.Mutex

Make sure only one goroutine can access a variable at a time to avoid conflicts

```go
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
```

## theory

### package

A package is a collection of source files in the same directory that are compiled together

### module

A module is a collection of related Go packages that are released together

## next

[go language specification](https://go.dev/ref/spec)

Effective GO

## channels 

Simple Wait to complete

```go
c := make(chan int)  // Allocate a channel.
// Start the sort in a goroutine; when it completes, signal on the channel.
go func() {
    list.Sort()
    c <- 1  // Send a signal; value does not matter.
}()
doSomethingForAWhile()
<-c   // Wait for sort to finish; discard sent value.
```

RPC

```go
type Request struct {
    args        []int
    f           func([]int) int
    resultChan  chan int
}

func sum(a []int) (s int) {
    for _, v := range a {
        s += v
    }
    return
}

request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
// Send request
clientRequests <- request
// Wait for response.
fmt.Printf("answer: %d\n", <-request.resultChan)

func handle(queue chan *Request) {
    for req := range queue {
        req.resultChan <- req.f(req.args)
    }
}
```

Multiple cpus

```go
type Vector []float64

// Apply the operation to v[i], v[i+1] ... up to v[n-1].
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
    for ; i < n; i++ {
        v[i] += u.Op(v[i])
    }
    c <- 1    // signal that this piece is done
}

const numCPU = runtime.NumCPU() // number of CPU cores

func (v Vector) DoAll(u Vector) {
    c := make(chan int, numCPU)  // Buffering optional but sensible.
    for i := 0; i < numCPU; i++ {
        go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
    }
    // Drain the channel.
    for i := 0; i < numCPU; i++ {
        <-c    // wait for one task to complete
    }
    // All done.
}
```

## recover (panic)

```go
func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```

## fmt

```
- %v default value  
- %d integer  
- %f float
- %t boolean 
- %c rune 
- %s string 
- %% literal percent 
- %T variable type
- %#v golang format, ready to paste  
- %p pointer  
- %q quoted string  

%.2f  
100.567 would print as 100.57

default right alingment  
%8.2f > "  100.57"  

left alignment  
%-8.2f > "100.57  "  

zero padding  
%08d > "00000123"
```

## generics

[https://go.dev/doc/tutorial/generics](https://go.dev/doc/tutorial/generics)

```golang
// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
fmt.Printf("Generic Sums: %v and %v\n",
   	SumIntsOrFloats[string, int64](ints),
   	SumIntsOrFloats[string, float64](floats))
```

types can be infered in this case

```golang
fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
   	SumIntsOrFloats(ints),
   	SumIntsOrFloats(floats))
```

type constraint

```golang
type Number interface {
    int64 | float64
}
// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

fmt.Printf("Generic Sums with Constraint: %v and %v\n",
   	SumNumbers(ints),
   	SumNumbers(floats))
```

## gopath

By default GOPATH value is $HOME/go but you can change it. Now that modules are produuction ready your code don't need to live in GOPATH/src but is still a good practice

```bash
export GOPATH=$Home/go:$Home/projects
```

Also is a good practice to include go/bin in your path

```bash
export PATH=$PATH:$GOPATH/bin
```

Test your installation

```bash
go version
```

## go modules

init go modules

```bash
go mod init github.com/mamcer/hello
```

go build, go test, and other package-building commands add new dependencies to go.mod as needed

list modules

```bash
go list -m all
```

upgrade a module

```bash
go get golang.org/x/text
```

list versions of a specific package

```
go list -m -versions rsc.io/sampler
```

get a specific version

```
go get rsc.io/sampler@v1.3.1
```

cleanup unused dependencies

```
go mod tidy
```

## project structure

Three related approaches

[https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

```bash
cmd/
docs/
internal/
pkg/
.gitignore
LICENSE.md
Makefile
README.md
go.mod
go.sum
```

```bash
$GOPATH/src/myproject

    ├──cmd/ -- this is where you compose several packages in to main package
    |  ├──foo -- an example would be `foo`
    |  |  ├──main.go
    ├──pkg/ -- this is where put your reusable packages 
    |  ├──pkg1 -- reusable package 1
    |  ├──pkg2 -- reusable package 2
    ├──otherpackage1
    |  ├── ...
    ├──otherpackage2
    |  ├── ...
```

```bash
$GOPATH/
   
    src/
        github.com/user/repo/
            mypkg/
                mysrc1.go
                mysrc2.go
            cmd/mycmd/
                main.go
    bin/
        mycmd	
```

## naming conventions

[https://talks.golang.org/2014/organizeio.slide#6](https://talks.golang.org/2014/organizeio.slide#6)

[https://blog.golang.org/package-names](https://blog.golang.org/package-names)

[https://talks.golang.org/2014/names.slide#1](https://talks.golang.org/2014/names.slide#1)

## go doc

show documentation in terminal for the current directory. Include unexported types

```
go doc --all -u		
```

## test

recursively look for tests

```
go test ./...
```

run tests in paralell

If the test does not have any dependency you can add

```
t.Paralell()
```

To explicitely say it can be run in paralell then:

```
go test mylib/...
```

Should be faster

List tests containing word. Basics in the example

```
go test -list Basics mylib/...
```

Run tests with word. Basics in the example

```
go test -run Basics mylib/...
```

Verbose option

```
go test -v mylib/...
```

Coverage

```
go test -coverprofile cover.out
go tool cover -func cover.out
go test -cover
```

html report

```
go test ./... -covermode=atomic -coverprofile=coverage.out -coverpkg=./... -count=1
go tool cover -html=coverage.out
```

Memory profile

```
go test -memprofile mem.out -memprofilerate 1 mylib
sudo apt install graphviz
go tool pprof
go tool pprof -web mylib.test mem.out
```

>  [http://graphviz.org](http://graphviz.org)

CPU profile 

```
go test -cpuprofile cpu.out -count 1000000 mylib
go tool pprof -web mylib.test cpu.out 
```

Benchmark 

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```

```
go test -bench=.
```

## clean workspace

obj/ test/  etc

```
go clean -x
```

## code examples

### concurrency 

example 01 

```go
package main

import (
	"fmt"
	"sync"
)

func runMe(s int) {
	fmt.Printf("hello there %v\n", s)
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(s int) {
			runMe(s)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

example 02

```go
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	vs := []string{"hello", "world", "this", "is", "us"}
	ch := make(chan int, len(vs))
	command := "echo"

	for i, a := range vs {
		go func(ch chan int, command string, a string) {
			fmt.Printf("executing %v: %v\n", i, a)
			cmd := exec.Command(command, a)
			if err := cmd.Run(); err != nil {
				fmt.Printf("error:%v\n", err)
				ch <- 0
			} else {
				ch <- 1
			}
		}(ch, command, a)
	}

	for i := 0; i < len(vs); i++ {
		r := <-ch
		if r == 0 {
			fmt.Printf("%v\n", vs[i])
		}
	}
}
```

### elapsed

```go
package main
import (
	"fmt"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func main() {
	defer elapsed("page")() // <-- The trailing () is the deferred call
	time.Sleep(time.Second * 2)
}
```

### run from api

```go
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"github.com/gin-gonic/gin"
)

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}

func ping(c *gin.Context) {
	cmd := exec.Command("mplayer", "-fs", "-vo", "xv", "-ao", "alsa:device=hdmi", "/home/mario/Videos/Nuovo Cinema Paradiso/	Nuovo.cinema.Paradiso.(1988).BDRip.720p.AC3.X264-CHD-Italian.mkv", "&")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"message": "pong",
	})

	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", out)
	//mplayer -fs -vo xv -ao alsa:device=hdmi /home/mario/Videos/Nuovo\ Cinema\ Paradiso/Nuovo.cinema.Paradiso.\(1988\).BDRip.	720p.AC3.X264-CHD-Italian.mkv
}

func main() {
	g := gin.Default()
	g.GET("/ping", ping)
	g.OPTIONS("/ping", preflight)
	g.Run(":5000")
}
```

### go file handler

```go
package main

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

type MyFile struct {
	Name string
}

func readDir() []MyFile {
	file, _ := os.Open("/data/data/com.termux/files/home/storage/pictures")
	var files []MyFile
	defer file.Close()
	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range list {
		files = append(files, MyFile{name})
	}
	return files
}

func main() {
	r := gin.Default()
	files := readDir()
	http.Handle("/", http.FileServer(http.Dir("/data/data/com.termux/files/home/storage/pictures")))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, files)
	})
	go r.Run()
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
```

### str format

```go
package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)
	fmt.Printf("struct2: %+v\n", p)
	fmt.Printf("struct3: %#v\n", p)
	fmt.Printf("type: %T\n", p)
	fmt.Printf("bool: %t\n", true)
	fmt.Printf("int: %d\n", 123)
	fmt.Printf("bin: %b\n", 14)
	fmt.Printf("char: %c\n", 64)
	fmt.Printf("hex: %x\n", 456)
	fmt.Printf("float1: %f\n", 78.9)
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)
	fmt.Printf("str1: %s\n", "\"string\"")
	fmt.Printf("str2: %q\n", "\"string\"")
	fmt.Printf("str3: %x\n", "hex this")
	fmt.Printf("pointer: %p\n", &p)
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
```
output

```
struct1: {1 2}
struct2: {x:1 y:2}
struct3: main.point{x:1, y:2}
type: main.point
bool: true
int: 123
bin: 1110
char: @
hex: 1c8
float1: 78.900000
float2: 1.234000e+08
float3: 1.234000E+08
str1: "string"
str2: "\"string\""
str3: 6865782074686973
pointer: 0xc00009e020
width1: |    12|   345|
width2: |  1.20|  3.45|
width3: |1.20  |3.45  |
width4: |   foo|     b|
width5: |foo   |b     |
sprintf: a string
io: an error
```

### linked list 

```go
package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	info interface{}
	next *Node
}

func insert(d interface{}, n *Node) *Node {
	nn := &Node{info: d, next: nil}
	if n != nil {
		p := n
		for p.next != nil {
			p = p.next
		}
		p.next = nn
	} else {
		n = nn
	}
	return n
}

func show(n *Node) {
	p := n
	for p != nil {
		fmt.Printf("-> %v ", p.info)
		p = p.next
	}
	fmt.Printf("\n")
}

func main() {
	var root *Node = nil
	for i := 0; i < 5; i++ {
		root = insert(rand.Intn(100), root)
	}
	show(root)
}
```

### uiprogress

```go
package main

import (
	"fmt"
	"time"
	"github.com/gosuri/uilive"
)

// https://github.com/gosuri/uilive
func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func main() {
	defer elapsed("page")() // <-- The trailing () is the deferred call
	writer := uilive.New()
	// start listening for updates and render
	writer.Start()
	for i := 0; i <= 100; i++ {
		fmt.Fprintf(writer, "Downloading.. (%d/%d) GB\n", i, 100)
		time.Sleep(time.Millisecond * 5)
	}
	fmt.Fprintln(writer, "Finished: Downloaded 100GB")
	writer.Stop() // flush and stop rendering
}
```

### year

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	var v string
	t := time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	yd := time.Now().YearDay()
	dc := t.YearDay()
	fmt.Printf("%v\n", time.Now().Year())
	for i := 1; i < dc; i++ {
		if i < yd {
			v = "x"
		} else if i == yd {
			v = "o"
		} else {
			v = "-"
		}

		if time.Date(time.Now().Year(), 1, 0, 0, 0, 0, 0, time.UTC).AddDate(0, 0, i).Weekday() == time.Sunday {
			fmt.Printf(" %v", v)
		} else {
			fmt.Printf("%v", v)
		}
	}
	fmt.Printf("\n")
	fmt.Printf("day %v of %v\n", yd, dc)
	fmt.Printf("week %v of %v\n", yd/7+1, dc/7)
	fmt.Printf("year completed: %.1f%%\n", float64(yd)/float64(dc)*100)
}
```

## basic http static file server	

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	p := "8081"
	fmt.Printf("hello handsome, port: %s\n", p)
	http.Handle("/",
		http.StripPrefix("/",
			http.FileServer(http.Dir("./"))))
	log.Fatal(http.ListenAndServe(":"+p, nil))
}
```

## libraries

decimal numbers:  
[github.com/shopspring/decimal](github.com/shopspring/decimal)

view validations:  
[github.com/go-playground/validator/v10](github.com/go-playground/validator/v10)

db:  
[github.com/jmoiron/sqlx](github.com/jmoiron/sqlx)

mysql  
[github.com/go-redis/redis/v8](github.com/go-redis/redis/v8)

test  
[github.com/DATA-DOG/go-sqlmock](github.com/DATA-DOG/go-sqlmock)	