$ErrorActionPreference = 'Stop'

$url = 'https://github.com/antybrowser/antybrowser-cli/releases/download/cli-v1.0.0/antybrowser-windows-amd64.zip'
$checksum = 'PLACEHOLDER_SHA256'
$checksumType = 'sha256'

Install-ChocolateyZipPackage -PackageName 'antybrowser-cli' -Url $url -UnzipLocation "$env:toolsDirectory" -Checksum $checksum -ChecksumType $checksumType

$binPath = Join-Path "$env:toolsDirectory" "antybrowser.exe"
Install-BinFile -Name 'antybrowser' -Path $binPath
