package gof1

// Circuit type contains details of a sepecific F1 racing circuit
type Circuit struct {
	CircuitID   string `json:"circuitId"`
	URL         string `json:"url"`
	CircuitName string `json:"circuitName"`
	Location    struct {
		Lat      string `json:"lat"`
		Long     string `json:"long"`
		Locality string `json:"locality"`
		Country  string `json:"country"`
	} `json:"Location"`
}

// Driver type details a specific F1 driver who has participated in 1 or more recorded sessions.
type Driver struct {
	DriverID        string `json:"driverId"`
	PermanentNumber string `json:"permanentNumber"`
	Code            string `json:"code"`
	URL             string `json:"url"`
	GivenName       string `json:"givenName"`
	FamilyName      string `json:"familyName"`
	DateOfBirth     string `json:"dateOfBirth"`
	Nationality     string `json:"nationality"`
}

// Contructor type details an F1 team.
type Constructor struct {
	ConstructorID string `json:"constructorId"`
	URL           string `json:"url"`
	Name          string `json:"name"`
	Nationality   string `json:"nationality"`
}

// Result type details the race result for a specific driver.
type Result struct {
	Number       string      `json:"number"`
	Position     string      `json:"position"`
	PositionText string      `json:"positionText"`
	Points       int16       `json:"points"`
	Driver       Driver      `json:"Driver"`
	Constructor  Constructor `json:"Constructor"`
	Grid         string      `json:"grid"`
	Laps         string      `json:"laps"`
	Status       string      `json:"status"`
	Time         struct {
		Millis string `json:"millis"`
		Time   string `json:"time"`
	} `json:"Time,omitempty"`
	FastestLap struct {
		Rank string `json:"rank"`
		Lap  string `json:"lap"`
		Time struct {
			Time string `json:"time"`
		} `json:"Time"`
		AverageSpeed struct {
			Units string `json:"units"`
			Speed string `json:"speed"`
		} `json:"AverageSpeed"`
	} `json:"FastestLap,omitempty"`
}

// Race type details a specific event in a specific season of F1
type Race struct {
	Season   string   `json:"season"`
	Round    string   `json:"round"`
	URL      string   `json:"url"`
	RaceName string   `json:"raceName"`
	Circuit  Circuit  `json:"Circuit"`
	Date     string   `json:"date"`
	Time     string   `json:"time"`
	Results  []Result `json:"Results"`
}

// F1 type is essentially a wrapper type :)
type F1 struct {
	MRData struct {
		RaceTable struct {
			Season string `json:"season"`
			Round  string `json:"round"`
			Races  []Race `json:"Races"`
		} `json:"RaceTable"`
	} `json:"MRData"`
}
