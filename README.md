# gn-term

Terminal-based [GeekNews](https://news.hada.io) reader for Korean tech news.

![Go Version](https://img.shields.io/github/go-mod/go-version/JunyeolYu/gn-term)
![License](https://img.shields.io/github/license/JunyeolYu/gn-term)
![Release](https://img.shields.io/github/v/release/JunyeolYu/gn-term)

## Installation

### Go

```bash
go install github.com/JunyeolYu/gn-term@latest
```

### Homebrew (macOS/Linux)

```bash
brew install JunyeolYu/tap/gn-term
```

### Chocolatey (Windows)

```bash
choco install gn-term
```

### Manual Download

Download the binary from [GitHub Releases](https://github.com/JunyeolYu/gn-term/releases).

## Usage

```bash
gn-term
```

### Keyboard Shortcuts

| Key | Action |
|-----|--------|
| `j` / `↓` | Move down |
| `k` / `↑` | Move up |
| `l` / `→` | View comments / article |
| `h` / `←` | Go back |
| `Space` | Open article in browser |
| `c` | Open comments in browser |
| `r` | Refresh |
| `q` / `Ctrl+C` | Quit |

### Navigation Flow

```
Article List → Comments → Article Content
     ←─────────────←──────────────←
```

- Press `l` or `→` on the article list to view comments
- Press `l` or `→` on comments to view the article content
- Press `h` or `←` to go back

## Version

```bash
gn-term -v
# or
gn-term --version
```

## License

MIT
