/*
Package slackdk contains the implementation of slack rtm for easy communication
with Slack
*/
package slackdk

// NewClient returns a new Slack client based on the given API Key
func NewClient(key string) Client {
	return &clientImpl{
		apiToken: key,
	}
}
