.PHONY: recipefiles
recipefiles:
	ifeq (, $(shell which jq))
		$(error "jq is not installed. Make sure to installed it with 'apt-get install jq' or 'brew install jq'")
	else
		jq -s . src/**/recipefile.json > src/app/recipefiles.json
	endif
