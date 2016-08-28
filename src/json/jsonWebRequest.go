package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type GitData struct {
	Name     string
	Language string
}

func main() {
	url := "https://api.github.com/users/mralexgray/repos"

	respdata := getJsonData(url)
	jsonOutPut := stringTojson(respdata)

	for _, jsondata := range jsonOutPut {
		fmt.Printf(" Name : %s     ->  language: %s  \n", jsondata.Name, jsondata.Language)
	}

}

func getJsonData(url string) string {

	resp, error := http.Get(url)

	handleError(error)
	defer resp.Body.Close()
	bytes, readError := ioutil.ReadAll(resp.Body)
	bodyContenet := string(bytes)
	handleError(readError)
	return bodyContenet
}

func handleError(ioError error) {
	if ioError != nil {
		panic(ioError)
	}
}

func stringTojson(rawData string) []GitData {

	gitdata := make([]GitData, 0, 25)
	decoder := json.NewDecoder(strings.NewReader(rawData))
	_, err := decoder.Token()
	handleError(err)

	var tmpgitdata GitData
	for decoder.More() {
		err := decoder.Decode(&tmpgitdata)
		handleError(err)
		gitdata = append(gitdata, tmpgitdata)
	}
	return gitdata
}
