package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	spellbookDir := ".."
	spellsDir := filepath.Join(spellbookDir, "spells")
	sidebarPath := filepath.Join(spellbookDir, "_sidebar.md")
	spellsIndexPath := filepath.Join(spellsDir, "index.md")

	// Get base files (excluding hidden and special docsify files)
	baseFiles, err := ioutil.ReadDir(spellbookDir)
	if err != nil {
		fmt.Printf("Error reading spellbook dir: %v\n", err)
		return
	}

	var sidebarLines []string
	sidebarLines = append(sidebarLines, "- [Home](README.md)")

	for _, file := range baseFiles {
		name := file.Name()
		if !file.IsDir() && strings.HasSuffix(name, ".md") && 
		   !strings.HasPrefix(name, "_") && name != "README.md" {
			title := strings.TrimSuffix(name, ".md")
			sidebarLines = append(sidebarLines, fmt.Sprintf("- [%s](%s)", strings.Title(title), name))
		}
	}

	// Process Spells
	spellFiles, err := ioutil.ReadDir(spellsDir)
	if err != nil {
		fmt.Printf("Error reading spells dir: %v\n", err)
		return
	}

	sidebarLines = append(sidebarLines, "- [Spells](spells/index.md)")
	var spellLinks []string
	var indexLinks []string

	for _, file := range spellFiles {
		name := file.Name()
		if !file.IsDir() && strings.HasSuffix(name, ".md") && name != "index.md" {
			title := strings.TrimSuffix(name, ".md")
			spellLinks = append(spellLinks, fmt.Sprintf("  - [%s](spells/%s)", title, name))
			indexLinks = append(indexLinks, fmt.Sprintf("[%s](%s)  ", title, name))
		}
	}
	
	sort.Strings(spellLinks)
	sort.Strings(indexLinks)

	sidebarLines = append(sidebarLines, spellLinks...)

	// Write _sidebar.md
	err = ioutil.WriteFile(sidebarPath, []byte(strings.Join(sidebarLines, "\n")), 0644)
	if err != nil {
		fmt.Printf("Error writing sidebar: %v\n", err)
		return
	}

	// Write spells/index.md
	indexContent := "# spells\n\n" + strings.Join(indexLinks, "\n")
	err = ioutil.WriteFile(spellsIndexPath, []byte(indexContent), 0644)
	if err != nil {
		fmt.Printf("Error writing spells index: %v\n", err)
		return
	}

	fmt.Println("Successfully updated _sidebar.md and spells/index.md")
}
