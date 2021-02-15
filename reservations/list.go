package reservations

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/JoseFMP/lmpm"
	"github.com/JoseFMP/lmpm/utils"
)

type LMPMReservationClient lmpm.Client

func (client LMPMReservationClient) List(updateAtFrom time.Time, updateAtTo time.Time, startFrom time.Time, startTo time.Time) ([]Reservation, error) {

	subpath := composePath(client.CompanyID)

	getParams := composeGetParams(updateAtFrom, updateAtTo, startFrom, startTo)

	req, errCreatingReq := utils.CreateGetReq(client.BaseUrl, subpath, client.APIKey, getParams)
	if errCreatingReq != nil {
		return nil, errCreatingReq
	}

	respPayload, errPayload, errDoingReq := utils.DoReq(req)
	if errDoingReq != nil {
		if errPayload != nil {
			log.Print(errPayload)
		}
		return nil, errDoingReq
	}

	var reservations []Reservation

	errUnmarshalling := json.Unmarshal(respPayload, &reservations)

	if errUnmarshalling != nil {
		return reservations, errUnmarshalling
	}
	return reservations, nil
}

func composePath(companyID lmpm.CompanyID) string {

	return fmt.Sprintf("companies/%s/reservations", string(companyID))
}
