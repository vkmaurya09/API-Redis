package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis"
)

type F2week struct {
	Low                 string `json:"low"`
	High                string `json:"high"`
	Low_change          string `json:"low_change"`
	High_change         string `json:"high_change"`
	Low_change_percent  string `json:"low_change_percent"`
	High_change_percent string `json:"high_change_percent"`
	Range               string `json:"range"`
}
type Data struct {
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Exchange       string `json:"exchange"`
	Mic_code       string `json:"mic_code"`
	Currency       string `json:"currency"`
	Datetime       string `json:"datetime"`
	Timestamp      int64  `json:"timestamp"`
	Open           string `json:"open"`
	High           string `json:"high"`
	Low            string `json:"low"`
	Close          string `json:"close"`
	Volume         string `json:"volume"`
	Previous_close string `json:"previous_close"`
	Change         string `json:"change"`
	Percent_change string `json:"percent_change"`
	Average_volume string `json:"average_volume"`
	Is_market_open bool   `json:"is_market_open"`
	Fifty_two_week F2week `json:"fifty_two_week"`
}

func main() {
	// redis connection
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	// api
	url := "https://twelve-data1.p.rapidapi.com/quote?symbol=AAPL&interval=1day&outputsize=30&format=json"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "8ac6baebddmsh6727062dd5d2c95p1e8e0djsn6845680a60fe")
	req.Header.Add("X-RapidAPI-Host", "twelve-data1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	// reading data
	data, _ := ioutil.ReadAll(res.Body)
	// unMarshalData
	unMarshalData := Data{}
	err := json.Unmarshal([]byte(data), &unMarshalData)
	if err != nil {
		panic(err)
	}

	// fmt.Println(unMarshalData)
	// set data
	err = client.Set("stock", data, 0).Err()
	if err != nil {
		panic(err)
	}
	// get data
	val, err := client.Get("unMarshalData").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

}
