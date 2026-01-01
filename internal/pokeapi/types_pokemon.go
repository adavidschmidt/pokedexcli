package pokeapi

type LocationPokemon struct {
	Name          string       `json:"name"`
	EncounterList []Encounters `json:"pokemon_encounters"`
}

type Encounters struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
}
