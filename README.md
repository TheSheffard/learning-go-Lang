# Learn Go Lang

A beginner-friendly collection of Go examples covering the core language concepts used to build small command-line programs.

## About

This project was created as a practical Go learning exercise. The examples are organized by topic and demonstrate how Go values, functions, collections, and control flow work through delivery, rider, player, and product scenarios.

The program prints examples to the terminal so that each concept can be observed immediately.

## Topics Covered

- Declaring and assigning variables
- Go's short variable declaration syntax
- Constants and zero values
- Strings, integers, floats, and booleans
- Type conversion
- Arithmetic and comparison operators
- Conditional statements
- Loops, `break`, and `continue`
- Functions with parameters and return values
- Multiple return values
- Arrays and slices
- Appending to slices and iterating with `range`
- Maps, updating entries, deleting entries, and checking keys
- Structs and struct fields
- Blank identifiers (`_`)

## Project Structure

```text
LearnGoLang/
├── main.go
└── README.md
```

## Prerequisites

Install the Go programming language:

- [Download Go](https://go.dev/dl/)
- Verify the installation:

```powershell
go version
```

Go 1.20 or newer is recommended.

## Running the Program

From a PowerShell terminal in the project directory, run:

```powershell
go run main.go
```

To compile the program into an executable, run:

```powershell
go build main.go
```

Then execute the generated Windows executable:

```powershell
.\main.exe
```

## Example

Some of the output from the program includes:

```text
Emmanuel 20
Joker 4
7

...

14
6
Saka 12
```

The exact output depends on the examples that are enabled in [main.go](main.go).

## How the Examples Are Organized

The `main` function calls different learning examples directly. Several of the learning functions print their results with `fmt.Println`.

Main examples include:

| Function | Purpose |
| --- | --- |
| `conditionalStatement` | Demonstrates `if`, `else if`, and `else` |
| `Loop` | Demonstrates counting loops, countdowns, `break`, and `continue` |
| `greetRider` | Demonstrates function parameters |
| `addGoals` | Demonstrates a function return value |
| `calculateScores` | Demonstrates multiple return values |
| `LearningArray` | Demonstrates fixed-length arrays |
| `LearingSlices` | Demonstrates slices and the `append` function |
| `learningMap` | Demonstrates map operations |
| `learningStruct` | Demonstrates custom struct types |

> Note: `LearingSlices` uses the current spelling from the source code. A future code cleanup could rename it to `learningSlices`.

## Learning Suggestions

Try making small changes to the examples:

1. Change a rider's age or delivery count.
2. Add another club to the array examples.
3. Append a new player or delivery earning to a slice.
4. Add a rider to the deliveries map.
5. Create another struct value.
6. Change a condition and observe how the output changes.
7. Comment or enable learning functions in `main`.

## Code Notes

- Go declarations can use an explicit type, such as `var username string`, or infer a type with `:=`.
- Numeric types must normally be converted explicitly. For example, an `int` is converted with `float64(distance)` before adding it to a `float64`.
- A slice has a flexible length, while an array has a fixed length declared in brackets.
- A map stores values by key.
- Structs group related values into one named type.
- Functions that return multiple values can discard one value with `_`.

## Recommended Go Resources

- [Go Documentation](https://go.dev/doc/)
- [A Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)

## License

This project is intended for personal learning and educational use.
