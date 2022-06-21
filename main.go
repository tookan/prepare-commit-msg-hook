package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	branch, _ := exec.Command("git", "symbolic-ref", "--short", "HEAD").Output()
	r, _ := regexp.Compile("^((feature|hotfix)\\/)?([a-zA-Z]+-[0-9]+)")
	patternMatches := r.FindStringSubmatch(string(branch))
	if len(patternMatches) > 3 && patternMatches[3] != "" {
		taskCode := patternMatches[3]
		commitMesFile := os.Args[1]
		commitMes, _ := os.ReadFile(commitMesFile)

		isCommitMesMatch, _ := regexp.MatchString("^[a-zA-Z]+-[0-9]+\\s", string(commitMes))
		if !isCommitMesMatch {
			result := fmt.Sprintf("%s %s", taskCode, commitMes)
			os.WriteFile(commitMesFile, []byte(result), 0644)
		}
	}

}
