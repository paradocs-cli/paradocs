$ErrorActionPreference = 'Stop';
$toolsDir   = "$(Split-Path -parent $MyInvocation.MyCommand.Definition)"
$url        = 'https://saparadocs22.blob.core.windows.net/exec/paradocs.zip'
$url64      = ''
$destLocation = "$($Env:SystemRoot)\System32\"

$packageArgs = @{
  packageName   = 'paradocs'
  unzipLocation = $toolsDir
  #fileType      = 'EXE_MSI_OR_MSU'
  url           = $url
  url64bit      = $url64
  #softwareName  = 'paradocs*'

  #checksum      = ''
  #checksumType  = 'sha256'
  checksum64    = 'e22e975c5758681e7a090aac7ea8b680ac1bb6657f2ff08a263d9e2ba9683cce'
  checksumType64= 'sha256'

  #silentArgs    = "/qn /norestart /l*v `"$($env:TEMP)\$($packageName).$($env:chocolateyPackageVersion).MsiInstall.log`""
 # validExitCodes= @(0, 3010, 1641)
}

#Install-ChocolateyPackage @packageArgs
Install-ChocolateyZipPackage @packageArgs
#Install-ChocolateyZipPackage -PackageName @packageName -Url64bit $url -Checksum64 $checksum -ChecksumType64 sha256 -UnzipLocation $destination










    








