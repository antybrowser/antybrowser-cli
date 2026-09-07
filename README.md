# Antybrowser

[![GitHub Release](https://img.shields.io/github/v/release/antybrowser/antybrowser-cli?style=flat-square&color=blue)](https://github.com/antybrowser/antybrowser-cli/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)

Tools for managing [Antybrowser](https://antybrowser) profiles, proxies, and automations.

## Products

| Product | Description | Install |
|---------|-------------|---------|
| **[Desktop App](https://github.com/antybrowser/antybrowser)** | Full GUI for profile management | See below |
| **[CLI](#cli-installation)** | Command-line interface | See below |
| **[SDK](https://github.com/antybrowser/SDK)** | Multi-language API client | See below |

---

## Desktop App Installation

### Windows

#### WinGet (recommended)
```powershell
winget install Antybrowser.Antybrowser
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

#### Manual Download
Download `Antybrowser Setup {version}.exe` from [GitHub Releases](https://github.com/antybrowser/antybrowser/releases/latest).

### Linux

#### Snap
```bash
sudo snap install antybrowser
```

#### Flatpak
```bash
flatpak install flathub io.antybrowser.Desktop
```

#### DEB (Debian/Ubuntu)
```bash
# Download .deb from GitHub Releases, then:
sudo dpkg -i antybrowser_*.deb
sudo apt-get install -f
```

#### RPM (Fedora/RHEL)
```bash
# Download .rpm from GitHub Releases, then:
sudo rpm -i antybrowser_*.rpm
```

#### AppImage (portable)
```bash
chmod +x Antybrowser_*.AppImage
./Antybrowser_*.AppImage
```

### macOS

#### Homebrew
```bash
brew install --cask antybrowser
```

#### Manual Download
Download `Antybrowser-{version}-mac-universal.dmg` from [GitHub Releases](https://github.com/antybrowser/antybrowser/releases/latest).

---

## CLI Installation

### macOS / Linux (Homebrew)

```bash
brew install antybrowser-cli
```

### Windows

#### WinGet
```powershell
winget install Antybrowser.AntybrowserCLI
```

#### Chocolatey
```powershell
choco install antybrowser-cli
```

### Go

```bash
go install github.com/antybrowser/antybrowser-cli/cmd/antybrowser@latest
```

### npm

```bash
npm install -g @antybrowser/cli
```

### pip

```bash
pip install antybrowser-cli
```

### Manual Download

Download the latest binary from [GitHub Releases](https://github.com/antybrowser/antybrowser-cli/releases/latest):

| Platform       | Archive                                          |
|----------------|--------------------------------------------------|
| macOS (ARM)    | `antybrowser-darwin-arm64.tar.gz`               |
| macOS (Intel)  | `antybrowser-darwin-amd64.tar.gz`               |
| Linux (ARM)    | `antybrowser-linux-arm64.tar.gz`                |
| Linux (x64)    | `antybrowser-linux-amd64.tar.gz`                |
| Windows (ARM)  | `antybrowser-windows-arm64.zip`                 |
| Windows (x64)  | `antybrowser-windows-amd64.zip`                 |

### Docker

```bash
docker pull antybrowser/cli:latest
docker run --rm antybrowser/cli antybrowser --help
```

---

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

---

## SDK Libraries

Official SDKs are available for 7 languages:

| Language   | Package                    | Repository |
|------------|----------------------------|------------|
| TypeScript | `@antybrowser/sdk`         | [SDK](https://github.com/antybrowser/SDK) |
| Python     | `antybrowser`              | [SDK](https://github.com/antybrowser/SDK) |
| C#         | `Antybrowser.SDK`          | [SDK](https://github.com/antybrowser/SDK) |
| Go         | `github.com/antybrowser/sdk-go` | [SDK](https://github.com/antybrowser/SDK) |
| PHP        | `antybrowser/sdk`          | [SDK](https://github.com/antybrowser/SDK) |
| Ruby       | `antybrowser`              | [SDK](https://github.com/antybrowser/SDK) |
| Java       | `com.antybrowser.sdk`      | [SDK](https://github.com/antybrowser/SDK) |

---

## Links

- [Website](https://antybrowser.com)
- [Documentation](https://docs.antybrowser.com)
- [Download](https://antybrowser.com/download)
- [Blog](https://antybrowser.com/blog)
- [Changelog](https://antybrowser.com/changelog)
- [Support](mailto:support@antybrowser.com)

## License

[MIT](LICENSE)
