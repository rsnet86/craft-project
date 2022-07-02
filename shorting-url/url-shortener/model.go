package main

type Request struct {
	LongURL string `json:"long_url" db:"long_url"`
	EmailID string `json:"email_id" db:"email_id"`
}

type Response struct {
	EmailID  string `json:"email_id" db:"email_id"`
	ShortURL string `json:"short_url" db:"short_url"`
	Expiry   string `json:"expiry" db:"expiry"`
}

type GetResponse struct {
	EmailID string `json:"email_id"`
	LongURL string `json:"long_url" `
	Expiry  string `json:"expiry"`
}

type DeleteResponse struct {
	Message string `json:"message"`
	LongUrl string `json:"long_url"`
}

// CREATE TABLE IF NOT EXISTS tiny_url
// (
//    long_url        TEXT PRIMARY KEY NOT NULL,
// 	short_url       TEXT DEFAULT '',
// 	email_id        TEXT  DEFAULT '',
// 	expiry          TEXT,
//    created_at 		TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
//    updated_at 		TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
// );
