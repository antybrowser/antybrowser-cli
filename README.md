# Antybrowser CLI

[![GitHub Release](https://img.shields.io/github/v/release/antybrowser/antybrowser-cli?style=flat-square&color=blue)](https://github.com/antybrowser/antybrowser-cli/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)

Command-line interface for managing [Antybrowser](https://antybrowser.com) profiles, proxies, and automations.

## Features

- Create, list, start, and stop browser profiles
- Manage proxy configurations
- Run browser automations from the terminal
- Full API access to all Antybrowser features
- Cross-platform: macOS, Linux, and Windows

## Installation

### macOS (Homebrew)

```bash
brew tap antybrowser/tap https://github.com/antybrowser/homebrew-tap
brew install antybrowser
```

Or directly:

```bash
brew install --formula https://raw.githubusercontent.com/antybrowser/antybrowser-cli/main/homebrew/antybrowser.rb
```

### Linux

#### Debian / Ubuntu (.deb)

```bash
# Download the latest .deb from GitHub Releases, then:
sudo dpkg -i antybrowser_1.0.0_amd64.deb
sudo apt-get install -f
```

#### Snap

```bash
sudo snap install antybrowser
```

#### Flatpak

```bash
flatpak install flathub io.antybrowser.CLI
```

#### Docker

```bash
docker pull antybrowser/cli:latest
docker run --rm antybrowser/cli antybrowser --help
```

### Windows

#### WinGet

```powershell
winget install antybrowser.Antybrowser
```

#### Chocolatey

```powershell
choco install antybrowser
```

#### Scoop

```powershell
scoop bucket add antybrowser https://github.com/antybrowser/scoop-bucket
scoop install antybrowser
```

### Manual Installation

Download the latest binary for your platform from [GitHub Releases](https://github.com/antybrowser/antybrowser-cli/releases/latest):

| Platform       | Archive                                          |
|----------------|--------------------------------------------------|
| macOS (ARM)    | `antybrowser-darwin-arm64.tar.gz`               |
| macOS (Intel)  | `antybrowser-darwin-amd64.tar.gz`               |
| Linux (ARM)    | `antybrowser-linux-arm64.tar.gz`                |
| Linux (x64)    | `antybrowser-linux-amd64.tar.gz`                |
| Windows (ARM)  | `antybrowser-windows-arm64.zip`                 |
| Windows (x64)  | `antybrowser-windows-amd64.zip`                 |

Extract and add to your `PATH`.

### Go

```bash
go install github.com/antybrowser/antybrowser-cli@latest
```

### npm

```bash
npm install -g @antybrowser/cli
```

### pip

```bash
pip install antybrowser-cli
```

## Quick Start

```bash
# Check installation
antybrowser --version

# Start the local API server
antybrowser serve

# Create a profile
antybrowser profiles create --name "My Profile" --os windows

# List profiles
antybrowser profiles list

# Start a profile
antybrowser profiles start <profile-id>

# Configure a proxy
antybrowser profiles update <profile-id> --proxy "socks5://user:pass@host:port"
```

## Configuration

Antybrowser CLI reads configuration from `~/.antybrowser/config.json` or environment variables:

| Variable              | Description                      | Default                     |
|-----------------------|----------------------------------|-----------------------------|
| `ANTYBROWSER_API_KEY` | API key for cloud features       | —                           |
| `ANTYBROWSER_API_URL` | Local API server URL             | `http://127.0.0.1:3001`     |
| `ANTYBROWSER_PROFILE` | Config profile to use            | `default`                   |

## SDK Libraries

Official SDKs are available for 7 languages:

| Language   | Package                    | Repository                                                |
|------------|----------------------------|-----------------------------------------------------------|
| TypeScript | `@antybrowser/sdk`         | [github.com/antybrowser/SDK](https://github.com/antybrowser/SDK) |
| Python     | `antybrowser`              | [github.com/antybrowser/SDK](https://github.com/antybrowser/SDK) |
| C#         | `Antybrowser.SDK`          | [github.com/antybrowser/SDK](https://github.com/antybrowser/SDK) |
| Go         | `github.com/antybrowser/sdk-go` | [github.com/antybrowser/SDK](https://github.com/antybrowser/SDK) |
| PHP        | `antybrowser/sdk`          | [github.com/antybrowser/SDK](https://github.com/antybrowser/SDK) |
| Ruby       | `antybrowser`              | [github.com/antybrowser/SDK](https://github.com/antybrowser/SDK) |
| Java       | `com.antybrowser.sdk`      | [github.com/antybrowser/SDK](https://github.com/antybrowser/SDK) |

## Links

- [Antybrowser Website](https://antybrowser.com)
- [Documentation](https://docs.antybrowser.com)
- [GitHub Releases](https://github.com/antybrowser/antybrowser-cli/releases)
- [SDK Repository](https://github.com/antybrowser/SDK)
- [Report Issues](https://github.com/antybrowser/antybrowser-cli/issues)

## License

[MIT](LICENSE)
