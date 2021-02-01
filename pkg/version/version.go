package version

var (
	version     string = "v0.0.0"
	gitCommitID string = "git-hash"
)

func Version() string {
	return version
}

func GitCommitID() string {
	return gitCommitID
}
