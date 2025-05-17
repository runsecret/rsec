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
  - [AWS Secrets Manager](#aws-secrets-manager)
  - [Azure Key Vault](#azure-key-vault)
  - [HashiCorp Vault](#hashicorp-vault)
  - [Upcoming Vault Support (Roadmap)](#upcoming-vault-support-roadmap)
- [Uninstalling](#uninstalling)
- [Contributing](#contributing)
- [License](#license)


## Why RunSecret?

Most engineering teams are solving local development secrets with git-ignored `.env` files. These files store secrets in plaintext, and are rife with problems:

- ❌ Manual bootstrapping the secret values for every new developer
- ❌ Regular manual updates need to be performed by every team member when secrets rotate
- ❌ There's a high risk of accidental secret commits, especially when you're managing multiple `.env` files
- ❌ Employees retain access to secrets when they rotate teams or leave the company

It's generally a pain for the developers, and a security risk for the company. RunSecret seeks to address these issues by providing a simple, secure way to manage secrets for local development.

## How It Works

RunSecret uses **secret references** to solve these problems. Think of these as pointers that replace the secrets themselves. To use RunSercet you simply:

1. Update your environment variabes or `.env` files with references to secrets (not the secrets themselves)
2. Use `rsec run` to inject actual secret values from your vault at runtime
3. Continue using your existing development workflow with no code changes!

**Benefits:**

- ✅ **Easy Setup**: Works with existing `.env` files and environment variables
- ✅ **Team-Friendly**: `.env` files can be safely committed to git, making team onboarding/offboarding a breeze
- ✅ **Vault Agnostic**: Works with multiple secret storage solutions, with more on the way
- ✅ **Secure**: Leverages your existing vault permissions so only the who need access to secrets can get them
- ✅ **Leak Prevention**: Automatically redacts secret values in logs and console output

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

   Ensure your local machine is authenticated with your secrets vault. This can vary by vault provider, but often requires loging in via a provided CLI.

2. **Replace static secrets with references**

   Instead of:
   ```
   DATABASE_PASSWORD=MyS3cretP@ssw0rd
   ```

   Use a secret reference (AWS Secrets Manager Example):
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

All secret references, regardless of the underlying vault, conform to the following format:

```bash
rsec://<vaultAddress>/<vaultType>/<secretNameOrPath>?<arguments>
```

where:
* `vaultAddress`: Is the minimum address required to reach out to the vault. The exact format of this value may vary from one vault provider to the next.
* `vaultType`: Is a specific string that tells `rsec` which vault provider this secret lives in.
* `secretNameOrPath`: Is the full path (where applicable) or name of the secret you want to access.
* `arguments`: Are optional arguments provided as query parameters to configure how to access the secret value. Note: Some arguments may only apply for certain secret vaults. These will be documented for each secret vault provider.

Tip: Use the `rsec ref` command to generate references from Vault Addresses or vice versa:

```bash
$ rsec ref arn:aws:secretsmanager:us-east-1:012345678912:secret:DatabasePassword
Secret Reference:  rsec://012345678912/sm.aws/MyTestSecret?region=us-east-1
Vault Address: arn:aws:secretsmanager:us-east-1:012345678912:secret:MyDatabasePassword
```

For specific examples for your vault provider, see the relevant [Supported Secret Vaults](#supported-secret-vaults) subsection.

## Complete Example

Here's a complete workflow example:

1. **Create a .env file with secret references**

   ```bash
   # myapp/.env
   DATABASE_URL=rsec://012345678912/sm.aws/DatabasePassword?region=us-east-1
   API_KEY=rsec://apiVault/kv.azure/ApiKey
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
   - Retrieves actual secrets from AWS Secrets Manager and Azure Key Vault (multiple vaults can be used simultaneously)
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

**Secret Reference Format**
All AWS Secrets Manager references are constructed with the following values:
```bash
rsec://<accountNumber>/sm.aws/<secretName>?region=<region>
```

The `region` attribute is currently only applicable to AWS Secrets Manager references, and will appear on all secret references for this provider.

Example:
```
rsec://012345678912/sm.aws/DatabasePassword?region=us-east-1
```

### Azure Key Vault

**Authentication:**
- Uses standard Azure SDK authentication
- Supports Azure CLI, Managed Identity, Service Principal

**Secret Reference Format**
For Azure Key Vault, the reference format is:
```bash
rsec://<vaultAddress>/kv.azure/<secretName>?<arguments>
```

Where:
* `vaultAddress`: The name of the Azure Key Vault (e.g. if using `https://myvault.vault.azure.net/` then `myvault` is the vault address)
* `secretName`: The name of the secret in the Azure Key Vault
* `arguments`: Optional arguments, such as `version` to specify a specific version of the secret.

Example:
```
rsec://myvault/kv.azure/MySecret?version=1.0
```

**Using Azure China/Azure Government/Azure Germany**
If you are using Azure China, Azure Government, or Azure Germany, you can use `rsec` in the same way described above, but will replace the `kv.azure` with a region-specific value. For example, for Azure China, you would use `kv.azure.cn` instead of `kv.azure`. Example secret references for each Azure region is as follows:
```bash
# Azure China
rsec://myvault/kv.azure.cn/MySecret?version=1.0

# Azure Government
rsec://myvault/kv.azure.us/MySecret?version=1.0

# Azure Germany
rsec

### HashiCorp Vault

Supports both versions of the KV secret engine (v1 and v2) and credential-based secret enginees such as database, aws, rabbitmq, etc.

**Authentication:**
- Currently supports token-based authentication via one of the following methods:
  - The `VAULT_TOKEN` environment variable
  - Tokens set in the `~/.vault-token` file (via userpass authentication)

**Limitations**
- Currently returns the entire secret object, not just the value. This will be addressed in [issue #8](https://github.com/runsecret/rsec/issues/8).

**Secret Reference Format**
For HashiCorp Vault, the reference format may look like one of the following:
```bash
rsec://<mountPath>/kv1.hashi/<secretName>?<arguments>
rsec://<mountPath>/kv2.hashi/<secretName>?<arguments>
rsec://<mountPath>/cred.hashi/<secretName>?<arguments>
```

Where:
* `mountPath`: The path where the secret or credential is mounted in HashiCorp Vault.
* `secretName`: The name of the secret in HashiCorp Vault
* `arguments`: Optional arguments
  * `endpoint`: The URL encoded endpoint to the HashiCorp Vault server (e.g. `https://vault.example.com`). If not set, will default to the `VAULT_ADDR` environment variable.


### Upcoming Vault Support (Roadmap)
- GCP Secret Manager
- AWS Parameter Store

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
