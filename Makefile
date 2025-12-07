.PHONY: clean
clean:
	@rm -rf ./build


.PHONY: build
build:
	@go build \
		-ldflags="-w -s" \
		-trimpath \
		-o ./build/api \
		./cmd


.PHONY: install
install:
	@go mod tidy
	@go install github.com/air-verse/air@latest


.PHONY: run
run:
	@if command -v air; then\
		air && rm -rf ./build/api;\
	fi
