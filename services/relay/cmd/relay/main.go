package main

import (
	"fmt"
	"lib/utils"
	"net/http"
	"os"
)

func main() {
	logger := utils.NewLogger()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("This is GolinuxCloud's websocket server")
	})

	logger.Debug(fmt.Sprintf("listening on port :%s", os.Getenv("RELAY_PORT")))

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("RELAY_PORT")), nil)
	if err != nil {
		logger.Error(err)
	}
}
