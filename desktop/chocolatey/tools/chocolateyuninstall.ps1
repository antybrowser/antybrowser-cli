$ErrorActionPreference = 'Stop'

$packageArgs = @{
  packageName            = 'antybrowser'
  silentArgs             = '/S'
  validExitCodes         = @(0, 1223)
  file                   = "$env:LOCALAPPDATA\Programs\antybrowser\Uninstall Antybrowser.exe"
  fileType               = 'EXE'
}

Uninstall-ChocolateyPackage @packageArgs
