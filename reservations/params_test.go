package reservations

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAddOneParam(t *testing.T) {

	// arrange
	mockMap := make(map[string]string)
	mockTimeStamp := time.Unix(1613380252, 0)
	mockKey := "lalala"

	// act
	addOneParam(mockMap, mockTimeStamp, mockKey)

	// verify
	require.Len(t, mockMap, 1)

	timeSerialized, exists := mockMap[mockKey]
	require.True(t, exists)

	timeParsed, errParsing := time.Parse(time.RFC3339, timeSerialized)
	require.Nil(t, errParsing)
	require.Equal(t, mockTimeStamp, timeParsed)
}
