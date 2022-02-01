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

// DumpConsensusResponseResultPeerStateRoundState struct for DumpConsensusResponseResultPeerStateRoundState
type DumpConsensusResponseResultPeerStateRoundState struct {
	Height string `json:"height"`
	Round string `json:"round"`
	Step int32 `json:"step"`
	StartTime string `json:"start_time"`
	Proposal bool `json:"proposal"`
	ProposalBlockPartsHeader DumpConsensusResponseResultPeerStateRoundStateProposalBlockPartsHeader `json:"proposal_block_parts_header"`
	ProposalPolRound NullableInt32 `json:"proposal_pol_round"`
	ProposalPol NullableString `json:"proposal_pol"`
	Prevotes NullableString `json:"prevotes"`
	Precommits NullableString `json:"precommits"`
	LastCommitRound NullableInt32 `json:"last_commit_round"`
	LastCommit NullableString `json:"last_commit"`
	CatchupCommitRound NullableInt32 `json:"catchup_commit_round"`
	CatchupCommit NullableString `json:"catchup_commit"`
}

// NewDumpConsensusResponseResultPeerStateRoundState instantiates a new DumpConsensusResponseResultPeerStateRoundState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDumpConsensusResponseResultPeerStateRoundState(height string, round string, step int32, startTime string, proposal bool, proposalBlockPartsHeader DumpConsensusResponseResultPeerStateRoundStateProposalBlockPartsHeader, proposalPolRound NullableInt32, proposalPol NullableString, prevotes NullableString, precommits NullableString, lastCommitRound NullableInt32, lastCommit NullableString, catchupCommitRound NullableInt32, catchupCommit NullableString) *DumpConsensusResponseResultPeerStateRoundState {
	this := DumpConsensusResponseResultPeerStateRoundState{}
	this.Height = height
	this.Round = round
	this.Step = step
	this.StartTime = startTime
	this.Proposal = proposal
	this.ProposalBlockPartsHeader = proposalBlockPartsHeader
	this.ProposalPolRound = proposalPolRound
	this.ProposalPol = proposalPol
	this.Prevotes = prevotes
	this.Precommits = precommits
	this.LastCommitRound = lastCommitRound
	this.LastCommit = lastCommit
	this.CatchupCommitRound = catchupCommitRound
	this.CatchupCommit = catchupCommit
	return &this
}

// NewDumpConsensusResponseResultPeerStateRoundStateWithDefaults instantiates a new DumpConsensusResponseResultPeerStateRoundState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDumpConsensusResponseResultPeerStateRoundStateWithDefaults() *DumpConsensusResponseResultPeerStateRoundState {
	this := DumpConsensusResponseResultPeerStateRoundState{}
	return &this
}

// GetHeight returns the Height field value
func (o *DumpConsensusResponseResultPeerStateRoundState) GetHeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultPeerStateRoundState) GetHeightOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *DumpConsensusResponseResultPeerStateRoundState) SetHeight(v string) {
	o.Height = v
}

// GetRound returns the Round field value
func (o *DumpConsensusResponseResultPeerStateRoundState) GetRound() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Round
}

// GetRoundOk returns a tuple with the Round field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultPeerStateRoundState) GetRoundOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Round, true
}

// SetRound sets field value
func (o *DumpConsensusResponseResultPeerStateRoundState) SetRound(v string) {
	o.Round = v
}

// GetStep returns the Step field value
func (o *DumpConsensusResponseResultPeerStateRoundState) GetStep() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Step
}

// GetStepOk returns a tuple with the Step field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultPeerStateRoundState) GetStepOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Step, true
}

// SetStep sets field value
func (o *DumpConsensusResponseResultPeerStateRoundState) SetStep(v int32) {
	o.Step = v
}

// GetStartTime returns the StartTime field value
func (o *DumpConsensusResponseResultPeerStateRoundState) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultPeerStateRoundState) GetStartTimeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *DumpConsensusResponseResultPeerStateRoundState) SetStartTime(v string) {
	o.StartTime = v
}

// GetProposal returns the Proposal field value
func (o *DumpConsensusResponseResultPeerStateRoundState) GetProposal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Proposal
}

// GetProposalOk returns a tuple with the Proposal field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultPeerStateRoundState) GetProposalOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Proposal, true
}

// SetProposal sets field value
func (o *DumpConsensusResponseResultPeerStateRoundState) SetProposal(v bool) {
	o.Proposal = v
}

// GetProposalBlockPartsHeader returns the ProposalBlockPartsHeader field value
func (o *DumpConsensusResponseResultPeerStateRoundState) GetProposalBlockPartsHeader() DumpConsensusResponseResultPeerStateRoundStateProposalBlockPartsHeader {
	if o == nil {
		var ret DumpConsensusResponseResultPeerStateRoundStateProposalBlockPartsHeader
		return ret
	}

	return o.ProposalBlockPartsHeader
}

// GetProposalBlockPartsHeaderOk returns a tuple with the ProposalBlockPartsHeader field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultPeerStateRoundState) GetProposalBlockPartsHeaderOk() (*DumpConsensusResponseResultPeerStateRoundStateProposalBlockPartsHeader, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProposalBlockPartsHeader, true
}

// SetProposalBlockPartsHeader sets field value
func (o *DumpConsensusResponseResultPeerStateRoundState) SetProposalBlockPartsHeader(v DumpConsensusResponseResultPeerStateRoundStateProposalBlockPartsHeader) {
	o.ProposalBlockPartsHeader = v
}

// GetProposalPolRound returns the ProposalPolRound field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetProposalPolRound() int32 {
	if o == nil || o.ProposalPolRound.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ProposalPolRound.Get()
}

// GetProposalPolRoundOk returns a tuple with the ProposalPolRound field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetProposalPolRoundOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProposalPolRound.Get(), o.ProposalPolRound.IsSet()
}

// SetProposalPolRound sets field value
func (o *DumpConsensusResponseResultPeerStateRoundState) SetProposalPolRound(v int32) {
	o.ProposalPolRound.Set(&v)
}

// GetProposalPol returns the ProposalPol field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetProposalPol() string {
	if o == nil || o.ProposalPol.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProposalPol.Get()
}

// GetProposalPolOk returns a tuple with the ProposalPol field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetProposalPolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProposalPol.Get(), o.ProposalPol.IsSet()
}

// SetProposalPol sets field value
func (o *DumpConsensusResponseResultPeerStateRoundState) SetProposalPol(v string) {
	o.ProposalPol.Set(&v)
}

// GetPrevotes returns the Prevotes field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetPrevotes() string {
	if o == nil || o.Prevotes.Get() == nil {
		var ret string
		return ret
	}

	return *o.Prevotes.Get()
}

// GetPrevotesOk returns a tuple with the Prevotes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetPrevotesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Prevotes.Get(), o.Prevotes.IsSet()
}

// SetPrevotes sets field value
func (o *DumpConsensusResponseResultPeerStateRoundState) SetPrevotes(v string) {
	o.Prevotes.Set(&v)
}

// GetPrecommits returns the Precommits field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetPrecommits() string {
	if o == nil || o.Precommits.Get() == nil {
		var ret string
		return ret
	}

	return *o.Precommits.Get()
}

// GetPrecommitsOk returns a tuple with the Precommits field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetPrecommitsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Precommits.Get(), o.Precommits.IsSet()
}

// SetPrecommits sets field value
func (o *DumpConsensusResponseResultPeerStateRoundState) SetPrecommits(v string) {
	o.Precommits.Set(&v)
}

// GetLastCommitRound returns the LastCommitRound field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetLastCommitRound() int32 {
	if o == nil || o.LastCommitRound.Get() == nil {
		var ret int32
		return ret
	}

	return *o.LastCommitRound.Get()
}

// GetLastCommitRoundOk returns a tuple with the LastCommitRound field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetLastCommitRoundOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastCommitRound.Get(), o.LastCommitRound.IsSet()
}

// SetLastCommitRound sets field value
func (o *DumpConsensusResponseResultPeerStateRoundState) SetLastCommitRound(v int32) {
	o.LastCommitRound.Set(&v)
}

// GetLastCommit returns the LastCommit field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetLastCommit() string {
	if o == nil || o.LastCommit.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastCommit.Get()
}

// GetLastCommitOk returns a tuple with the LastCommit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetLastCommitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastCommit.Get(), o.LastCommit.IsSet()
}

// SetLastCommit sets field value
func (o *DumpConsensusResponseResultPeerStateRoundState) SetLastCommit(v string) {
	o.LastCommit.Set(&v)
}

// GetCatchupCommitRound returns the CatchupCommitRound field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetCatchupCommitRound() int32 {
	if o == nil || o.CatchupCommitRound.Get() == nil {
		var ret int32
		return ret
	}

	return *o.CatchupCommitRound.Get()
}

// GetCatchupCommitRoundOk returns a tuple with the CatchupCommitRound field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetCatchupCommitRoundOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CatchupCommitRound.Get(), o.CatchupCommitRound.IsSet()
}

// SetCatchupCommitRound sets field value
func (o *DumpConsensusResponseResultPeerStateRoundState) SetCatchupCommitRound(v int32) {
	o.CatchupCommitRound.Set(&v)
}

// GetCatchupCommit returns the CatchupCommit field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetCatchupCommit() string {
	if o == nil || o.CatchupCommit.Get() == nil {
		var ret string
		return ret
	}

	return *o.CatchupCommit.Get()
}

// GetCatchupCommitOk returns a tuple with the CatchupCommit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DumpConsensusResponseResultPeerStateRoundState) GetCatchupCommitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CatchupCommit.Get(), o.CatchupCommit.IsSet()
}

// SetCatchupCommit sets field value
func (o *DumpConsensusResponseResultPeerStateRoundState) SetCatchupCommit(v string) {
	o.CatchupCommit.Set(&v)
}

func (o DumpConsensusResponseResultPeerStateRoundState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["height"] = o.Height
	}
	if true {
		toSerialize["round"] = o.Round
	}
	if true {
		toSerialize["step"] = o.Step
	}
	if true {
		toSerialize["start_time"] = o.StartTime
	}
	if true {
		toSerialize["proposal"] = o.Proposal
	}
	if true {
		toSerialize["proposal_block_parts_header"] = o.ProposalBlockPartsHeader
	}
	if true {
		toSerialize["proposal_pol_round"] = o.ProposalPolRound.Get()
	}
	if true {
		toSerialize["proposal_pol"] = o.ProposalPol.Get()
	}
	if true {
		toSerialize["prevotes"] = o.Prevotes.Get()
	}
	if true {
		toSerialize["precommits"] = o.Precommits.Get()
	}
	if true {
		toSerialize["last_commit_round"] = o.LastCommitRound.Get()
	}
	if true {
		toSerialize["last_commit"] = o.LastCommit.Get()
	}
	if true {
		toSerialize["catchup_commit_round"] = o.CatchupCommitRound.Get()
	}
	if true {
		toSerialize["catchup_commit"] = o.CatchupCommit.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDumpConsensusResponseResultPeerStateRoundState struct {
	value *DumpConsensusResponseResultPeerStateRoundState
	isSet bool
}

func (v NullableDumpConsensusResponseResultPeerStateRoundState) Get() *DumpConsensusResponseResultPeerStateRoundState {
	return v.value
}

func (v *NullableDumpConsensusResponseResultPeerStateRoundState) Set(val *DumpConsensusResponseResultPeerStateRoundState) {
	v.value = val
	v.isSet = true
}

func (v NullableDumpConsensusResponseResultPeerStateRoundState) IsSet() bool {
	return v.isSet
}

func (v *NullableDumpConsensusResponseResultPeerStateRoundState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDumpConsensusResponseResultPeerStateRoundState(val *DumpConsensusResponseResultPeerStateRoundState) *NullableDumpConsensusResponseResultPeerStateRoundState {
	return &NullableDumpConsensusResponseResultPeerStateRoundState{value: val, isSet: true}
}

func (v NullableDumpConsensusResponseResultPeerStateRoundState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDumpConsensusResponseResultPeerStateRoundState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

