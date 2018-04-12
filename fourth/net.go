package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"bytes"
	"net/http"
)

func main() {
	data := make(map[string]interface{})
	data["ActionId"] = 2308
	data["skip"] = 2
	data["AppId"] = 1010
	data["limit"] = 10

	bytesData, _ := json.Marshal(data)

	reader := bytes.NewReader(bytesData)

	resp, err := http.Post("http://127.0.0.1:60003/frontend/json", "application/json", reader)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}
