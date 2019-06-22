# Simple logout/shutdown/reboot screen for linux

Exactly what the title says. Written in Go and html/js/css.

# Instructions

0. Download Go from [here](https://golang.org/)
1. Run `go get github.com/zserge/webview`
2. Clone this repo and run `go build`
3. Fill out the logout/shutdown/reboot scripts with code that works for your setup
4. Run executable with the path to the cloned dir as the first argument and optionally the port to run it on as the second argument(defaults to port 3141), e.g. `homescreen /path/to/dir/with/scripts/and/images port`
