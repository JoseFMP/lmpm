package reservations

import (
	"time"

	"github.com/JoseFMP/lmpm"
	"github.com/JoseFMP/lmpm/fees"
	"github.com/JoseFMP/lmpm/utils"
)

type Reservation struct {
	ID               ReservationID     `json:"id"`
	CompanyID        lmpm.CompanyID    `json:"company_id"`        //: "b2e6a1c3-1a5e-44ae-a8fd-81f76fd715cf",
	TripID           string            `json:"trip_id"`           //: "756bca85-e98d-4cde-a99e-d64c85c74e6f",
	Start            time.Time         `json:"start"`             //: "2019-08-24T14:15:22Z",
	End              time.Time         `json:"end"`               //: "2019-08-24T14:15:22Z",
	Timezone         string            `json:"timezone"`          //: "string",
	BookingType      string            `json:"booking_type"`      //: "renter",
	BookingSource    string            `json:"booking_source"`    //: "string",
	CheckedIn        bool              `json:"checked_in"`        //: true,
	CheckedInAt      time.Time         `json:"checked_in_at"`     //: "2019-08-24T14:15:22Z",
	CheckedOut       bool              `json:"checked_out"`       //: true,
	CheckedOutAt     time.Time         `json:"checked_out_at"`    //: "2019-08-24T14:15:22Z",
	DoorCode         string            `json:"door_code"`         //: "string",
	CheckoutCode     string            `json:"checkout_code"`     //: "string",
	HousekeepingNote string            `json:"housekeeping_note"` //: "string",
	Adults           int8              `json:"adults"`            //: 0,
	Children         int8              `json:"children"`          //: 0,
	Pets             int8              `json:"pets"`              //: 0,
	PropertyId       string            `json:"property_id"`       //: "05003a8a-8f3c-454b-8884-a906ec46f5f5",
	Status           ReservationStatus `json:"status"`            //: "booked",
	CreatedAt        time.Time         `json:"created_at"`        //: "2019-08-24T14:15:22Z",
	UpdatedAt        time.Time         `json:"updated_at"`        //: "2019-08-24T14:15:22Z",
	BookedAt         time.Time         `json:"booked_at"`         //: "2019-08-24T14:15:22Z",
	RentTotal        float64           `json:"rent_total"`        //: 0,
	FeesTotal        float64           `json:"fees_total"`        //: 0,
	TaxTotal         float64           `json:"tax_total"`         //: 0,
	Guest            lmpm.Guest        `json:"guest"`             //: {},
	NightlyRates     []NightlyRate     `json:"nightly_rates"`     //: [],
	FeeItems         []fees.FeeItem    `json:"fee_items"`         //: []

}

type ReservationID string
type ReservationStatus string

type NightlyRate struct {
	DateBangkok utils.CalendarDay `json:"date"`
	Rate        float64           `json:"rate"`
}
