package entities

type ServerInformationModel struct {
  Address string
  SslGrade string
  Country string
  Owner string
}

type DomainModel struct {
  Servers []ServerInformationModel
  ServersChanged bool
  SslGrade string
  PreviousSslGrade string
  Logo string
  Title string
  Is_down bool
}

func MakeServerInformationModel( _address string, _sslgrade string, _country string, _owner string ) ( ServerInformationModel ) {
  return ServerInformationModel{ _address, _sslgrade, _country, _owner}
}
