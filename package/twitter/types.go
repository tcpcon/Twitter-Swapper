package twitter

type Account struct {
    Authorization string
    AuthToken     string
    XCsrfToken    string
}

type SettingsResponse struct {
    ScreenName string `json:"screen_name"`
}

type SwapStatistics struct {
    TimeElapsed             int 
    RareAccountNewName      string
    ReceivingAccountNewName string
}
