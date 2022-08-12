package inventorysrv

import (
	"sync"

	muxrouter "github.com/osalomon89/go-bookstore/inventory/infrastructure/muxrouter"
	"github.com/osalomon89/go-bookstore/inventory/infrastructure/mysqldb"
	http "github.com/osalomon89/go-bookstore/inventory/interfaces/controllers/http"
	repository "github.com/osalomon89/go-bookstore/inventory/interfaces/repository"
	usecases "github.com/osalomon89/go-bookstore/inventory/usecases"
)

var (
	httpMuxRouter muxrouter.MuxRouter = *muxrouter.NewMuxRouter()
)

func InventoryService(wg *sync.WaitGroup) {
	defer wg.Done()

	inventoryControllers := getInventoryControllers()

	httpMuxRouter.POST("/inventory/add", inventoryControllers.Add)
	httpMuxRouter.GET("/inventory/all", inventoryControllers.GetAll)
	httpMuxRouter.SERVE(":8888")
}

func getInventoryControllers() http.HTTPInteractor {
	mySqlDBRepository, _ := mysqldb.NewMySqlDB()

	inventoryRepository := repository.NewRepositoryInteractor(mySqlDBRepository)
	inventoryUseCases := usecases.NewUseCasesInteractor(inventoryRepository)
	inventoryControllers, _ := http.NewHTTPInteractor(inventoryUseCases)

	return inventoryControllers
}
