// +build deploy

package engine

import (
	"bytes"
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

func init() {
	deploy = true
}

var vars map[string]string

func newPath(parseString string) string {
	tmpl := template.New("t")
	tmpl, err := tmpl.Parse(parseString)
	if err != nil {
		log.Fatal("Parse: ", err)
		return ""
	}

	var tpl bytes.Buffer
	err1 := tmpl.Execute(&tpl, vars)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
		return ""
	}
	return tpl.String()
}

// TODO: This function is still a mess. Need to make Deploy system in order to tidy it up.
func initDeploy(n string) {
	_, filename, _, ok := runtime.Caller(2)
	if !ok {
		panic("No caller information")
	}

	// p := path.Dir(filename)

	vars = map[string]string{
		"seperator": string(os.PathSeparator),
		"fullpath":  filepath.FromSlash(path.Dir(filename)),
		"gopath":    build.Default.GOPATH,
		"name":      n,
	}
	vars["package"] = strings.Replace(vars["fullpath"], newPath("{{.gopath}}{{.seperator}}src{{.seperator}}"), "", -1)
	fmt.Println(vars["package"])

	fmt.Printf("> Engine: Deploying...\n")
	errs := false
	for name, platform := range platforms {
		valid, deps := validate(name)
		if valid {
			vars["ext"] = platform.DeployFileExtension
			vars["platform"] = name

			fmt.Printf("\n> Engine: Deploying platform '%v'\n", name)
			if platform.BuildDependency != "" {
				fmt.Printf("> Engine: Resolving build dependency '%v'\n", platform.BuildDependency)

				// consider manually forcing a update check via user interface
				// cmdDep := exec.Command("go", "get", "-u", platform.BuildDependency)

				cmdDep := exec.Command("go",
					"get",
					// "-v",
					platform.BuildDependency,
				)
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
				// "-v",
				newPath("-o={{.gopath}}{{.seperator}}bin{{.seperator}}{{.package}}{{.seperator}}{{.platform}}{{.seperator}}{{.name}}{{.ext}}"),
				fmt.Sprintf("%v=%v", platform.TagsArg, strings.Trim(fmt.Sprintf("%v", platform.Tags), "[]")),
			)
			fmt.Printf("> Engine: Deploy tags %v\n", fmt.Sprintf("%v=%v", platform.TagsArg, strings.Trim(fmt.Sprintf("%v", platform.Tags), "[]")))
			fmt.Printf("> Engine: Deploy source %v\n", newPath("{{.fullpath}}"))
			fmt.Printf("> Engine: Deploy destination %v\n", newPath("{{.gopath}}{{.seperator}}bin{{.seperator}}{{.package}}{{.seperator}}{{.platform}}{{.seperator}}{{.name}}{{.ext}}"))
			platform.Args = append(platform.Args, newPath("{{.package}}"))

			cmd := exec.Command(platform.Command, platform.Args...)
			// fmt.Printf("> GOPATH: %v\n", os.Getenv("GOPATH"))
			// cmd.Env = os.Environ()

			// Bring back this code later when we cross compile again
			// if cgo {
			// 	cmd.Env = append(cmd.Env, "CGO_ENABLED=1")
			// }
			if len(platform.ARCH) > 0 {
				cmd.Env = append(cmd.Env, fmt.Sprintf("GOARCH=%v", platform.ARCH))
				// } else {
				// 	cmd.Env = append(cmd.Env, "GOARCH=386")
				// 	cmd.Env = append(cmd.Env, "GOARCH=amd64")
			}

			if len(platform.CC) > 0 {
				cmd.Env = append(cmd.Env, platform.CC)
			}
			// if platform.GOOS != "" {
			// 	cmd.Env = append(cmd.Env, fmt.Sprintf("GOOS=%v", platform.GOOS))
			// 	cmd.Env = append(cmd.Env, fmt.Sprintf("GOPATH=%v", build.Default.GOPATH))
			// }
			// cmd.Env = append(cmd.Env, fmt.Sprintf("GOOS=%v", "linux"))
			cmdErr, _ := cmd.StderrPipe()

			startErr := cmd.Start()
			if startErr != nil {
				fmt.Printf("> Engine: Error 0 during deploy: %v\n", startErr)
				continue
				// os.Exit(0)
			}

			errOutput, _ := ioutil.ReadAll(cmdErr)
			if len(errOutput) > 0 {
				fmt.Println("> Engine: Error 1 during deploy.")
				fmt.Printf("          PLATFORM: %v\n", name)
				wd, err := os.Getwd()
				if err != nil {
					fmt.Printf("> Engine: Unable to get working directory for %v platform\n", name)
				} else {
					fmt.Printf("          CWD: %v\n", wd)
				}
				fmt.Printf("          Error: %v", string(errOutput))
				errs = true
				continue
			}

			// Copy assets
			fmt.Printf("> Engine: Deploying %v assets: %v\n", name, len(assets))
			assetDest := newPath("{{.gopath}}{{.seperator}}bin{{.seperator}}{{.package}}{{.seperator}}{{.platform}}")
			assetSrc := newPath("{{.gopath}}{{.seperator}}src{{.seperator}}{{.package}}{{.seperator}}")
			for _, asset := range assets {
				asset = strings.Replace(asset, "/", "\\", -1)
				assetPath := fmt.Sprintf("%v%v%v", assetDest, string(os.PathSeparator), asset)
				assetPathSplit := strings.Split(assetPath, string(os.PathSeparator))
				assetDir := strings.Join(assetPathSplit[:len(assetPathSplit)-1], string(os.PathSeparator))
				if _, err := os.Stat(assetDir); os.IsNotExist(err) {
					os.MkdirAll(assetDir, os.ModePerm)
				}
				fmt.Printf("> Engine: Copying asset to '%v' from '%v'...\n", assetPath, fmt.Sprintf("%v||%v", assetSrc, asset))
				err := cp(assetPath, fmt.Sprintf("%v%v", assetSrc, asset))
				if err != nil {
					fmt.Printf("> Engine: Deploying asset '%v' failed - %v\n", asset, err)
					errs = true
					continue
					// os.Exit(0)
				}
			}

		} else {
			fmt.Printf("> Engine: Deploying platform '%v' has the following unresolved dependencies %v...\n", name, deps)
			os.Exit(0)
		}
	}

	if errs {
		fmt.Printf("> Engine: Deployment completeled with errors.\n")
	} else {
		fmt.Printf("> Engine: Deployment completeled successfully.\n")
	}
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
