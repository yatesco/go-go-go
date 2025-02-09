package domain

import (
	"math/rand"
	"time"
)

type StubFundAllocatorValidator struct {
	Passes bool
}

func (v StubFundAllocatorValidator) IsWithinThreshold(accountId *AccountId, allocations *PendingAllocations) bool {
	return v.Passes
}

var _ FundAllocatorValidator = StubFundAllocatorValidator{}
var _ FundAllocatorValidator = (*StubFundAllocatorValidator)(nil)

func RandomValidAccountAndAllocations() (accountId AccountId, pendingAllocations PendingAllocations) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// TODO - use domain services, but I don't know how to do that in go....
	var allocations []PendingAllocation
	amount, err := NewHackyAmount(uint(r.Uint32()), Pound)
	if err != nil {
		panic("logic error - failed to create amount")
	}

	allocation := PendingAllocation{
		Id:     FundId(r.Uint32()),
		Amount: *amount,
	}
	allocations = append(allocations, allocation)
	validator := StubFundAllocatorValidator{
		Passes: true,
	}

	result, err := NewAllocations(validator, &accountId, allocations)
	if err != nil {
		panic("logic error - failed to create allocationst")
	}

	return accountId, *result
}
