package fetcher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetGreetingFromURL(fetchTarget string) (ElasticGreeting, error) {
	greeting := ElasticGreeting{}

	resp, err := http.Get(fetchTarget)
	if err != nil {
		return greeting, err
	}

	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return greeting, err
	}
	if err := json.Unmarshal(jsonBytes, &greeting); err != nil {
		return greeting, err
	}

	return greeting, nil
}
