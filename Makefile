build:
		go build -o dezge cmd/dezge/main.go
run:
		go run cmd/dezge/main.go -c cmd/dezge/config.toml
client:
		go run cmd/dezgectl/main.go -c cmd/dezge/config.toml
