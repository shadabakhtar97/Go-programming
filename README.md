# Go-programming
We will cover golang or go programming  from basic to advance

# Error Handling in Go Programming
### What is an Error in go programming ?

In Go programming, an error is a value that represents an abnormal or unexpected condition that has occurred during the execution of a program. Errors are used to signal and handle exceptional situations, such as file not found, network connection issues, division by zero, or any other circumstance that can prevent a program from executing as intended.

In Go, errors are typically represented by values of the built-in `error` interface type. The `error` interface is defined as follows:

```go
type error interface {
    Error() string
}
```

Any type that implements this interface by providing an `Error` method can be used as an error value in Go. The `Error` method should return a string that describes the error.

Here's a simple example of how errors are used in Go:

```go
package main

import (
    "fmt"
    "errors"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }

    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

In this example, the `divide` function returns an error if the second argument is zero, indicating a division by zero error. The `main` function checks for errors returned by `divide` and handles them accordingly.

### -------------------------------------------------------------------------------------------------------------------------------------
#### How to handle errors in Go programming language ?
In Go programming, error handling is a crucial part of writing reliable and robust code. Go encourages explicit error handling, and there are several common patterns and techniques for handling errors effectively:

1. Return Errors: Functions that can encounter errors should return both a result and an error value. This is a common practice in Go.

   ```go
   func divide(a, b int) (int, error) {
       if b == 0 {
           return 0, errors.New("division by zero")
       }
       return a / b, nil
   }
   ```

   When calling this function, you can check the error value to determine if an error occurred:

   ```go
   result, err := divide(10, 0)
   if err != nil {
       // Handle the error
       fmt.Println("Error:", err)
   } else {
       fmt.Println("Result:", result)
   }
   ```

2. Use `errors.New` or `fmt.Errorf`: You can create custom error messages using `errors.New` or `fmt.Errorf` to provide meaningful information about the error.

   ```go
   import "errors"

   func openFile(filename string) (file *File, err error) {
       // ...
       if fileNotFound {
           return nil, errors.New("file not found")
       }
       // ...
   }
   ```

3. Error Wrapping: You can use the `fmt.Errorf` function to wrap errors with additional context, making it easier to trace the source of an error:

   ```go
   import "fmt"

   func processFile(filename string) error {
       // ...
       if err := openFile(filename); err != nil {
           return fmt.Errorf("failed to open file: %w", err)
       }
       // ...
   }
   ```

   This allows you to use `%w` to wrap an error while preserving the original error information.

4. Custom Error Types: You can define your custom error types by creating structs that implement the `error` interface. This can be useful for more structured error handling.

   ```go
   type MyError struct {
       ErrorCode    int
       ErrorMessage string
   }

   func (e *MyError) Error() string {
       return fmt.Sprintf("Error %d: %s", e.ErrorCode, e.ErrorMessage)
   }
   ```

5. `panic` and `recover`: In exceptional cases where an error is catastrophic and should halt the program, you can use `panic` to stop the normal flow of execution. You can then recover from the panic using the `recover` function to handle it gracefully.

   ```go
   func someFunction() {
       defer func() {
           if r := recover(); r != nil {
               fmt.Println("Recovered from panic:", r)
           }
       }()
   
       // ...
       if somethingBadHappened {
           panic("Something bad happened")
       }
       // ...
   }
   ```

6. Error Propagation: In a larger application, you might need to propagate errors up the call stack to handle them at a higher level. This involves returning errors from lower-level functions and checking them in higher-level functions.

Remember that the appropriate error-handling strategy depends on the specific needs of your application and the nature of the errors you expect to encounter. Go's philosophy is to make error handling explicit and clear, so you should aim to provide informative error messages and handle errors in a way that makes your code more reliable and maintainable.



Go encourages explicit error handling to ensure robust and reliable code. Developers are expected to check for errors and decide how to handle them, whether that involves logging the error, retrying an operation, or taking some other appropriate action.

### -----------------------------------------------------------------------------------------------------------------------------
### Blocks used to handle errors in Go programming lanuage ?
In Go programming, error handling is typically done using conditional statements and blocks, such as `if` and `switch`, to handle different error scenarios. Here are some common blocks and constructs used to handle errors in Go:

1. `if` Statements:
   - `if` statements are commonly used to check for and handle errors. You can use them to conditionally execute code based on whether an error is `nil` (indicating no error) or not.

   ```go
   result, err := someFunction()
   if err != nil {
       // Handle the error
   } else {
       // Handle the result
   }
   ```

2. `switch` Statements:
   - You can use `switch` statements to handle multiple error cases with different behaviors. This is especially useful when you want to handle different error types differently.

   ```go
   switch err := someFunction(); err {
   case nil:
       // No error, continue
   case specificErrorType:
       // Handle specific error type
   default:
       // Handle other errors
   }
   ```

3. `defer` Statements:
   - The `defer` statement can be used to ensure that cleanup or error handling code is executed before a function returns, even if an error occurs later in the function.

   ```go
   func openFile(filename string) (file *File, err error) {
       file, err = os.Open(filename)
       if err != nil {
           return nil, err
       }
       defer file.Close() // Ensure the file is closed regardless of error
       // ...
   }
   ```

4. `panic` and `recover`:
   - You can use `panic` and `recover` to handle exceptional, unrecoverable errors in a controlled manner. `panic` is used to trigger a panic, and `recover` is used to catch and handle the panic.

   ```go
   func someFunction() {
       defer func() {
           if r := recover(); r != nil {
               // Handle the panic
           }
       }()
   
       // ...
       if somethingBadHappened {
           panic("Something bad happened")
       }
       // ...
   }
   ```

5. Custom Error Handling Blocks:
   - Depending on the complexity of your code, you may create custom blocks or functions to handle specific types of errors or error scenarios. This can make your error handling code more modular and easier to manage.

   ```go
   func handleDatabaseError(err error) {
       // Custom logic for handling database errors
   }

   // ...

   db, err := connectToDatabase()
   if err != nil {
       handleDatabaseError(err)
   }
   ```

Remember that the specific approach to error handling will vary depending on the requirements of your application and the nature of the errors you expect to encounter. The key is to make error handling explicit, informative, and organized to ensure the reliability and maintainability of your code.

