# devops-security-cli

A CLI for [devops.security](https://devops.security) by [Kakugo](https://kakugo.ch/).

## Getting started

### From Releases

1. Download the release for your OS / Architecture on the 
   [Releases](https://github.com/denysvitali/devops-security-cli/releases) page.
2. Move the binary to your `$PATH` and start using the CLI

### From Source

#### Requirements

- Go 1.19+
- Make

#### Building

```
make build
```

You can then find your build in `./builds`

### Configuring the CLI

Go to `$XDG_CONFIG` (see [`xdg`](https://github.com/adrg/xdg) for reference) and configure the CLI via
the YAML file:

```
$ cat ~/.config/devops-security/config.yaml
```

This file should have at least a `token` entry:

```yaml
token: "your-token-goes-here"
```


## Usage

### `magic-links` - Get Magic Links

```
$ devops-security magic-links
```

Results in:

```
 ID   First Chars   Created At            Expires At            Permissions                                                                             Status 
 14   b77eb         2023-02-04 09:34:46   2023-12-31 23:59:59   consumeLicenses,listMagicLinks,createMagicLinks,revokeMagicLinks,viewUsageInformation   valid  
 15   46633         2023-02-04 18:31:56   2023-02-13 23:59:59   consumeLicenses,listMagicLinks,createMagicLinks,revokeMagicLinks,viewUsageInformation   valid
```