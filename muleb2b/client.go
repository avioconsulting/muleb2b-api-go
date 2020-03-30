package muleb2b

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://anypoint.mulesoft.com/"
	partnerApiPath = "partnermanager/partners/api/v1/"
)

type Client struct {
	BaseURL         *url.URL
	client          *http.Client
	accessToken     *string
	orgId           *string
	envId           *string
	hostPartnerId   *string
	hostPartnerName *string
	PartnerBaseURL  *url.URL
}

func NewClient(baseURL, orgId *string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	var myBase string
	if baseURL == nil || len(*baseURL) < 1 {
		myBase = defaultBaseURL
	} else {
		myBase = *baseURL
	}
	base, err := url.Parse(myBase)
	if err != nil {
		return nil, err
	}

	rel := &url.URL{Path: fmt.Sprintf("%s", partnerApiPath)}
	partnerApiUrl := base.ResolveReference(rel)

	cli := &Client{
		BaseURL:        base,
		PartnerBaseURL: partnerApiUrl,
		orgId:          orgId,
		client:         httpClient,
	}

	return cli, nil
}

func (cli *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(cli.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", cli.BaseURL)
	}
	u, err := cli.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")

	if cli.accessToken != nil {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *cli.accessToken))
	}

	return req, nil
}

func (cli *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := cli.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		if v != nil {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		return resp, fmt.Errorf("error code (%d) received from service, response body: %s\n", resp.StatusCode, string(bodyBytes))
	}
	return resp, err
}

func String(str string) *string {
	return &str
}

func Boolean(val bool) *bool {
	return &val
}

func Integer(val int) *int {
	return &val
}
