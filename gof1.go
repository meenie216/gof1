package gof1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const baseURL = "http://ergast.com/api/f1"

// GetRacesInSeason returns details of all races in the season specified.  Results are not returned.
func GetRacesInSeason(year int) []Race {
	url := fmt.Sprintf("%s/%v/schedule.json", baseURL, year)
	result := makeAPICall(url)
	return result.MRData.RaceTable.Races
}

// GetRacesWithResults queries the Ergast api and returns the details of every completed race in the F1 season specified
func GetRacesWithResults(year int) []Race {
	url := fmt.Sprintf("%s/%v/results.json?limit=1000&offset=0", baseURL, year)
	result := makeAPICall(url)
	return result.MRData.RaceTable.Races
}

// GetRaceWithResults retrieves details and results of the race specified by year and number
func GetRaceWithResults(year int, raceNumber int) Race {
	url := fmt.Sprintf("%s/%v/%v/results.json?limit=1000&offset=0", baseURL, year, raceNumber)
	result := makeAPICall(url)

	var r Race
	if len(result.MRData.RaceTable.Races) != 1 {
		return r
	}

	return result.MRData.RaceTable.Races[0]
}

func makeAPICall(url string) F1 {
	var result F1

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("NewRequest:", err)
		return result
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Do:", err)
		return result
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println(err)
	}

	return result
}
