#!/bin/sh

get_snippet(){
    snippet=$(sed -n "/$1/,/$2/p" "$3") 
    snippet=$(echo "$snippet" | grep -v "$1")
    snippet=$(echo "$snippet" | grep -v "$2")
    echo "$snippet" # to trim the lines | awk '{$1=$1};1'
}

echo "/*Recipe OpenTelemetry*/" > my_recipe.js
echo "/*Dependencies:*/" >> my_recipe.js
get_snippet otel_dependencies_start otel_dependencies_end tracing.js >> my_recipe.js

echo "/*Instrumentation:*/" >> my_recipe.js
get_snippet otel_instrumentation_start otel_instrumentation_end tracing.js >> my_recipe.js

echo "" >> my_recipe.js
echo "" >> my_recipe.js
echo "/*How to create a manual span*/" >> my_recipe.js
echo "/*Dependencies:*/" >> my_recipe.js
get_snippet span_dependencies_start span_dependencies_end app.js >> my_recipe.js

echo "" >> my_recipe.js
echo "/*Creation:*/" >> my_recipe.js
get_snippet span_creation_start span_creation_end app.js >> my_recipe.js