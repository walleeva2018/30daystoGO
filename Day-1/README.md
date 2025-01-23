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

### Loops 

```
// Traditional C-style for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Go's version of while loop
count := 0
for count < 5 {
    fmt.Println(count)
    count++
}

// Infinite loop (with break)
for {
    fmt.Println("Infinite loop")
    // Use break to exit
    break
}


// Iterating over slices
fruits := []string{"apple", "banana", "orange"}
for index, value := range fruits {
    fmt.Printf("Index: %d, Value: %s\n", index, value)
}

// Iterating over maps
scores := map[string]int{"Alice": 95, "Bob": 89}
for key, value := range scores {
    fmt.Printf("%s scored %d\n", key, value)
}


// Using continue to skip iterations
for i := 0; i < 5; i++ {
    if i == 2 {
        continue // Skip iteration when i is 2
    }
    fmt.Println(i)
}

// Using break to exit loop
for i := 0; i < 5; i++ {
    if i == 3 {
        break // Exit loop when i is 3
    }
    fmt.Println(i)
}
```

### Best Practices

Always handle I/O errors appropriately.
Use the initialization statement in if conditions when the variable is only needed in that scope.
Prefer switch statements over long if-else chains for better readability.
Use range when iterating over collections for cleaner code.
Consider using labeled break and continue statements in nested loops when necessary.
