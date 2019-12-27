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
  Id int
  Id_domainconsult int
  Address string
  Ssl_grade string
  Country string
  Owner string
}
