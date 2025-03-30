# RunSecret (aka - rsec)

RunSecret (rsec) is a CLI that addresss the pain of safely managing secrets for local development with one simple philosophy: **Never store secrets statically**.


## Why?

Managing secrets for local development is, frankly, broken and it's a pain that only gets worse the larger your team is. Most commonly teams are managing secret values within git-ignored `.env` files, that every developer must bootstrap themselves and then remember to manage everytime one of those secrets gets rotated. The secrets are usually stored as plaintext which means they are just one accidental file-rename away from being committed (trust me I've seen it), and when a developer leaves the team they take those secrets with them. This process not only leaves gaps in your security posture, but it takes time away from your developers to focus on what they do best - writing code.

## How is rsec different?

`rsec` is built upon one fundamental concept: "secret references". Instead of storing secrets statically, `rsec` allows you to reference secrets stored in your team's secret vault (ex: AWS Secrets Manager) and inject them into your local environment _at runtime_ for your application to consume. These secret references are simple to understand (and can be generated by rsec!) and provide several benefits:
- **Easy to Setup**: Secret references can be stored in ENV VARs, or loaded from `.env` files. When you start your application using rsec, any provided `.env` files are loaded into the ENV VARs for that process and replaced with their actual secret values. This means no changes to your application code are required!
- **Simple for Teams**: Secret references are not secrets themselves, so `.env` files containing them can be safely stored in your git repository. This means that when a developer leaves the team, they take nothing with them, and new developers have what they need without any manual setup.
- **Vault Agnostic**: `rsec` is built to be vault agnostic, meaning you can use it with any secret vault that you choose. As a v0.1.0 product, `rsec` currently supports AWS Secrets Manager, but support for other vaults such as GCP Secrets Manager, Azure Key Vault, Hashicorp Vault and more are on the roadmap.
- **Secure**: Secret references point to your existing secrets vault, and access to them is controlled by the permissions you already have set up in that vault. No additional permissions are required for `rsec` to work, and if you remove a developer's access from the vault they will no longer be able to access thoe secrets.
- **Prevent Leaks**: Since `rsec` knows about the secrets it injects into your environment, it also has the ability to detect and redact accidental leaks of those secrets in your logs! rsec scans stdout and stderr for any secret values and replaces them with `*****` before they are printed to the console, reducing the risk of accidental leaks.

Excited? Let's get started!

# Installation

`rsec` is distributed as a single binary, written in Go, and can be installed using `go install`:

```bash
go install github.com/runsecret/rsec@latest
```

Package managers and other installation methods are on the roadmap, but for now this is the easiest way to get started.

# Quick Start

`rsec` is designed to be simple to use, and require as few changes as possible to get setup.

The first thing you need to do is update any current instances of staticly stored secrets, either in ENV VARS or a `.env` file to secret references. Also make sure you local machine is properly authenticated to the referenced secrets vault so `rsec` is able to rerieve your secrets.

Once you've done so, you're ready to make magic happen!To run your application with `rsec` simply run:

```bash
rsec run -- <your usual application comamnd>
```

Or, if you store your secret references in a `.env` file:

```bash
rsec run -e <you.env> -- <your usual application comamnd>
```

One thing that might seem strange if you haven't come across it before is that double dash. This is a standard unix command to specify that we've reached the end of our command options for `rsec`. This is required so that you can safely provide the command flags and arguments you need for your application, without `rsec` thinking those are for it. For example you could run something like the following, and the double dash will ensure that `rsec` doesn ot attempt to parse the `--prefix` flag:

```bash
rsec run -e .env -- npm run dev --prefix ./my_app
```

## Redacting secrets


And that's it! Below you'll find a couple other helper commands, as well as details on the secret vaults supported today. I hope you (and your SecOps team!) enjoy `rsec`. If you hit any problems, or have ideas on how to make `rsec` even better, please don't hesitate to create [an issue](https://github.com/runsecret/rsec/issues/new).

# Commands

Today `rsec` has only 3 simple commands.

## `rsec run`

This is where all the magic happens. Any command run with `rsec run` will have all ENV VAR (or `.env`) secret references replaced with actual secret values before running. Additionally, `stdout` and `stderr` will filter out any cases of those secrets with a redacted `*****` string.

### Arguments
`rsec run` requires only one argument, which is the command you want to run (usually the command you use today to run your application locally). It is highly recommended to use this command after the standard unix `--` operator, as this will ensure `rsec` does not attempt to parse any flags or arguments you provide to your command. For example:

```bash
rsec run -- npm run dev --prefix ./my_app
```

### Flags
`rsec run` accepts the following optional flags:
- `-e, --env string   Env file to read env vars from`
- `-h, --help         help for run`


## `rsec copy`

Sometimes you need a secret value for something other than local development. This could be a password for your database client, or a GUI login. `rsec` makes grabbing these secrets from your vault painless and as secure as possible, with the `rsec copy` command. Rooted in our **Never store secrets statically** philosphy, the `rsec copy` command will neither print nor store the requested secret value locally, and instead copies the secret value directly to your clipboard. This process uses the OSC 52 escape sequence to be easily cross platform, and accessible over SSH.

### Arguments
`rsec copy` requires only one argument, which is the _secret reference_ for the secret you want to copy.

```bash
rsec copy aws://us-east-1/012345678912/MyTestSecret
```

### Flags
`rsec copy` accepts the following optional flags:
- `-h, --help         help for copy`

## `rsec ref`

The `rsec ref` is a helper method that accepts a secret identifier (either a secret reference, or a vault address) and outputs both the vault address and secret reference for the provided value. This command is primarily used to ease the generation of accurate secret references, but can also be used when the original vault address has been lost

### Arguments
`rsec ref` requires only one argument, which is a secret identifier.

```bash
rsec ref arn:aws:secretsmanager:us-east-1:012345678912:secret:MyTestSecret
```

### Flags
`rsec ref` accepts the following optional flags:
- `-h, --help         help for ref`

# Supported Vaults

## AWS Secrets Manager

### Secret Reference Format

`rsec` fully supports retrieving secrets from [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/). Typicall the vault address for AWS Secrets Manager comes in the form of an "ARN" that looks something like this:

`arn:aws:secretsmanager:us-east-1:012345678912:secret:MyTestSecret`

This address is necessarily long for AWS, as this format is used across all AWS services for a variety of needs. For our purposes however, using the ARN presents 2 problems. First, it's long and redundant - we know we are reaching out to secrets manager to find a secret when we use `rsec`. Second, and more importantly, it's very possible you have other ARNs in your ENV VARs that are not secrets but are required by your application. For these reasons, `rsec` uses a different format for secret references that point to AWS Secrets Manager:

`aws://<region>/<account_number>/<secret_name>`

This format removes the redunant information, and is uniquely `rsec` so we know when we are looking at a secret that needs replacing. For example, the above arn example converted into an `rsec` secret reference, now looks like:

`aws://us-east-1/012345678912/MyTestSecret`

Not so bad, right? And don't worry, you don't have to remember this. Pass an arn into the `rsec ref` command and it'll generate the correct secret reference for you!

### Authentication

`rsec` authenticates to AWS Secrets Managert using the Golang AWS SDK. This means all the normal ways you authenticate with AWS from your local machine should "Just Work ™️". If you are having trouble however, feel free to open an issue!
