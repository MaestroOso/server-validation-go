package whois

import (
	"net/http" //Http handling
	"encoding/json" //Json parsing
	"io/ioutil"
	"log"
)

func WhoIs( domain string ) ( WhoIsRequestModel , error ) {
	log.Println("Whois to get info on", domain)
	request, err := http.Get( WhoIsApiUrl + domain )

	if err != nil {
		return WhoIsRequestModel{}, err
	}

	defer request.Body.Close();

	data, ioerror := ioutil.ReadAll( request.Body )

	if ioerror != nil {
		return WhoIsRequestModel{}, ioerror
	}

	model := WhoIsRequestModel{}
	parseerror := json.Unmarshal(data, &model)

	if parseerror != nil {
		return WhoIsRequestModel{}, parseerror
	}

	return model, nil
}
