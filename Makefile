

.PHONY: run
run: 
	go run cmd/aereq/main.go


.PHONY: t
t: 
	go clean -testcache
	go test ./... -count=2 -cover -race