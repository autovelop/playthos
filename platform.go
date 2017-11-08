package engine

type Platform struct {
	Command             string
	Args                []string
	TagsArg             string
	Tags                []string
	DeployFileExtension string
	BuildDependency     string
}
