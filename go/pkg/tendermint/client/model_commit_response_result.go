/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CommitResponseResult struct for CommitResponseResult
type CommitResponseResult struct {
	SignedHeader CommitResponseResultSignedHeader `json:"signed_header"`
	Canonical bool `json:"canonical"`
}

// NewCommitResponseResult instantiates a new CommitResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitResponseResult(signedHeader CommitResponseResultSignedHeader, canonical bool) *CommitResponseResult {
	this := CommitResponseResult{}
	this.SignedHeader = signedHeader
	this.Canonical = canonical
	return &this
}

// NewCommitResponseResultWithDefaults instantiates a new CommitResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitResponseResultWithDefaults() *CommitResponseResult {
	this := CommitResponseResult{}
	return &this
}

// GetSignedHeader returns the SignedHeader field value
func (o *CommitResponseResult) GetSignedHeader() CommitResponseResultSignedHeader {
	if o == nil {
		var ret CommitResponseResultSignedHeader
		return ret
	}

	return o.SignedHeader
}

// GetSignedHeaderOk returns a tuple with the SignedHeader field value
// and a boolean to check if the value has been set.
func (o *CommitResponseResult) GetSignedHeaderOk() (*CommitResponseResultSignedHeader, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SignedHeader, true
}

// SetSignedHeader sets field value
func (o *CommitResponseResult) SetSignedHeader(v CommitResponseResultSignedHeader) {
	o.SignedHeader = v
}

// GetCanonical returns the Canonical field value
func (o *CommitResponseResult) GetCanonical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *CommitResponseResult) GetCanonicalOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *CommitResponseResult) SetCanonical(v bool) {
	o.Canonical = v
}

func (o CommitResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["signed_header"] = o.SignedHeader
	}
	if true {
		toSerialize["canonical"] = o.Canonical
	}
	return json.Marshal(toSerialize)
}

type NullableCommitResponseResult struct {
	value *CommitResponseResult
	isSet bool
}

func (v NullableCommitResponseResult) Get() *CommitResponseResult {
	return v.value
}

func (v *NullableCommitResponseResult) Set(val *CommitResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitResponseResult(val *CommitResponseResult) *NullableCommitResponseResult {
	return &NullableCommitResponseResult{value: val, isSet: true}
}

func (v NullableCommitResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

