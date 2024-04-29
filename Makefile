# Bundles all recipe files into a single one for the website
.PHONY: recipedb
recipedb:
	@command -v jq >/dev/null 2>&1 || { echo >&2 "jq is not installed. Please install jq to continue."; exit 1; }
	@find src -name 'recipefile.json' -exec jq -s -c . {} + > ./src/site/src/lib/store/data.json

# Prod build of the website
.PHONY: site
site:
	npm --prefix src/site run build
	cd src/site && zip -r build.zip build
