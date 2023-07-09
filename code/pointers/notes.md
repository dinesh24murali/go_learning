# Pointers

GO is a pass by value language

```
jimPointer := &jim
```
- here we are creating a pointer to the jim variable by getting the address
- The works different from pointers in C or C++

```
(*pointerToPerson).firstName
```
- Putting asterisk (*) before a pointer will turn it into a variable. Here we are converting the pointer to a variable and accessing its value

## Note:

- for some reason when I printed the pointer 
```
fmt.Println("Pointer", pointerToPerson)
```
it did not print out the memory address number. It printed this
```
Pointer &{jimmy brown {jim@gmail.com 94000}}
```
- Here pointerToPerson is of the type pointer to the `person` struct
```
pointerToPerson *person
```

# Slices

- When you are modfying a slice from another function shown in `gotchas.go` it will work even though we are not passing a reference
- This is because Slice in go is like a struct on top of Array, so when we pass a slice to a function we are passing the copy of the memory addresses to an array