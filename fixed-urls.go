package shopeego

func init() {
	availablePaths["GetAccessToken"] = "/api/v2/auth/token/get"
	availablePaths["RefreshAccessToken"] = "/api/v2/auth/access_token/get"
}
