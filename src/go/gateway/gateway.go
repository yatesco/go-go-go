package gateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mydomain.com/domain"
)

const HACK_ACCOUNT = domain.AccountId(10)

func Router(
	repository domain.AllocationRepository,
	recorder domain.AllocationRecorder,
) *gin.Engine {
	// bootstrap the system
	r := gin.Default()
	r.GET("/allocations", func(c *gin.Context) {
		getPendingAllocations(repository, c)
	})
	r.POST("/allocations", func(c *gin.Context) {
		recordPendingAllocations(recorder, c)
	})
	return r
}

// TODO - test this
func getPendingAllocations(repository domain.AllocationRepository, c *gin.Context) {
	// HACK TODO - where does the safe accountId come from?
	accountId := HACK_ACCOUNT
	allocations := repository.ListPendingAllocations(&accountId)
	c.IndentedJSON(http.StatusOK, allocations)
}

// TODO - test this
func recordPendingAllocations(recorder domain.AllocationRecorder, c *gin.Context) {
	// NOTE: decide whether JSON serialisation will handle ensure
	// domain objects are valid or whether the gateway should expose something
	// like a PotentialAllocation of type map[Currency, NonNullUInt]

	// HACK TODO - where does the safe accountId come from?
	accountId := HACK_ACCOUNT
	// HACK TODO - JSON serialisation into allocations
	_, allocations := domain.RandomValidAccountAndAllocations()

	recorder.RecordAllocations(&accountId, &allocations)
	c.IndentedJSON(http.StatusOK, nil)
}
