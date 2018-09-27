.PHONY: run
run:
	@docker run --rm --name goworker -p 8000:8000 goworker-img

.PHONY: build
build:
	@docker build -t goworker-img .
