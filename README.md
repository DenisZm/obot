# obot - DevOps Helper Telegram Bot

[![License](https://img.shields.io/github/license/deniszm/obot)](./LICENSE)

## Description

`obot` is a Telegram bot designed to assist DevOps engineers by providing CIDR calculator functionality. The bot processes user input in the format of an IP address with CIDR notation (e.g., `192.168.0.1/24`) and responds with subnet information, mask, and IP address range.

## Features

- Calculate subnet information from CIDR notation
- Determine subnet mask in dotted decimal format
- Calculate IP address range
- Count number of usable hosts
- Handle special cases (`/31` and `/32` subnets)

## Installation and Setup

### Prerequisites

- Go (latest stable version)
- Telegram bot token from @BotFather

### Installation

1. Clone the repository:
```bash
git clone https://github.com/deniszm/obot.git
cd obot
```

2. Install dependencies:
```bash
go get github.com/spf13/cobra
go get gopkg.in/telebot.v4
```

3. Build the application:
```bash
go build -ldflags "-X=github.com/deniszm/obot/cmd.appVersion=v1.1.0"
```

### Running

1. Export your Telegram bot token:
```bash
export TELE_TOKEN="your_bot_token"
```

2. Run the bot:
```bash
./obot start
```

## Usage

After starting the bot, you can interact with it on Telegram using the following commands:

- `/start` - Display welcome message and instructions
- `/subnet <IP>/<CIDR>` - Calculate subnet information

Example:
```
/subnet 192.168.0.1/24
```

## Docker image build for different platforms

To build the Docker image for ARM64 platform, use the following command:

```sh
docker buildx build \
  --platform linux/arm64 \
  --tag ghcr.io/deniszm/obot:1.0.0-arm \
  --build-arg VERSION=1.0.0 \
  --load \
  .
```

You can change the `--platform` and `--tag` values to build for other platforms and versions.

## Contact

- Telegram Bot: [t.me/deniszm_obot](https://t.me/deniszm_obot)
- GitHub: [github.com/deniszm/obot](https://github.com/deniszm/obot)

## License

Licensed under Apache License 2.0. See [LICENSE](./LICENSE) file for more information.
