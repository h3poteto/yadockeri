package helm

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

// Status gets a status of helm package.
func (d *Deploy) Status() (*release.Release, error) {
	client := action.NewStatus(d.config)
	return client.Run(d.StackName)
}

// PrintStatus genrates status text from struct.
func PrintStatus(rel *release.Release) (string, error) {
	if rel == nil {
		return "", errors.New("release does not exist")
	}
	output := fmt.Sprintf("NAME:    %s\n", rel.Name)

	if rel.Info.LastDeployed.IsZero() {
		output += fmt.Sprintf("LAST DEPLOYED: %s\n", rel.Info.LastDeployed.Format(time.ANSIC))
	}
	output += fmt.Sprintf("NAMESPACE: %s\n", rel.Namespace)
	output += fmt.Sprintf("STATUS: %s\n", rel.Info.Status.String())
	output += fmt.Sprintf("REVISION: %d\n", rel.Version)

	// Tests
	executions := executionsByHookEvent(rel)
	if tests, ok := executions[release.HookTest]; !ok || len(tests) == 0 {
		output += fmt.Sprintf("TEST SUITE: None\n")
	} else {
		for _, h := range tests {
			if h.LastRun.StartedAt.IsZero() {
				continue
			}
			output += fmt.Sprintf("TEST SUITE: %s\n%s\n%s\n%s\n",
				h.Name,
				fmt.Sprintf("Last Started:   %s", h.LastRun.StartedAt.Format(time.ANSIC)),
				fmt.Sprintf("Last Completed: %s", h.LastRun.CompletedAt.Format(time.ANSIC)),
				fmt.Sprintf("Phase:          %s", h.LastRun.Phase))
		}
	}

	if strings.EqualFold(rel.Info.Description, "Dry run complete") {
		output += fmt.Sprintf("HOOKS:")
		for _, h := range rel.Hooks {
			output += fmt.Sprintf("---\n# Source: %s\n%s\n", h.Path, h.Manifest)
		}
		output += fmt.Sprintf("MANIFEST:\n%s\n", rel.Manifest)
	}

	// Notes
	if len(rel.Info.Notes) > 0 {
		output += fmt.Sprintf("NOTEST:\n%s\n", rel.Info.Notes)
	}

	return output, nil
}

func executionsByHookEvent(rel *release.Release) map[release.HookEvent][]*release.Hook {
	result := make(map[release.HookEvent][]*release.Hook)
	for _, h := range rel.Hooks {
		for _, e := range h.Events {
			executions, ok := result[e]
			if !ok {
				executions = []*release.Hook{}
			}
			result[e] = append(executions, h)
		}
	}
	return result
}
