package core

var client Client

type Client struct {
	AccessToken string
	AppInfo     AccessTokenServer
}

func GetClient() Client {
	return client
}
