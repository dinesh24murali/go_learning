## How to use a function declared in a different file present in the same package

### Way 1:

```bash
go run main.go state.go
```

### Way 2:

Convert the current directory into a go module

```
go mod init <module name>
```
Then you can run all the files
```bash
go run .
```

## Running functions from a different module:

In this folder there is a directory called `hello`. This is a separate module. The functions in the file `hello/main.go` belong to `hello` package. To access the functions present in this package.

**Step 1:**
Run the following from the root of the application where `main` package is initialized
```
go mod init <module name>
```
**Step 2:**
Now you can import the `hello` package using:
```go
import "13_module_example/hello"
```
Then you can use the functions inside the `hello` package:
```go
hello.PrintState()
```
