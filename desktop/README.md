# Antybrowser Desktop - Package Manager Manifests

This directory contains package manager configurations for the **Antybrowser Desktop app** (Electron).

These manifests are separate from the CLI manifests in the root directory.

## Directory Structure

```
desktop/
├── winget/           # WinGet manifest (Windows)
│   └── antybrowser.yaml
├── chocolatey/       # Chocolatey package (Windows)
│   ├── antybrowser.nuspec
│   └── tools/
│       ├── chocolateyinstall.ps1
│       └── chocolateyuninstall.ps1
├── scoop/            # Scoop manifest (Windows)
│   └── antybrowser.json
├── homebrew/         # Homebrew cask (macOS)
│   └── antybrowser.rb
├── snap/             # Snap package (Linux)
│   └── snapcraft.yaml
└── flatpak/          # Flatpak manifest (Linux)
    ├── io.antybrowser.Desktop.yml
    ├── io.antybrowser.Desktop.desktop
    ├── io.antybrowser.Desktop.metainfo.xml
    ├── antybrowser-run.sh
    └── antybrowser.svg
```

## Package Managers

| Manager | Platform | Install Command |
|---------|----------|-----------------|
| WinGet | Windows | `winget install Antybrowser.Antybrowser` |
| Chocolatey | Windows | `choco install antybrowser` |
| Scoop | Windows | `scoop install antybrowser` |
| Homebrew | macOS | `brew install --cask antybrowser` |
| Snap | Linux | `snap install antybrowser` |
| Flathub | Linux | `flatpak install flathub io.antybrowser.Desktop` |

## Updating Versions

When releasing a new version, update the following in each manifest:

1. **Version number** - The version string
2. **Download URLs** - Point to the new release assets
3. **SHA256 checksums** - Calculate from release artifacts

The GitHub Actions workflow (`.github/workflows/release-electron.yml`) automates this process.

## Submission Process

### WinGet
1. Update manifest with new version
2. Fork `microsoft/winget-pkgs`
3. Add manifest to `manifests/a/Antybrowser/Antybrowser/{version}/`
4. Submit PR

### Chocolatey
1. `choco pack antybrowser.nuspec`
2. `choco push antybrowser.{version}.nupkg --source="https://push.chocolatey.org/"`

### Scoop
1. Update manifest in `antybrowser/scoop-bucket` repo
2. Or submit PR to `ScoopInstaller/Extras`

### Snap Store
1. `snapcraft upload --release=stable antybrowser_{version}_amd64.snap`

### Flathub
1. Fork `flathub/flathub`
2. Create branch with updated manifest
3. Submit PR against `new-pr` branch

### Homebrew
1. Update cask in `antybrowser/homebrew-tap` repo
2. Or submit PR to `homebrew/homebrew-cask`
