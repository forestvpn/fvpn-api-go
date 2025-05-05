package subscriber

import (
	"context"
	"encoding/base64"
	"github.com/coder/websocket/wsjson"
	"github.com/forestvpn/fvpn-api-go"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

const subProtocol = "message-queue-v1"

type ClientOption func(s *Client)

func WithAPIKey(apiKey fvpn.APIKey) ClientOption {
	return func(s *Client) {
		s.headers.Set(apiKey.Prefix, apiKey.Key)
	}
}

func WithBasicAuth(username, password string) ClientOption {
	return func(s *Client) {
		s.headers.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(username+":"+password)))
	}
}

func WithHeader(key, val string) ClientOption {
	return func(s *Client) {
		s.headers.Set(key, val)
	}
}

// Client is a utility for receiving wireguard key events from a message-queue server
type Client struct {
	url     string
	headers http.Header
}

func New(url string, options ...ClientOption) *Client {
	s := &Client{url: url, headers: http.Header{}}

	for _, option := range options {
		option(s)
	}

	return s
}

// Subscribe establishes a websocket connection for a message-queue channel, and emits messages on the given channel
func (s *Client) Subscribe(ctx context.Context, channel chan<- Event) error {
	err := s.connect(ctx, channel)

	if err != nil {
		return err
	}

	return nil
}

func (s *Client) connect(ctx context.Context, channel chan<- Event) error {
	conn, _, err := websocket.Dial(ctx, s.url, &websocket.DialOptions{
		Subprotocols: []string{subProtocol},
		HTTPHeader:   s.headers,
	})

	if err != nil {
		return err
	}

	go s.read(ctx, channel, conn)

	return nil
}

func (s *Client) read(ctx context.Context, channel chan<- Event, conn *websocket.Conn) {
	for {
		v := Event{}
		err := wsjson.Read(ctx, conn, &v)
		if err != nil {
			log.Println("error reading from websocket, reconnecting", err)

			// Make sure the connection is closed
			conn.Close(websocket.StatusInternalError, "")

			// Start attempting to reconnect
			go s.reconnect(ctx, channel)

			return
		}

		channel <- v
	}
}

func (s *Client) reconnect(ctx context.Context, channel chan<- Event) {
	// Sleep
	time.Sleep(time.Second)

	// Attempt to create a new connection
	err := s.connect(ctx, channel)
	if err != nil {
		go s.reconnect(ctx, channel)
	} else {
		log.Println("successfully reconnected to websocket")
	}
}
