# Technical Specification: DevOps Helper Telegram Bot (obot)

## 1. Overview
The **obot** is a Telegram bot designed to assist DevOps engineers by providing a CIDR calculator functionality. The bot processes user input in the format of an IP address with a CIDR notation (e.g., `192.168.0.1/24`) and responds with the subnet mask and the range of IP addresses for the specified subnet. The bot is implemented in Go, using the `github.com/spf13/cobra` library for command-line interface management and `gopkg.in/telebot.v4` for interaction with the Telegram Bot API.

## 2. Objectives
- Develop a Telegram bot that handles CIDR calculation requests.
- Provide a user-friendly interface for DevOps engineers to retrieve subnet information.
- Ensure minimal external dependencies, relying only on specified libraries.
- Practice Go programming, Telegram Bot API integration, and CLI application development.

## 3. Functional Requirements
### 3.1 CIDR Calculator
- **Command**: `/subnet <IP>/<CIDR>`
- **Input Format**: An IP address with CIDR notation (e.g., `192.168.0.1/24`).
- **Output**:
  - Subnet address in CIDR notation (e.g., `192.168.0.0/24`).
  - Subnet mask in dotted decimal format (e.g., `255.255.255.0`).
  - Range of usable IP addresses (e.g., `192.168.0.1 - 192.168.0.254` for `/24`).
  - Number of usable hosts (e.g., `254` for `/24`).
- **Error Handling**:
  - Validate input format; respond with an error message for invalid IP/CIDR (e.g., `Invalid IP/CIDR format`).
  - Handle edge cases, such as `/31` or `/32` subnets with limited or no usable hosts.

### 3.2 Start Command
- **Command**: `/start`
- **Output**: A welcome message with instructions on how to use the bot, including the `/subnet` command format.

## 4. Technical Requirements
### 4.1 Programming Language and Libraries
- **Language**: Go (latest stable version).
- **Libraries**:
  - `github.com/spf13/cobra`: For creating a CLI interface to start the bot.
  - `gopkg.in/telebot.v4`: For interacting with the Telegram Bot API.
  - Standard Go libraries (e.g., `net`, `os`, `strings`) for CIDR calculations and environment variable management.
- **Dependencies**: No additional external dependencies beyond the specified libraries.

### 4.2 Bot Configuration
- The bot retrieves its Telegram API token from an environment variable named `TELE_TOKEN`.
- If `TELE_TOKEN` is not set, the bot should exit with an error message.

### 4.3 CLI Interface
- The program is named **obot** and is executed via a single CLI command: `obot`.
- The `cobra` library is used to define the `obot` command, which initializes and starts the Telegram bot.

### 4.4 Telegram Bot Setup
- The bot is created using `telebot.NewBot()` with a `LongPoller` for message handling (timeout: 10 seconds).
- Handlers are registered for:
  - `/start`: Displays the welcome message.
  - `/subnet`: Processes CIDR calculation requests.

## 5. Implementation Details
### 5.1 File Structure
- **main.go**: Contains the main program logic, including CLI setup with `cobra`, bot initialization, and message handlers.
- **README.md**: Project documentation with setup instructions, bot link, and usage examples.

### 5.2 CIDR Calculation Logic
- Use the Go `net` package to parse CIDR input (`net.ParseCIDR`).
- Calculate:
  - Subnet mask using `ipNet.Mask`.
  - Address range by applying the mask to the IP and calculating the first and last usable IPs.
  - Number of hosts using the formula `2^(32 - CIDR) - 2` (subtracting network and broadcast addresses for subnets with CIDR ≤ 30).
- Handle special cases (e.g., `/31` and `/32` subnets).

### 5.3 Error Handling
- Validate CIDR input and return user-friendly error messages for invalid formats.
- Log errors (e.g., failure to create bot or connect to Telegram API) to the console for debugging.

## 6. Setup Instructions
1. **Install Go**: Ensure the latest stable version of Go is installed (https://golang.org/doc/install).
2. **Create Telegram Bot**:
   - Use `@BotFather` in Telegram to create a new bot.
   - Obtain the bot token and store it in the `TELE_TOKEN` environment variable.
3. **Install Dependencies**:
   ```bash
   go get github.com/spf13/cobra
   go get gopkg.in/telebot.v4
   ```
4. **Run the Bot**:
   ```bash
   export TELE_TOKEN="your_bot_token"
   go run main.go
   ```
5. **Test the Bot**:
   - Interact with the bot via Telegram using `/start` and `/subnet` commands.
   - Example: `/subnet 192.168.0.1/24`.

## 7. Deliverables
- **Source Code**:
  - `main.go`: Complete implementation of the bot.
  - `README.md`: Documentation with project description, setup instructions, bot link (e.g., `t.me/Obot`), and usage examples.
- **GitHub Repository**:
  - Create a public GitHub repository for the project.
  - Push all code and documentation to the repository.
  - Include a link to the repository as the submission.
- **Bot Link**: Provide a Telegram link to the bot (e.g., `t.me/Obot`).

## 8. Acceptance Criteria
- The bot successfully starts with the `obot` command and connects to Telegram using the provided `TELE_TOKEN`.
- The `/start` command displays a welcome message with usage instructions.
- The `/subnet` command correctly processes valid CIDR input (e.g., `192.168.0.1/24`) and returns:
  - Subnet address.
  - Subnet mask.
  - IP address range.
  - Number of usable hosts.
- Invalid CIDR input is handled with clear error messages.
- The code is well-documented, with comments explaining key functions.
- The GitHub repository contains all required files (`main.go`, `README.md`) and is publicly accessible.
- The `README.md` includes setup instructions, bot link, and example commands.

## 9. Constraints
- The bot must not rely on external APIs or services beyond the Telegram Bot API.
- The implementation must use only the specified libraries (`cobra`, `telebot.v4`) and standard Go packages.
- The bot should handle Telegram messages efficiently using `LongPoller`.

## 10. Future Enhancements (Out of Scope)
- Additional DevOps tools (e.g., storage unit conversion, DNS lookup).
- Persistent storage for user data or query history.
- Advanced CIDR calculations (e.g., splitting subnets).

## 11. References
- Telegram Bot API: https://core.telegram.org/bots/api
- Cobra Documentation: https://github.com/spf13/cobra
- Telebot.v4 Documentation: https://pkg.go.dev/gopkg.in/telebot.v4
- Example Repository: https://github.com/den-vasyliev/kbot