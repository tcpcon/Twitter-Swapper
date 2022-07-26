package twitter

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
)

func VerifyAccount(account Account) (string, error) {
	var client = &fasthttp.Client{}

	req := fasthttp.AcquireRequest()

	req.SetRequestURI("https://twitter.com/i/api/1.1/account/settings.json")

	req.Header.SetMethod(fasthttp.MethodGet)
	
	req.Header.Set("authorization", account.Authorization)
	req.Header.Set("cookie", "auth_token=" + account.AuthToken + "; ct0=" + account.XCsrfToken + ";")
	req.Header.Set("x-csrf-token", account.XCsrfToken)

	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)

	if err != nil {
		SaveToLogFile(err.Error(), "ERROR")
	}
	
	var settingsResponse SettingsResponse
	json.Unmarshal(resp.Body(), &settingsResponse)

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)

	return settingsResponse.ScreenName, err
}

func CheckRatelimit(account Account) bool {
	var client = &fasthttp.Client{}

	req := fasthttp.AcquireRequest()

	req.SetRequestURI("https://twitter.com/i/api/1.1/account/settings.json")

	req.Header.SetMethod(fasthttp.MethodPost)
	
	req.Header.Set("authorization", account.Authorization)
	req.Header.Set("cookie", "auth_token=" + account.AuthToken + "; ct0=" + account.XCsrfToken + ";")
	req.Header.Set("x-csrf-token", account.XCsrfToken)

	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err != nil {
		SaveToLogFile(err.Error(), "ERROR")
	}
	
	if resp.StatusCode() == 429 {
		return true
	}

	return false
}
