package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/urfave/cli/v3"

	"pb-llm/internal/scraper"
	"pb-llm/internal/types"
)

type config struct {
	debugAmount int
	target      string
	targets     map[types.DocumentCategory]struct{}
	output      string
	folderMode  bool
}

func main() {
	cmd := &cli.Command{
		Name:  "pb-llm",
		Usage: "Scrape PocketBase documentation and generate LLM-friendly outputs",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Debug mode - number of websites per category to fetch (0 = disabled)",
				Value:   0,
			},
			&cli.StringFlag{
				Name:    "target",
				Aliases: []string{"t"},
				Usage:   "Target category: all|general|api|go|js (comma-separated supported)",
				Value:   "all",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output type: all|txt|md",
				Value:   "all",
			},
			&cli.BoolFlag{
				Name:    "folder",
				Aliases: []string{"f"},
				Usage:   "Save one file per page into a variation folder",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			cfg, err := parseConfig(c)
			if err != nil {
				return err
			}
			return runScraper(cfg)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("❌ %v\n", err)
		os.Exit(1)
	}
}

func parseConfig(c *cli.Command) (config, error) {
	var cfg config

	cfg.debugAmount = c.Int("debug")
	cfg.target = c.String("target")
	cfg.output = strings.ToLower(strings.TrimSpace(c.String("output")))
	cfg.folderMode = c.Bool("folder")

	if cfg.debugAmount < 0 {
		return cfg, fmt.Errorf("debug amount must be >= 0")
	}

	targets, err := parseTargets(cfg.target)
	if err != nil {
		return cfg, err
	}
	cfg.targets = targets

	if _, err := parseOutputFormats(cfg.output); err != nil {
		return cfg, err
	}

	if len(c.Args().Slice()) > 0 {
		return cfg, fmt.Errorf("unexpected arguments: %s", strings.Join(c.Args().Slice(), " "))
	}

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

func runScraper(cfg config) error {
	fmt.Println("🚀 PocketBase Documentation Scraper for LLMs")
	fmt.Println("===========================================")
	if cfg.debugAmount > 0 {
		fmt.Printf("🐛 DEBUG MODE - FETCHING FIRST %d WEBSITES PER CATEGORY\n", cfg.debugAmount)
	}
	fmt.Printf("🎯 Target categories: %s\n", formatTargets(cfg.targets))
	fmt.Println("📦 Generating 4 variations: Full, Go-only, JS-only, Core-only")
	fmt.Printf("📦 Output format(s): %s\n", cfg.output)
	if cfg.folderMode {
		fmt.Println("📁 Folder mode: enabled (one file per page)")
	}

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

	formats, err := parseOutputFormats(cfg.output)
	if err != nil {
		return err
	}

	for _, variation := range variations {
		fmt.Printf("🎯 Processing %s variation (%s)...\n", variation.name, variation.desc)

		filteredDocs := s.FilterDocsByExtensions(allDocs, variation.extension)
		fmt.Printf("   📊 %d sections included\n", len(filteredDocs))

		for _, format := range formats {
			if cfg.folderMode {
				if err := s.SaveDocsToFolder(filteredDocs, sessionDir, variation.name, format); err != nil {
					fmt.Printf("⚠️ Failed to save %s folder (%s): %v\n", variation.name, format, err)
				} else {
					fmt.Printf("   ✅ %s/ (%s)\n", variation.name, format)
				}
				continue
			}

			outputFile := fmt.Sprintf("pocketbase_docs_%s.%s", variation.name, format)
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
	fmt.Printf("📄 Output generated for formats: %s\n", strings.Join(formats, ", "))
	if cfg.folderMode {
		fmt.Printf("   • full/, go/, js/, core/ - One file per page\n")
	} else {
		fmt.Printf("   • pocketbase_docs_full.*, pocketbase_docs_go.*, pocketbase_docs_js.*, pocketbase_docs_core.*\n")
	}
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

func parseOutputFormats(output string) ([]string, error) {
	switch output {
	case "", "all":
		return []string{"md", "txt"}, nil
	case "md", "txt":
		return []string{output}, nil
	default:
		return nil, fmt.Errorf("invalid output %q (use all|txt|md)", output)
	}
}
