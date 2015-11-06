## log: a proposal to make the current log package in the standard library more extensible

This is a proposal to make the current Go Logging package more extensible in a fully backawards
compatible way by allowing `logger.Output()` call a custom logging implementation to write logs 
referencing [this golang-nuts discussion](https://groups.google.com/forum/#!topic/golang-nuts/R7ryo7RdBPY).

To do this, we define a function type called `OutputFn` with the same signature as the current
`Logger.Ouptut()` and move its current implementation into a function called `Loggger.DefOutputFn`.
We introduce a mutex protected variable `Logger.outputfn` which by default points to `Logger.DefOutputFn`
but can be swapped for any other implementation by calling Logger.SetOutputFn()

For the standard logger, we introduce `log.SetOutputFn()` and `log.SetDefOutputFn()` to set and
reset custom loggers

The example file customlogger_test.go shows how we can use this to swap a logger that outputs
logs in tab separated variable file format. 
