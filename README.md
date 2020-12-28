Todo: Implement

# How to install:

## Prerequisites
- Go 1.13+
- Node / NPM

On mac, make sure your xcode-select is up to date:
- xcode-select --install

## Installing dependencies
- run go get -u -d ./... in the root of the project to ensure latest dependencies
- go get -u github.com/wailsapp/wails/cmd/wails

# Build
- In the root directory of the project run `wails build`
- Run the executable in the build folder

## Regenerating protos:
TODO: write

After generating proto files, in js generated code:

/* eslint-disable */ before generated code comment
