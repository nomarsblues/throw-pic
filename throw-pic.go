package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var api = "https://sm.ms/api/v2/upload"

func upload(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("cannot open file", err)
		return
	}
	data, _ := ioutil.ReadAll(file)
	boundary := "----WebKitFormBoundarypBaTXQhCOVePZw7g"
	body := "--" + boundary + "\n"
	body = body + "Content-Disposition: form-data; name=\"smfile\"; filename=\"" + path + "\"\n"
	body = body + "Content-Type: image/png\n\n"
	body = body + string(data) + "\n"
	body = body + "--" + boundary + "--"
	request, _ := http.NewRequest(http.MethodPost, api, strings.NewReader(body))
	request.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)
	response, _ := http.DefaultClient.Do(request)
	responseBody, _ := ioutil.ReadAll(response.Body)
	var result Result
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		fmt.Println("from json err ", err)
	}
	if result.Success {
		fmt.Println(result.Data.Url)
	} else {
		fmt.Println(result.Images)
	}
}

func main() {
	var path string
	flag.StringVar(&path, "p", "", "path")
	flag.Parse()
	upload(path)
}

type Result struct {
	Images string
	Success bool
	Data Data
}

type Data struct {
	Url string
}
