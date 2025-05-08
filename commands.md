## Init go module
```bash
go mod init github.com/deniszm/obot
```

## Install cobra cli
```bash
go install github.com/spf13/cobra-cli@latest
```

## Scaffold application
```bash
cobra-cli init
```

## Genarate "version" command code
```
cobra-cli add version
```

## Build application with defined variable
```bash
go build -ldflags "-X=github.com/deniszm/obot/cmd.appVersion=v1.0.0"
```