# PocketBase Docs Scraper for LLMs

```bash
# Clone the repository
git clone https://github.com/magooney-loon/pb-llm
cd pb-llm

# Run the scraper
go run cmd/main.go
```

```bash
PocketBase Documentation Scraper for LLM Usage
	=============================================

	DESCRIPTION:
	  Scrapes PocketBase documentation and automatically generates 4 variations:
	  • Full - Complete documentation with all extensions
	  • Go-only - Go extensions only (backend development)
	  • JS-only - JavaScript extensions only (frontend development)
	  • Core-only - Core PocketBase without any extensions

	  Each variation is generated in ultra-compact LLM-optimized and plain text formats.

	USAGE:
	  go run cmd/main.go [OPTIONS]

	OPTIONS:
	  -h, --help
	        Show this help message
	  -d, --debug AMOUNT
	        Debug mode - fetch AMOUNT websites per selected category (0 = disabled)
	  -t, --target CATEGORY
	        Target categories: all|general|api|go|js
	        Comma-separated values are supported (example: api,go)
	  -o, --output TYPE
	        Output type: all|txt|md
	  -f, --folder
	        Save one file per page in per-variation folders with sanitized page titles
	        (example: js/Introduction.md, go/Go_Overview.md)

	OUTPUT FORMATS:
	  • .llm.md - Ultra-compact LLM format for maximum token efficiency
	  • .txt - Plain text format for general use

	FEATURES:
	  🤖 LLM-optimized output format
	  📊 Token counting and estimation
	  📈 Context window usage analysis
	  🔧 AI training dataset structure
	  📝 Comprehensive LLM usage statistics
	  📄 Plain text backup format
	  🎯 Automatic generation of all variations
	  📦 Pick exactly what you need

	OUTPUT (4 variations × 2 formats = 8 documentation files):
	  • pocketbase_docs_full.llm.md/.txt - Complete documentation
	  • pocketbase_docs_go.llm.md/.txt - Go extensions only
	  • pocketbase_docs_js.llm.md/.txt - JavaScript extensions only
	  • pocketbase_docs_core.llm.md/.txt - Core PocketBase only
	  • summary_*.txt - Individual statistics for each variation

	EXAMPLE:
	  go run cmd/main.go                      # Generates all 4 variations
	  go run cmd/main.go -d 2                 # Debug mode - 2 per selected category
	  go run cmd/main.go -t api               # Target API category only
	  go run cmd/main.go -t api,go -d 2       # Target API + Go, debug mode
	  go run cmd/main.go -o md                # Only markdown output
	  go run cmd/main.go -o txt               # Only text output
	  go run cmd/main.go -o md -f -t js       # Folder mode with titled files (js/Introduction.md, js/Authentication.md, ...)

	All files saved in timestamped docs/session_YYYY-MM-DD_HH-MM-SS.mmm/ directory
```
