# Makefile

default: help

.PHONY: help
help: ## Show this help.
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

.PHONY: gen
gen: ## Generate the package.
	rm goswucards.go
	rm -rf images
	mkdir -p images
	cd internal && go run main.go