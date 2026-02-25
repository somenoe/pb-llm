# PocketBase Docs Scraper for LLMs

Scrapes PocketBase documentation and generates LLM-optimized output in multiple formats and variations.

## Quick Start

```bash
git clone https://github.com/magooney-loon/pb-llm
cd pb-llm
go run cmd/main.go
```

## Description

Generates 4 documentation variations:

- **Full** - Complete documentation with all extensions
- **Go** - Go extensions only (backend development)
- **JS** - JavaScript extensions only (frontend development)
- **Core** - Core PocketBase without extensions

Each variation is available in two formats:

- `.llm.md` - Ultra-compact LLM-optimized format
- `.txt` - Plain text format

## Usage

```bash
go run cmd/main.go [OPTIONS]
```

### Options

| Option                  | Description                                                              |
| ----------------------- | ------------------------------------------------------------------------ |
| `-h, --help`            | Show help message                                                        |
| `-d, --debug AMOUNT`    | Debug mode - fetch AMOUNT websites per category (0 = disabled)           |
| `-t, --target CATEGORY` | Target categories: `all`, `general`, `api`, `go`, `js` (comma-separated) |
| `-o, --output TYPE`     | Output type: `all`, `txt`, `md`                                          |
| `-f, --folder`          | Save one file per page in folders with sanitized titles                  |

### Examples

```bash
go run cmd/main.go                   # Generate all variations
go run cmd/main.go -d 2              # Debug mode - 2 per category
go run cmd/main.go -t api            # Target API only
go run cmd/main.go -t api,go -d 2    # Target API + Go, debug mode
go run cmd/main.go -o md -f -t js    # Generate JS variation with folder structure
```

## Output

Files are saved in timestamped directories: `docs/session_YYYY-MM-DD_HH-MM-SS.mmm/`

Generated files include:

- `pocketbase_docs_[variation].[format]` - Documentation files
- `summary_[variation].txt` - Statistics for each variation
