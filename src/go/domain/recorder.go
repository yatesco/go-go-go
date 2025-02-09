package domain

type SimpleAllocationRecorder struct {
	Repository AllocationRepository
}

func (v SimpleAllocationRecorder) RecordAllocations(accountId *AccountId, allocations *PendingAllocations) bool {
	v.Repository.RecordPendingAllocations(accountId, allocations)
	return true
}

var _ AllocationRecorder = SimpleAllocationRecorder{}
var _ AllocationRecorder = (*SimpleAllocationRecorder)(nil)
