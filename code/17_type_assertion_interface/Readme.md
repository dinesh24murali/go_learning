Here’s an example of how to perform a **type assertion** for a custom interface in Go:

### Example: Type Assertion for a Custom Interface

```go
package main

import "fmt"

// Define a custom interface
type Animal interface {
    Speak() string
}

// Define types that implement the Animal interface
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    // Create an Animal variable
    var a Animal = Dog{} // a contains a value of type Dog

    // Perform a type assertion to check if it's a Dog
    if dog, ok := a.(Dog); ok {
        fmt.Println("It's a Dog!")
        fmt.Println("Dog says:", dog.Speak())
    } else {
        fmt.Println("It's not a Dog.")
    }

    // Perform a type assertion to check if it's a Cat
    if cat, ok := a.(Cat); ok {
        fmt.Println("It's a Cat!")
        fmt.Println("Cat says:", cat.Speak())
    } else {
        fmt.Println("It's not a Cat.")
    }

    // Direct type assertion (no safety check)
    // This will panic if `a` is not a Cat
    // Uncomment the line below to see the panic
    // cat := a.(Cat)
    // fmt.Println(cat.Speak())
}
```

### Output:
```
It's a Dog!
Dog says: Woof!
It's not a Cat.
```

---

### Explanation of Key Points:

1. **Custom Interface**: 
   - `Animal` is a custom interface with one method, `Speak()`.

2. **Implementations of the Interface**: 
   - Both `Dog` and `Cat` types implement the `Animal` interface by defining the `Speak()` method.

3. **Type Assertion**:
   - The line `dog, ok := a.(Dog)` checks if the underlying value of `a` is of type `Dog`.
   - If true, `dog` contains the value of type `Dog`, and `ok` is `true`.
   - If false, `ok` is `false`, and `dog` is the zero value for `Dog`.

4. **Without Safety Check**:
   - The direct type assertion `cat := a.(Cat)` will cause a runtime panic if the type of `a` is not `Cat`.

---

### Practical Use Case:
Type assertions are commonly used when you have an interface that can hold values of different concrete types, and you need to perform specific operations based on the actual type. Using the **comma-ok idiom** ensures safety and prevents runtime panics.