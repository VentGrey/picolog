/*
Picolog is a very tiny logger made for my own use. It's designed to be
as minimal as I needed it to be, and it's not meant to replace any
other logger. It's just a simple tool that I use to print debugging
information for my own projects.

As such, I don't expect anyone to use it, I uploaded it to github
just in case someone finds it useful and for me to be able to import
it as a go module in my other projects.

Picolog is considered finished, and I won't be adding any new
features to it. Only bug fixes or improvements with a very good
reasoning behind them. I don't want to bloat it.


Getting started is as easy as creating a new Logger, providing a package name,
minimum log level and a flag if you want your logs to be colored or not:

   // Creating a new logger without color support
   NewLogger("example/main", Info, false)

   // Customizing the log level and disabling colored output
   logger := picolog.NewLogger("my_app", picolog.Warning, true)

   // Picolog is Thread Safe
   go logger.Info("Logging from goroutine 1")
   go logger.Info("Logging from goroutine 2")

   By default, Picolog uses ANSI color codes for log levels.
   This can be enabled by setting the third parameter of NewLogger to true.
   The default value for this feature is: false


VentGrey - 2023 (omar@laesquinagris.com)
*/
package picolog
