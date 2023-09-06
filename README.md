# 🌲 Picolog 🌲

Picolog is my personal minimalistic logging library written in Go. Designed for simplicity and ease of use, it's perfect for very small projects where you need a quick and easy logging format without the hassle. Think of it as a tiny reliable friend for minimal debug and information tracking! 🐞

## 🌟 Features

> Originally not built with concurrency in mind. Found out the hard way this is needed when you try to log from multiple goroutines. Enjoy.

- **Simple:** Easy to setup and start using. No bloat, no unnecessary features.
- **Lightweight:** A very (no, really) very small codebase that won't bloat your project.
- **Thread-Safe:** ~~Now~~ you can feel free to log from multiple goroutines.
- **Customizable Log Levels**: Comes with several log levels, enough for non-complex applications (Info, Debug, Warning, Error, Ok) for better message categorization.
- **Optional colored output**: Log messages can come in different colors based on their level for better visibility.

> **Note:** The colored output uses ANSI escape codes, which are not supported by all terminals. If you're using a terminal that doesn't support ANSI escape codes the output might mess up your terminal output.

## 📦 Installation

This package is not published on the Go package registry, so you'll have to install it manually.

```bash
go get -u github.com/VentGrey/picolog
```

## 🧪 Running tests

To run this project's tests just run:

```bash
go test
```

If you wish to read a richer output:

```bash
go test -v
```

## 📚 Usage

Here's how you can start using Picolog in your project.

```go
package main

import "github.com/VentGrey/picolog"

func main() {
    // picolog has colored output disabled by default.
    logger := picolog.NewLogger("main", picolog.Info, false)
    
    logger.Info("This is an info message")
    logger.Debug("This is a debug message")
    logger.Warning("This is a warning message")
    logger.Error("This is an error message")
    logger.Ok("This is an ok message")
}
```

## 🚀 Advanced Usage
### 📝 Configuring Log Level and Display Name

You can set a minimum log level and control whether to display the logger name:

```go
logger := picolog.NewLogger("main", picolog.Warning, false)
```

With this configuration, only `Warning`, `Error` and `Ok` messages will be displayed.

#### Format Strings are unsupported, why?

Picolog does not natively support format strings like `fmt.Printf` in individual log functions. However, you can use `fmt.Sprintf` to prepare formatted strings before logging:

```go
formattedMessage := fmt.Sprintf("Hello, %s!", "world")
logger.Info(formattedMessage)
```

This feature is not supported by default because it would require the logger to parse the format string and the arguments, which would add unnecessary complexity to the codebase.

### 🎨 Colored Output and Terminal Support

This feature is set to `false` by default.

Picolog uses ANSI color codes to differentiate log levels with colours. While this works well on most modern terminals, some environments might not support ANSI color codes.

To enable colored output, set the `colored` parameter to `true` when creating a new logger.

```go
logger := picolog.NewLogger("main", picolog.Info, true)
```

This will disable the ANSI color codes, providing a plain text output suitable for log files or unsupported terminals.

### 🧵 Thread-Safety

Picolog is designed to be thread-safe. You can safely use the same logger instance across multiple goroutines:

``` go
go logger.Info("Hello from goroutine 1")
go logger.Info("Hello from goroutine 2")
```

Or a slightly more complex example using `sync.WaitGroup` and anonymous functions to simulate multiple goroutines logging at the same time:

```go
package main

import (
    "github.com/VentGrey/picolog"
    "sync"
)

func main() {
    logger := picolog.NewLogger("main", picolog.Info, false)
    
    var wg sync.WaitGroup
    wg.Add(2)
    
    go func() {
        defer wg.Done()
        logger.Info("Hello from goroutine 1")
    }()
    
    go func() {
        defer wg.Done()
        logger.Info("Hello from goroutine 2")
    }()
    
    wg.Wait()
}
```

## 🤝 Contributing

Picolog is considered feature-complete won't be actively developed. However, if you find a bug or think of a useful feature, feel free to create an issue or submit a pull request.

## 📝 License

This project is licensed under the GPL-3+ License. See the [LICENSE](LICENSE) file for details.

---

Made with ❤️ and Go.
