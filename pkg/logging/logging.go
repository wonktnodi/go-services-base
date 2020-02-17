package logging

const (
  TRACE int = iota
  // DEBUG represents debug log level.
  DEBUG
  // INFO represents info log level.
  INFO
  // WARN represents warn log level.
  WARN
  // ERROR represents error log level.
  ERROR
  // FATAL represents fatal log level.
  FATAL
)

type Logger interface {
  Trace(args ...interface{})
  Tracef(format string, args ...interface{})
  Debug(args ...interface{})
  Debugf(format string, args ...interface{})
  Warn(args ...interface{})
  Warnf(format string, args ...interface{})
  Info(args ...interface{})
  Infof(format string, args ...interface{})
  Error(args ...interface{})
  Errorf(format string, args ...interface{})
  Fatal(args ...interface{})
  Fatalf(format string, args ...interface{})
  
  SetLevel(level int)
}

var logger Logger = nil

func Trace(args ...interface{}) {
  logger.Trace(args...)
}

func Tracef(format string, args ...interface{}) {
  logger.Tracef(format, args...)
}

func Debug(args ...interface{}) {
  logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
  logger.Debugf(format, args...)
}

func Info(args ...interface{}) {
  logger.Debug(args...)
}

func Infof(format string, args ...interface{}) {
  logger.Debugf(format, args...)
}

func Warn(args ...interface{}) {
  logger.Debug(args...)
}

func Warnf(format string, args ...interface{}) {
  logger.Debugf(format, args...)
}

func Error(args ...interface{}) {
  logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
  logger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
  logger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
  logger.Fatalf(format, args...)
}

func SetLevel(level int) {
  logger.SetLevel(level)
}

func GetLogger() Logger {
  return logger
}