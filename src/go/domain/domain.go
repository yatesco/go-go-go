package domain

import (
	"errors"
	"fmt"
)

type HackyCurrency int

const (
	Pound HackyCurrency = iota + 1
)

// I *refuse* to believe this isn't a solved problem already, but for now
// this will suffice....
// NOTE: `P`ublic members to keep JSON here. Prefer to move it out to DTOs but that is questionable.
type HackyAmount struct {
	// nothing bad has ever come from doing this...
	// e.g, for Pound this would be pence
	NumberOfSmallestWholeUnit uint          `json:"units"`
	Currency                  HackyCurrency `json:"currency"`
}

func NewHackyAmount(units uint, currency HackyCurrency) (amount *HackyAmount, err error) {
	if units <= 0 {
		// TODO - internationalise errors
		return nil, errors.New(fmt.Sprintf("units must be positive. You provided [%v]", units))
	}

	return &HackyAmount{
		NumberOfSmallestWholeUnit: units,
		Currency:                  currency,
	}, nil
}

type (
	AccountId int
	FundId    int
)

// NOTE: not needed yet
// type Status int

// const (
// 	Pending Status = iota + 1
// 	Applied
// 	// TODO: discriminated union to accept code and payload
// 	Failed
// )

type PendingAllocation = struct {
	Id     FundId
	Amount HackyAmount
}

type PendingAllocations = struct {
	// TODO - this is much more accurately modelled as a map[FundId]HackyAmount
	allocations []PendingAllocation
}

// TODO: test
func NewAllocations(validator FundAllocatorValidator, accountId *AccountId, inAllocations []PendingAllocation) (result *PendingAllocations, err error) {
	// sanity
	if validator == nil {
		return nil, errors.New("logic error - validator is nil!")
	}

	if inAllocations == nil {
		// lack of Option monad means dealing with nils. Sigh
		return nil, nil
	} else if len(inAllocations) == 0 {
		// TODO - internationalise errors
		return nil, errors.New(fmt.Sprintf("there must be at least one allocation"))
	}

	// assert allocations != nil
	result = &PendingAllocations{
		allocations: inAllocations,
	}

	// optimistic validation (i.e. the validity of this validation is transient)
	if !validator.IsWithinThreshold(accountId, result) {
		// TODO - internationalise errors
		return nil, errors.New(fmt.Sprintf("total allocation exceeds allowed threshold"))
	}

	return
}

// NOTE not needed yet
// type Allocation = struct {
// 	id     FundId
// 	amount HackyAmount
// 	status Status
// }

type FundAllocatorValidator interface {
	// optimisitic validator (i.e. _not_ authoritative as updates could be happening in parallel)
	IsWithinThreshold(accountId *AccountId, allocations *PendingAllocations) bool
}

type AllocationRecorder interface {
	RecordAllocations(accountId *AccountId, allocations *PendingAllocations) bool
}

type AllocationRepository interface {
	RecordPendingAllocations(accountId *AccountId, allocations *PendingAllocations)
	ListPendingAllocations(accountId *AccountId) *PendingAllocations
}
