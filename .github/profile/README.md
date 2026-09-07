# Antybrowser

> Anti-detect browser for multi-account management, affiliate marketing, and web scraping.

---

## Products

### [Antybrowser Desktop](https://github.com/antybrowser/antybrowser)
Full-featured desktop app with browser profile management, fingerprint isolation, proxy support, and extension management.

**Install on Windows:**
```powershell
# WinGet (recommended)
winget install Antybrowser.Antybrowser

# Chocolatey
choco install antybrowser

# Scoop
scoop install antybrowser
```

**Install on Linux:**
```bash
# Snap
sudo snap install antybrowser

# Flatpak
flatpak install flathub io.antybrowser.Desktop

# DEB (Debian/Ubuntu)
sudo dpkg -i antybrowser_*.deb

# RPM (Fedora/RHEL)
sudo rpm -i antybrowser_*.rpm

# AppImage (portable)
chmod +x Antybrowser_*.AppImage && ./Antybrowser_*.AppImage
```

**Install on macOS:**
```bash
# Homebrew
brew install --cask antybrowser
```

---

### [Antybrowser CLI](https://github.com/antybrowser/antybrowser-cli)
Command-line interface for automation and scripting.

```bash
# Go
go install github.com/antybrowser/antybrowser-cli/cmd/antybrowser@latest

# npm
npm install -g @antybrowser/cli

# pip
pip install antybrowser-cli

# Homebrew
brew install antybrowser-cli
```

---

### [Antybrowser SDK](https://github.com/antybrowser/SDK)
Multi-language SDK for programmatic access to the Antybrowser API.

| Language | Package |
|----------|---------|
| TypeScript/JavaScript | `npm install @antybrowser/sdk` |
| Python | `pip install antybrowser-sdk` |
| Go | `go get github.com/antybrowser/SDK/go` |
| C# | `dotnet add package Antybrowser.SDK` |
| Ruby | `gem install antybrowser-sdk` |
| Java | Maven/Gradle via GitHub Packages |
| PHP | `composer require antybrowser/sdk` |

---

## Links

| Resource | URL |
|----------|-----|
| Website | [antybrowser.com](https://antybrowser.com) |
| Documentation | [docs.antybrowser.com](https://docs.antybrowser.com) |
| Downloads | [antybrowser.com/download](https://antybrowser.com/download) |
| Blog | [antybrowser.com/blog](https://antybrowser.com/blog) |
| Changelog | [antybrowser.com/changelog](https://antybrowser.com/changelog) |
| Support | [support@antybrowser.com](mailto:support@antybrowser.com) |

---

## Repository Structure

| Repository | Description |
|------------|-------------|
| [`antybrowser/antybrowser`](https://github.com/antybrowser/antybrowser) | Electron desktop app |
| [`antybrowser/antybrowser-cli`](https://github.com/antybrowser/antybrowser-cli) | CLI tool + package manager configs |
| [`antybrowser/SDK`](https://github.com/antybrowser/SDK) | Multi-language SDK |
| [`antybrowser/homebrew-tap`](https://github.com/antybrowser/homebrew-tap) | Homebrew tap |
| [`antybrowser/scoop-bucket`](https://github.com/antybrowser/scoop-bucket) | Scoop bucket |
| [`antybrowser/.github`](https://github.com/antybrowser/.github) | Org profile (you are here) |

---

## Platform Support

| Platform | Desktop App | CLI | SDK |
|----------|:-----------:|:---:|:---:|
| Windows x64 | ✅ | ✅ | ✅ |
| Windows ARM64 | 🔜 | ✅ | ✅ |
| macOS ARM64 | 🔜 | ✅ | ✅ |
| macOS x64 | 🔜 | ✅ | ✅ |
| Linux x64 | ✅ | ✅ | ✅ |
| Linux ARM64 | 🔜 | ✅ | ✅ |

✅ Available now &nbsp; 🔜 Coming soon

---

## Contributing

We welcome contributions! Please see our [Contributing Guide](https://antybrowser.com/docs/contributing) for details.

## License

MIT License - see [LICENSE](LICENSE) for details.
