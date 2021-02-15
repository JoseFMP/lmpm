package reservations

import (
	"encoding/json"
	"testing"

	"github.com/JoseFMP/lmpm"
	"github.com/stretchr/testify/require"
)

func TestParseReservation(t *testing.T) {

	// arrange
	var reservation Reservation

	// act
	errUnmarshalling := json.Unmarshal([]byte(mockRawReservation), &reservation)

	// verify

	require.Nil(t, errUnmarshalling)
	require.NotNil(t, reservation)
	require.Equal(t, reservation.ID, ReservationID("497f6eca-6276-4993-bfeb-53cbbbba6f08"))
	require.Equal(t, reservation.CompanyID, lmpm.CompanyID("b2e6a1c3-1a5e-44ae-a8fd-81f76fd715cf"))
	require.Equal(t, reservation.TripID, "756bca85-e98d-4cde-a99e-d64c85c74e6f")

	require.Equal(t, reservation.RentTotal, 90.0)
	require.Equal(t, reservation.FeesTotal, 10.0)
	require.Equal(t, reservation.TaxTotal, 7.0)

	require.Len(t, reservation.FeeItems, 1)
	require.NotNil(t, reservation.FeeItems[0])
	require.Equal(t, reservation.FeeItems[0].Name, "Horse riding")
	require.Equal(t, reservation.FeeItems[0].Rate, 10.0)

}

const mockRawReservation = `
{
	"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
	"company_id": "b2e6a1c3-1a5e-44ae-a8fd-81f76fd715cf",
	"trip_id": "756bca85-e98d-4cde-a99e-d64c85c74e6f",
	"start": "2019-08-24T14:15:22Z",
	"end": "2019-08-24T14:15:22Z",
	"timezone": "string",
	"booking_type": "renter",
	"booking_source": "string",
	"checked_in": true,
	"checked_in_at": "2019-08-24T14:15:22Z",
	"checked_out": true,
	"checked_out_at": "2019-08-24T14:15:22Z",
	"door_code": "string",
	"checkout_code": "string",
	"housekeeping_note": "string",
	"adults": 0,
	"children": 0,
	"pets": 0,
	"property_id": "05003a8a-8f3c-454b-8884-a906ec46f5f5",
	"status": "booked",
	"created_at": "2019-08-24T14:15:22Z",
	"updated_at": "2019-08-24T14:15:22Z",
	"booked_at": "2019-08-24T14:15:22Z",
	"rent_total": 90,
	"fees_total": 10,
	"tax_total": 7,
	"guest": {
	"name": "string",
	"phone": "string",
	"email": "string",
	"location": {
	"country": "string",
	"state": "string",
	"city": "string",
	"postal_code": "string"
	}
	},
	"nightly_rates": [
	{
	"date": "2019-08-24",
	"rate": 0
	}
	],
	"fee_items": [
	{
	"name": "Horse riding",
	"rate": 10.0,
	"housekeeping_note": "string"
	}
	]
	}
`
