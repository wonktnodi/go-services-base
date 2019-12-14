package databases

import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/wonktnodi/go-services-base/pkg/config"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "github.com/wonktnodi/go-utils/log"
  "xorm.io/xorm"
)

func InitMysql(setting *config.Database, debug bool) (dbInst *xorm.Engine) {
  logger := log.NewAdaptor(10)
  var err error
  connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&charset=utf8",
    setting.User, setting.Password, setting.Host, setting.Port, setting.Name)
  dbInst, err = xorm.NewEngine("mysql", connectString)
  if err != nil {
    logging.Errorf("failed to connect database: %s", err)
    return
  }
  dbInst.SetMaxOpenConns(20)
  
  dbLogger := xorm.NewSimpleLogger(logger)
  if dbLogger != nil {
    dbInst.SetLogger(dbLogger)
  }
  if debug {
    dbInst.ShowSQL(true)
  }
  return dbInst
}
