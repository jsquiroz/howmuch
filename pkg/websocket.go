package pkg

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	streamURL = "wss://ws.coincap.io/prices?assets=%s"
)

var wg sync.WaitGroup

func listen(id string, price chan float64) {
	url := fmt.Sprintf(streamURL, id)
	c, _, err := websocket.DefaultDialer.DialContext(context.Background(), url, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	msg := make(map[string]string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			c.ReadJSON(&msg)

			p, err := strconv.ParseFloat(msg[id], 64)
			if err != nil {
				continue
			}

			price <- p
		}

	}()

	wg.Wait()
}
