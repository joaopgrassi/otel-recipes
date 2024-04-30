.PHONY: recipedb
recipedb:
	@command -v jq >/dev/null 2>&1 || { echo >&2 "jq is not installed. Please install jq to continue."; exit 1; }
	@find src -name 'recipefile.json' -exec jq -s -c . {} + > ./src/site/src/lib/store/data.json

.PHONY: tidy-modules
tidy-modules:
	@find . -type d \( -name build -prune \) -o -name go.mod -print | while read -r gomod_path; do \
		dir_path=$$(dirname "$$gomod_path"); \
		echo "Executing 'go mod tidy' in directory: $$dir_path"; \
		(cd "$$dir_path"  && GOPROXY=$(GOPROXY) go get -u ./... && GOPROXY=$(GOPROXY) go mod tidy) || exit 1; \
	done
