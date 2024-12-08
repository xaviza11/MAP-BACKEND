package database

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Country struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	LAT  float64 `json:"lat"`
	LON  float64 `json:"lon"`
}

type Faction struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Support string `json:"support"`
}

type Export struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Import struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Port struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Railway struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CountryInfo struct {
	ID        int       `json:"id"`
	CountryID int       `json:"country_id"`
	Factions  []Faction `json:"factions"`
	Notes     string    `json:"notes"`
	Exports   []Export  `json:"exports"`
	Imports   []Import  `json:"imports"`
	Ports     []Port    `json:"ports"`
	Railway   []Railway `json:"railway"`
}
