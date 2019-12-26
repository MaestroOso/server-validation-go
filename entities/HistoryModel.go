package entities

import (
  "time"
)

type HistoryServerModel struct {
  Address string
  Ssl_grade string
  Country string
  Owner string
}

type HistoryDomainConsultModel struct {
  Consult_time time.Time
  Ssl_grade string
  Title string
  Logo string
  Is_down bool
  Server []HistoryServerModel
}

type HistoryDomainModel struct {
  Host string
  Servers []HistoryDomainConsultModel
}

type HistoryModel struct {
  Items []HistoryDomainModel
}
