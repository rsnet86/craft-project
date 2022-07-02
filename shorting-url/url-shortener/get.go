package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTinyURl(w http.ResponseWriter, r *http.Request) {

	//longUrl := r.URL.Query().Get("long_url")
	//fmt.Println(longUrl)
	shorturl := r.URL.Query().Get("short_url")
	fmt.Println(shorturl)

	var shortUrl, longurl, emailId, expiry string
	var err error

	// get data from redis
	longurl = GetDataFromRedisByShortUrl(shorturl)
	//if not found in redis then get from db
	if longurl == "" {
		shortUrl, longurl, emailId, expiry, err = GetDataFromDatabase("", shorturl)
	}

	fmt.Println(shortUrl)
	//send response to frontend
	resp := GetResponse{
		EmailID: emailId,
		LongURL: longurl,
		Expiry:  expiry,
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusFound)

	_, err = w.Write(bytes)
	if err != nil {
		return
	}
	return

}
