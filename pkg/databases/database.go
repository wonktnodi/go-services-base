package databases

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/wonktnodi/go-services-base/pkg/config"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "github.com/wonktnodi/go-utils/log"
  "time"
  "xorm.io/xorm"
)

func InitMysql(setting *config.Database, localTime, debug bool) (dbInst *xorm.Engine) {
  logger := log.NewAdaptor(10)
  var err error
  connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&charset=utf8",
    setting.User, setting.Password, setting.Host, setting.Port, setting.Name)
  if localTime {
    //connectString += "&loc=UTC"
    connectString += "&serverTimezone=UTC"
  }
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
  //testTime(connectString)
  return dbInst
}

func testTime(connectString string) {
  db, err := sql.Open("mysql", connectString)
  var myTime time.Time
  rows, err := db.Query("SELECT current_timestamp()")
  fmt.Println(time.Now())
  if rows.Next() {
    if err = rows.Scan(&myTime); err != nil {
      panic(err)
    }
  }
  
  fmt.Println(myTime)
}
