package bmkg

var (
	testClient *Client
)

func setup() {
	testClient = NewClient(nil)
	// testClient.BaseURL = testServer.URL
}
