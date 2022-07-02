package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DeleteTinyURl(w http.ResponseWriter, r *http.Request) {

	//longUrl := r.URL.Query().Get("long_url")
	//fmt.Println(longUrl)
	shortUrl := r.URL.Query().Get("short_url")
	fmt.Println(shortUrl)

	// delete data from redis
	isDeleteRedis := DeleteDataFromRedisByShortUrl(shortUrl)
	//delete from db
	isDeleteDB := DeleteDataFromDatabase(shortUrl)

	fmt.Println(isDeleteRedis, isDeleteDB)
	//send response to frontend
	resp := DeleteResponse{
		Message: "",
		LongUrl: shortUrl,
	}

	if isDeleteRedis == 1 && isDeleteDB {
		resp.Message = "data deleted successfully"
	} else {
		resp.Message = "data not found"
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(bytes)
	if err != nil {
		return
	}
	return
}
