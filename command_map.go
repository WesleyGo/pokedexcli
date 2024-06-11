package main

import (
	"errors"
	"fmt"
)

func commandListLocations(conf *config) error {
	locationsResp, err := conf.pokeapiClient.ListLocations(conf.next)
	if err != nil {
		return err
	}

	conf.next = locationsResp.Next
	conf.previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandListPreviousLocations(conf *config) error {
	if conf.previous == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := conf.pokeapiClient.ListLocations(conf.previous)
	if err != nil {
		return err
	}

	conf.next = locationsResp.Next
	conf.previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
