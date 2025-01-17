First of install go from [Here](https://go.dev/doc/install)

In old way to run main.go file 
`go run main.go` 

you have to run it where the enviroment is aware of GOPATH 

now to run go code from anywhere there are few ways. We will go with most used one using GO MOD

#### What Is a Go Module?
A Go module is a collection of Go packages stored in a specific directory with a go.mod file at its root. It is the standard way to manage dependencies and versioning in modern Go development.

After that you will need main func to get an executable  standalone go file to run 

#### Input/Output in Go
Go provides robust I/O operations through its standard packages. The most commonly used package for basic I/O operations is `fmt`.

```
// Reading input from the console
var name string
fmt.Print("Enter your name: ")
fmt.Scanln(&name)

// Reading formatted input
var age int
var height float64
fmt.Scanf("%d %f", &age, &height)

// Basic printing
fmt.Print("Hello without newline")
fmt.Println("Hello with newline")

// Formatted printing
name := "Alice"
age := 25
fmt.Printf("Name: %s, Age: %d\n", name, age)

// String formatting
formatted := fmt.Sprintf("Name: %s, Age: %d", name, age)
```

#### Conditions

one weird thing about go the bracket in condition  can't start from new line it has to be beside the condition weird 

```
age := 18
if age >= 18 {
    fmt.Println("You are an adult")
} else if age >= 13 {
    fmt.Println("You are a teenager")
} else {
    fmt.Println("You are a child")
}

// Go also supports a unique initialization statement in if conditions: 
// suppose calculateScore() return a score 

if score := calculateScore(); score > 100 {
    fmt.Println("High score!")
} else {
    fmt.Println("Keep trying!")
}


day := "Monday"
switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Friday":
    fmt.Println("TGIF!")
default:
    fmt.Println("Regular day")
}

// Switch with no expression
switch {
case time.Hour() < 12:
    fmt.Println("Good morning!")
case time.Hour() < 17:
    fmt.Println("Good afternoon!")
default:
    fmt.Println("Good evening!")
}
```