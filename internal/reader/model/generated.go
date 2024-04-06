// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type ConvenientFilter struct {
	Field *string             `json:"field,omitempty"`
	Eq    *string             `json:"eq,omitempty"`
	Ne    *string             `json:"ne,omitempty"`
	Gt    *string             `json:"gt,omitempty"`
	Gte   *string             `json:"gte,omitempty"`
	Lt    *string             `json:"lt,omitempty"`
	Lte   *string             `json:"lte,omitempty"`
	In    []*string           `json:"in,omitempty"`
	Nin   []*string           `json:"nin,omitempty"`
	And   []*ConvenientFilter `json:"and,omitempty"`
	Or    []*ConvenientFilter `json:"or,omitempty"`
}

// Filter object to restrict results depending on input properties
type InputFilter struct {
	// Filter only inputs with index lower than a given value
	IndexLowerThan *int `json:"indexLowerThan,omitempty"`
	// Filter only inputs with index greater than a given value
	IndexGreaterThan *int `json:"indexGreaterThan,omitempty"`
}

// Validity proof for an output
type OutputValidityProof struct {
	// Local input index within the context of the related epoch
	InputIndexWithinEpoch int `json:"inputIndexWithinEpoch"`
	// Output index within the context of the input that produced it
	OutputIndexWithinInput int `json:"outputIndexWithinInput"`
	// Merkle root of all output hashes of the related input, given in Ethereum hex binary format (32 bytes), starting with '0x'
	OutputHashesRootHash string `json:"outputHashesRootHash"`
	// Merkle root of all voucher hashes of the related epoch, given in Ethereum hex binary format (32 bytes), starting with '0x'
	VouchersEpochRootHash string `json:"vouchersEpochRootHash"`
	// Merkle root of all notice hashes of the related epoch, given in Ethereum hex binary format (32 bytes), starting with '0x'
	NoticesEpochRootHash string `json:"noticesEpochRootHash"`
	// Hash of the machine state claimed for the related epoch, given in Ethereum hex binary format (32 bytes), starting with '0x'
	MachineStateHash string `json:"machineStateHash"`
	// Proof that this output hash is in the output-hashes merkle tree. This array of siblings is bottom-up ordered (from the leaf to the root). Each hash is given in Ethereum hex binary format (32 bytes), starting with '0x'.
	OutputHashInOutputHashesSiblings []string `json:"outputHashInOutputHashesSiblings"`
	// Proof that this output-hashes root hash is in epoch's output merkle tree. This array of siblings is bottom-up ordered (from the leaf to the root). Each hash is given in Ethereum hex binary format (32 bytes), starting with '0x'.
	OutputHashesInEpochSiblings []string `json:"outputHashesInEpochSiblings"`
}

// Page metadata for the cursor-based Connection pagination pattern
type PageInfo struct {
	// Cursor pointing to the first entry of the page
	StartCursor *string `json:"startCursor,omitempty"`
	// Cursor pointing to the last entry of the page
	EndCursor *string `json:"endCursor,omitempty"`
	// Indicates if there are additional entries after the end curs
	HasNextPage bool `json:"hasNextPage"`
	// Indicates if there are additional entries before the start curs
	HasPreviousPage bool `json:"hasPreviousPage"`
}

// Data that can be used as proof to validate notices and execute vouchers on the base layer blockchain
type Proof struct {
	// Validity proof for an output
	Validity *OutputValidityProof `json:"validity"`
	// Data that allows the validity proof to be contextualized within submitted claims, given as a payload in Ethereum hex binary format, starting with '0x'
	Context string `json:"context"`
}

type CompletionStatus string

const (
	CompletionStatusUnprocessed                CompletionStatus = "UNPROCESSED"
	CompletionStatusAccepted                   CompletionStatus = "ACCEPTED"
	CompletionStatusRejected                   CompletionStatus = "REJECTED"
	CompletionStatusException                  CompletionStatus = "EXCEPTION"
	CompletionStatusMachineHalted              CompletionStatus = "MACHINE_HALTED"
	CompletionStatusCycleLimitExceeded         CompletionStatus = "CYCLE_LIMIT_EXCEEDED"
	CompletionStatusTimeLimitExceeded          CompletionStatus = "TIME_LIMIT_EXCEEDED"
	CompletionStatusPayloadLengthLimitExceeded CompletionStatus = "PAYLOAD_LENGTH_LIMIT_EXCEEDED"
)

var AllCompletionStatus = []CompletionStatus{
	CompletionStatusUnprocessed,
	CompletionStatusAccepted,
	CompletionStatusRejected,
	CompletionStatusException,
	CompletionStatusMachineHalted,
	CompletionStatusCycleLimitExceeded,
	CompletionStatusTimeLimitExceeded,
	CompletionStatusPayloadLengthLimitExceeded,
}

func (e CompletionStatus) IsValid() bool {
	switch e {
	case CompletionStatusUnprocessed, CompletionStatusAccepted, CompletionStatusRejected, CompletionStatusException, CompletionStatusMachineHalted, CompletionStatusCycleLimitExceeded, CompletionStatusTimeLimitExceeded, CompletionStatusPayloadLengthLimitExceeded:
		return true
	}
	return false
}

func (e CompletionStatus) String() string {
	return string(e)
}

func (e *CompletionStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CompletionStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CompletionStatus", str)
	}
	return nil
}

func (e CompletionStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
