package muleb2b

import (
	"fmt"
	"net/url"
)

type EdiFormat struct {
	Id          *string `json:"id"`
	FormatType  *string `json:"formatType"`
	Description *string `json:"description"`
	Label       *string `json:"label"`
}

type EdiFormatVersion struct {
	Id          *string `json:"id"`
	FormatType  *string `json:"formatType"`
	Version     *string `json:"version"`
	Description *string `json:"description"`
	Label       *string `json:"label"`
}

func (cli *Client) ListEdiFormats() (*[]EdiFormat, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/ediFormats", *cli.orgId, *cli.envId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var ediResp []EdiFormat

	_, err = cli.Do(req, &ediResp)
	if err != nil {
		return nil, err
	}

	return &ediResp, nil
}

func (cli *Client) GetEdiFormatByFormat(format string) (*EdiFormat, error) {
	ediFormats, err := cli.ListEdiFormats()

	if err != nil {
		return nil, err
	}

	if ediFormats != nil {
		for _, ediFormat := range *ediFormats {
			if *ediFormat.FormatType == format {
				return &ediFormat, nil
			}
		}
	}

	return nil, nil
}

func (cli *Client) GetEdiFormatByLabel(label string) (*EdiFormat, error) {
	ediFormats, err := cli.ListEdiFormats()

	if err != nil {
		return nil, err
	}

	if ediFormats != nil {
		for _, ediFormat := range *ediFormats {
			if *ediFormat.Label == label {
				return &ediFormat, nil
			}
		}
	}

	return nil, nil
}

func (cli *Client) GetEdiFormatById(id string) (*EdiFormat, error) {
	ediFormats, err := cli.ListEdiFormats()

	if err != nil {
		return nil, err
	}

	if ediFormats != nil {
		for _, ediFormat := range *ediFormats {
			if *ediFormat.Id == id {
				return &ediFormat, nil
			}
		}
	}

	return nil, nil
}

func (cli *Client) ListEdiFormatVersions(formatType string) (*[]EdiFormatVersion, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/ediFormats/%s/ediFormatVersions", *cli.orgId, *cli.envId, formatType)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var versions []EdiFormatVersion

	_, err = cli.Do(req, &versions)
	if err != nil {
		return nil, err
	}

	return &versions, nil
}

func (cli *Client) GetEdiFormatVersionByLabel(formatType, label string) (*EdiFormatVersion, error) {
	versions, err := cli.ListEdiFormatVersions(formatType)

	if err != nil {
		return nil, err
	}

	if versions != nil {
		for _, version := range *versions {
			if *version.Label == label {
				return &version, nil
			}
		}
	}
	return nil, nil
}

func (cli *Client) GetEdiFormatVersionByVersion(formatType, version string) (*EdiFormatVersion, error) {
	versions, err := cli.ListEdiFormatVersions(formatType)

	if err != nil {
		return nil, err
	}

	if versions != nil {
		for _, v := range *versions {
			if *v.Version == version {
				return &v, nil
			}
		}
	}
	return nil, nil
}

func (cli *Client) GetEdiFormatVersionById(formatType, id string) (*EdiFormatVersion, error) {
	versions, err := cli.ListEdiFormatVersions(formatType)

	if err != nil {
		return nil, err
	}

	if versions != nil {
		for _, version := range *versions {
			if *version.Id == id {
				return &version, nil
			}
		}
	}
	return nil, nil
}
