package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"pb-llm/internal/scraper"
)

func main() {
	var (
		help  = flag.Bool("help", false, "Show help message")
		debug = flag.Bool("d", false, "Debug mode - only fetch first 3 websites per category (12 total)")
	)
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	runScraper(*debug)
}

func runScraper(debug bool) {
	fmt.Println("🚀 PocketBase Documentation Scraper for LLMs")
	fmt.Println("===========================================")
	if debug {
		fmt.Println("🐛 DEBUG MODE - ONLY FETCHING FIRST 3 WEBSITES PER CATEGORY (12 TOTAL)")
	}
	fmt.Println("📦 Generating 4 variations: Full, Go-only, JS-only, Core-only")
	fmt.Println("📦 Each in 2 formats: MD (ultra-compact) and TXT")

	s := scraper.New()

	fmt.Println("📥 Scraping all sections once (smart optimization)...")
	allDocs, err := s.ScrapeAll("both", debug)
	if err != nil {
		log.Fatalf("❌ Scraping failed: %v", err)
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
				log.Printf("⚠️ Failed to save %s %s format: %v", variation.name, format, err)
			} else {
				fmt.Printf("   ✅ %s\n", outputFile)
			}
		}

		summaryFile := fmt.Sprintf("summary_%s.txt", variation.name)
		if err := s.SaveSummaryToFile(filteredDocs, sessionDir, summaryFile); err != nil {
			log.Printf("⚠️ Failed to save %s summary: %v", variation.name, err)
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
	  -help
	        Show this help message
	  -d
	        Debug mode - only fetch first 3 websites per category (12 total)

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
	  go run cmd/main.go -d                   # Debug mode - 3 per category (12 total)

	All files saved in timestamped docs/session_YYYY-MM-DD_HH-MM-SS.mmm/ directory`

	fmt.Println(helpText)
}
