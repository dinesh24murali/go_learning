# Setup a new project

Like how we have `package.json` in nodejs, there is a file called `go.mod`.

**Step 1:**

In the terminal go to the folder you want to setup the go project and run the following code:
```bash
go mod init <package name>
```
For this example let's call the package name `learn`
```bash
go mod init learn
```

This will create the `go.mod` file with the following content:

```mod
module learn

go 1.23
```

**Step 2:**

Create a file called `main.go`. Paste the following content into it:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
```

**Step 3:**

To create a new `package`, create a new folder inside the current directory. The folder name should be same as package name. Let's call this package `greeting`

```bash
mkdir greeting
cd greeting
```
Create a new file inside the folder. The file name can be any thing. Let's call it `greet.go`. Paste the following content into it:

```go
package greeting

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```
Now you have created a package.

**Step 4:**

To import files from the package into `main.go` file, make the following modification:

```go
package main

import "fmt"
import (
	"learn/greeting" // We need to import the package
)

func main() {
	text := greeting.Hello("Jim") // we can use it this way
	fmt.Println(text)
}
```

This is a basic setup for a simple go project.

