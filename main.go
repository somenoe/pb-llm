package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"pb-llm/internal/scraper"
	"pb-llm/internal/types"
)

type config struct {
	help        bool
	debugAmount int
	target      string
	targets     map[types.DocumentCategory]struct{}
}

type scrapeDoneMsg struct {
	err error
}

type model struct {
	config config
	err    error
}

func main() {
	cfg, err := parseFlags(os.Args[1:])
	if err != nil {
		fmt.Printf("❌ %v\n\n", err)
		printHelp()
		os.Exit(2)
	}

	if cfg.help {
		printHelp()
		return
	}

	p := tea.NewProgram(model{config: cfg}, tea.WithoutRenderer())
	finalModel, runErr := p.Run()
	if runErr != nil {
		fmt.Printf("❌ Bubble Tea runtime error: %v\n", runErr)
		os.Exit(1)
	}

	appModel, ok := finalModel.(model)
	if !ok {
		fmt.Println("❌ Unexpected program model type")
		os.Exit(1)
	}

	if appModel.err != nil {
		fmt.Printf("❌ %v\n", appModel.err)
		os.Exit(1)
	}
}

func parseFlags(args []string) (config, error) {
	var cfg config

	fs := flag.NewFlagSet("pb-llm", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	fs.BoolVar(&cfg.help, "h", false, "Show help message")
	fs.BoolVar(&cfg.help, "help", false, "Show help message")
	fs.IntVar(&cfg.debugAmount, "d", 0, "Debug mode - number of websites per category to fetch (0 = disabled)")
	fs.IntVar(&cfg.debugAmount, "debug", 0, "Debug mode - number of websites per category to fetch (0 = disabled)")
	fs.StringVar(&cfg.target, "t", "all", "Target category: all|general|api|go|js")
	fs.StringVar(&cfg.target, "target", "all", "Target category: all|general|api|go|js")

	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			cfg.help = true
			return cfg, nil
		}
		return cfg, err
	}

	if len(fs.Args()) > 0 {
		return cfg, fmt.Errorf("unexpected arguments: %s", strings.Join(fs.Args(), " "))
	}

	if cfg.debugAmount < 0 {
		return cfg, fmt.Errorf("debug amount must be >= 0")
	}

	targets, err := parseTargets(cfg.target)
	if err != nil {
		return cfg, err
	}
	cfg.targets = targets

	return cfg, nil
}

func parseTargets(targetArg string) (map[types.DocumentCategory]struct{}, error) {
	allTargets := map[types.DocumentCategory]struct{}{
		types.CategoryGeneral:      {},
		types.CategoryAPI:          {},
		types.CategoryGoExtensions: {},
		types.CategoryJSExtensions: {},
	}

	input := strings.TrimSpace(strings.ToLower(targetArg))
	if input == "" || input == "all" {
		return allTargets, nil
	}

	targets := make(map[types.DocumentCategory]struct{})
	parts := strings.Split(input, ",")
	for _, raw := range parts {
		part := strings.TrimSpace(raw)
		switch part {
		case "all":
			return allTargets, nil
		case "general":
			targets[types.CategoryGeneral] = struct{}{}
		case "api":
			targets[types.CategoryAPI] = struct{}{}
		case "go", "go-extensions":
			targets[types.CategoryGoExtensions] = struct{}{}
		case "js", "js-extensions", "javascript":
			targets[types.CategoryJSExtensions] = struct{}{}
		default:
			return nil, fmt.Errorf("invalid target %q (use all|general|api|go|js, comma-separated allowed)", part)
		}
	}

	if len(targets) == 0 {
		return nil, fmt.Errorf("target must include at least one category")
	}

	return targets, nil
}

func (m model) Init() tea.Cmd {
	return runScraperCmd(m.config)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch typed := msg.(type) {
	case scrapeDoneMsg:
		m.err = typed.err
		return m, tea.Quit
	case tea.KeyMsg:
		switch typed.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return ""
}

func runScraperCmd(cfg config) tea.Cmd {
	return func() tea.Msg {
		return scrapeDoneMsg{err: runScraper(cfg)}
	}
}

func runScraper(cfg config) error {
	fmt.Println("🚀 PocketBase Documentation Scraper for LLMs")
	fmt.Println("===========================================")
	if cfg.debugAmount > 0 {
		fmt.Printf("🐛 DEBUG MODE - FETCHING FIRST %d WEBSITES PER CATEGORY\n", cfg.debugAmount)
	}
	fmt.Printf("🎯 Target categories: %s\n", formatTargets(cfg.targets))
	fmt.Println("📦 Generating 4 variations: Full, Go-only, JS-only, Core-only")
	fmt.Println("📦 Each in 2 formats: MD (ultra-compact) and TXT")

	s := scraper.New()

	fmt.Println("📥 Scraping all sections once (smart optimization)...")
	allDocs, err := s.ScrapeAll("both", cfg.debugAmount, cfg.targets)
	if err != nil {
		return fmt.Errorf("scraping failed: %w", err)
	}

	variations := []struct {
		name      string
		extension string
		desc      string
	}{
		{"full", "both", "Complete documentation with all extensions"},
		{"go", "go", "Go extensions only (backend development)"},
		{"js", "js", "JavaScript extensions only (frontend development)"},
		{"core", "none", "Core PocketBase without extensions"},
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05.000")
	sessionDir := fmt.Sprintf("session_%s", timestamp)

	fmt.Printf("💾 Saving all variations to: docs/%s/\n\n", sessionDir)

	for _, variation := range variations {
		fmt.Printf("🎯 Processing %s variation (%s)...\n", variation.name, variation.desc)

		filteredDocs := s.FilterDocsByExtensions(allDocs, variation.extension)
		fmt.Printf("   📊 %d sections included\n", len(filteredDocs))

		formats := []string{"md", "txt"}
		fileExtensions := []string{".md", ".txt"}

		for i, format := range formats {
			outputFile := fmt.Sprintf("pocketbase_docs_%s%s", variation.name, fileExtensions[i])
			if err := s.SaveToFile(filteredDocs, sessionDir, outputFile, format); err != nil {
				fmt.Printf("⚠️ Failed to save %s %s format: %v\n", variation.name, format, err)
			} else {
				fmt.Printf("   ✅ %s\n", outputFile)
			}
		}

		summaryFile := fmt.Sprintf("summary_%s.txt", variation.name)
		if err := s.SaveSummaryToFile(filteredDocs, sessionDir, summaryFile); err != nil {
			fmt.Printf("⚠️ Failed to save %s summary: %v\n", variation.name, err)
		} else {
			fmt.Printf("   ✅ %s\n", summaryFile)
		}

		fmt.Println()
	}

	fmt.Printf("🎉 All variations generated successfully!\n")
	fmt.Printf("📁 Session directory: docs/%s/\n\n", sessionDir)
	fmt.Printf("📄 Available files:\n")
	fmt.Printf("   • pocketbase_docs_full.md/.txt - Complete documentation (ultra-compact)\n")
	fmt.Printf("   • pocketbase_docs_go.md/.txt - Go extensions only (ultra-compact)\n")
	fmt.Printf("   • pocketbase_docs_js.md/.txt - JavaScript extensions only (ultra-compact)\n")
	fmt.Printf("   • pocketbase_docs_core.md/.txt - Core PocketBase only (ultra-compact)\n")
	fmt.Printf("   • summary_*.txt - Individual statistics for each variation\n\n")
	fmt.Printf("🤖 Pick the variation that matches your needs!\n")
	fmt.Printf("💡 .md format is now ultra-compact for maximum token efficiency!\n")

	return nil
}

func formatTargets(targets map[types.DocumentCategory]struct{}) string {
	if len(targets) == 0 {
		return "none"
	}

	ordered := []types.DocumentCategory{
		types.CategoryGeneral,
		types.CategoryAPI,
		types.CategoryGoExtensions,
		types.CategoryJSExtensions,
	}

	result := make([]string, 0, len(targets))
	for _, category := range ordered {
		if _, ok := targets[category]; ok {
			result = append(result, string(category))
		}
	}

	sort.Strings(result)
	return strings.Join(result, ", ")
}

func printHelp() {
	const helpText = `PocketBase Documentation Scraper for LLM Usage
	=============================================

	DESCRIPTION:
	  Scrapes PocketBase documentation and automatically generates 4 variations:
	  • Full - Complete documentation with all extensions
	  • Go-only - Go extensions only (backend development)
	  • JS-only - JavaScript extensions only (frontend development)
	  • Core-only - Core PocketBase without any extensions

	  Each variation is generated in ultra-compact markdown and plain text formats.

	USAGE:
	  go run cmd/main.go [OPTIONS]

	OPTIONS:
	  -h, --help
	        Show this help message
	  -d, --debug AMOUNT
	        Debug mode - fetch AMOUNT websites per category (0 = disabled, default: 0)
	  -t, --target CATEGORY
	        Target categories: all|general|api|go|js
	        Comma-separated values are supported (for example: api,go)

	OUTPUT FORMATS:
	  • .md - Ultra-compact markdown format optimized for LLM token efficiency
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
	  go run cmd/main.go -d 1                 # Debug mode - 1 per selected category
	  go run cmd/main.go -t api               # Scrape API docs only
	  go run cmd/main.go -t api,go -d 2       # API + Go targets, debug=2

	All files saved in timestamped docs/session_YYYY-MM-DD_HH-MM-SS.mmm/ directory`

	fmt.Println(helpText)
}
