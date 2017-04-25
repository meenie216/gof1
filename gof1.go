package gof1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// GetRaces queries the Ergast api and returns the details of every completed race in the F1 season specified
func GetRaces(year int) []Race {
	url := fmt.Sprintf("http://ergast.com/api/f1/%v/results.json?limit=1000&offset=0", year)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("NewRequest:", err)
		return nil
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Do:", err)
		return nil
	}
	defer resp.Body.Close()

	var result F1

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println(err)
	}

	return result.MRData.RaceTable.Races
}
