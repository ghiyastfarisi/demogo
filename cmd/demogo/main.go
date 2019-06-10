package demogo

import (
	internalInit "github.com/ghiyastfarisi/demogo/internal/init"
)

// Main function entrypoint
func Main() {
	internalInit.Init()
	router()
}
