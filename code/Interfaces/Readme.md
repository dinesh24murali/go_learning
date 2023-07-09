# Interfaces

- In this example any `type` in this program that has a the `getGreeting()` method associated to it as a receiver it will automatically be part of the interface `bot`
- In this case both `englishBot` and `spanishBot` both has individual `getGreeting()` methods associated to it in the same program. So in the `printGreeting()` function we can use `(b bot)` as a parameter which will call the respective `getGreeting()` function
- Interfaces are not concrete types, meaning we cannot create an value with it

## Rules

1. Interfaces are not generic types
2. Interface are 'implicit', we don't need to explicit link the types
3. Intefaces are a contract to help us manage a type
4. Interfaces are tough. Step #1 is understanding how to read them. Interfaces are not mandatory