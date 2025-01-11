package structure

import (
	"api/domain/requests"
	"api/lib/utils/formatstring"
	"net/http"
	"bytes"
	"encoding/json"
	"io"
)

func ConstructorInstance() *requests.MethodAPI {
	return &requests.MethodAPI{
		POST:    "POST",
		PUT:     "PUT",
		GET:     "GET",
		DELETE:  "DELETE",
		OPTIONS: "OPTIONS",
		Headers: make(map[string]string),
	}
}

func RESTApi(method, url string, headersAPI *requests.MethodAPI, body interface{}) (string, error) {
	headers := headersAPI.Headers
	var bodyBytes *bytes.Buffer
	if body != nil {
		bodyData, err := json.Marshal(body)
		if err != nil {
			return "", formatstring.FormatStringErrorWithDetails("ErrJsonEncode", err)
		}
		bodyBytes = bytes.NewBuffer(bodyData)
	} else {
		bodyBytes = nil
	}
	req, err := http.NewRequest(method, url, bodyBytes)
	if err != nil {
		return "", formatstring.FormatStringErrorWithDetails("ErrRequestCreate", err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", formatstring.FormatStringErrorWithDetails("ErrRequestSending", err)
	}
	defer resp.Body.Close()
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", formatstring.FormatStringErrorWithDetails("ErrReadResponseBody", err)
	}
	return string(bodyResp), nil
}