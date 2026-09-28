GO = go
OUTPUT = pokedexcli

build:
	@$(GO) build -o $(OUTPUT) && ./$(OUTPUT)

test:
	@$(GO) test ./...

run:
	bootdev run

submit:
	bootdev submit
