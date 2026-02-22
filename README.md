# rd-cli

A command-line tool for downloading files through the [Real-Debrid](https://real-debrid.com) service.

Provide a link from a supported file hoster (e.g. Rapidgator, 1Fichier) and rd-cli will unrestrict it via the Real-Debrid API and download the file at full speed.

## Requirements

- A Real-Debrid account with an API token
- Go 1.25+

## Installation

Build from source:

```sh
git clone https://github.com/yourname/rd-cli
cd rd-cli
go build -o rd-cli .
```

## Usage

### Configure your API token

```sh
rd-cli set-token                  # interactive prompt
rd-cli set-token YOUR_TOKEN_HERE  # pass token directly
```

Your token is saved to the OS user config directory (`~/.config/rd-cli/` on Linux/macOS).

### Download a file

```sh
rd-cli download <url>
```

## Commands

| Command | Description |
|---|---|
| `set-token [token]` | Save your Real-Debrid API token |
| `get-config` | Display current configuration |
| `download <url>` | Unrestrict and download a file |

## License

MIT
