package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jayanthdeejay/mining/producer/binhex"
	"github.com/jayanthdeejay/mining/producer/necklace"
	"github.com/jayanthdeejay/mining/rabbitmq"
	"github.com/jayanthdeejay/mining/store"
	amqp "github.com/rabbitmq/amqp091-go"
)

// For string of length n, you need MAX = n + 2
// For 256 bit strings, set MAX = 258
const MAX = 258
const KEYLEN uint16 = 256

var (
	ch      *amqp.Channel
	queue   *amqp.Queue
	ctx     context.Context
	db      *sql.DB
	initial string
)

func init() {
	db, initial = store.DbOpen()
	ch, queue, ctx = rabbitmq.BunnyOpen()
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
	fmt.Println(initial, a)
	for {
		pk := binhex.BinaryToHex(a[1 : MAX-1])
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(pk),
		}
		if err := ch.PublishWithContext(ctx, "", queue.Name, false, false, message); err != nil {
			panic(err)
		}

		// Save key to db once every million iterations
		count++
		if count == 1000000 {
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
