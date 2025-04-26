# PowerShell uninstallation script for rsec

$ErrorActionPreference = 'Stop'

# Define variables
$BinaryName = "rsec"
$InstallDir = "$env:LOCALAPPDATA\$BinaryName"
$BinaryPath = Join-Path $InstallDir "$BinaryName.exe"

# Check if the binary exists
if (-not (Test-Path $BinaryPath)) {
    Write-Host "$BinaryName is not installed at $BinaryPath"
    exit 0
}

# Remove the binary and installation directory
try {
    # Remove the binary file
    Remove-Item -Path $BinaryPath -Force
    Write-Host "Removed $BinaryName executable"

    # Remove the directory if it's empty
    $dirItems = Get-ChildItem -Path $InstallDir -Force -ErrorAction SilentlyContinue
    if (-not $dirItems -or $dirItems.Count -eq 0) {
        Remove-Item -Path $InstallDir -Force
        Write-Host "Removed installation directory: $InstallDir"
    }

    # Remove from PATH
    $currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
    if ($currentPath -and $currentPath.Contains($InstallDir)) {
        $pathParts = $currentPath.Split(';')
        $newPathParts = $pathParts | Where-Object { $_ -ne $InstallDir }
        $newPath = $newPathParts -join ';'
        [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
        Write-Host "Removed $InstallDir from PATH"
    }

    Write-Host "$BinaryName has been successfully uninstalled."
} catch {
    # Catch and display error without using $_ directly in strings
    $errorMsg = $_.Exception.Message
    Write-Host ("Failed to uninstall " + $BinaryName + ": " + $errorMsg)
    exit 1
}
