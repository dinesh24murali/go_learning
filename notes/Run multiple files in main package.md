## From Chatgpt

### Question 1:
I am learning Golang. I already know NodeJS. I have doubts regarding how to run a Go application. In NodeJS to use a function that is declared in a different file, we use import and export keywords. In Go we don't need import and export keywords. Let's take the same scenario in Go. Where I want to use a function declared in a different file, let's call these files "main.go" and "helper.go". When I run it I need to pass both the file names in the terminal "go run main.go helper.go" which is different from NodeJS where we just run the root file. Am I doing something wrong here? or is this the only way to run multiple files in GoLang?

### Answer:

In Go, when working with multiple files in the same package, you don't need to specify all file names while running the application. Here's the correct way to handle this:

1. **Same Package Approach**: Ensure all files (`main.go` and `helper.go`) are part of the same package, typically `package main` for executables. You can then simply run:

   ```bash
   go run .
   ```

   This tells Go to run all `.go` files in the current directory that belong to the `main` package.

2. **Example Structure**:
   - `main.go`:
     ```go
     package main

     import "fmt"

     func main() {
         result := HelperFunction()
         fmt.Println(result)
     }
     ```

   - `helper.go`:
     ```go
     package main

     func HelperFunction() string {
         return "Hello from HelperFunction"
     }
     ```

   Running `go run .` will compile and execute the program, considering all files in the directory as part of the `main` package.

3. **Important Notes**:
   - All files in the same package must share the same `package` declaration (e.g., `package main` for executables).
   - If files are in different packages, you'll need to use `import` to bring the package into the main file, and you should run `go run main.go` after organizing your modules properly (for example, in a `helper` package).

Let me know if you'd like an example using different packages!