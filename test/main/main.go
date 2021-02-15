package main

import (
	"log"
	"time"

	"github.com/JoseFMP/lmpm"
	"github.com/JoseFMP/lmpm/reservations"
	"github.com/JoseFMP/lmpm/utils"
)

func main() {

	conf := readTestConfig()

	lmpmClient, errInitializing := lmpm.Init(string(conf.CompanyID), conf.APIKey, conf.BaseUrl)

	if errInitializing != nil {
		panic(errInitializing)
	}

	var reservatitionsClient *reservations.LMPMReservationClient = (*reservations.LMPMReservationClient)(lmpmClient)

	now := time.Now()

	updatedAt := utils.Period{
		From: now.Add(-time.Hour * 24 * 100),
		To:   now.Add(-time.Hour * 24 * 1),
	}

	start := utils.Period{
		From: now.Add(-time.Hour * 24 * 100),
		To:   now.Add(-time.Hour * 24 * 1),
	}
	reservations, errGettingRes := reservatitionsClient.List(updatedAt, start)
	if errGettingRes != nil {
		panic(errGettingRes)
	}
	log.Print(reservations)
}
