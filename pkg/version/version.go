package version

import (
	"encoding/json"
	"runtime"
	"time"
)

var (
	versionFile = "/src/github.com/plombardi89/klusternaut/VERSION"

	// The GitSha of the current commit (automatically set at compile time)
	GitSha string

	// The Version of the program from the VERSION file (automatically set at compile time)
	// Assume version is master so we can fetch versions from tests.
	// ldflags will automatically override this string.
	KlusternautVersion = "master"
)

type Version struct {
	Version   string
	GitCommit string
	BuildDate string
	GoVersion string
	GOOS      string
	GOArch    string
}

func GetVersion() *Version {
	return &Version{
		Version:   KlusternautVersion,
		GitCommit: GitSha,
		BuildDate: time.Now().UTC().String(),
		GoVersion: runtime.Version(),
		GOOS:      runtime.GOOS,
		GOArch:    runtime.GOARCH,
	}
}

func GetVersionJSON() string {
	verBytes, _ := json.Marshal(GetVersion())
	return string(verBytes)
}
