GO = go
OUTPUT = pokedexcli

build:
	@$(GO) build -o $(OUTPUT) && ./$(OUTPUT)

test:
	@$(GO) test ./...
