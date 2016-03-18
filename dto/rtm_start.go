package dto

// RtmStart is the dto structure used to decode Slack rtm start response
type RtmStart struct {
	URL  string `json:"url"`
	Self struct {
		ID string `json:"id"`
	} `json:"self"`

	OK    bool   `json:"ok"`
	Error string `json:"error"`
}
