# 🌲 Picolog 🌲

[![Go CI](https://github.com/VentGrey/picolog/actions/workflows/go.yml/badge.svg)](https://github.com/VentGrey/picolog/actions/workflows/go.yml)

[![Go Reference](https://pkg.go.dev/badge/github.com/VentGrey/picolog.svg)](https://pkg.go.dev/github.com/VentGrey/picolog)

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

To install picolog, simply run this command in your Go project's route:

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

### ⚙ Examples

You can find some examples in the [examples](examples) directory.

These examples include:

#### Log Faker using picolog

`fake_log_generator.go` is a simple log generator that uses picolog to generate fake logs. It's a good example of how you can use picolog in your project.

#### Python script to plot error messages

A simple python script that reads a generated logfile and plots the number of error ocurrences. It's a good example of how you can use picolog to generate logs and then use them in other applications.

#### Python script to plot log level ocurrences

Another simple python script that reads a generated logfile and plots the number of ocurrences of each log level. It's a good example of how you can use picolog to generate logs and then use them in other applications.

Try extending those scripts o make a Go parser that sends that data to a database, your preffered metrics tool or observability platform.

#### A pre-generated fake log file

`fake_log.txt` is a pre-generated fake log file that you can use to test the python scripts.
You can generate your own log file by running `fake_log_generator.go` or using Picolog in your own project.

## 📝 Documentation

To view the documentation, run:

```bash
godoc -http=:6060
```

Then open your browser and go to `http://localhost:6060/pkg/github.com/VentGrey/picolog/`.

Alternatively you can view the documentation in your terminal by running:

```bash
go doc github.com/VentGrey/picolog
```

## 📚 Usage

Here's how you can start using Picolog in your project.

```go
package main

import "github.com/VentGrey/picolog"

func main() {
    // picolog has colored output disabled by default.
    logger := picolog.NewLogger("main", picolog.Info, false)
    
    logger.Log(picolog.Info, "This is an info message")
    logger.Log(picolog.Debug, "This is a debug message")
    logger.Log(picolog.Warning, "This is a warning message")
    logger.Log(picolog.Error, "This is an error message")
    logger.Log(picolog.Ok, "This is an ok message")
}
```

The output of this program will be:

```
[INFO] - 2023-09-06 20:28:59 : This is an info message - At package: main
[DEBUG] - 2023-09-06 20:28:59 : This is a debug message - At package: main
[WARNING] - 2023-09-06 20:28:59 : This is a warning message - At package: main
[ERROR] - 2023-09-06 20:28:59 : This is an error message - At package: main
[OK] - 2023-09-06 20:28:59 : This is an ok message - At package: main
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
