package main

import (
	"lib/utils"
	"net/http"
)

func main() {
	logger := utils.NewLoggerWithLevel(utils.LogLevelDebug)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("This is GolinuxCloud's websocket server")
	})

	logger.Debug("listening on ports :3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		logger.Error(err)
	}
}
