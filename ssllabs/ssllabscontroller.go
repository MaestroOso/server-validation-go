package ssllabs

import (
	"net/http" //Http handling
	"encoding/json" //Json parsing
	"fmt"
	"io/ioutil"
)

func SslLabs( domain string ) ( SslLabsRequestModel , error ) {

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

	fmt.Println(model)
	return model, nil
}
