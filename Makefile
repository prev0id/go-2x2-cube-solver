BATCH_SIZE=100
BATCHES=1
SCRAMBLE_LENGTH=20

.PHONY: run
run:
	go run cmd/main.go

.PHONY: test
test:
	go test -v -parallel $(BATCHES) ./pkg/solver -args -batches=$(BATCHES) -batch_size=$(BATCH_SIZE) -scramble_length=$(SCRAMBLE_LENGTH)

.PHONY: benchmark
benchmark:
	go test -bench=. ./pkg/solver -benchmem -run notest  -args -scramble_length=$(SCRAMBLE_LENGTH)