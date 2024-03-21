package coingecko

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"go.uber.org/ratelimit"
)

// Client struct
type Client struct {
	httpClient  *http.Client
	rateLimiter ratelimit.Limiter
}

// NewClient create new client object
func NewClient(httpClient *http.Client) *Client {
	return newClientWithLimiter(httpClient, ratelimit.New(30, ratelimit.Per(time.Minute)))
}

func newClientWithLimiter(cli *http.Client, limiter ratelimit.Limiter) *Client {
	if cli == nil {
		cli = http.DefaultClient
	}
	if limiter == nil {
		limiter = ratelimit.NewUnlimited()
	}

	return &Client{
		httpClient:  cli,
		rateLimiter: limiter,
	}
}

func NewLimitClient(b int, every time.Duration, cli *http.Client) *Client {
	var limiter ratelimit.Limiter
	if every > 0 && b > 0 {
		limiter = ratelimit.New(b, ratelimit.Per(every))
	}
	return newClientWithLimiter(cli, limiter)
}

// helper
// doReq HTTP client
func doReq(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// MakeReq HTTP request helper
func (c *Client) MakeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	c.rateLimiter.Take()

	resp, err := doReq(req, c.httpClient)
	if err != nil {
		return nil, err
	}
	return resp, err
}
