$ErrorActionPreference = 'Stop'

$url = 'https://github.com/antybrowser/antybrowser-cli/releases/download/cli-v1.0.0/antybrowser-windows-amd64.zip'
$checksum = '656a6401b20ba7541d90f986942816c703e6bb98d966469fd7a6f45e6b36525b'
$checksumType = 'sha256'

Install-ChocolateyZipPackage -PackageName 'antybrowser-cli' -Url $url -UnzipLocation "$env:toolsDirectory" -Checksum $checksum -ChecksumType $checksumType

$extractedBin = Join-Path "$env:toolsDirectory" "antybrowser-windows-amd64.exe"
$targetBin = Join-Path "$env:toolsDirectory" "antybrowser.exe"

if (Test-Path $extractedBin) {
    Copy-Item $extractedBin $targetBin -Force
    Remove-Item $extractedBin -Force -ErrorAction SilentlyContinue
}

Install-BinFile -Name 'antybrowser' -Path $targetBin
