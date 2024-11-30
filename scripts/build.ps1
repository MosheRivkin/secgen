
$version = "v1.0.0"
$platforms = @(
    @{GOOS="windows"; GOARCH="amd64"; Suffix=".exe"},
    @{GOOS="windows"; GOARCH="386"; Suffix=".exe"},
    @{GOOS="linux"; GOARCH="amd64"; Suffix=""},
    @{GOOS="linux"; GOARCH="386"; Suffix=""},
    @{GOOS="darwin"; GOARCH="amd64"; Suffix=""},
    @{GOOS="darwin"; GOARCH="arm64"; Suffix=""}
)

New-Item -ItemType Directory -Force -Path "release"

foreach ($platform in $platforms) {
    $env:GOOS = $platform.GOOS
    $env:GOARCH = $platform.GOARCH
    
    $output = "release/secgen_${version}_$($platform.GOOS)_$($platform.GOARCH)$($platform.Suffix)"
    Write-Host "Building for $($platform.GOOS) $($platform.GOARCH)..."
    
    go build -o $output main.go
}

Set-Location release
Get-ChildItem -Filter "secgen_*" | ForEach-Object {
    $hash = Get-FileHash -Algorithm SHA256 $_.Name
    "$($hash.Hash)  $($_.Name)" | Out-File -Append -FilePath "checksums.txt"
}