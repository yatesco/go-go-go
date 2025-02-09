// API enforcement test
//
// This ensures that any provided implementation behaves as an
// domain.AllocationRepository.
//
// Unfortunately, I don't know how to do this in GO!
package infrastructure

import (
	"testing"

	"mydomain.com/domain"
)

// TODO: figure out how to isolate these tests so they can be provided with
// the _implementation_ being tests. In other words, make `newSut` a strategy
// somehow so these tests can be called against the `InMemoryRepository`, the
// `DatabaseRepository` etc.
func newSut() domain.AllocationRepository {
	return NewInMemoryRepository()
}

func TestListAccountsWithNoPendingAllocationsReturnsEmpty(t *testing.T) {
	sut := newSut()

	accountId := domain.AccountId(10)
	actual := sut.ListPendingAllocations(&accountId)
	// expected := nil

	if actual != nil {
		// t.Fatalf(`incorrect allocations. Expected: %v but received %v`, expected, actual)
		t.Fatalf(`incorrect allocations. Expected: NIL but received %v`, actual)
	}
}

func TestRecordWithNoExistingAllocationsRemembersAllocations(t *testing.T) {
	sut := newSut()
	accountId, expected := domain.RandomValidAccountAndAllocations()

	sut.RecordPendingAllocations(&accountId, &expected)

	actual := sut.ListPendingAllocations(&accountId)

	if actual != &expected {
		t.Fatalf(`incorrect allocations. Expected: %v but received %v`, expected, actual)
	}
}

func TestRecordWithExistingAllocationsRemembersAllocationsWhenAllocationsAreNotNil(t *testing.T) {
	sut := newSut()

	// GIVEN
	accountId, existing := domain.RandomValidAccountAndAllocations()
	sut.RecordPendingAllocations(&accountId, &existing)
	_, expected := domain.RandomValidAccountAndAllocations()

	// WHEN
	sut.RecordPendingAllocations(&accountId, &expected)
	if &existing == &expected {
		t.Fatalf(`LOGIC ERROR - insufficient uniqueness in test data: %v`, expected)
	}

	// THEN
	actual := sut.ListPendingAllocations(&accountId)
	if actual != &expected {
		t.Fatalf(`incorrect allocations. Expected: %v but received %v`, expected, actual)
	}
}

func TestRecordWithExistingAllocationsRemembersAllocationsWhenAllocationsAreNil(t *testing.T) {
	sut := newSut()

	// GIVEN
	accountId, existing := domain.RandomValidAccountAndAllocations()
	sut.RecordPendingAllocations(&accountId, &existing)

	// WHEN
	sut.RecordPendingAllocations(&accountId, nil)

	// THEN
	actual := sut.ListPendingAllocations(&accountId)
	if actual != nil {
		t.Fatalf(`incorrect allocations. Expected: NIL but received %v`, actual)
	}
}
