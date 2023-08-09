package atomicasset

import (
	"github.com/eosswedenorg-go/unixtime"
)

// Types

type ChainHealth struct {
	Status    string        `json:"status"`
	HeadBlock int64         `json:"head_block"`
	HeadTime  unixtime.Time `json:"head_time"`
}

type RedisHealth struct {
	Status string `json:"status"`
}

type PostgresHealth struct {
	Status  string                   `json:"status"`
	Readers []map[string]interface{} `json:"readers"`
}

type HealthData struct {
	Version  string         `json:"version"`
	Postgres PostgresHealth `json:"postgres"`
	Redis    RedisHealth    `json:"redis"`
	Chain    ChainHealth    `json:"chain"`
}

// Responses

type Health struct {
	APIResponse
	Data HealthData
}

// Client API Functions

// GetHealth fetches "/health" from API
func (c *Client) GetHealth() (Health, error) {
	var health Health

	r, err := c.fetch("GET", "/health", nil, &health.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&health)
	}
	return health, err
}
