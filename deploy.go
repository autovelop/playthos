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
		if valid {
			cmdrm := exec.Command("rm", fmt.Sprintf("bin/%v_%v%v", strings.Replace(strings.ToLower(n), " ", "_", -1), name, platform.BinaryFileExtension))
			cmdrm.Start()

			platform.Tags = append(platform.Tags, deps...)

			// Make sure all go dependencies are installed
			cmdinstall := exec.Command("go", "install", fmt.Sprintf("-tags=%v", strings.Trim(fmt.Sprintf("%v", platform.Tags), "[]")), p)
			cmdinstall.Start()

			platform.Args = append(platform.Args, "-o",
				fmt.Sprintf("bin/%v_%v%v", strings.Replace(strings.ToLower(n), " ", "_", -1), name, platform.BinaryFileExtension),
				fmt.Sprintf("%v=%v", platform.TagsArg, strings.Trim(fmt.Sprintf("%v", platform.Tags), "[]")),
			)
			platform.Args = append(platform.Args, p)
			//fmt.Println(platform)

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

			cmdErr, _ := cmd.StderrPipe()

			startErr := cmd.Start()
			if startErr != nil {
				return
			}
			errOutput, _ := ioutil.ReadAll(cmdErr)
			fmt.Printf("%s", errOutput)
		} else {
			fmt.Printf("> Engine: Platform '%v' has the following unresolved dependencies %v...\n", name, deps)
		}
	}
	os.Exit(0)

	// Objectives:
	// only needs deploy tag to trigger
	// deploy concurrently
	// use cgo, gopherjs, or just go (if no opengl/glfw is required)
	// optionally package assets inside binaries

	// 	c := bindata.NewConfig()
	// 	c.Input = []bindata.InputConfig{bindata.InputConfig{
	// 		Path:      filepath.Clean("assets"),
	// 		Recursive: true,
	// 	}}
	// 	c.Package = "engine"
	// 	c.Tags = "deployed"
	// 	c.Output = fmt.Sprintf("%v/assets.go", "../playthos")
	// 	bindata.Translate(c)

	// 	for _, platform := range platforms {
	// 		simpleName := "linux"
	// 		fileExtension := ""
	// 		cgo := false
	// 		var cc string
	// 		arch386 := false
	// 		switch platform {
	// 		case PlatformLinux:
	// 			fmt.Printf("- Linux\n- Requirements: libgl1-mesa-dev, xorg-dev\n\n")
	// 			break
	// 		case PlatformMacOS:
	// 			fmt.Printf("- MacOS\n- Requirements: xcode 7.3, cmake, libxml2, fuse, osxcross\n- Full details: https://github.com/tpoechtrager/osxcross#packaging-the-sdk\n\n")
	// 			simpleName = "darwin"
	// 			cgo = true
	// 			cc = "CC=o32-clang"
	// 		case PlatformWindows:
	// 			fmt.Printf("- Windows (32-bit only)\n- Requirements: mingw-w64-gcc\n\n")
	// 			simpleName = "windows"
	// 			cgo = true
	// 			arch386 = true
	// 			fileExtension = ".exe"
	// 			cc = "CC=i686-w64-mingw32-gcc -fno-stack-protector -D_FORTIFY_SOURCE=0 -lssp"
	// 			break
	// 		default:
	// 			continue
	// 			break
	// 		}

	// 		cmdArgs := []string{
	// 			"build",
	// 			"-v",
	// 			"-o",
	// 			fmt.Sprintf("bin/%v_%v%v", strings.ToLower(e.gameName), simpleName, fileExtension),
	// 			"-tags",
	// 			fmt.Sprintf("deployed %v %v %v", platform, simpleName, GetTags()),
	// 			e.gamePackage,
	// 		}
	// 		cmd := exec.Command("go", cmdArgs...)
	// 		cmd.Env = os.Environ()

	// 		if cgo {
	// 			cmd.Env = append(cmd.Env, "CGO_ENABLED=1")
	// 		}

	// 		if arch386 {
	// 			cmd.Env = append(cmd.Env, "GOARCH=386")
	// 		} else {
	// 			cmd.Env = append(cmd.Env, "GOARCH=amd64")
	// 		}

	// 		if len(cc) > 0 {
	// 			cmd.Env = append(cmd.Env, cc)
	// 		}
	// 		cmd.Env = append(cmd.Env, fmt.Sprintf("GOOS=%v", simpleName))

	// 		cmdErr, _ := cmd.StderrPipe()

	// 		startErr := cmd.Start()
	// 		if startErr != nil {
	// 			return
	// 		}
	// 		errOutput, _ := ioutil.ReadAll(cmdErr)
	// 		fmt.Printf("%s", errOutput)
	// 	}
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
