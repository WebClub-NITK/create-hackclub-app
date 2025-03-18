Following commands to cross-compile your Go CLI tool:

bash
Copy
Edit
# For Linux
GOOS=linux GOARCH=amd64 go build -o cli-tool-linux main.go

# For macOS
GOOS=darwin GOARCH=amd64 go build -o cli-tool-macos main.go

# For Windows
GOOS=windows GOARCH=amd64 go build -o cli-tool.exe 