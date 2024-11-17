# GO packages

1. In go, each file belongs to a package. The executable package is the main package.
2. You can have multiple files that belong to the main package. But the convention in go is to keep all the files belong to a package in the same directory. Meaning it is not good practice to keep a file belonging main package to a nested directory when you have main package already declared in the current directory.



## From Chatgpt:

### Question 1:

In my current directory I have the main.go file belonging to the main package. I have a directory called helper in my current directory. There is  file called helper.go which belongs to the main package. The helper.go has a few functions. How do use the functions that are present in helper.go in the main.go file?

### Answer:

If `helper.go` belongs to the `main` package but is located in a subdirectory (e.g., `helper`), Go's compiler won't automatically recognize it as part of the `main` package. Unlike files in the same directory, files in subdirectories are treated as separate packages.

To use functions from `helper.go` in `main.go` while keeping `helper.go` in the `helper` subdirectory, you have two options:

---

### **Option 1: Move `helper.go` to the Root Directory**
If both files belong to the `main` package, you can move `helper.go` to the same directory as `main.go`. Then, you can directly use the functions in `helper.go` without any import.

#### Example:
- Directory structure:
  ```
  myapp/
  ├── main.go
  ├── helper.go
  ```

- `main.go`:
  ```go
  package main

  import "fmt"

  func main() {
      fmt.Println(HelperFunction())
  }
  ```

- `helper.go`:
  ```go
  package main

  func HelperFunction() string {
      return "Hello from HelperFunction"
  }
  ```

Run the program with:
```bash
go run .
```

---

### **Option 2: Treat `helper` as a Separate Package**
If you want to keep `helper.go` in the `helper` subdirectory, you need to change its package to `helper` and import it into `main.go`.

#### Steps:
1. **Set up directory structure**:
   ```
   myapp/
   ├── main.go
   ├── helper/
   │   └── helper.go
   ```

2. **Modify the code**:
   - `main.go`:
     ```go
     package main

     import (
         "fmt"
         "myapp/helper"
     )

     func main() {
         fmt.Println(helper.HelperFunction())
     }
     ```

   - `helper/helper.go`:
     ```go
     package helper

     func HelperFunction() string {
         return "Hello from HelperFunction"
     }
     ```

3. **Run the program**:
   ```bash
   go mod init myapp
   go run .
   ```

---

### Why Keep `helper.go` in `main` Package?
If you’re intentionally keeping `helper.go` in the `main` package, move it to the root directory. Subdirectories in Go are generally used to define distinct packages, so putting `helper.go` in `helper/` implies it’s a separate package.

