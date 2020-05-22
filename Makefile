
all: attributes.go

attributes.go: attributes.txt attributes_generate.go
	go run -tags generate attributes_generate.go attributes.txt > attributes.go
