# Antybrowser CLI

[![GitHub Release](https://img.shields.io/github/v/release/antybrowser/antybrowser-cli?style=flat-square&color=blue)](https://github.com/antybrowser/antybrowser-cli/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Go Report](https://goreportcard.com/badge/github.com/antybrowser/antybrowser-cli?style=flat-square)](https://goreportcard.com/report/github.com/antybrowser/antybrowser-cli)

Command-line interface for [Antybrowser](https://antybrowser.com) — the anti-detect browser for multi-account management, affiliate marketing, web scraping, and privacy.

Manage browser profiles, proxies, extensions, and automations directly from your terminal. Connects to the Antybrowser desktop app's local API server.

## Features

- **Profile Management** — Create, list, start, stop, duplicate, and delete browser profiles with unique fingerprints
- **Proxy Support** — Manage HTTP, SOCKS4, and SOCKS5 proxies with bulk checking
- **Extension Management** — List and manage browser extensions per profile
- **Group Organization** — Organize profiles into groups with custom colors
- **Automation** — Run browser automations from the command line
- **JSON Output** — Machine-readable output for scripting and CI/CD pipelines
- **Persistent Config** — Save API URL and key to disk, override with flags or environment variables
- **Cross-Platform** — Works on macOS (ARM/Intel), Linux (ARM/x64), and Windows (ARM/x64)

## Installation

### Windows

```powershell
# WinGet (recommended)
winget install Antybrowser.AntybrowserCLI

# Chocolatey
choco install antybrowser-cli

# Scoop
scoop bucket add antybrowser https://github.com/antybrowser/scoop-bucket
scoop install antybrowser-cli
```

### macOS

```bash
# Homebrew
brew install antybrowser-cli
```

### Linux

```bash
# Homebrew (Linux)
brew install antybrowser-cli

# AUR (Arch Linux)
yay -S antybrowser-cli-bin

# Nix
nix-env -iA nixpkgs.antybrowser-cli
```

### Go

```bash
go install github.com/antybrowser/antybrowser-cli/cmd/antybrowser@latest
```

### Docker

```bash
docker pull antybrowser/cli:latest
docker run --rm antybrowser/cli antybrowser --help
```

### Manual Download

Download the latest binary from [GitHub Releases](https://github.com/antybrowser/antybrowser-cli/releases/latest):

| Platform | Archive |
|----------|---------|
| macOS (ARM) | `antybrowser-darwin-arm64.tar.gz` |
| macOS (Intel) | `antybrowser-darwin-amd64.tar.gz` |
| Linux (ARM) | `antybrowser-linux-arm64.tar.gz` |
| Linux (x64) | `antybrowser-linux-amd64.tar.gz` |
| Windows (ARM) | `antybrowser-windows-arm64.zip` |
| Windows (x64) | `antybrowser-windows-amd64.zip` |

## Quick Start

```bash
# Check connection to Antybrowser desktop app
antybrowser status

# Configure custom API URL (if not running on default port)
antybrowser config set --api-url http://192.168.1.100:5173

# Set API key for authenticated access
antybrowser config set --api-key your-api-key

# Create a profile
antybrowser profiles create --name "My Profile" --os windows --browser Chrome

# List all profiles
antybrowser profiles list

# Start a profile (launches browser)
antybrowser profiles start 123

# Stop a running profile
antybrowser profiles stop 123

# Duplicate a profile
antybrowser profiles duplicate 123 --name "Profile Copy"

# List proxies
antybrowser proxies list

# Add a proxy
antybrowser proxies create --name "US Proxy" --host proxy.example.com --port 1080 --type socks5

# Check a proxy
antybrowser proxies check --host proxy.example.com --port 1080 --type socks5

# List extensions
antybrowser extensions list

# List groups
antybrowser groups list

# Run an automation
antybrowser automations run 1 --profile-id 123

# JSON output for scripting
antybrowser --output json profiles list
```

## Commands

| Command | Description |
|---------|-------------|
| `antybrowser status` | Check connection to Antybrowser |
| `antybrowser profiles list` | List all browser profiles |
| `antybrowser profiles create` | Create a new profile |
| `antybrowser profiles get <id>` | Get profile details |
| `antybrowser profiles update <id>` | Update a profile |
| `antybrowser profiles delete <id>` | Delete a profile |
| `antybrowser profiles start <id>` | Start a browser profile |
| `antybrowser profiles stop <id>` | Stop a running profile |
| `antybrowser profiles duplicate <id>` | Duplicate a profile |
| `antybrowser proxies list` | List all proxies |
| `antybrowser proxies create` | Create a new proxy |
| `antybrowser proxies check` | Check proxy connectivity |
| `antybrowser proxies delete <id>` | Delete a proxy |
| `antybrowser extensions list` | List all extensions |
| `antybrowser extensions delete <id>` | Delete an extension |
| `antybrowser groups list` | List all groups |
| `antybrowser groups create` | Create a new group |
| `antybrowser groups update <id>` | Update a group |
| `antybrowser groups delete <id>` | Delete a group |
| `antybrowser automations list` | List all automations |
| `antybrowser automations run <id>` | Run an automation |
| `antybrowser config show` | Show current configuration |
| `antybrowser config set` | Set and save configuration |
| `antybrowser config reset` | Reset configuration to defaults |
| `antybrowser version` | Print version info |

## Configuration

The CLI stores configuration at:
- **Windows**: `%APPDATA%\Antybrowser\CLI\config.json`
- **macOS**: `~/.config/antybrowser/config.json`
- **Linux**: `~/.config/antybrowser/config.json`

Priority order:
1. Command-line flags (`--api-url`, `--api-key`)
2. Environment variables (`ANTYBROWSER_API_URL`, `ANTYBROWSER_API_KEY`)
3. Config file (`antybrowser config set`)

## Requirements

The Antybrowser desktop app must be running for the CLI to work. The CLI connects to the local API server on `http://127.0.0.1:5173` by default.

## SDK Libraries

Official SDKs are available for 7 languages:

| Language | Package | Repository |
|----------|---------|------------|
| TypeScript | `@antybrowser/sdk` | [SDK](https://github.com/antybrowser/SDK) |
| Python | `antybrowser` | [SDK](https://github.com/antybrowser/SDK) |
| C# | `Antybrowser.SDK` | [SDK](https://github.com/antybrowser/SDK) |
| Go | `github.com/antybrowser/sdk-go` | [SDK](https://github.com/antybrowser/SDK) |
| PHP | `antybrowser/sdk` | [SDK](https://github.com/antybrowser/SDK) |
| Ruby | `antybrowser` | [SDK](https://github.com/antybrowser/SDK) |
| Java | `com.antybrowser.sdk` | [SDK](https://github.com/antybrowser/SDK) |

## Links

- [Website](https://antybrowser.com)
- [Documentation](https://docs.antybrowser.com)
- [Desktop App](https://github.com/antybrowser/antybrowser)
- [SDK](https://github.com/antybrowser/SDK)
- [GitHub Releases](https://github.com/antybrowser/antybrowser-cli/releases)
- [Report Issues](https://github.com/antybrowser/antybrowser-cli/issues)

## License

[MIT](LICENSE)
