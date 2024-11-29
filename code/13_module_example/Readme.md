## How to use a function declared in a different file present in the same package

### Way 1:

```bash
go run main.go state.go
```

### Way 2:

Convert the current directory into a go module

```
go mod <module name>
```
Then you can run all the files
```bash
go run .
```