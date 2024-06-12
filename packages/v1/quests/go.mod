module quests

go 1.20

require (
	github.com/algorand/go-algorand-sdk v1.24.0
	golang.org/x/exp v0.0.0-20240604190554-fc45aab8b7f8
	lib v0.0.0
)

require (
	github.com/algorand/go-codec/codec v1.1.10 // indirect
	github.com/fatih/color v1.17.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace lib v0.0.0 => ../../../lib
