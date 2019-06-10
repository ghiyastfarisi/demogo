package init

import (
	"log"

	"github.com/ghiyastfarisi/demogo/configs"
)

// Init ...
func Init() {
	log.Println("[LOG] running init.go > Init()")
	// initiate config
	configs.Init()
}
