package models

type DomainDbModel struct {
  Id int
  Host string
  Ssl_grade string
  Is_down bool
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
