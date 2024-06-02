# Makefile to run Go program
#
# Variables
MAIN=main.go

# Targets
.PHONY: run

# Run the application
run:
	go run $(MAIN)

test: 
	go test ./rpc/rpc_test.go
