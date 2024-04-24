.PHONY: recipedb
recipedb:
	find src -name 'recipefile.json' -exec jq -s -c . {} + > ./src/site/src/lib/store/data.json
