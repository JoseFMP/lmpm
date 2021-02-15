package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/JoseFMP/lmpm"
)

func DoReq(req *http.Request) ([]byte, *LMPMErrorPayload, error) {

	httpClient := http.Client{}

	resp, errDoingReq := httpClient.Do(req)

	if errDoingReq != nil {
		return nil, nil, errDoingReq
	}

	payload, errReadingPayload := ioutil.ReadAll(resp.Body)
	if errReadingPayload != nil {
		return nil, nil, errReadingPayload
	}

	if resp.StatusCode != http.StatusOK {
		errorMsg := fmt.Sprintf("Error making request: %d - %s", resp.StatusCode, resp.Status)

		var payloadPointer *LMPMErrorPayload
		if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusBadRequest {

			var errorPayload LMPMErrorPayload
			errUnmarshalling := json.Unmarshal(payload, &errorPayload)
			if errUnmarshalling == nil {
				payloadPointer = &errorPayload
			}
		}

		return nil, payloadPointer, fmt.Errorf(errorMsg)
	}

	return payload, nil, nil
}

func CreateGetReq(environment lmpm.Environment, subpath string, apiKey string, urlParams map[string]string) (*http.Request, error) {

	return createReq(http.MethodGet, environment, subpath, apiKey, urlParams)
}

func createReq(method string, environment lmpm.Environment, subpath string, apiKey string, urlParams map[string]string) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", environment, subpath)

	req, errCreatingReq := http.NewRequest(method, url, nil)

	if errCreatingReq != nil {
		return nil, errCreatingReq
	}

	req.Header.Add(xApiKeyHeaderName, apiKey)

	if urlParams != nil {
		q := req.URL.Query()
		for k, v := range urlParams {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	return req, errCreatingReq
}

const xApiKeyHeaderName = `x-api-key`
