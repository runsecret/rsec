# PowerShell installation script for rsec

$ErrorActionPreference = 'Stop'

# Define variables
$GithubRepo = "runsecret/rsec"
$BinaryName = "rsec"
$InstallDir = "$env:LOCALAPPDATA\$BinaryName"

# Create installation directory if it doesn't exist
if (-not (Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
    Write-Host "Created installation directory: $InstallDir"
}

# Add to PATH if not already there
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
if (-not $currentPath.Contains($InstallDir)) {
    [Environment]::SetEnvironmentVariable("Path", "$currentPath;$InstallDir", "User")
    Write-Host "Added $InstallDir to PATH"
}

# Determine system architecture
$arch = if ([Environment]::Is64BitOperatingSystem) {
    if ([System.Runtime.InteropServices.RuntimeInformation]::ProcessArchitecture -eq [System.Runtime.InteropServices.Architecture]::Arm64) {
        "arm64"
    } else {
        "x86_64"
    }
} else {
    "386"
}

# Get latest release info
Write-Host "Fetching the latest release of $BinaryName..."
$latestRelease = Invoke-RestMethod -Uri "https://api.github.com/repos/$GithubRepo/releases/latest"
$version = $latestRelease.tag_name
Write-Host "Latest release: $version"

# Construct download URL
$downloadUrl = "https://github.com/$GithubRepo/releases/download/$version/${BinaryName}_Windows_${arch}.tar.gz"
Write-Host "Downloading from: $downloadUrl"

# Create temporary folder
$tempFolder = Join-Path $env:TEMP ([Guid]::NewGuid().ToString())
New-Item -ItemType Directory -Path $tempFolder -Force | Out-Null

# Download the release
$tarGzPath = Join-Path $tempFolder "release.tar.gz"
Invoke-WebRequest -Uri $downloadUrl -OutFile $tarGzPath

# Extract the archive
Write-Host "Extracting..."
# PowerShell 5.1+ has Expand-Archive, but we need to handle tar.gz
tar -xzf $tarGzPath -C $tempFolder

# Install the binary
$sourceBinary = Join-Path $tempFolder "$BinaryName.exe"
$targetBinary = Join-Path $InstallDir "$BinaryName.exe"
Copy-Item -Path $sourceBinary -Destination $targetBinary -Force
Write-Host "Installed $BinaryName to $targetBinary"

# Clean up
Remove-Item -Recurse -Force $tempFolder

Write-Host "$BinaryName $version has been installed successfully!"
Write-Host "Run '$BinaryName --help' to get started."
