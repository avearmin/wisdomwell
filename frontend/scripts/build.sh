#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Create or clean the dist directory
echo "Creating dist directory..."
rm -rf dist
mkdir dist

# Compile Elm code
echo "Compiling Elm to dist/elm.js..."
elm make src/Main.elm --output=dist/elm.js

# Copy static files (index.html and config.json) to dist
echo "Copying index.html and config.json to dist..."
cp src/index.html dist/
cp src/config.json dist/

# Done
echo "Build completed. Your app is ready to be served from the dist/ directory."
