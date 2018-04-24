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
	data["ActionId"] = "Test"
	data["DevicesId"] = "123"

	bytesData, _ := json.Marshal(data)

	reader := bytes.NewReader(bytesData)

	resp, err := http.Post("http://192.168.13.55:60003/frontend/json", "application/json", reader)
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
