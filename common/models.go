package common

type Configs struct {
	CloudFlareApiEmail string
	CloudFlareApiKey   string
}

type Error struct {
	Code    int
	Status  string
	Message string
}
