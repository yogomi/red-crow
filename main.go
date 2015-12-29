package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	url := "http://api-sandbox.oanda.com/v1/prices?instruments=EUR_USD"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Accept-Datetime-Format", "UNIX")

	if err != nil {
		fmt.Println(err)
	}

	resp, _ := client.Do(req)

	byteArray, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("%#v\n", resp.Header)
	etag := resp.Header.Get("Etag")
	fmt.Println(etag)
	fmt.Println(resp.Status)

	req.Header.Add("If-None-Match", etag)
	fmt.Printf("%#v\n", req.Header)

	resp, _ = client.Do(req)
	fmt.Println(resp.Status)
	defer resp.Body.Close()

	byteArray, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}
