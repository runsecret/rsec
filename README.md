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

```bash
go install github.com/runsecret/rsec@latest
```

Make sure your `GOPATH/bin` directory is in your system path.

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
   DATABASE_PASSWORD=aws://us-east-1/012345678912/MyDatabasePassword
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

Secret references point to secrets in your vault. The format depends on the vault provider.

For **AWS Secrets Manager**:

```
aws://<region>/<account_number>/<secret_name>
```

Example:
```
aws://us-east-1/012345678912/MyDatabasePassword
```

Use the `rsec ref` command to generate references from ARNs:

```bash
$ rsec ref arn:aws:secretsmanager:us-east-1:012345678912:secret:MyDatabasePassword
Secret Reference:  aws://us-east-1/012345678912/MyTestSecret
Vault Address: arn:aws:secretsmanager:us-east-1:012345678912:secret:MyDatabasePassword
```

## Complete Example

Here's a complete workflow example:

1. **Create a .env file with secret references**

   ```bash
   # myapp/.env
   DATABASE_URL=postgres://postgres:aws://us-east-1/012345678912/DatabasePassword@localhost:5432/myapp
   API_KEY=aws://us-east-1/012345678912/ApiKey
   ```

2. **Commit this file to your repository**

   Since it contains only references, not actual secrets, it's safe to commit.

3. **Run your application using rsec**

   ```bash
   rsec run -e myapp/.env -- npm start
   ```

4. **What happens behind the scenes:**
   - RunSecret loads the .env file
   - Retrieves actual secrets from AWS Secrets Manager
   - Injects them into the environment for your application
   - Monitors stdout/stderr to redact any accidental leaks

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

Copies a secret value to your clipboard securely.

```bash
rsec copy <secret-reference>
```

Example:
```bash
rsec copy aws://us-east-1/012345678912/MyApiKey
```

### `rsec ref`

Generates a secret reference from a vault address.

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
- Configure using AWS CLI: `aws configure`

**Limitations:**
- Only supports string values (not binary)
- Does not support secret rotation

### Upcoming Vault Support (Roadmap)
- GCP Secret Manager
- Azure Key Vault
- HashiCorp Vault

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
