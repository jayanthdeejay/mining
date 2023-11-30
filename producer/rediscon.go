package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/jayanthdeejay/mining/address"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "172.18.0.2"
	port     = 5432
	user     = "postgres"
	password = "VeryLongPassword"
	dbname   = "nidhi"
)

func init() {
	var err error
	uriString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", uriString)
	if err != nil {
		log.Fatal("Invalid DB config:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("DB unreachable:", err)
	}
}

func main() {
	// Connect to the Redis server
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Read messages from the mining channel
	count := 0
	for {
		res, err := rdb.LPop(ctx, "debruijn").Result()
		if err != nil {
			log.Fatalf("Failed to pop key from list: %v", err)
		}

		//fmt.Println(res)
		ProcessKey(string(res))
		if count == 100000 {
			fmt.Println("Received key:", res)
			count = 0
		}
		count++
	}
}

func ProcessKey(key string) {
	add := address.GetEthAddress(key)
	Checkaddress(key, add)
	pubKeyHash := address.HexToPubKeyHash(key)
	p2pkh_add := address.GetP2PKHAddress(pubKeyHash)
	Checkaddress(key, p2pkh_add)

	bech32_add := address.GetBech32Address(pubKeyHash)
	Checkaddress(key, bech32_add)

	p2sh_add := address.P2shAddress(key)
	Checkaddress(key, p2sh_add)
}

func Checkaddress(key, add string) {
	var exists bool
	row := db.QueryRow(`
        SELECT EXISTS(SELECT 1 FROM ethereum WHERE address = $1)
               OR EXISTS(SELECT 1 FROM bitcoin WHERE address = $1)`, add)
	err := row.Scan(&exists)

	if err != nil {
		log.Fatalf("Failed to check if address exists: %v", err)
	}

	if exists {
		_, err = db.Exec("INSERT INTO found (key, address) VALUES ($1, $2)", key, add)
		if err != nil {
			log.Fatalf("Failed to save key to database: %v", err)
		}
		fmt.Println("Key saved to database")
	}
}
