# RunSecret (rsec)

[![Go Report Card](https://goreportcard.com/badge/github.com/runsecret/rsec)](https://goreportcard.com/report/github.com/runsecret/rsec)
[![GitHub release](https://img.shields.io/github/release/runsecret/rsec.svg)](https://github.com/runsecret/rsec/releases/latest)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

> Securely manage secrets for local development with one simple philosophy: **Never store secrets statically**

RunSecret (rsec) is a CLI tool that simplifies secret management for local development by injecting secrets from your team's vault at runtime.

## Table of Contents
- [Why RunSecret?](#why-runsecret)
- [How It Works](#how-it-works)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Secret Reference Format](#secret-reference-format)
- [Complete Example](#complete-example)
- [Commands Reference](#commands-reference)
- [Supported Secret Vaults](#supported-secret-vaults)
- [Troubleshooting](#troubleshooting)
- [Uninstalling](#uninstalling)
- [Contributing](#contributing)
- [License](#license)

## Why RunSecret?

Managing secrets for local development is fundamentally broken:

- ❌ Git-ignored `.env` files with plaintext secrets
- ❌ Manual bootstrapping for every new developer
- ❌ Regular manual updates when secrets rotate
- ❌ High risk of accidental secret commits
- ❌ Ex-employees retain access to secrets

These practices compromise security and waste developer time.

## How It Works

RunSecret uses **secret references** to solve these problems:

1. Store references to secrets (not the secrets themselves) in your environment or `.env` files
2. Use `rsec run` to inject actual secret values from your vault at runtime
3. Continue using your existing development workflow with no code changes

**Benefits:**

- ✅ **Easy Setup**: Works with existing `.env` files and environment variables
- ✅ **Team-Friendly**: References can be safely committed to git
- ✅ **Vault Agnostic**: Works with multiple secret storage solutions
- ✅ **Secure**: Leverages your existing vault permissions
- ✅ **Leak Prevention**: Automatically redacts secrets in logs and console output

## Installation

### macOS (Homebrew)
If you use Homebrew, you can install RunSecret with the following command:
```bash
brew install runsecret/tap/rsec
```

### Linux/macOS (curl)
For Linux and macOS, you can use the following command to install RunSecret:
```bash
curl -sSL https://raw.githubusercontent.com/runsecret/rsec/main/scripts/install.sh | bash
```

### Linux/macOS (wget)
For Linux and macOS, you can use the following command to install RunSecret:
```bash
wget -qO- https://raw.githubusercontent.com/runsecret/rsec/main/scripts/install.sh | bash
```

### Windows (PowerShell):
For Windows, you can use the following command to install RunSecret:
```powershell
iwr -useb https://raw.githubusercontent.com/runsecret/rsec/main/scripts/install.ps1 | iex
```

## Quick Start

1. **Authenticate with your secret vault**

   Ensure your local machine is authenticated with your secrets vault. For AWS, this typically means configuring your AWS credentials.

2. **Replace static secrets with references**

   Instead of:
   ```
   DATABASE_PASSWORD=MyS3cretP@ssw0rd
   ```

   Use a secret reference:
   ```
   DATABASE_PASSWORD=rsec://012345678912/sm.aws/MyDatabasePassword?region=us-east-1
   ```

3. **Run your application with rsec**

   ```bash
   rsec run -- npm start
   ```

   Or with a .env file:
   ```bash
   rsec run -e .env -- npm start
   ```

## Secret Reference Format

Secret references point to secrets in your vault. The format is consistent across all supported secret vaults, though optional arguments may only apply to specific Secret Vaults which support or require thoe features.

### General Secret Reference Format

All secret references conform to the following format:

```bash
rsec://<vaultProviderAddress>/<vaultType>/<secretNameOrPath>?<arguments>
```

where:
* `vaultProviderAddress`: Is the minimum address required to reach out to the vault. For AWS this is simply the account number you want to use. For other, self-hosted vaults, it will be the full hostname used to reach the vault.
* `vaultType`: Is a specific string that tells `rsec` which vault provider this secret lives in.
* `secretNameOrPath`: Is the full path (where applicable) or name of the secret you want to access.
* `arguments`: Are optional arguments provided as query parameters to configure how to access the secret value. Note: Some arguments may only apply for certain secret vaults. These will be documented for each secret vault provider.

### AWS Secrets Manager

Being more specific from the general format defined above, all AWS Secrets Manager references are constructed with the following values:
```bash
rsec://<accountNumber>/sm.aws/<secretName>?region=<region>
```

The `region` attribute is currently only applicable to AWS Secrets Manager references, and will appear on all secret references for this provider.

Example:
```
rsec://012345678912/sm.aws/DatabasePassword?region=us-east-1
```


Tip: Use the `rsec ref` command to generate references from ARNs or vice versa:

```bash
$ rsec ref arn:aws:secretsmanager:us-east-1:012345678912:secret:DatabasePassword
Secret Reference:  rsec://012345678912/sm.aws/MyTestSecret?region=us-east-1
Vault Address: arn:aws:secretsmanager:us-east-1:012345678912:secret:MyDatabasePassword
```

## Complete Example

Here's a complete workflow example:

1. **Create a .env file with secret references**

   ```bash
   # myapp/.env
   DATABASE_URL=rsec://012345678912/sm.aws/DatabasePassword?region=us-east-1
   API_KEY=rsec://012345678912/sm.aws/ApiKey?region=us-east-1
   ```

2. **Commit this file to your repository**

   Since it contains only references, not actual secrets, it's safe to commit and share with your team.

3. **Run your application using rsec**

  For example:
   ```bash
   rsec run -e .env -- npm start
   ```

4. **What happens behind the scenes:**
   - RunSecret loads the .env file
   - Retrieves actual secrets from AWS Secrets Manager
   - Injects them into the environment **only** for your application
   - Monitors stdout/stderr and redacts any accidental leaks of those secrets

## Commands Reference

### `rsec run`

Runs a command with secrets injected into the environment.

```bash
rsec run [-e|--env <.env file>] -- <command> [arguments]
```

Example:
```bash
rsec run -e .env -- npm run dev --prefix ./my_app
```

### `rsec copy`

Retrieves and copies a secret value directly to your clipboard securely.

```bash
rsec copy <secret-reference>
```

Example:
```bash
rsec copy rsec://012345678912/sm.aws/MyApiKey?region=us-east-1
```

### `rsec ref`

Generates a secret reference from a vault address, or vice versa.

```bash
rsec ref <vault-address>
```

Example:
```bash
rsec ref arn:aws:secretsmanager:us-east-1:012345678912:secret:MyApiKey
```

## Supported Secret Vaults

### AWS Secrets Manager

**Authentication:**
- Uses standard AWS SDK authentication
- Supports environment variables, shared credentials file, IAM roles

**Limitations:**
- Only supports string values (not binary)

### Upcoming Vault Support (Roadmap)
- Azure Key Vault
- GCP Secret Manager
- AWS Parameter Store
- HashiCorp Vault

## Uninstalling
We recommend using the same method you used to install RunSecret to uninstall it. For example, if you installed it with Homebrew, you can uninstall it with:
```bash
brew uninstall runsecret/tap/rsec
```
If you installed it with the install script, you can uninstall it with:
```bash
curl -sSL https://raw.githubusercontent.com/runsecret/rsec/main/scripts/uninstall.sh | bash
```
Or with wget:
```bash
wget -qO- https://raw.githubusercontent.com/runsecret/rsec/main/scripts/uninstall.sh | bash
```
Or with PowerShell:
```powershell
iwr -useb https://raw.githubusercontent.com/runsecret/rsec/main/scripts/uninstall.ps1 | iex
```



## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feat/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feat/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
