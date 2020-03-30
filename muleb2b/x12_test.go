package muleb2b

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

const (
	getPartnerX12Response = `[
    {
        "envelopeHeaders": {
            "authInfoQualifierISA01": null,
            "authInfoISA02": null,
            "securityInfoQualifierISA03": null,
            "securityInfoISA04": null,
            "interchangeReceiverIdQualifierISA07": null,
            "interchangeReceiverIdISA07": null,
            "repetitionSeparatorCharacterISA11": null,
            "repetitionInterchangeAcknowledmentsISA14": null,
            "defaultInterchangeUsageIndicatorISA15": null,
            "componentElementSeparator": null
        },
        "parserSettings": {
            "failDocumentWhenValueLengthOutsideAllowedRange": true,
            "failDocumentWhenInvalidCharacterInValue": true,
            "failDocumentIfValueIsRepeatedTooManyTimes": true,
            "failDocumentIfUnknownSegmentsAreUsed": true,
            "failDocumentWhenSegmentsAreOutOfOrder": true,
            "failDocumentWhenTooManyRepeatsOfSegment": true,
            "failDocumentWhenUnusedSegmentsAreIncluded": true,
            "require997": false,
            "generate999": false,
            "generateTA1": false,
            "checkDuplicateDays": 30,
            "ackEndpointId": null
        },
        "characterSetAndEncoding": {
            "characterSer": "EXTENDED",
            "characterEncoding": null,
            "lineEndingBetweenSegments": null
        },
        "controlNumberSettings": {
            "initialInterchangeControlNumber": "00",
            "initialGSControlNumber": "00",
            "initialTransactionSetControlNumber": "00",
            "requireUniqueGSControlNumbers": null,
            "requireUniqueTransactionSetControlNumber": null,
            "requireUniqueISAcontrolNumbersISA13": true,
            "requireUniqueGSControlNumbersGS06": false,
            "requireUniqueTransactionSetControlNumbersST02": false
        },
        "id": "51a1e4f3-fa83-4183-b101-6499cb8e36f7",
        "configType": "READ",
        "formatType": "X12InboundConfig",
        "formatTypeId": "25c1bc8a-801f-4337-a2a6-7721ef971460",
        "partnerId": "ae5e69af-9de2-4355-8502-f2a330326ad2"
    },
    {
        "envelopeHeaders": {
            "authInfoQualifierISA01": "00",
            "authInfoISA02": "",
            "securityInfoQualifierISA03": "00",
            "securityInfoISA04": "",
            "interchangeReceiverIdQualifierISA07": "01",
            "interchangeReceiverIdISA07": "",
            "repetitionSeparatorCharacterISA11": "U",
            "repetitionInterchangeAcknowledmentsISA14": null,
            "defaultInterchangeUsageIndicatorISA15": null,
            "componentElementSeparator": "-"
        },
        "terminatorDelimiter": {
            "segmentTerminatorCharacter": "*",
            "dataElementDelimiter": "~",
            "stringSubstituionCharacter": "."
        },
        "characterSetEncoding": {
            "characterSer": "EXTENDED",
            "characterEncoding": "UTF8",
            "lineEndingBetweenSegments": "LFCR"
        },
        "controlNumberSettings": {
            "initialInterchangeControlNumber": "000000001",
            "initialGSControlNumber": "000000001",
            "initialTransactionSetControlNumber": "00000001",
            "requireUniqueGSControlNumbers": true,
            "requireUniqueTransactionSetControlNumber": true,
            "requireUniqueISAcontrolNumbersISA13": null,
            "requireUniqueGSControlNumbersGS06": null,
            "requireUniqueTransactionSetControlNumbersST02": null
        },
        "id": "25c1bc8a-801f-4337-a2a6-7721ef971466",
        "configType": "WRITE",
        "formatType": "X12OutboundConfig",
        "formatTypeId": "25c1bc8a-801f-4337-a2a6-7721ef971460",
        "partnerId": null
    }
]`
)

func TestListPartnerX12Configurations(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "ae5e69af-9de2-4355-8502-f2a330326ad2")) // Partner ID

		w.Write([]byte(getPartnerX12Response))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	c.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	x12List, err := c.ListPartnerX12Configurations("ae5e69af-9de2-4355-8502-f2a330326ad2")

	assert.Nil(t, err)
	assert.Equal(t, 2, len(x12List))
}

func TestGetPartnerInboundX12Configuration(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "ae5e69af-9de2-4355-8502-f2a330326ad2")) // Partner ID

		w.Write([]byte(getPartnerX12Response))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	c.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	x12, err := c.GetPartnerInboundX12Configuration("ae5e69af-9de2-4355-8502-f2a330326ad2")

	assert.Nil(t, err)
	assert.Equal(t, "X12InboundConfig", *x12.FormatType)
}

func TestGetPartnerOutboundX12Configuration(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "ae5e69af-9de2-4355-8502-f2a330326ad2")) // Partner ID

		w.Write([]byte(getPartnerX12Response))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	c.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	x12, err := c.GetPartnerOutboundX12Configuration("ae5e69af-9de2-4355-8502-f2a330326ad2")

	assert.Nil(t, err)
	assert.Equal(t, "X12OutboundConfig", *x12.FormatType)
}

func TestGetPartnerX12ConfigurationById(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "ae5e69af-9de2-4355-8502-f2a330326ad2")) // Partner ID

		w.Write([]byte(getPartnerX12Response))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	c.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	x12, err := c.GetPartnerX12ConfigurationById("ae5e69af-9de2-4355-8502-f2a330326ad2", "51a1e4f3-fa83-4183-b101-6499cb8e36f7")

	assert.Nil(t, err)
	assert.Equal(t, "X12InboundConfig", *x12.FormatType)
}

func TestUpdatePartnerInboundX12Configuration(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "ae5e69af-9de2-4355-8502-f2a330326ad2")) // Partner ID
		assert.True(t, strings.Contains(r.URL.String(), "51a1e4f3-fa83-4183-b101-6499cb8e36f7")) // X12 Config ID

		w.Write([]byte(""))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	c.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	parserSettings := X12ParserSettings{
		FailDocumentWhenValueLengthOutsideAllowedRange: Boolean(true),
		FailDocumentWhenInvalidCharacterInValue:        Boolean(true),
		FailDocumentIfValueIsRepeatedTooManyTimes:      Boolean(true),
		FailDocumentIfUnknownSegmentsAreUsed:           Boolean(true),
		FailDocumentWhenSegmentsAreOutOfOrder:          Boolean(true),
		FailDocumentWhenTooManyRepeatsOfSegment:        Boolean(true),
		FailDocumentWhenUnusedSegmentsAreIncluded:      Boolean(true),
		Require997:  Boolean(false),
		Generate999: Boolean(false),
		GenerateTA1: Boolean(false),
	}

	characterEncoding := X12CharacterSetAndEncoding{
		CharacterSet: String("EXTENDED"),
	}

	controlNumberSettings := X12ControlNumberSettings{
		InitialInterchangeControlNumber:               String("00"),
		InitialGSControlNumber:                        String("00"),
		InitialTransactionSetControlNumber:            String("00"),
		RequireUniqueISAcontrolNumbersISA13:           Boolean(true),
		RequireUniqueGSControlNumbersGS06:             Boolean(false),
		RequireUniqueTransactionSetControlNumbersST02: Boolean(false),
	}

	x12 := X12{
		Id:                      String("51a1e4f3-fa83-4183-b101-6499cb8e36f7"),
		ConfigType:              String("READ"),
		FormatType:              String("X12InboundConfig"),
		FormatTypeId:            String("25c1bc8a-801f-4337-a2a6-7721ef971460"),
		ParserSettings:          &parserSettings,
		CharacterSetAndEncoding: &characterEncoding,
		ControlNumberSettings:   &controlNumberSettings,
	}

	err = c.UpdatePartnerX12Configuration("ae5e69af-9de2-4355-8502-f2a330326ad2", &x12)

	assert.Nil(t, err)
}
