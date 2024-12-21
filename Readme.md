I am hoping u have installed go on your system , ensure by running 'go version'

# CLI Tool  
## Features  
- Accepts parameters like project name, description, and template type (web2, web3, or AI).  
- Interactive chat-like interface for input.  
- Uses `promptui` for a user-friendly selection experience.  


## Installation  
1. Install Cobra CLI:  
   ```bash  
   go install github.com/spf13/cobra-cli@latest  
   ```  

2. Install PromptUI:  
   ```bash  
   go get github.com/manifoldco/promptui  
   ```  

## Build  
To build the CLI tool, run:  
```bash  
go build -o cli-tool  
```  

## Usage  
To create a new project, execute:  
```bash  
./cli-tool create  
```  