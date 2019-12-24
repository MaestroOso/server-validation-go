package ssllabs

import (
	"net/http" //Http handling
	"encoding/json" //Json parsing
	"log"
	"io/ioutil"
)

func SslLabs( domain string ) ( SslLabsRequestModel , error ) {
	log.Println( "SslLabs to get info on", domain )
	request, err := http.Get( SslLabsApiUrl + domain )

	if err != nil {
		return SslLabsRequestModel{}, err
	}

	defer request.Body.Close();

	data, ioerror := ioutil.ReadAll( request.Body )

	if ioerror != nil {
		return SslLabsRequestModel{}, ioerror
	}

	model := SslLabsRequestModel{}
	parseerror := json.Unmarshal(data, &model)

	if parseerror != nil {
		return SslLabsRequestModel{}, parseerror
	}

	return model, nil
}
