package twitter

type Account struct {
    Authorization string
    AuthToken     string
    XCsrfToken    string
}

type SettingsResponse struct {
    ScreenName string `json:"screen_name"`
}
