package muleb2b

import (
	"net/url"
)

// Authentication request type
type authRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

// Authentication response type
type authResponse struct {
	AccessToken *string `json:"access_token"`
	TokenType   *string `json:"token_type"`
}

func (cli *Client) Login(username, password string) error {
	rel := &url.URL{Path: "accounts/login"}
	u := cli.BaseURL.ResolveReference(rel)

	authReq := authRequest{&username, &password}

	req, err := cli.NewRequest("POST", u.String(), &authReq)
	if err != nil {
		return err
	}

	var authResponse authResponse
	_, err = cli.Do(req, &authResponse)
	if err != nil {
		return err
	}
	cli.accessToken = authResponse.AccessToken
	return nil
}
