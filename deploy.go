// +build deploy

package engine

import (
	"fmt"
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
		fmt.Println(platform)
		if valid {
			if platform.BuildDependency != "" {
				cmdDep := exec.Command("go", "install", "-i", platform.BuildDependency)
				cmdErrDep, _ := cmdDep.StderrPipe()

				err := cmdDep.Start()
				if err != nil {
					fmt.Printf("> Engine: Error during deploy (installing build depencency) - %v\n", err)
					os.Exit(0)
				}
				errOutput, _ := ioutil.ReadAll(cmdErrDep)
				fmt.Printf("%s", errOutput)
			}

			platform.Tags = append(platform.Tags, deps...)
			platform.Args = append(platform.Args,
				fmt.Sprintf("%v=%v", platform.TagsArg, strings.Trim(fmt.Sprintf("%v", platform.Tags), "[]")),
			)
			platform.Args = append(platform.Args, p)

			cmd := exec.Command(platform.Command, platform.Args...)
			cmd.Env = os.Environ()

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

			// cmdErr, _ := cmd.StderrPipe()

			startErr := cmd.Start()
			if startErr != nil {
				fmt.Printf("> Engine: Error during deploy - %v\n", startErr)
				os.Exit(0)
			}
			// errOutput, _ := ioutil.ReadAll(cmdErr)
			// fmt.Printf("%s", errOutput)
		} else {
			fmt.Printf("> Engine: Deploying platform '%v' has the following unresolved dependencies %v...\n", name, deps)
			os.Exit(0)
		}
	}
	fmt.Printf("> Engine: Deployed\n")
	os.Exit(0)
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
