package transaction

import (
	"testing"

	"github.com/Axway/agent-sdk/pkg/traceability/redaction"
	"github.com/stretchr/testify/assert"
)

func createHTTPProtocol(uri, method, reqHeaders, resHeaders string, status, reqLen, resLen int) (TransportProtocol, error) {
	redaction.SetupGlobalRedaction(redaction.Config{})
	return NewHTTPProtocolBuilder().
		SetURI(uri).
		SetVersion("1.1").
		SetArgs(`{"param1": ["date"], "param2": ["day, time"]}`).
		SetMethod(method).
		SetStatus(status, "statusTxt").
		SetUserAgent("userAgent").
		SetHost("host").
		SetByteLength(reqLen, resLen).
		SetRemoteAddress("remoteName", "remoteAddr", 2222).
		SetLocalAddress("localAddr", 1111).
		SetAuthSubjectID("authsubject").
		SetSSLProperties("TLS1.1", "sslServer", "sslSubject").
		SetHeaders(reqHeaders, resHeaders).
		SetIndexedHeaders(`{"indexedrequest": "value", "x-amplify-indexed": "random", "x-amplify-indexedagain": "else"}`,
			`{"indexedresponse": "value", "x-indexedresponse": "random", "x-indexed": "test"}`).
		SetPayload("requestPayload", "responsePayload").
		SetWAFStatus(1).
		Build()
}

func TestHTTPProtocolBuilder(t *testing.T) {
	httpProtocol, err := createHTTPProtocol("/testuri", "GET", `{"request": "value", "x-amplify-something": "random", "x-amplify-somethingelse": "else"}`,
		`{"response": "value", "x-response": "random", "x-value": "test"}`, 200, 10, 10)
	assert.Nil(t, err)
	assert.NotNil(t, httpProtocol)

	httpProtocolBuilder := NewHTTPProtocolBuilder()

	httpProtocol, err = httpProtocolBuilder.Build()
	assert.Nil(t, httpProtocol)
	assert.NotNil(t, err)
	assert.Equal(t, "uri property not set in HTTP protocol details", err.Error())

	httpProtocol, err = httpProtocolBuilder.
		SetURI("/test").
		Build()
	assert.Nil(t, httpProtocol)
	assert.NotNil(t, err)
	assert.Equal(t, "method property not set in HTTP protocol details", err.Error())
	httpProtocol, err = httpProtocolBuilder.
		SetURI("/test").
		SetMethod("GET").
		Build()
	assert.Nil(t, httpProtocol)
	assert.NotNil(t, err)
	assert.Equal(t, "host property not set in HTTP protocol details", err.Error())

	httpProtocol, err = httpProtocolBuilder.
		SetURI("/test").
		SetMethod("GET").
		SetHost("host").
		SetStatus(20, "OK").
		Build()
	assert.Nil(t, httpProtocol)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid status code set in HTTP protocol details", err.Error())

	httpProtocol, err = httpProtocolBuilder.
		SetURI("/test").
		SetMethod("GET").
		SetHost("host").
		SetStatus(200, "OK").
		Build()
	assert.Nil(t, err)
	assert.NotNil(t, httpProtocol)

	httpProtocol, err = httpProtocolBuilder.
		SetURI("/test").
		SetMethod("GET").
		SetHost("host").
		SetStatus(200, "OK").
		AddArg("newarg", []string{"one", "two"}).
		Build()
	assert.Nil(t, err)
	assert.NotNil(t, httpProtocol)

	httpProtocol, err = httpProtocolBuilder.
		SetURI("/test").
		SetMethod("GET").
		SetHost("host").
		SetStatus(200, "OK").
		SetArgsMap(map[string][]string{"test": {"one", "two"}}).
		Build()
	assert.Nil(t, err)
	assert.NotNil(t, httpProtocol)

	httpProtocol, err = httpProtocolBuilder.
		SetURI("/test").
		SetMethod("GET").
		SetHost("host").
		SetStatus(200, "OK").
		SetRequestHeaders(map[string]string{"reqHead": "one"}).
		SetResponseHeaders(map[string]string{"rspHead": "two"}).
		Build()
	assert.Nil(t, err)
	assert.NotNil(t, httpProtocol)

	httpProtocol, err = httpProtocolBuilder.
		SetURI("/test").
		SetMethod("GET").
		SetHost("host").
		SetStatus(200, "OK").
		AddRequestHeader("key", "two").
		AddResponseHeader("key", "two").
		Build()
	assert.Nil(t, err)
	assert.NotNil(t, httpProtocol)

	httpProtocol, err = httpProtocolBuilder.
		SetURI("/test").
		SetMethod("GET").
		SetHost("host").
		SetStatus(200, "OK").
		SetIndexedRequestHeaders(map[string]string{"test": "one"}).
		SetIndexedResponseHeaders(map[string]string{"test": "two"}).
		Build()
	assert.Nil(t, err)
	assert.NotNil(t, httpProtocol)

	httpProtocol, err = httpProtocolBuilder.
		SetURI("/test").
		SetMethod("GET").
		SetHost("host").
		SetStatus(200, "OK").
		AddIndexedRequestHeader("key", "one").
		AddIndexedResponseHeader("key", "two").
		Build()
	assert.Nil(t, err)
	assert.NotNil(t, httpProtocol)
}
