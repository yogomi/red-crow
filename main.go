package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
)

const token_length = 65

func main() {
	url := "https://api-fxtrade.oanda.com/v1/accounts"
	client := &http.Client{}

	file, err := os.Open("token.txt")
	if err != nil {
		fmt.Println(err)
	}
	data := make([]byte, token_length)
	count, err := file.Read(data)
	fmt.Printf("%d %q\n", count, data);
	token := string(data)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Accept-Datetime-Format", "UNIX")
	req.Header.Set("Accept-Encoding", "")
	req.Header.Set("Authorization", "Bearer " + token)

	if err != nil {
		fmt.Println(err)
	}

	resp, _ := client.Do(req)

	byteArray, _ := ioutil.ReadAll(resp.Body)
	etag := resp.Header.Get("Etag")
	fmt.Println(etag)
	fmt.Println(resp.Status)
	defer resp.Body.Close()

	fmt.Println(string(byteArray))
}
