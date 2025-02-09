package infrastructure

import (
	"mydomain.com/domain"
)

type InMemoryRepository struct {
	accountIdToLatestPendingAllocations map[*domain.AccountId]*domain.PendingAllocations
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		accountIdToLatestPendingAllocations: make(map[*domain.AccountId]*domain.PendingAllocations),
	}
}

var _ domain.AllocationRepository = InMemoryRepository{}
var _ domain.AllocationRepository = (*InMemoryRepository)(nil) // Verify that *T implements I.

// assumption is that this will never fail except for infrastructure errors
func (repo InMemoryRepository) ListPendingAllocations(accountId *domain.AccountId) *domain.PendingAllocations {
	return repo.accountIdToLatestPendingAllocations[accountId]
}

// assumption is that this will never fail except for infrastructure errors
func (repo InMemoryRepository) RecordPendingAllocations(accountId *domain.AccountId, allocations *domain.PendingAllocations) {
	repo.accountIdToLatestPendingAllocations[accountId] = allocations
}
