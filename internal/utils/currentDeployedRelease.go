package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type StatusResponse struct {
	Hostname    string `json:"hostname"`
	Environment string `json:"environment"`
	Ok          bool   `json:"ok"`
	Version     string `json:"version"`
}

func GetCurrentDeployedRelease(url string) string {

	var statusResponse StatusResponse

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&statusResponse)

	return statusResponse.Version
}
