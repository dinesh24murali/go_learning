# Read interface

- All http types like: file, image, string, web socket will use this interface to convert data to byte type

# Writer interface

- This is the opposite of `Read` interface. It will take a `[]byte` type and convert it to the desired output: Outgoing http request, text to hard drive, image file, Terminal