# Antybrowser Distribution Setup Script
# Run this after authenticating with `gh auth login`

Write-Host "=== Antybrowser Distribution Setup ===" -ForegroundColor Cyan
Write-Host ""

# Check if gh is authenticated
$authStatus = gh auth status 2>&1
if ($LASTEXITCODE -ne 0) {
    Write-Host "ERROR: GitHub CLI not authenticated." -ForegroundColor Red
    Write-Host "Run: gh auth login" -ForegroundColor Yellow
    exit 1
}

Write-Host "[1/6] Pushing antybrowser-cli changes..." -ForegroundColor Green
Set-Location "C:\Users\antybrowser\Desktop\antybrowser-cli"
git push origin master
if ($LASTEXITCODE -eq 0) {
    Write-Host "  Pushed successfully" -ForegroundColor Green
} else {
    Write-Host "  Push failed - check your credentials" -ForegroundColor Red
}

Write-Host ""
Write-Host "[2/6] Creating antybrowser/antybrowser repo..." -ForegroundColor Green
$existingRepo = gh repo view antybrowser/antybrowser --json name 2>&1
if ($LASTEXITCODE -ne 0) {
    gh repo create antybrowser/antybrowser --public --description "Antybrowser desktop app - anti-detect browser for multi-account management"
    Write-Host "  Created antybrowser/antybrowser" -ForegroundColor Green
} else {
    Write-Host "  Repo already exists" -ForegroundColor Yellow
}

Write-Host ""
Write-Host "[3/6] Creating antybrowser/.github repo..." -ForegroundColor Green
$existingOrg = gh repo view antybrowser/.github --json name 2>&1
if ($LASTEXITCODE -ne 0) {
    # Clone, add profile, push
    $tempDir = Join-Path $env:TEMP "antybrowser-org-profile"
    if (Test-Path $tempDir) { Remove-Item $tempDir -Recurse -Force }
    
    gh repo create antybrowser/.github --public
    git clone "https://github.com/antybrowser/.github.git" $tempDir
    Set-Location $tempDir
    git checkout -b main 2>$null
    
    New-Item -ItemType Directory -Force -Path "profile" | Out-Null
    Copy-Item "C:\Users\antybrowser\Desktop\antybrowser-cli\.github\profile\README.md" "profile\README.md"
    
    git add -A
    git commit -m "feat: add org profile page"
    git push origin main
    
    Set-Location "C:\Users\antybrowser\Desktop\antybrowser-cli"
    Remove-Item $tempDir -Recurse -Force
    Write-Host "  Created antybrowser/.github with profile page" -ForegroundColor Green
} else {
    Write-Host "  Repo already exists" -ForegroundColor Yellow
}

Write-Host ""
Write-Host "[4/6] Setting up repository secrets..." -ForegroundColor Green
Write-Host "  You need to set these secrets manually:" -ForegroundColor Yellow
Write-Host "    1. Go to https://github.com/antybrowser/antybrowser-cli/settings/secrets" -ForegroundColor Cyan
Write-Host "    2. Add 'TAP_REPO_TOKEN' - a PAT with repo scope for homebrew-tap and scoop-bucket" -ForegroundColor Cyan
Write-Host "    3. Go to https://github.com/antybrowser/antybrowser-cli/settings/secrets/actions" -ForegroundColor Cyan
Write-Host "    4. Add 'SNAP_TOKEN' - from snapcraft.io (snapcraft export-login --snaps=antybrowser)" -ForegroundColor Cyan

Write-Host ""
Write-Host "[5/6] Registering snap package name..." -ForegroundColor Green
Write-Host "  Run these commands:" -ForegroundColor Yellow
Write-Host "    snapcraft login" -ForegroundColor Cyan
Write-Host "    snapcraft register antybrowser" -ForegroundColor Cyan

Write-Host ""
Write-Host "[6/6] Submitting to Flathub (optional)..." -ForegroundColor Green
Write-Host "  Run these commands:" -ForegroundColor Yellow
Write-Host "    1. Fork https://github.com/flathub/flathub on GitHub" -ForegroundColor Cyan
Write-Host "    2. Clone with: git clone -b new-pr https://github.com/YOUR_USERNAME/flathub.git" -ForegroundColor Cyan
Write-Host "    3. Create branch: git checkout -b antybrowser" -ForegroundColor Cyan
Write-Host "    4. Copy desktop/flatpak/* files into the repo" -ForegroundColor Cyan
Write-Host "    5. Commit and push, then create PR against new-pr branch" -ForegroundColor Cyan

Write-Host ""
Write-Host "=== Setup Complete ===" -ForegroundColor Cyan
Write-Host ""
Write-Host "After setting up secrets, test the pipeline:" -ForegroundColor Green
Write-Host "  cd C:\Users\antybrowser\Desktop\octologin.com" -ForegroundColor Cyan
Write-Host "  git tag v2.1.0" -ForegroundColor Cyan
Write-Host "  git push origin v2.1.0" -ForegroundColor Cyan
Write-Host ""
Write-Host "This will trigger the full build + publish pipeline." -ForegroundColor Green
