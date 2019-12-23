package ssllabs

import (

)

type Endpoint struct {
  IpAddress string `json:"ipAddress"`
  Grade string `json:"grade"`
}

type SslLabsRequestModel struct {
  Endpoints []Endpoint `json: endpoints`
}
