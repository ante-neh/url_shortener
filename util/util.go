package util

import (
	"os"
	"strings"
)

func RemoveDomainError(url string) bool {
	domain := os.Getenv("DOMAIN") 

	if (url == domain){
		return false 
	} 

	newUrl := strings.Replace(url, "http://", "", 1)
	newUrl = strings.Replace(newUrl, "https://", "", 1)
	newUrl = strings.Replace(newUrl, "www.", "", 1)
	newUrl = strings.Split(newUrl,"/")[0] 

	if (newUrl == domain){
		return false 
	} 

	return true
}

func EnforeHTTP(url string) string {

	if url[:4] != "http" {
		return "http://" + url
	}
	
	return url
}