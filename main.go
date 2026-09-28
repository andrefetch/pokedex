package main

func main() {
	cfg := &config{commands: getCommands()}
	startRepl(cfg)
}
