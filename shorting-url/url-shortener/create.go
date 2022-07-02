package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func CreateTinyURl(w http.ResponseWriter, r *http.Request) {

	var request Request

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = dbPsql.Ping()
	if err != nil {
		return
	}

	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		return
	}
	fmt.Printf("reqBody data in newEvent:%#v\n ", request)

	// validation from redis
	isPresent := ValidateData(request.LongURL)
	if isPresent == 1 {
		shortUrl := GetDataFromRedisByLongUrl(request.LongURL)
		resp := Response{
			EmailID:  request.EmailID,
			ShortURL: shortUrl,
			Expiry:   "",
		}

		bytes, err := json.Marshal(resp)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusCreated)

		_, err = w.Write(bytes)
		if err != nil {
			return
		}
		return
	}

	// logic to create short URl
	shorturl := TinyURLLogic(7)

	// store this request in redis with short url
	StoreDataToRedisByLongUrl(request.LongURL, shorturl)
	StoreDataToRedisByShortUrl(request.LongURL, shorturl)

	// store in db
	expiry := time.Now().AddDate(0, 0, 1)
	StoreDataToDatabase(request.LongURL, shorturl, request.EmailID, expiry.String())

	//send response to frontend
	resp := Response{
		EmailID:  request.EmailID,
		ShortURL: shorturl,
		Expiry:   expiry.String(),
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(bytes)
	if err != nil {
		return
	}
	return
}

func ValidateData(longurl string) int64 {
	isPresent := IsDataExistToRedisByLongUrl(longurl)
	return isPresent
}
