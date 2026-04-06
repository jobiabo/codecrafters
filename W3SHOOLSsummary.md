🟢 GO (GOLANG)

## 1. Definition

Go is a compiled, statically typed language developed by Google.
Designed for performance, simplicity, and concurrency.


## 2. Key Features

Fast compilation
Garbage collection
Strong typing
Built-in concurrency (goroutines)
Cross-platform


## 3. Program Structure

```go
package main
import "fmt"

func main() {
    fmt.Println("Hello")
}
```

package main → entry package
main() → starting point


## 4. Variables
```go
var x int = 10
y := 20
var → explicit
:= → shorthand (inside functions only)
```

## 5. Constants

```go
const PI = 3.14
Immutable values
```

## 6. Data Types
``go
int → integers
float64 → decimals
string → text
bool → true/false
```

## 7. Operators
Arithmetic: + - * / %
Comparison: == != > <
Logical: && || !


## 8. Control Statements
```go
If-Else
if x > 0 {
}
Switch
switch x {
case 1:
default:
}
```


## 9. Loops
```go
for i := 0; i < 5; i++ {
}
Only loop in Go is for

```

## 10. Functions

```go
func add(a int, b int) int {
    return a + b
}
Supports multiple return values

```

## 11. Arrays
```go
var arr [3]int = [3]int{1,2,3}
Fixed size
```


## 12. Slices

s := []int{1,2,3}
Dynamic arrays



## 13. Maps
m := map[string]int{"a":1}
Key-value pairs



## 14. Structs

```go
type Person struct {
    name string
    age int
}
Custom data types

```

## 15. Concurrency
Goroutines → lightweight threads
Channels → communication between goroutines



## 16. Error Handling
```go
val, err := func()
if err != nil {
}
No exceptions
Errors handled explicitly
```


## 17. Important Notes
No classes → uses structs
No inheritance
No operator overloading
No exceptions



## 18. Common Packages
fmt → input/output
math → math functions
net/http → web



## 19. Uses of Go
Backend development
Cloud systems
Networking tools



## 20. One-Line Summary

Go = simple + fast + concurrent programming language