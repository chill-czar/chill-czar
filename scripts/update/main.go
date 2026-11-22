package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type PRItem struct {
	Number int    `json:"number"`
	State  string `json:"state"`
	URL    string `json:"url"`
}

type TrackedRepo struct {
	Repo            string
	Focus           string
	AdditionalRepos []string
	FallbackStatus  string
}

func fetchRepoPRs(username, repo string) ([]PRItem, error) {
	cmd := exec.Command("gh", "search", "prs", "--author", username, "--repo", repo, "--limit", "50", "--json", "number,state,url")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("gh search failed: %v (stderr: %s)", err, stderr.String())
	}

	var items []PRItem
	if err := json.Unmarshal(stdout.Bytes(), &items); err != nil {
		return nil, err
	}
	return items, nil
}

func formatRepoStatus(merged, open []PRItem) string {
	if len(merged) == 0 && len(open) == 0 {
		return "Up to date"
	}

	var parts []string

	if len(merged) > 0 {
		sort.Slice(merged, func(i, j int) bool {
			return merged[i].Number > merged[j].Number
		})
		var links []string
		for _, pr := range merged {
			links = append(links, fmt.Sprintf("[#%d](%s)", pr.Number, pr.URL))
		}
		if len(merged) == 1 {
			parts = append(parts, fmt.Sprintf("1 merged PR (%s)", links[0]))
		} else {
			parts = append(parts, fmt.Sprintf("%d merged PRs (%s)", len(merged), strings.Join(links, ", ")))
		}
	}

	if len(open) > 0 {
		sort.Slice(open, func(i, j int) bool {
			return open[i].Number > open[j].Number
		})
		var links []string
		for _, pr := range open {
			links = append(links, fmt.Sprintf("[#%d](%s)", pr.Number, pr.URL))
		}
		parts = append(parts, fmt.Sprintf("Active: %s", strings.Join(links, ", ")))
	}

	return strings.Join(parts, ", ")
}

func buildWorkTable(username string, repos []TrackedRepo) (string, int, int) {
	var sb strings.Builder
	sb.WriteString("| Repository | Focus | Upstream Status |\n")
	sb.WriteString("| :--- | :--- | :--- |\n")

	totalTrackedOpen := 0
	totalTrackedMerged := 0

	for _, tr := range repos {
		var allPRs []PRItem
		prs, err := fetchRepoPRs(username, tr.Repo)
		if err == nil {
			allPRs = append(allPRs, prs...)
		}
		for _, extraRepo := range tr.AdditionalRepos {
			extraPRs, extraErr := fetchRepoPRs(username, extraRepo)
			if extraErr == nil {
				allPRs = append(allPRs, extraPRs...)
			}
		}

		var status string
		if len(allPRs) > 0 {
			var merged []PRItem
			var open []PRItem
			for _, item := range allPRs {
				st := strings.ToLower(item.State)
				if st == "merged" {
					merged = append(merged, item)
				} else if st == "open" {
					open = append(open, item)
				}
			}
			totalTrackedMerged += len(merged)
			totalTrackedOpen += len(open)
			status = formatRepoStatus(merged, open)
		} else {
			status = tr.FallbackStatus
		}

		repoURL := fmt.Sprintf("https://github.com/%s", tr.Repo)
		sb.WriteString(fmt.Sprintf("| [%s](%s) | %s | %s |\n", tr.Repo, repoURL, tr.Focus, status))
	}

	return strings.TrimSpace(sb.String()), totalTrackedOpen, totalTrackedMerged
}

func getPRCounts(trackedOpen, trackedMerged int) (openCount int, mergedCount int) {
	if trackedMerged > 0 {
		mergedCount = trackedMerged
	} else {
		mergedCount = 14 // fallback verified open-source merged PRs
	}

	if trackedOpen > 0 {
		openCount = trackedOpen
	} else {
		openCount = 5 // fallback verified open-source in-flight PRs
	}

	return openCount, mergedCount
}

func replaceSection(content, startMarker, endMarker, replacement string) (string, error) {
	startIdx := strings.Index(content, startMarker)
	if startIdx == -1 {
		return content, fmt.Errorf("marker not found: %s", startMarker)
	}
	endIdx := strings.Index(content, endMarker)
	if endIdx == -1 {
		return content, fmt.Errorf("marker not found: %s", endMarker)
	}
	if startIdx >= endIdx {
		return content, fmt.Errorf("invalid marker positions for %s and %s", startMarker, endMarker)
	}

	before := content[:startIdx+len(startMarker)]
	after := content[endIdx:]
	return before + "\n" + strings.TrimSpace(replacement) + "\n" + after, nil
}

func parseReadingList(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var items []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.Contains(line, " - http") {
			parts := strings.SplitN(line, " - ", 2)
			title := strings.TrimSpace(parts[0])
			url := strings.TrimSpace(parts[1])
			items = append(items, fmt.Sprintf("- [%s](%s)", title, url))
		} else if strings.HasPrefix(line, "http") {
			items = append(items, fmt.Sprintf("- [%s](%s)", line, line))
		} else {
			items = append(items, fmt.Sprintf("- %s", line))
		}
	}

	if len(items) == 0 {
		return "_Nothing queued right now._", nil
	}
	return strings.Join(items, "\n"), nil
}

func main() {
	username := os.Getenv("GH_USER")
	if username == "" {
		username = "chill-czar"
	}

	readmePath := "README.md"
	if len(os.Args) > 1 {
		readmePath = os.Args[1]
	}

	contentBytes, err := os.ReadFile(readmePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read %s: %v\n", readmePath, err)
		os.Exit(1)
	}
	content := string(contentBytes)

	// Define tracked open-source repositories exclusively
	trackedRepos := []TrackedRepo{
		{
			Repo:           "lima-vm/lima",
			Focus:          "Linux VMs on macOS/Linux, container runtimes",
			FallbackStatus: "1 merged PR ([#5837](https://github.com/lima-vm/lima/pull/5837)), Active: [#5834](https://github.com/lima-vm/lima/pull/5834)",
		},
		{
			Repo:           "kubevirt/kubevirt",
			Focus:          "Kubernetes virtualization API and runtime",
			FallbackStatus: "Active: [#19982](https://github.com/kubevirt/kubevirt/pull/19982), [#19971](https://github.com/kubevirt/kubevirt/pull/19971)",
		},
		{
			Repo:           "argoproj/argo-cd",
			Focus:          "Declarative GitOps continuous delivery for Kubernetes",
			FallbackStatus: "1 merged PR ([#31132](https://github.com/argoproj/argo-cd/pull/31132)), Active: [#31131](https://github.com/argoproj/argo-cd/pull/31131)",
		},
		{
			Repo:           "google/adk-go",
			Focus:          "Go Agent Development Kit",
			FallbackStatus: "5 merged PRs ([#1470](https://github.com/google/adk-go/pull/1470), [#1353](https://github.com/google/adk-go/pull/1353), [#1351](https://github.com/google/adk-go/pull/1351), [#1349](https://github.com/google/adk-go/pull/1349), [#1320](https://github.com/google/adk-go/pull/1320))",
		},
		{
			Repo:           "Tracer-Cloud/opensre",
			Focus:          "CI/CD reliability agent runtime",
			FallbackStatus: "5 merged PRs ([#7028](https://github.com/Tracer-Cloud/opensre/pull/7028), [#7025](https://github.com/Tracer-Cloud/opensre/pull/7025), [#7018](https://github.com/Tracer-Cloud/opensre/pull/7018), [#7001](https://github.com/Tracer-Cloud/opensre/pull/7001), [#6997](https://github.com/Tracer-Cloud/opensre/pull/6997))",
		},
		{
			Repo:           "e2b-dev/runtime",
			Focus:          "MicroVM sandbox code execution",
			FallbackStatus: "1 merged PR ([#3419](https://github.com/e2b-dev/runtime/pull/3419))",
		},
		{
			Repo:            "superserve-ai/superserve",
			Focus:           "Agent sandboxing platform",
			AdditionalRepos: []string{"superserve-ai/canaries"},
			FallbackStatus:  "1 merged PR ([#20](https://github.com/superserve-ai/canaries/pull/20)), Active: [#332](https://github.com/superserve-ai/superserve/pull/332)",
		},
	}

	// 1. Where the work goes table (dynamic query of open-source projects)
	workBlock, trackedOpen, trackedMerged := buildWorkTable(username, trackedRepos)
	content, err = replaceSection(content, "<!--WORK:START-->", "<!--WORK:END-->", workBlock)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error updating work table: %v\n", err)
		os.Exit(1)
	}

	// 2. Stats block (strictly open-source upstream contributions)
	openPRs, mergedPRs := getPRCounts(trackedOpen, trackedMerged)
	statsBlock := fmt.Sprintf("| Metric | Count |\n| :--- | :--- |\n| Merged Upstream PRs | %d |\n| In-flight Upstream PRs | %d |\n| Tracked Ecosystems | Go, Postgres, Linux/KVM, Kubernetes |", mergedPRs, openPRs)
	content, err = replaceSection(content, "<!--STATS:START-->", "<!--STATS:END-->", statsBlock)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error updating stats: %v\n", err)
		os.Exit(1)
	}

	// 3. Reading list
	readingPath := "reading.txt"
	if _, err := os.Stat(readingPath); os.IsNotExist(err) {
		readingPath = filepath.Join(filepath.Dir(readmePath), "reading.txt")
	}
	readingList, err := parseReadingList(readingPath)
	if err != nil {
		readingList = "_Failed to parse reading.txt_"
	}
	content, err = replaceSection(content, "<!--READING:START-->", "<!--READING:END-->", readingList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error updating reading list: %v\n", err)
		os.Exit(1)
	}

	// 4. Sync timestamp
	syncTimestamp := fmt.Sprintf("Last synced: %s UTC", time.Now().UTC().Format("2006-01-02 15:04"))
	content, err = replaceSection(content, "<!--SYNC:START-->", "<!--SYNC:END-->", syncTimestamp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error updating sync timestamp: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(readmePath, []byte(content), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write %s: %v\n", readmePath, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully updated %s for %s (Merged: %d, Open: %d)\n", readmePath, username, mergedPRs, openPRs)
}
