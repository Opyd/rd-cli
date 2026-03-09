# rd-cli

A command-line tool for downloading files through the [Real-Debrid](https://real-debrid.com) service.

Provide a link from a supported file hoster (e.g. Rapidgator, 1Fichier) and rd-cli will unrestrict it via the Real-Debrid API and download the file at full speed.

## Requirements

- A Real-Debrid account with an API token
- Go 1.25+

## Installation

Build from source:

```sh
git clone https://github.com/Opyd/rd-cli
cd rd-cli
go build -o rd-cli .
```

## Usage

### Configure your API token

```sh
rd-cli set-token                  # interactive prompt
rd-cli set-token YOUR_TOKEN_HERE  # pass token directly
```

Your token is saved to `user-config.json` in the OS user config directory:

| OS | Path |
|---|---|
| Linux | `~/.config/rd-cli/user-config.json` |
| macOS | `~/Library/Application Support/rd-cli/user-config.json` |
| Windows | `%AppData%\rd-cli\user-config.json` |

### Download a file

```sh
rd-cli download <url>                          # download to current directory
rd-cli download <url> -n filename.mkv          # custom filename
rd-cli download <url> -p /path/to/dir          # download to specific directory
rd-cli download <url> -p /path/to/dir -n file  # custom path and filename
```

### Batch download from file

Create a text file with one URL per line:

```
https://rapidgator.net/file/abc
https://1fichier.com/?xyz
```

```sh
rd-cli download -f links.txt                   # batch download to current directory
rd-cli download -f links.txt -p /path/to/dir   # batch download to specific directory
```

Failed downloads are logged to stderr and the remaining links continue downloading.

## Commands

| Command | Description |
|---|---|
| `set-token [token]` | Save your Real-Debrid API token |
| `get-config` | Display current configuration |
| `download <url>` | Unrestrict and download a file |

## Roadmap

- [x] Unrestrict link via Real-Debrid API
- [x] Download file with progress bar
- [x] Custom filename via `--name` flag
- [x] Custom output path as second argument
- [x] `--path` flag instead of positional argument
- [x] Batch downloading (multiple links from file)
- [ ] Retry on network error with exponential backoff
- [ ] Resume interrupted downloads (HTTP Range header)

## License

MIT
