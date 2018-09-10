package engine

// Platform defines how a platform is identified, deployed against, and packaged.
type Platform struct {
	Command             string
	Args                []string
	TagsArg             string
	Tags                []string
	GOOS                string
	CC                  string
	ARCH                string
	DeployFileExtension string
	BuildDependency     string
}
