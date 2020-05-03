package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type Negotiation struct {
	URL                     string  `json:"Url"`
	ConnectionToken         string  `json:"ConnectionToken"`
	ConnectionID            string  `json:"ConnectionId"`
	KeepAliveTimeout        float64 `json:"KeepAliveTimeout"`
	DisconnectTimeout       float64 `json:"DisconnectTimeout"`
	ConnectionTimeout       float64 `json:"ConnectionTimeout"`
	TryWebSockets           bool    `json:"TryWebSockets"`
	ProtocolVersion         string  `json:"ProtocolVersion"`
	TransportConnectTimeout float64 `json:"TransportConnectTimeout"`
	LongPollDelay           float64 `json:"LongPollDelay"`
}

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	simplepushkey := os.Getenv("SIMPLEPUSH_KEY")
	token := ""
	if resp, err := http.Get("https://www.coinpanic.com/signalr/negotiate?clientProtocol=1.5&connectionData=%5B%7B%22name%22%3A%22notificationhub%22%7D%5D"); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var negotiation Negotiation
			if err := json.Unmarshal(body, &negotiation); err == nil {
				token = url.QueryEscape(negotiation.ConnectionToken)
			} else {
				log.Fatal("negotiation:", err)
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
	c, _, err := websocket.DefaultDialer.Dial("wss://www.coinpanic.com/signalr/connect?transport=webSockets&clientProtocol=1.5&connectionToken="+token+"&connectionData=%5B%7B%22name%22%3A%22notificationhub%22%7D%5D&tid=7", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			if strings.Contains(string(message), "Deposit") {
				fmt.Println(time.Now(), "FreeMoney!")
				if _, err := http.Get("https://api.simplepush.io/send/" + simplepushkey + "/Community Jar/New Deposit"); err != nil {
					log.Println("notify: SIMPLEPUSH_KEY:", simplepushkey, "ERROR:", err)
				}
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, nil)
			if err != nil {
				log.Println("write("+t.String()+"):", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
