package slackdk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync/atomic"

	"golang.org/x/net/websocket"

	"github.com/alienchow/slackdk/dto"
)

const (
	slackStartURL = "https://slack.com/api/rtm.start?token="
)

// Client is the exported slackdk interface for communciations with Slack
type Client interface {
	// Connect it to initiate a websocket connection to Slacj
	Connect() error

	// Close is to disconnect the Slack client and stop any running jobs
	Close() error

	// Send sends a message to Slack
	Send(*Message) error
}

// clientImpl is the actual logic implementation of the Client interface
type clientImpl struct {
	apiToken string
	id       string
	ws       *websocket.Conn
	wsURL    string
	counter  uint64
}

// Connect implements the Client interface
func (c *clientImpl) Connect() error {
	var err error
	if err = c.execRtmStart(); err != nil {
		return err
	}

	c.ws, err = websocket.Dial(c.wsURL, "", "https://api.slack.com/")
	if err != nil {
		return err
	}
	return nil
}

// execRtmStart retrieves websocket url
func (c *clientImpl) execRtmStart() error {
	resp, err := http.Get(slackStartURL + c.apiToken)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	rtmResp := &dto.RtmStart{}
	err = json.Unmarshal(body, rtmResp)
	if err != nil {
		return err
	}

	if !rtmResp.OK {
		return errors.New(rtmResp.Error)
	}

	c.id = rtmResp.Self.ID
	c.wsURL = rtmResp.URL
	return nil
}

// Close implements the Client interface
func (c *clientImpl) Close() error {
	return c.ws.Close()
}

// Send implements the Client interface
func (c *clientImpl) Send(m *Message) error {
	m.ID = c.getCounter()
	return websocket.JSON.Send(c.ws, m)
}

func (c *clientImpl) getCounter() uint64 {
	return atomic.AddUint64(&c.counter, 1)
}
