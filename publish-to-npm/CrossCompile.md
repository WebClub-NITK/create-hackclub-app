Following commands to cross-compile the Go CLI tool (Execute this in the root directory and paste the built files in publish-to-npm/bin folder):


# For Linux
GOOS=linux GOARCH=amd64 go build -o cli-tool-linux main.go

# For macOS
GOOS=darwin GOARCH=amd64 go build -o cli-tool-macos main.go

# For Windows
GOOS=windows GOARCH=amd64 go build -o cli-tool.exe 