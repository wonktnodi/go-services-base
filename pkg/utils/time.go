package utils

import "time"

func TimeTruncateDay(t time.Time) (ret time.Time) {
  ret = time.Date(t.Year(), t.Month(), t.Day(),
    0, 0, 0, 0, t.Location())
  return
}

func TimeTruncateWeek(t time.Time, firstDay time.Weekday) (ret time.Time) {
  diff := firstDay - t.Weekday()
  ret = TimeTruncateDay(t).AddDate(0, 0, int(diff))
  return
}

func TimeTruncateMonth(t time.Time) (ret time.Time) {
  ret = time.Date(t.Year(), t.Month(), 1,
    0, 0, 0, 0, t.Location())
  return
}
