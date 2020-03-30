package muleb2b

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

type X12 struct {
	Id                      *string                     `json:"id,omitempty"`
	ConfigType              *string                     `json:"configType"`
	FormatType              *string                     `json:"formatType"`
	FormatTypeId            *string                     `json:"formatTypeId"`
	PartnerId               *string                     `json:"partnerId"`
	IsTemplate              *bool                       `json:"isTemplate,omitempty"`
	EnvelopeHeaders         *X12EnvelopeHeaders         `json:"envelopeHeaders"`
	ParserSettings          *X12ParserSettings          `json:"parserSettings,omitempty"`
	CharacterSetAndEncoding *X12CharacterSetAndEncoding `json:"characterSetAndEncoding,omitempty"`
	ControlNumberSettings   *X12ControlNumberSettings   `json:"controlNumberSettings,omitempty"`
	TerminatorDelimiter     *X12TerminatorDelimiter     `json:"terminatorDelimiter,omitempty"`
}

type X12EnvelopeHeaders struct {
	AuthInfoQualifierISA01                   *string `json:"authInfoQualifierISA01,omitempty"`
	AuthInfoISA02                            *string `json:"authInfoISA02,omitempty"`
	SecurityInfoQualifierISA03               *string `json:"securityInfoQualifierISA03,omitempty"`
	SecurityInfoISA04                        *string `json:"securityInfoISA04,omitempty"`
	InterchangeReceiverIdQualifierISA07      *string `json:"interchangeReceiverIdQualifierISA07,omitempty"`
	InterchangeReceiverIdISA07               *string `json:"interchangeReceiverIdISA07,omitempty"`
	RepetitionSeparatorCharacterISA11        *string `json:"repetitionSeparatorCharacterISA11,omitempty"`
	RepetitionInterchangeAcknowledmentsISA14 *string `json:"repetitionInterchangeAcknowledmentsISA14,omitempty"`
	DefaultInterchangeUsageIndicatorISA15    *string `json:"defaultInterchangeUsageIndicatorISA15,omitempty"`
	ComponentElementSeparator                *string `json:"componentElementSeparator,omitempty"`
}

type X12ParserSettings struct {
	FailDocumentWhenValueLengthOutsideAllowedRange *bool   `json:"failDocumentWhenValueLengthOutsideAllowedRange,omitempty"`
	FailDocumentWhenInvalidCharacterInValue        *bool   `json:"failDocumentWhenInvalidCharacterInValue,omitempty"`
	FailDocumentIfValueIsRepeatedTooManyTimes      *bool   `json:"failDocumentIfValueIsRepeatedTooManyTimes,omitempty"`
	FailDocumentIfUnknownSegmentsAreUsed           *bool   `json:"failDocumentIfUnknownSegmentsAreUsed,omitempty"`
	FailDocumentWhenSegmentsAreOutOfOrder          *bool   `json:"failDocumentWhenSegmentsAreOutOfOrder,omitempty"`
	FailDocumentWhenTooManyRepeatsOfSegment        *bool   `json:"failDocumentWhenTooManyRepeatsOfSegment,omitempty"`
	FailDocumentWhenUnusedSegmentsAreIncluded      *bool   `json:"failDocumentWhenUnusedSegmentsAreIncluded,omitempty"`
	Require997                                     *bool   `json:"require997,omitempty"`
	Generate999                                    *bool   `json:"generate999,omitempty"`
	GenerateTA1                                    *bool   `json:"generateTA1,omitempty"`
	CheckDuplicateDays                             *int    `json:"checkDuplicateDays,omitempty"`
	AckEndpointId                                  *string `json:"ackEndpointId"`
}

type X12CharacterSetAndEncoding struct {
	CharacterSet              *string `json:"characterSer"`
	CharacterEncoding         *string `json:"characterEncoding"`
	LineEndingBetweenSegments *string `json:"lineEndingBetweenSegments"`
}

type X12ControlNumberSettings struct {
	InitialInterchangeControlNumber               *string `json:"initialInterchangeControlNumber,omitempty"`
	InitialGSControlNumber                        *string `json:"initialGSControlNumber,omitempty"`
	InitialTransactionSetControlNumber            *string `json:"initialTransactionSetControlNumber"`
	RequireUniqueGSControlNumbers                 *bool   `json:"requireUniqueGSControlNumbers"`
	RequireUniqueTransactionSetControlNumber      *bool   `json:"requireUniqueTransactionSetControlNumber"`
	RequireUniqueISAcontrolNumbersISA13           *bool   `json:"requireUniqueISAcontrolNumbersISA13"`
	RequireUniqueGSControlNumbersGS06             *bool   `json:"requireUniqueGSControlNumbersGS06"`
	RequireUniqueTransactionSetControlNumbersST02 *bool   `json:"requireUniqueTransactionSetControlNumbersST02"`
}

type X12TerminatorDelimiter struct {
	SegmentTerminatorCharacter  *string `json:"segmentTerminatorCharacter"`
	DataElementDelimiter        *string `json:"dataElementDelimiter"`
	StringSubstitutionCharacter *string `json:"stringSubstituionCharacter"`
}

// List the X12 Configurations for a partner
func (cli *Client) ListPartnerX12Configurations(partnerId string) ([]X12, error) {
	if cli.envId == nil {
		return nil, errors.New("Environment must be set before Partner details can be retrieved")
	}

	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/ediFormats/X12/configurations", *cli.orgId, *cli.envId, partnerId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp []X12

	_, err = cli.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (cli *Client) GetPartnerInboundX12Configuration(partnerId string) (*X12, error) {
	x12s, err := cli.ListPartnerX12Configurations(partnerId)

	if err != nil {
		return nil, err
	}

	for _, x12 := range x12s {
		if *x12.FormatType == "X12InboundConfig" {
			return &x12, nil
		}
	}

	return nil, errors.New("X12InboundConfig was not found")
}

func (cli *Client) GetPartnerOutboundX12Configuration(partnerId string) (*X12, error) {
	x12s, err := cli.ListPartnerX12Configurations(partnerId)
	if err != nil {
		return nil, err
	}
	for _, x12 := range x12s {
		if *x12.FormatType == "X12OutboundConfig" {
			return &x12, nil
		}
	}
	return nil, errors.New("X12OutboundConfig was not found")
}

/*
 * TODO: Change this implementation after the API is implemented
 */
func (cli *Client) GetPartnerX12ConfigurationById(partnerId, configId string) (*X12, error) {
	x12s, err := cli.ListPartnerX12Configurations(partnerId)
	if err != nil {
		return nil, err
	}
	for _, x12 := range x12s {
		if x12.Id != nil && *x12.Id == configId {
			return &x12, nil
		}
	}
	return nil, nil
}

func (cli *Client) CreatePartnerX12Configuration(partnerId string, x12 *X12) error {
	if cli.envId == nil {
		return errors.New("Environment must be set before Partner details can be updated")
	}
	if x12 == nil {
		return fmt.Errorf("cannot create a nil X12")
	}

	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/ediFormats/X12/configurations", *cli.orgId, *cli.envId, partnerId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	x12.IsTemplate = Boolean(true)
	req, err := cli.NewRequest("POST", u.String(), x12)
	if err != nil {
		return err
	}

	_, err = cli.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

// Updates a partner Inbound X12 Configuration (Outbound is not allowed yet)
// Only update is available at this point
func (cli *Client) UpdatePartnerX12Configuration(partnerId string, x12 *X12) error {
	if cli.envId == nil {
		return errors.New("Environment must be set before Partner details can be updated")
	}
	if x12 == nil {
		return fmt.Errorf("cannot update with a nil X12")
	}

	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/ediFormats/X12/configurations/%s", *cli.orgId, *cli.envId, partnerId, *x12.Id)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	x12.IsTemplate = Boolean(false)
	req, err := cli.NewRequest("PUT", u.String(), x12)
	if err != nil {
		return err
	}

	_, err = cli.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (x12 *X12) String() string {
	j, _ := json.Marshal(x12)
	return string(j)
}
