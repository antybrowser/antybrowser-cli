$ErrorActionPreference = 'Stop'

$packageArgs = @{
  packageName            = 'antybrowser'
  installerType          = 'exe'
  url                    = 'https://github.com/antybrowser/antybrowser/releases/download/v2.0.9/Antybrowser%20Setup%202.0.9.exe'
  silentArgs             = '/S'
  validExitCodes         = @(0, 1223)
  checksum               = 'PLACEHOLDER_SHA256'
  checksumType           = 'sha256'
  url64bit               = 'https://github.com/antybrowser/antybrowser/releases/download/v2.0.9/Antybrowser%20Setup%202.0.9.exe'
  checksum64             = 'PLACEHOLDER_SHA256'
  checksumType64         = 'sha256'
}

Install-ChocolateyPackage @packageArgs
