package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const coinMarketCapAPI = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"
const apiKey = "54877cdc-3ef1-4d21-844d-a0b5c80fa866"

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func getCoinPrice(coinName string) (float64, error) {
	req, err := http.NewRequest("GET", coinMarketCapAPI, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("X-CMC_PRO_API_KEY", apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	data, ok := result["data"].([]interface{})
	if !ok {
		return 0, fmt.Errorf("Invalid data format")
	}

	for _, v := range data {
		entry, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if entry["name"].(string) == coinName {
			quote := entry["quote"].(map[string]interface{})
			usd := quote["USD"].(map[string]interface{})
			return usd["price"].(float64), nil
		}
	}

	return 0, fmt.Errorf("Coin not found")
}

type CoinRequest struct {
	CoinName string `json:"coinName"`
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	log.Println("Connection established")

	for {

		var request CoinRequest

		if err := conn.ReadJSON(&request); err != nil {
			log.Println(err)
			break
		}

		log.Println(request)
		coinName := request.CoinName

		log.Println(coinName)
		price, err := getCoinPrice(coinName)
		if err != nil {
			log.Println(err)
			break
		}

		response := map[string]float64{
			coinName: price,
		}

		if err := conn.WriteJSON(response); err != nil {
			log.Println(err)
			break
		}
		time.Sleep(30 * time.Second) // Pobieraj co 30 sekund.
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
