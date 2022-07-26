package twitter

import (
	"github.com/valyala/fasthttp"
	"swapper/package/globals"
	"encoding/json"
	"math/rand"
	"time"
)

var (
	updateTimer bool = true
	hasClaimed  bool = false
)

func timer() {
	for updateTimer {
		globals.TimeElapsed++
		time.Sleep(1 * time.Millisecond)
	}
}

func randomLetterNumberString(n int) string {
	var characterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, n)
    for i := range b {
        b[i] = characterRunes[rand.Intn(len(characterRunes))]
    }
    return string(b)
}

func releaseUsername(newName string, rareAccount Account) bool {
	var client = &fasthttp.Client{}

	req := fasthttp.AcquireRequest()

	req.SetRequestURI("https://twitter.com/i/api/1.1/account/settings.json?screen_name=" + newName)
	req.Header.SetMethod(fasthttp.MethodPost)
	
	req.Header.Set("authorization", rareAccount.Authorization)
	req.Header.Set("cookie", "auth_token=" + rareAccount.AuthToken + "; ct0=" + rareAccount.XCsrfToken + ";")
	req.Header.Set("x-csrf-token", rareAccount.XCsrfToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err != nil {
		globals.Errors++
		SaveToLogFile(err.Error(), "ERROR")
		return false
	}
	
	var settingsResponse SettingsResponse
	json.Unmarshal(resp.Body(), &settingsResponse)

	if resp.StatusCode() == 429 {
		globals.Ratelimits++; return false
	} else if resp.StatusCode() == 200 && settingsResponse.ScreenName == newName {
		 return true
	} else {
		return false
	}
}

func claimUsername(target string, recievingAccount Account) {
	var client = &fasthttp.Client{}

	req := fasthttp.AcquireRequest()

	req.SetRequestURI("https://twitter.com/i/api/1.1/account/settings.json?screen_name=" + target)
	req.Header.SetMethod(fasthttp.MethodPost)
	
	req.Header.Set("authorization", recievingAccount.Authorization)
	req.Header.Set("cookie", "auth_token=" + recievingAccount.AuthToken + "; ct0=" + recievingAccount.XCsrfToken + ";")
	req.Header.Set("x-csrf-token", recievingAccount.XCsrfToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)

	fasthttp.ReleaseRequest(req)
	
	if err != nil {
		globals.Errors++
		SaveToLogFile(err.Error(), "ERROR")
		return;
	}

	go handleResponse(target, resp)
}

func handleResponse(target string, resp *fasthttp.Response) {
	defer fasthttp.ReleaseResponse(resp)

	var settingsResponse SettingsResponse
	json.Unmarshal(resp.Body(), &settingsResponse)

	if resp.StatusCode() == 429 {
		globals.Ratelimits++
	} else if resp.StatusCode() == 200 && settingsResponse.ScreenName == target { 
		hasClaimed = true
		updateTimer = false
	} else {
		globals.Requests++
	}
}

func Swap(target string, receivingAccount Account, rareAccount Account, rareAccountNewName *string) {
	sem := make(chan int, 2) // 2 Concurrent Goroutines

	go func() {
		for !hasClaimed {
			sem <- 1
		
			go func (target string, recivingAccount Account) {
				defer func() {
					<-sem
				}()
				claimUsername(target, receivingAccount)
			
			time.Sleep(100 * time.Microsecond)
			
			}(target, receivingAccount)
		}
	}()
	
	go func() {
		for {
			*rareAccountNewName = randomLetterNumberString(10)
			completed := releaseUsername(*rareAccountNewName, rareAccount)
			if completed { go timer(); break }
		}
	}()
	
	for {
		if hasClaimed {
			break
		}
		time.Sleep(10 * time.Microsecond)
	}
}
