package utils

import (
	"net/http" //Http handling
	"log"
	"io/ioutil"
  "regexp"
)

func GetWebpageInfo( domain string ) ( string, string, error ) {
	log.Println( "Website to get info on", domain )
  title := ""
  icon := ""

	request, err := http.Get( "http://" + domain )

	if err != nil {
		return "", "", err
	}

	defer request.Body.Close();

	data, ioerror := ioutil.ReadAll( request.Body )

	if ioerror != nil {
		return "", "", ioerror
	}

  //Get Icon
  iconvalue := string( GetWebPageData( GetWebPageData( data, `<link[^>]*rel="shortcut icon"[^>]*>` ), `href="[^"]*"` ) )
  if iconvalue != "" {
    icon = iconvalue
  }

  //Get Title
  titlevalue := string( GetWebPageData( GetWebPageData( data, `<title>[^<]*</title>` ), `>.*<` ) )
  if titlevalue != "" {
    title = titlevalue
  }

	return title, icon, nil
}

func GetWebPageData( webpage []byte , expression string ) ( []byte ) {
  compiledRegex := regexp.MustCompile( expression )
  return compiledRegex.Find( webpage )
}
