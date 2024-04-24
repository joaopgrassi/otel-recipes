JQ_VERSION := $(shell jq --version 2>/dev/null)

.PHONY: recipedb
recipedb:
	@command -v jq >/dev/null 2>&1 || { echo >&2 "jq is not installed. Please install jq to continue."; exit 1; }
	@find src -name 'recipefile.json' -exec jq -s -c . {} + > ./src/site/src/lib/store/data.json

