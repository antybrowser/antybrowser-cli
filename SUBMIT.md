# Antybrowser CLI — Package Manager Submission Guide

This guide walks you through submitting the Antybrowser CLI to each package manager.

## Prerequisites

1. GitHub account: `antybrowser`
2. The release is live: https://github.com/antybrowser/antybrowser-cli/releases/tag/cli-v1.0.0

---

## 1. WinGet (Windows)

**URL after submission:** https://github.com/microsoft/winget-pkgs

### Steps:
1. Fork `microsoft/winget-pkgs` on GitHub
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/winget-pkgs.git
   cd winget-pkgs
   ```
3. Create the manifest directory:
   ```bash
   mkdir -p manifests/a/Antybrowser/AntybrowserCLI/1.0.0
   ```
4. Copy the 3 manifest files from `winget/antybrowser-cli.yaml` (it's a singleton, so you only need one file):
   ```bash
   cp ../antybrowser-cli/winget/antybrowser-cli.yaml manifests/a/Antybrowser/AntybrowserCLI/1.0.0/Antybrowser.AntybrowserCLI.yaml
   ```
5. Commit and push:
   ```bash
   git add -A
   git commit -m "Add Antybrowser CLI v1.0.0"
   git push
   ```
6. Create PR against `microsoft/winget-pkgs` master branch

### Validation:
```bash
# Install wingetcreate (one-time)
winget install wingetcreate

# Or validate locally
winget validate manifests/a/Antybrowser/AntybrowserCLI/1.0.0/
```

---

## 2. Chocolatey (Windows)

**URL after submission:** https://community.chocolatey.org/packages/antybrowser-cli

### Steps:
1. Create account at https://community.chocolatey.org/account
2. Generate API key at https://community.chocolatey.org/account/apikeys
3. Install Chocolatey (one-time):
   ```powershell
   Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
   ```
4. Build the package:
   ```powershell
   cd C:\Users\antybrowser\Desktop\antybrowser-cli\chocolatey
   choco pack antybrowser-cli.nuspec
   ```
5. Push to Chocolatey:
   ```powershell
   choco push antybrowser-cli.1.0.0.nupkg --source="https://push.chocolatey.org/" --api-key="YOUR_API_KEY"
   ```

---

## 3. Homebrew (macOS/Linux)

**URL after submission:** https://formulae.brew.sh/formula/antybrowser-cli

### Option A: Homebrew Tap (easier, immediate)
1. Fork `antybrowser/homebrew-tap` on GitHub (already exists)
2. Add the formula:
   ```bash
   git clone https://github.com/antybrowser/homebrew-tap.git
   cd homebrew-tap
   cp ../antybrowser-cli/homebrew/antybrowser-cli.rb Formula/
   git add -A
   git commit -m "Add antybrowser-cli v1.0.0"
   git push
   ```
3. Users install with:
   ```bash
   brew tap antybrowser/tap https://github.com/antybrowser/homebrew-tap
   brew install antybrowser-cli
   ```

### Option B: Homebrew Core (better discoverability)
1. Fork `Homebrew/homebrew-core` on GitHub
2. Add formula to `Formula/a/antybrowser-cli.rb`
3. Submit PR
4. Must pass CI: `brew audit --new antybrowser-cli`

---

## 4. Docker Hub

**URL after submission:** https://hub.docker.com/r/antybrowser/cli

### Steps:
1. Create account at https://hub.docker.com (or use existing)
2. Create repository `antybrowser/cli`
3. Login:
   ```bash
   docker login
   ```
4. Build:
   ```bash
   cd C:\Users\antybrowser\Desktop\antybrowser-cli
   docker build -t antybrowser/cli:latest -t antybrowser/cli:1.0.0 .
   ```
5. Push:
   ```bash
   docker push antybrowser/cli:latest
   docker push antybrowser/cli:1.0.0
   ```

---

## 5. AUR (Arch Linux)

**URL after submission:** https://aur.archlinux.org/packages/antybrowser-cli-bin

### Steps:
1. Create AUR account at https://aur.archlinux.org
2. Add your SSH public key to account (Settings → SSH Keys)
   - Your key is at: `C:\Users\antybrowser\.ssh\id_ed25519.pub`
3. Clone the AUR repo:
   ```bash
   git clone ssh://aur@aur.archlinux.org/antybrowser-cli-bin.git
   cd antybrowser-cli-bin
   ```
4. Copy the PKGBUILD:
   ```bash
   cp ../antybrowser-cli/aur/PKGBUILD-bin PKGBUILD
   ```
5. Test build:
   ```bash
   makepkg -si
   ```
6. Publish:
   ```bash
   git add -A
   git commit -m "Add antybrowser-cli-bin v1.0.0"
   git push
   ```

---

## 6. Scoop (Windows)

**URL after submission:** https://github.com/ScoopInstaller/Extras

### Option A: Custom Bucket (easier)
1. Fork `antybrowser/scoop-bucket` (already exists)
2. Add `bucket/antybrowser-cli.json`
3. Users install with:
   ```powershell
   scoop bucket add antybrowser https://github.com/antybrowser/scoop-bucket
   scoop install antybrowser-cli
   ```

### Option B: Scoop Extras (better discoverability)
1. Fork `ScoopInstaller/Extras`
2. Add `bucket/antybrowser-cli.json`
3. Submit PR

---

## 7. Nixpkgs (Nix/NixOS)

**URL after submission:** https://search.nixos.org/packages?query=antybrowser-cli

### Steps:
1. Fork `NixOS/nixpkgs` on GitHub
2. Add `pkgs/by-name/an/antybrowser-cli/package.nix` with contents from `nix/default.nix`
3. Submit PR
4. Must pass ofborg eval

---

## Summary Checklist

| Platform | Account Needed | Submit Via | Time to Live |
|----------|---------------|------------|--------------|
| WinGet | GitHub | PR to winget-pkgs | 1-7 days |
| Chocolatey | chocolatey.org | `choco push` | 1-3 days |
| Homebrew | GitHub | PR to homebrew-core or tap | Tap: instant, Core: 1-7 days |
| Docker Hub | hub.docker.com | `docker push` | Instant |
| AUR | aur.archlinux.org | SSH push | Instant |
| Scoop | GitHub | PR to ScoopInstaller/Extras | 1-7 days |
| Nixpkgs | GitHub | PR to NixOS/nixpkgs | 1-14 days |
