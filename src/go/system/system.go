package system

import (
	"mydomain.com/domain"
	"mydomain.com/gateway"
	"mydomain.com/infrastructure"
)

// Bootstrap the entire thing
func main() {
	// choose the appropriate infrastructure (event buses, repositories etc.)
	// pass them into the Gin handlers
	// start the polling scheduler/event listeners
	// profit!

	repository := domain.AllocationRepository(infrastructure.NewInMemoryRepository())
	recorder := domain.AllocationRecorder(domain.SimpleAllocationRecorder{
		Repository: repository,
	})
	router := gateway.Router(repository, recorder)
	router.Run()

}
