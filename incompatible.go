package incompatible

import (
	"fmt"
	"go/token"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/mod/modfile"
)

const (
	modEnvKey = "GOMOD"

	incompatibleVersionSuffix = "+incompatible"
)

// Result contains information about a go.mod import that has the +incompatible suffix.
type Result struct {
	Start  token.Position
	End    token.Position
	Reason string
}

// Analyse locates and analyses the projects go.mod for any +incompatible imports.
func Analyse() ([]Result, error) {
	p, err := modPath()
	if err != nil {
		return nil, fmt.Errorf("getting go.mod path: %w", err)
	}

	results, err := AnalyseMod(p)
	if err != nil {
		return nil, fmt.Errorf("analyzing %s file: %w", p, err)
	}

	return results, nil
}

// AnalyseMod analyses the provided go.mod for any +incompatible imports.
func AnalyseMod(path string) ([]Result, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}

	mod, err := modfile.Parse(path, f, nil)
	if err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}

	var results []Result
	for _, r := range mod.Require {
		if r.Indirect { // skip indirect imports
			continue
		}

		if strings.HasSuffix(r.Mod.Version, incompatibleVersionSuffix) { // if the require is +incompatible
			results = append(results, Result{
				Start:  token.Position{Filename: mod.Syntax.Name, Line: r.Syntax.Start.Line, Column: r.Syntax.Start.LineRune},
				End:    token.Position{Filename: mod.Syntax.Name, Line: r.Syntax.End.Line, Column: r.Syntax.End.LineRune},
				Reason: fmt.Sprintf("%s is a package and Go may update this version regardless of breaking changes.", r.Mod.String()),
			})
		}
	}

	return results, nil
}

func modPath() (string, error) {
	out, err := exec.Command("go", "env", modEnvKey).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("locating go.mod via go env %s: %w", modEnvKey, err)
	}
	return strings.TrimSpace(string(out)), nil
}
