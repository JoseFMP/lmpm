package main

import (
	"log"
	"time"

	"github.com/JoseFMP/lmpm"
	"github.com/JoseFMP/lmpm/reservations"
)

func main() {

	conf := readTestConfig()

	lmpmClient, errInitializing := lmpm.Init(string(conf.CompanyID), conf.APIKey, conf.BaseUrl)

	if errInitializing != nil {
		panic(errInitializing)
	}

	var reservatitionsClient *reservations.LMPMReservationClient = (*reservations.LMPMReservationClient)(lmpmClient)

	now := time.Now()

	updateAtFrom := now.Add(-time.Hour * 24 * 10)
	updateAtTo := now.Add(-time.Hour * 24 * 1)
	startFrom := now.Add(-time.Hour * 24 * 100)
	startTo := now.Add(-time.Hour * 24 * 1)
	reservations, errGettingRes := reservatitionsClient.List(updateAtFrom, updateAtTo, startFrom, startTo)
	if errGettingRes != nil {
		panic(errGettingRes)
	}
	log.Print(reservations)
}
