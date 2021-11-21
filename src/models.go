package src

type Cities struct {
	City        string
	RuCity      string
	Code        string
	Country     string
	RuCountry   string
	CountryCode string
}

type Airlines struct {
	Name string
	Code string
}

type Airports struct {
	CityCode   string
	Code       string
	Flightable bool
	Name       string
	RuName     string
}
