module quests

go 1.21

toolchain go1.21.4

require (
	github.com/algorand/go-algorand-sdk/v2 v2.6.0
	golang.org/x/exp v0.0.0-20240604190554-fc45aab8b7f8
	lib v0.0.0
)

require (
	github.com/algorand/go-codec/codec v1.1.10 // indirect
	github.com/fatih/color v1.17.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/sys v0.18.0 // indirect
)

replace lib v0.0.0 => ../../../lib
