package main

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

var dbPsql *sql.DB
var redisClient *redis.Client

const (
	host     = "tinyurl.ctlsr2ublpsl.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "Rihansh#12"
	dbname   = "tinyurl"
)

func CreatePostgres() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("db connected successfully!")
	return db
}

func CreateRedis() *redis.Client {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "10.0.102.161:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong, err)
	fmt.Println("redis connected successfully!")

	return redisClient
}

func StoreDataToDatabase(longurl, shorturl, emailid, expiry string) {

	sqlStatement := `INSERT INTO tiny_url (long_url, email_id, short_url, expiry) VALUES ($1, $2, $3, $4)`

	row, err := dbPsql.Exec(sqlStatement, longurl, emailid, shorturl, expiry)
	if err != nil {
		panic(err)
	}
	n, err := row.RowsAffected()
	if err != nil {
		panic(err)
	}
	println("n", n)
}

func GetDataFromDatabase(longurl, shorturl string) (string, string, string, string, error) {

	var shortUrl, longUrl, email, expiry string

	var err error
	if shortUrl == "" {
		sqlStatement := `SELECT short_url, long_url, email_id, expiry FROM tiny_url WHERE long_url =$1;`
		err = dbPsql.QueryRow(sqlStatement, longurl).Scan(&shorturl, &longUrl, &email, &expiry)
	} else {
		sqlStatement := `SELECT short_url, long_url, email_id, expiry FROM tiny_url WHERE short_url =$1;`
		err = dbPsql.QueryRow(sqlStatement, shorturl).Scan(&shorturl, &longUrl, &email, &expiry)
	}

	if err != nil {
		return shorturl, longurl, expiry, email, err
	}
	return shorturl, longurl, expiry, email, nil
}

func DeleteDataFromDatabase(shorturl string) bool {

	sqlStatement := `DELETE FROM tiny_url WHERE short_url =$1;`

	row, err := dbPsql.Exec(sqlStatement, shorturl)
	if err != nil {
		panic(err)
	}
	n, err := row.RowsAffected()
	if err != nil {
		panic(err)
	}
	println("n", n)
	if n > 0 {
		return true
	}
	return false
}

func StoreDataToRedisByLongUrl(longurl, shorturl string) {
	redisClient.Set(longurl, shorturl, 0)
}

func StoreDataToRedisByShortUrl(longurl, shorturl string) {
	redisClient.Set(shorturl, longurl, 0)
}

func IsDataExistToRedisByLongUrl(longurl string) int64 {
	isPresent := redisClient.Exists(longurl).Val()
	return isPresent
}

func IsDataExistToRedisByShortUrl(shorturl string) int64 {
	isPresent := redisClient.Exists(shorturl).Val()
	return isPresent
}

func GetDataFromRedisByLongUrl(longurl string) string {
	shorturl := redisClient.Get(longurl).Val()
	return shorturl
}

func GetDataFromRedisByShortUrl(shortUrl string) string {
	longurl := redisClient.Get(shortUrl).Val()
	return longurl
}

func DeleteDataFromRedisByLongUrl(longurl string) int64 {
	isDeleted := redisClient.Del(longurl).Val()
	return isDeleted
}

func DeleteDataFromRedisByShortUrl(shorturl string) int64 {
	isDeleted := redisClient.Del(shorturl).Val()
	return isDeleted
}
