package reservations

import "time"

const updateAtFromGetParamKey = `updatedAtFrom`
const updateAtToGetParamKey = `updatedAtTo`
const startFromGetParamKey = `startFrom`
const startToGetParamKey = `startTo`

func composeGetParams(updateAtFrom time.Time, updateAtTo time.Time, startFrom time.Time, startTo time.Time) map[string]string {
	getParams := make(map[string]string)

	addOneParam(getParams, updateAtFrom, updateAtFromGetParamKey)
	addOneParam(getParams, updateAtTo, updateAtToGetParamKey)
	addOneParam(getParams, startFrom, startFromGetParamKey)
	addOneParam(getParams, startTo, startToGetParamKey)

	if len(getParams) > 0 {
		return getParams
	}
	return nil
}

func addOneParam(targetMap map[string]string, timestamp time.Time, getParamKey string) {
	if !timestamp.IsZero() {
		asText, errMarshalling := timestamp.MarshalText()
		if errMarshalling != nil {
			panic("Error marshalling updateAtFrom")
		}
		targetMap[getParamKey] = string(asText)
	}
}
