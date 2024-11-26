# Go Conventions

1. In go all files belonging to a single package should be present in the same directory.

2. You don't need to use import statements to access functions that are in different files but in the same package.

3. In go to make a function private (so that other packages cannot access the package) you have to start the function name in small letter

1. When you are running functions that are part of a different file, in NodeJS you need to import it and use it. When running the main/root file, the complier is intelligent enough to know that one more file is involved and it will compile that as well. In go this is not straight forward. To explain how imports work in GO you need to understand what packages are
