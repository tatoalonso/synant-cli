package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	wordscli "github.com/tatoalonso/synant-cli/internal"
)

const (
	synonymsEndpoint = "/sinonimos"
	antonymsEndpoint = "/antonimos"
	format           = "/json/"
	apiURL           = "http://sesat.fdi.ucm.es:8080/servicios/rest"
)

type wordRepo struct {
	url string
}

// NewAPIRepository fetch beers from api
func NewAPIRepository() wordscli.WordsRepo {
	return &wordRepo{url: apiURL}
}

func (w *wordRepo) GetWords(typeOfWord, word string) (list wordscli.ListOfWords, err error) {

	var endPoint string

	switch typeOfWord {
	case "synonym":
		endPoint = fmt.Sprintf("%v%v%v%v", apiURL, synonymsEndpoint, format, word)
	case "antonym":
		endPoint = fmt.Sprintf("%v%v%v%v", apiURL, antonymsEndpoint, format, word)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", endPoint, nil)

	if err != nil {
		return wordscli.ListOfWords{nil, nil}, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return wordscli.ListOfWords{nil, nil}, err
	}

	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return wordscli.ListOfWords{nil, nil}, err
	}

	err = json.Unmarshal(contents, &list)

	if err != nil {
		return wordscli.ListOfWords{nil, nil}, err
	}

	return
}
