package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Platform struct {
	Name string
	BuildDirective string
	LinkageDirective string
}

func main() {
	platformsDirPath, _ := os.Args[1], os.Args[2]
	platformsPath := path.Join(platformsDirPath, "platforms.json")
	platformsContent, _ := ioutil.ReadFile(platformsPath)
	var platforms map[string][]Platform
	json.Unmarshal(platformsContent, &platforms)

	reposDirPath := path.Join(platformsDirPath, "repos")
	goModTemplate, _ := ioutil.ReadFile(path.Join(platformsDirPath, "go.mod.template"))
	for pkg, platformsForPkg := range platforms {
		repoPath := path.Join(reposDirPath, "celo-bls-go-" + pkg)
		goMod := strings.Replace(string(goModTemplate), "{PACKAGE}", pkg, 1)
		goModPath := path.Join(repoPath, "go.mod")
		ioutil.WriteFile(goModPath, []byte(goMod), 0644)
		for _, platform := range platformsForPkg {
			fmt.Printf("%s: %+v", pkg, platform)
		}
	}

}