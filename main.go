package main

var locationUrl = "https://pokeapi.co/api/v2/location-area"

func main() {
	config := &config{
		Next:     &locationUrl,
		Previous: nil,
	}

	initRepl(config)
}
