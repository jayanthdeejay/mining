package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jayanthdeejay/mining/producer/binhex"
	"github.com/jayanthdeejay/mining/producer/necklace"
	"github.com/jayanthdeejay/mining/store"
)

// For string of length n, you need MAX = n + 2
// For 256 bit strings, set MAX = 258
const MAX = 258
const KEYLEN uint16 = 256

var (
	rdb     *redis.Client
	ctx     context.Context
	db      *sql.DB
	initial string
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx = context.Background()
	db, initial = store.DbOpen()
	fmt.Println(initial)
}

// =====================================================================
// Generate de Bruijn sequences by iteratively applying a successor rule
// =====================================================================
func Produce() {
	var a [MAX]uint8
	var n = KEYLEN
	var count = 0
	var i uint16 = 1
	a = binhex.HexToBinary(initial)
	// fmt.Println(initial, a)
	defer rdb.Close()
	defer db.Close()
	for {
		pk := binhex.BinaryToHex(a[1 : MAX-1])
		err := rdb.RPush(ctx, "debruijn", pk).Err()
		if err != nil {
			panic(err)
		}
		
		// Loop until the length of the list goes below 50000
		for {
			length, err := rdb.LLen(context.Background(), "debruijn").Result()
			if err != nil {
				panic(err)
			}
			if length > 500000 {
				fmt.Println("Length of debruijn list:", length)
			}
			if length < 500000 {
				break
			}

			// Sleep for 1 second before checking again
			time.Sleep(5 * time.Second)
		}

		
		// Save key to db once every million iterations
		count++
		if count == 1000000 {
			fmt.Println(pk)
			err := store.Save(db, pk)
			if err != nil {
				panic(err)
			}
			count = 0
		}

		new_bit := necklace.Granddaddy(a, n)
		i = 1
		for i <= n {
			a[i] = a[i+1]
			i++
		}
		a[n] = new_bit
		if necklace.Zeros(a, n) {
			break
		}
	}
}

// ===========================================
func main() {
	Produce()
}

