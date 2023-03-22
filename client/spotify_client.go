package spotify_client

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"log"
	"os"
	"bytes"

)


type Config struct {
	BaseUrl		string	`json:"base_url"`
	Username	string	`json:"username"`
	AuthEndpoint	string	`json:"auth"`

}


type ApiKeys struct {
	ClientId	string	`json:"client_id"`
	ClientSecret	string	`json:"client_secret"`
}


func readConfig() Config {
	// Reading the config file 'config.json'
	configFile, err := os.Open("client/config.json")
	if err != nil {
		log.Fatal("Error reading config file: ", err)

	}
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)

	var config Config

	json.Unmarshal(byteValue, &config)

	return config

}


func readKeys() ApiKeys {
	// Reading the keys from file 'spotify.json'
	credFile, err := os.Open("client/spotify.json")
	if err != nil {
		log.Fatal("Error reading config file: ", err)

	}
	defer credFile.Close()

	byteValue, _ := ioutil.ReadAll(credFile)

	var keys ApiKeys

	json.Unmarshal(byteValue, &keys)

	return keys
}


func GetAuthToken()	string {
	//getting an auth token
	keys := readKeys()
	config := readConfig()

	spotifyAuth := config.BaseUrl + config.AuthEndpoint
	jsonBody := []byte(`{"grant_type": "client_credentials"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, spotifyAuth, bodyReader)
	if err != nil {
		log.Fatal("error building request: ", err)

	}

	var basic_client_creds string = "Basic " + keys.ClientId + ":" + keys.ClientSecret

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", basic_client_creds)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("error performing request: ", err)
	}

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("error reading: ", err)
	}
	returnString := string(response[:])
	return returnString

}


