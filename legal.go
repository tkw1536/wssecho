package main

import _ "embed"

//go:generate gogenlicense -p main -n LegalNotices -d legal_notices.go github.com/tkw1536/wssecho

//go:embed LICENSE
var License string

// LegalText returns legal text to be included in human-readable output using huelio.
func LegalText() string {
	return `
================================================================================
wssecho - Quick web socket echo server
================================================================================
` + License + "\n" + LegalNotices
}
