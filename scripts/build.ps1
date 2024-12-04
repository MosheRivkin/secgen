$version = "v0.0.1"
$platforms = @(
    @{GOOS="windows"; GOARCH="386"; Suffix=".exe"},
    @{GOOS="linux"; GOARCH="amd64"; Suffix=""},
    @{GOOS="linux"; GOARCH="386"; Suffix=""},
    @{GOOS="darwin"; GOARCH="amd64"; Suffix=""}, 
    @{GOOS="darwin"; GOARCH="arm64"; Suffix=""},
    @{GOOS="windows"; GOARCH="amd64"; Suffix=".exe"}
)

if (Test-Path "release") {
    Remove-Item -Path "release" -Recurse -Force
}
New-Item -ItemType Directory -Force -Path "release"
New-Item -ItemType Directory -Force -Path "temp"

foreach ($platform in $platforms) {
    $env:GOOS = $platform.GOOS
    $env:GOARCH = $platform.GOARCH
    
    $binaryName = "secgen$($platform.Suffix)"
    $tempPath = "temp/$binaryName"
    Write-Host "Building for $($platform.GOOS) $($platform.GOARCH)..."
    
    go build -o $tempPath main.go
    
    $zipName = "secgen_${version}_$($platform.GOOS)_$($platform.GOARCH).zip"
    Compress-Archive -Path $tempPath -DestinationPath "release/$zipName" -Force
}

Remove-Item -Path "temp" -Recurse

Set-Location release
Get-ChildItem -Filter "*.zip" | ForEach-Object {
    $hash = Get-FileHash -Algorithm SHA256 $_.Name
    "$($hash.Hash)  $($_.Name)" | Out-File -Append -FilePath "checksums.txt"
}
