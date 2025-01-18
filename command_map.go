package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapf(c *config) error {
	locations, err := getLocations(c.Next)
	if err != nil {
		return err
	}

	c.Previous = locations.Previous
	c.Next = locations.Next
	printLocations(locations.Results)

	return nil
}

func commandMapb(c *config) error {
	if c.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := getLocations(c.Previous)
	if err != nil {
		return err
	}

	c.Previous = locations.Previous
	c.Next = locations.Next
	printLocations(locations.Results)

	return nil
}

func getLocations(url *string) (pokeApiResponse, error) {
	res, err := http.Get(*url)
	if err != nil {
		return pokeApiResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokeApiResponse{}, err
	}

	var response pokeApiResponse

	if err := json.Unmarshal(data, &response); err != nil {
		return pokeApiResponse{}, err
	}

	return response, nil
}

func printLocations(a []Areas) {
	for _, a := range a {
		fmt.Println(a.Name)
	}
}

type pokeApiResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []Areas `json:"results"`
}

type Areas struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
