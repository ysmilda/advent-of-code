year=2025
day=0
session-token=

generate:
	go run ./cmd/generate -year=$(year) -day=$(day)

update:
	go run ./cmd/update

fetch:
	go run ./cmd/fetch -session-token=$(session-token)

all: generate update fetch

run: 
	go run . -year=$(year) -day=$(day)