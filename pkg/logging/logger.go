package logging

import (
  "github.com/wonktnodi/go-utils/log"
  "github.com/wonktnodi/go-services-base/pkg/debug"
)

type Log struct {
  instance *log.LogAdaptor
}

func NewLogger(runMode int) {
  var inst *log.LogAdaptor
  if runMode == debug.DEBUG_MODE {
    inst = log.Start(log.LogFilePath("./log", ""), log.EveryHour,
      log.AlsoStdout, log.LogFlags(log.Lfunc|log.Lfile|log.Lline))
  } else {
    inst = log.Start(log.LogFilePath("./log", ""), log.EveryHour,
      log.LogFlags(log.Lfunc|log.Lfile|log.Lline))
  }
  
  inst.SetCallDepth(5)
  logger = &Log{
    instance: inst,
  }
}

func NewAdaptor(path, name string, runMode int) *Log {
  var logger log.Logger
  if runMode == debug.DEBUG_MODE {
    logger = log.NewLogInstance(log.LogFilePath(path, name), log.EveryHour, log.AlsoStdout)
  } else {
    logger = log.NewLogInstance(log.LogFilePath(path, name), log.EveryHour)
  }
  
  adaptor := log.NewAdaptorFromInstance(&logger, 3)
  return &Log{
    instance: adaptor,
  }
}

func (l *Log) Trace(args ...interface{}) {
  l.instance.Traceln(args...)
}

func (l *Log) Tracef(format string, args ...interface{}) {
  l.instance.Tracef(format, args...)
}

func (l *Log) Debug(args ...interface{}) {
  l.instance.Debugln(args...)
}

func (l *Log) Debugf(format string, args ...interface{}) {
  l.instance.Debugf(format, args...)
}

func (l *Log) Info(args ...interface{}) {
  l.instance.Infoln(args...)
}

func (l *Log) Infof(format string, args ...interface{}) {
  l.instance.Infof(format, args...)
}

func (l *Log) Warn(args ...interface{}) {
  l.instance.Warnln(args...)
}

func (l *Log) Warnf(format string, args ...interface{}) {
  l.instance.Warnf(format, args...)
}

func (l *Log) Error(args ...interface{}) {
  l.instance.Errorln(args...)
}

func (l *Log) Errorf(format string, args ...interface{}) {
  l.instance.Errorf(format, args...)
}

func (l *Log) Fatal(args ...interface{}) {
  l.instance.Fatalln(args...)
}

func (l *Log) Fatalf(format string, args ...interface{}) {
  l.instance.Fatalf(format, args...)
}

func (l *Log) SetLevel(level int) {
  l.instance.SetLevel(log.LogLevel(level))
}
