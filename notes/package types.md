In Go we can have go code in files that can be executes/run and not executable, they can be shared

## Executable package:

They should start with

```
package main
```
- Only these will be build with `go build` command
- Other packages names (Eg: `package testing`) will not generate a build file
- Needs to have a fuv called `main`

## Reusable package:

- `package <package name>`: Defines a package that can be used as a dependency (helper code)
