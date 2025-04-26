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
    if ((Get-ChildItem -Path $InstallDir -Force | Measure-Object).Count -eq 0) {
        Remove-Item -Path $InstallDir -Force
        Write-Host "Removed installation directory: $InstallDir"
    }

    # Remove from PATH
    $currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
    if ($currentPath.Contains($InstallDir)) {
        $newPath = ($currentPath.Split(';') | Where-Object { $_ -ne $InstallDir }) -join ';'
        [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
        Write-Host "Removed $InstallDir from PATH"
    }

    Write-Host "$BinaryName has been successfully uninstalled."
} catch {
    # Properly capture and display the error
    $errorMessage = $_.Exception.Message
    Write-Host "Failed to uninstall $BinaryName: $errorMessage"
    exit 1
}
