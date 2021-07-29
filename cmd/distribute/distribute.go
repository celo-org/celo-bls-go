package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const WRITE_PERMS = 0644

type Platform struct {
	Name string
	BuildDirective string
	LinkageDirective string
}

func main() {
	sourceDir, _ := os.Args[1], os.Args[2]
	platformsDirPath := path.Join(sourceDir, "platforms")
	platformsPath := path.Join(platformsDirPath, "platforms.json")
	platformsContent, err := ioutil.ReadFile(platformsPath)
	panicIfError(err)
	var platforms map[string][]Platform
	err = json.Unmarshal(platformsContent, &platforms)
	panicIfError(err)

	blsGoOriginalPath := path.Join(platformsDirPath, "bls.go.template")
	blsHOriginalPath := path.Join(platformsDirPath, "bls.h")

	reposDirPath := path.Join(platformsDirPath, "repos")
	goModTemplate, err := ioutil.ReadFile(path.Join(platformsDirPath, "go.mod.template"))
	panicIfError(err)
	blsPlatformTemplate, err := ioutil.ReadFile(path.Join(platformsDirPath, "bls_platform.template"))
	panicIfError(err)
	for pkg, platformsForPkg := range platforms {
		pkgBuildDirectives := []string{}
		repoPath := path.Join(reposDirPath, fmt.Sprintf("celo-bls-go-%s", pkg))
		goMod := strings.Replace(string(goModTemplate), "{PACKAGE}", pkg, 1)
		goModPath := path.Join(repoPath, "go.mod")
		err = writeFile(goModPath, []byte(goMod))
		panicIfError(err)
		blsDirPath := path.Join(repoPath, "bls")
		blsGoPath := path.Join(blsDirPath, "bls.go")
		err = copyFile(blsGoOriginalPath, blsGoPath)
		panicIfError(err)
		blsHPath := path.Join(blsDirPath, "bls.h")
		err = copyFile(blsHOriginalPath, blsHPath)
		panicIfError(err)
		for _, platform := range platformsForPkg {
			blsPlatformPath := path.Join(blsDirPath, fmt.Sprintf("bls_%s.go", platform.Name))
			blsPlatform := strings.Replace(string(blsPlatformTemplate), "{BUILD_DIRECTIVE}", platform.BuildDirective, 1)
			blsPlatform = strings.Replace(blsPlatform, "{LINKAGE_DIRECTIVE}", platform.LinkageDirective, 1)
			err = writeFile(blsPlatformPath, []byte(blsPlatform))
			panicIfError(err)
			pkgBuildDirectives = append(pkgBuildDirectives, platform.BuildDirective)
		}
		pkgRouterTemplatePath := path.Join(platformsDirPath, "bls_router.go.template")
		pkgRouterTemplate, err := ioutil.ReadFile(pkgRouterTemplatePath)
		panicIfError(err)
		pkgRouter := strings.Replace(string(pkgRouterTemplate), "{PACKAGE}", pkg, 1)
		pkgRouter = strings.Replace(pkgRouter, "{BUILD_DIRECTIVE}", strings.Join(pkgBuildDirectives, " "), 1)

		blsPath := path.Join(sourceDir, "bls")
		pkgRouterPath := path.Join(blsPath, fmt.Sprintf("bls_%s.go", pkg))
		err = writeFile(pkgRouterPath, []byte(pkgRouter))
	}
}

func copyFile(src, dest string) error {
	fileContents, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = mkDirForFileIfNeeded(dest)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dest, fileContents, WRITE_PERMS)
}

func writeFile(path string, contents []byte) error {
	err := mkDirForFileIfNeeded(path)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, contents, WRITE_PERMS)
}

func mkDirForFileIfNeeded(filePath string) error {
	dirPath := path.Dir(filePath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.MkdirAll(dirPath, WRITE_PERMS)
	}

	return nil
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}