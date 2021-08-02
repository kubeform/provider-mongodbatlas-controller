// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth // import "go.mongodb.org/realm/auth"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultAuthURL = "https://realm.mongodb.com/api/admin/v3.0/auth/providers/mongodb-cloud/login"
	jsonMediaType  = "application/json"
)

type Config struct {
	client  *http.Client
	AuthURL *url.URL
}

func NewConfig(httpClient *http.Client) *Config {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultAuthURL)

	c := &Config{
		client:  httpClient,
		AuthURL: baseURL,
	}
	return c
}

func (c *Config) NewTokenFromCredentials(ctx context.Context, username, password string) (*Token, error) {
	v := &authenticateRequest{
		Username: username,
		APIKey:   password,
	}

	return c.auth(ctx, v)
}

func NewClient(src TokenSource) *http.Client {
	if src == nil {
		return http.DefaultClient
	}
	return &http.Client{
		Transport: &Transport{
			Base:   http.DefaultTransport,
			Source: src,
		},
	}
}

func BasicTokenSource(t *Token) TokenSource {
	return basicTokenSource{t}
}

// basicTokenSource is a TokenSource that always returns the same Token.
type basicTokenSource struct {
	t *Token
}

func (s basicTokenSource) Token() (*Token, error) {
	return s.t, nil
}

type authenticateRequest struct {
	Username string `json:"username"`
	APIKey   string `json:"apiKey"`
}

func (c *Config) auth(ctx context.Context, v *authenticateRequest) (*Token, error) {
	req, err := c.newAuthRequest(ctx, v)
	if err != nil {
		return nil, err
	}
	return c.doAuthRoundTrip(req)
}

func (c *Config) newAuthRequest(ctx context.Context, v *authenticateRequest) (*http.Request, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.AuthURL.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", jsonMediaType)
	req.Header.Set("Accept", jsonMediaType)

	return req, nil
}

func (c *Config) doAuthRoundTrip(req *http.Request) (*Token, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	const maxBodySlurpSize = 1 << 20
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, maxBodySlurpSize))
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("auth: cannot fetch token: %w", err)
	}
	if code := resp.StatusCode; code < 200 || code > 299 {
		return nil, &RetrieveError{
			Response: resp,
			Body:     body,
		}
	}

	var token *Token
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, err
	}
	if token.AccessToken == "" {
		return nil, errors.New("server response missing access_token")
	}
	return token, nil
}
