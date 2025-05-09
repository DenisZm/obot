## Initialize Go Module
Initialize a new Go module for the project.
```bash
go mod init github.com/deniszm/obot
```

## Install Cobra CLI
Install the Cobra CLI tool for generating and managing commands in a Go application.
```bash
go install github.com/spf13/cobra-cli@latest
```

## Scaffold Application
Create the initial structure of the application using Cobra CLI.
```bash
cobra-cli init
```

## Generate "Version" Command Code
Add a "version" command to the application using Cobra CLI.
```bash
cobra-cli add version
```

## Build Application with Version Variable
Build the application and embed the version information as a variable. The `-ldflags` option is used to pass linker flags to the Go compiler. In this case, the `-X` flag sets the value of the `appVersion` variable in the `github.com/deniszm/obot/cmd` package to `v1.0.0`.
```bash
go build -ldflags "-X=github.com/deniszm/obot/cmd.appVersion=v1.0.0"
```

## Format Code
Format the Go code in the project to ensure consistency and readability.
```bash
gofmt -s -w ./
```

## Download Dependencies
Download and install the dependencies required for the project.
```bash
go get
```

## Securely Set Environment Variable
Securely set the `TELE_TOKEN` environment variable by reading it as a hidden input and exporting it for use in the application.
```bash
read -s TELE_TOKEN
export TELE_TOKEN
```