package gokinde

// KindeURLs holds the URLs for redirection
type KindeURLs struct {
	SiteUrl         string
	RedirectUrl     string
	UnAuthorisedUrl string
}

type KindeCredentials struct {
	IssuerBaseUrl   string
	RedirectUrl     string
	SiteUrl         string
	Secret          string
	UnAuthorisedUrl string
	ClientID        string
}
