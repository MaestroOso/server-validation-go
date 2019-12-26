package models

import (
  "time"
)

type DomainDbModel struct {
  Id int
  Host string
  Ssl_grade string
  Is_down bool
}

type DomainConsultDbModel struct {
  Id int
  Id_domain int
  Consult_time time.Time
  Ssl_grade string
  Title string
  Logo string
  Is_down bool
}

type ServerDbModel struct {
  id int
  address string
  ssl_grade string
  country string
  owner string
  infoserver_id string
}
