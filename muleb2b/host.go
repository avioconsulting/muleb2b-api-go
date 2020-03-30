package muleb2b

import (
	"errors"
	"fmt"
	"net/url"
)

type hostResponse struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
}

func (cli *Client) GetHostDetails() error {

	if cli.envId == nil {
		return errors.New("Environment must be set before host details can be retrieved")
	}

	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partnerprofiles/host", *cli.orgId, *cli.envId)}
	u := cli.BaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	var hostResp hostResponse

	_, err = cli.Do(req, &hostResp)
	if err != nil {
		return err
	}

	cli.hostPartnerId = hostResp.Id
	cli.hostPartnerName = hostResp.Name

	return nil
}
