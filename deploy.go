// +build deploy

package engine

import (
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func init() {
	deploy = true
}

// TODO: This function is still a mess. Need to make Deploy system in order to tidy it up.
func initDeploy(n string, p string) {

	fmt.Printf("> Engine: Deploying...\n")
	for name, platform := range platforms {
		valid, deps := validate(name)
		if valid {
			fmt.Printf("> Engine: Deploying platform '%v'\n", name)
			if platform.BuildDependency != "" {
				fmt.Printf("> Engine: Resolving build dependency '%v'\n", platform.BuildDependency)

				// consider manually forcing a update check via user interface
				// cmdDep := exec.Command("go", "get", "-u", platform.BuildDependency)

				cmdDep := exec.Command("go", "get", "-v", platform.BuildDependency)
				cmdErrDep, _ := cmdDep.StderrPipe()

				err := cmdDep.Start()
				if err != nil {
					fmt.Printf("> Engine: Error during deploy (installing build dependency): %v\n", err)
					os.Exit(0)
				}
				errOutput, _ := ioutil.ReadAll(cmdErrDep)
				fmt.Printf("%s", errOutput)
			}

			platform.Tags = append(platform.Tags, deps...)
			platform.Args = append(platform.Args,
				"-v",
				fmt.Sprintf("-o=%v/bin/%v/%v%v", build.Default.GOPATH, p, n, platform.DeployFileExtension),
				fmt.Sprintf("%v=%v", platform.TagsArg, strings.Trim(fmt.Sprintf("%v", platform.Tags), "[]")),
			)
			fmt.Printf("> Engine: Deploy tags %v\n", fmt.Sprintf("%v=%v", platform.TagsArg, strings.Trim(fmt.Sprintf("%v", platform.Tags), "[]")))
			fmt.Printf("> Engine: Deploy source %v\n", p)
			fmt.Printf("> Engine: Deploy destination %v\n", fmt.Sprintf("%v/bin/%v/%v%v", build.Default.GOPATH, p, n, platform.DeployFileExtension))
			platform.Args = append(platform.Args, p)

			cmd := exec.Command(platform.Command, platform.Args...)
			// fmt.Printf("> GOPATH: %v\n", os.Getenv("GOPATH"))
			// cmd.Env = os.Environ()

			// Bring back this code later when we cross compile again
			// if cgo {
			// 	cmd.Env = append(cmd.Env, "CGO_ENABLED=1")
			// }
			// if arch386 {
			// 	cmd.Env = append(cmd.Env, "GOARCH=386")
			// } else {
			// 	cmd.Env = append(cmd.Env, "GOARCH=amd64")
			// }
			// if len(cc) > 0 {
			// 	cmd.Env = append(cmd.Env, cc)
			// }
			// cmd.Env = append(cmd.Env, fmt.Sprintf("GOOS=%v", simpleName))
			cmdErr, _ := cmd.StderrPipe()

			startErr := cmd.Start()
			if startErr != nil {
				fmt.Printf("> Engine: Error 0 during deploy: %v\n", startErr)
				os.Exit(0)
			}

			errOutput, _ := ioutil.ReadAll(cmdErr)
			if len(errOutput) > 0 {
				fmt.Printf("> Engine: Error 1 during deploy: %v\n", string(errOutput))
			}

		} else {
			fmt.Printf("> Engine: Deploying platform '%v' has the following unresolved dependencies %v...\n", name, deps)
			os.Exit(0)
		}
	}

	// Copy assets
	fmt.Printf("> Engine: Deploying assets: %v\n", len(assets))
	assetDest := fmt.Sprintf("%v/bin/%v", build.Default.GOPATH, p)
	assetSrc := fmt.Sprintf("%v/src/%v/", build.Default.GOPATH, p)
	for _, asset := range assets {
		assetPath := fmt.Sprintf("%v/%v", assetDest, asset)
		assetPathSplit := strings.Split(assetPath, "/")
		assetDir := strings.Join(assetPathSplit[:len(assetPathSplit)-1], "/")
		if _, err := os.Stat(assetDir); os.IsNotExist(err) {
			os.MkdirAll(assetDir, os.ModePerm)
		}
		fmt.Printf("> Engine: Copying asset to '%v' from '%v'...\n", assetPath, fmt.Sprintf("%v%v", assetSrc, asset))
		err := cp(assetPath, fmt.Sprintf("%v%v", assetSrc, asset))
		if err != nil {
			fmt.Printf("> Engine: Deploying asset '%v' failed - %v\n", asset, err)
			os.Exit(0)
		}
	}

	fmt.Printf("> Engine: Deployment completely successfully\n")
	os.Exit(0)
}

func cp(dst, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	// no need to check errors on read only file, we already got everything
	// we need from the filesystem, so nothing can go wrong now.
	defer s.Close()
	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}

func validate(o string) (bool, []string) {
	dependencies := []string{}
	resolves := []string{}
	tags := []string{}
	for _, pckg := range packages {
		for _, platform := range pckg.Platforms {
			if platform == o || platform == "generic" {
				for _, dependency := range pckg.Dependencies {
					dependencies = append(dependencies, dependency)
				}
				for _, resolve := range pckg.Resolves {
					resolves = append(resolves, resolve)
					tags = append(tags, pckg.Name)
				}
			}
		}
	}
	dependencies = removeDuplicates(dependencies)
	resolves = removeDuplicates(resolves)
	for d := len(dependencies) - 1; d >= 0; d-- {
		dependency := dependencies[d]
		for _, resolve := range resolves {
			if dependency == resolve {
				dependencies = dependencies[:d+copy(dependencies[d:], dependencies[d+1:])]
				break
			}
		}
	}
	if len(dependencies) > 0 {
		return false, dependencies
	}
	tags = removeDuplicates(tags)
	return true, tags
}

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}
	for v := range elements {
		if encountered[elements[v]] == true {
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}
