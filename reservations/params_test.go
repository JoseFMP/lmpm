package reservations

import (
	"testing"
	"time"

	"github.com/JoseFMP/lmpm/utils"
	"github.com/stretchr/testify/require"
)

func TestAddOneParam(t *testing.T) {

	// arrange
	mockMap := make(map[string]string)
	mockTimeStamp := time.Unix(1613380252, 0).In(utils.GetThailandTimeZone())
	mockKey := "lalala"

	// act
	addOneParam(mockMap, mockTimeStamp, mockKey)

	// verify
	require.Len(t, mockMap, 1)

	timeSerialized, exists := mockMap[mockKey]
	require.True(t, exists)

	timeParsed, errParsing := time.Parse(time.RFC3339, timeSerialized)
	require.Nil(t, errParsing)
	timeParsed = timeParsed.In(mockTimeStamp.Location())
	require.Equal(t, mockTimeStamp, timeParsed)
}

func TestComposeGetParams(t *testing.T) {

	// arrange
	updateAtFrom, errParsingUpdateAtFrom := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if errParsingUpdateAtFrom != nil {
		panic("Wrong updateAtFrom")
	}
	updatedAtTo, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	startFrom, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	startTo, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

	// act
	params := composeGetParams(updateAtFrom, updatedAtTo, startFrom, startTo)

	// verify
	require.NotNil(t, params)
	require.Len(t, params, 4)

}
