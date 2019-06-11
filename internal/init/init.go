package init

import (
	"log"

	internalConfig "github.com/ghiyastfarisi/demogo/internal/config"
)

// Init ...
func Init() {
	log.Println("[LOG] Running pkg internal > init.go > Init()")
	// initiate config file
	internalConfig.Init()
}
