# Pokedex 

An interactive command-line Pokedex built in Go that uses HTTP requests to fetch Pokemon & location data from PokeAPI. You can explore locations, catch pokemon, and grow your pokedex! Uses an in-memory caching system to reduce redundant API requests to improve lower latency and faster command execution.

## Installation

Requires Go, install @ [go.dev](https://go.dev/doc/install)

```sh
go install github.com/andrefetch/pokedex@latest
pokedex
```

## Usage & Quick Start
- Quite simple to get start, run `pokedex` in your terminal
- You should get `pokedex > ` displayed in your terminal.
- Run `help` in that prompt to get a display of all commands!
- Use `map` and `explore` to find pokemon, catch them all!

## Roadmap
- [x] Multi-type pokeballs + inventory
- [ ] BubbleTea Integration + TUI Customization
- [ ] Simulate Battles + Parties
- [ ] Persist pokeballs + pokemon instead of initializing a map for data
- [ ] Allow pokemon to evolve
- [ ] More unit tests
