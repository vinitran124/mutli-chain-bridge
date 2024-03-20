package backend

import (
	"fmt"
	"io"
	"runtime"
)

var (
	Version   = "v0.1.0"
	GitRev    = "undefined"
	GitBranch = "undefined"
	BuildDate = "Fri, 21 March 2024 00:16:00"
)

// PrintVersion prints version info into the provided io.Writer.
func PrintVersion(w io.Writer) {
	fmt.Fprintf(w, "Version:      %s\n", Version)
	fmt.Fprintf(w, "Git revision: %s\n", GitRev)
	fmt.Fprintf(w, "Git branch:   %s\n", GitBranch)
	fmt.Fprintf(w, "Go version:   %s\n", runtime.Version())
	fmt.Fprintf(w, "Built:        %s\n", BuildDate)
	fmt.Fprintf(w, "OS/Arch:      %s/%s\n", runtime.GOOS, runtime.GOARCH)
}
