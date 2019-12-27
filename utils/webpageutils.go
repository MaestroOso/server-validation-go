package utils

import (
	"net/http" //Http handling
	"log"
	"io/ioutil"
  "regexp"
	"strings"
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
	iconvalue = RemoveUnwantedString ( RemoveUnwantedString( iconvalue, `href="` ), `"` );
  if iconvalue != "" {
    icon = iconvalue
  }

  //Get Title
	titlevalue := string( GetWebPageData( GetWebPageData( data, `<title>[^<]*</title>` ), `>.*<` ) )
  titlevalue = RemoveUnwantedString( RemoveUnwantedString( titlevalue, `<` ) , `>` )
  if titlevalue != "" {
    title = titlevalue
  }

	return title, icon, nil
}

func GetWebPageData( webpage []byte , expression string ) ( []byte ) {
	if len( webpage ) == 0{
			return []byte{}
	}
  compiledRegex := regexp.MustCompile( expression )
  return compiledRegex.Find( webpage )
}

func RemoveUnwantedString( webpage string, substring string  ) ( string ){
	return strings.Replace( webpage, substring, "", 1)
}
