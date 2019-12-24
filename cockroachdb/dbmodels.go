package cockroachdb

import (
  "time"
)

type DomainDbModel struct {
  id int
  consulted_time time.Time
  host string
}

type InfoServerDbModel struct {
  id int
  servers_changed bool
  ssl_grade string
  previous_ssl_grade string
  logo string
  title string
  is_down bool
  domain_id int
}

type ServerDbModel struct {
  id int
  address string
  ssl_grade string
  country string
  owner string
  infoserver_id string
}
