package helm

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"text/tabwriter"

	"github.com/gosuri/uitable"
	"github.com/gosuri/uitable/util/strutil"
	"k8s.io/helm/pkg/proto/hapi/release"
	"k8s.io/helm/pkg/proto/hapi/services"
	"k8s.io/helm/pkg/timeconv"
)

func (d *Deploy) Status() (*services.GetReleaseStatusResponse, error) {
	return d.client.ReleaseStatus(d.StackName)
}

func PrintStatus(res *services.GetReleaseStatusResponse) (string, error) {
	if res == nil {
		return "", errors.New("release does not exist")
	}
	output := fmt.Sprintf("NAME:    %s\n", res.Name)

	if res.Info.LastDeployed != nil {
		output += fmt.Sprintf("LAST DEPLOYED: %s\n", timeconv.String(res.Info.LastDeployed))
	}
	output += fmt.Sprintf("NAMESPACE: %s\n", res.Namespace)
	output += fmt.Sprintf("STATUS: %s\n", res.Info.Status.Code)
	output += fmt.Sprintf("\n")

	// Resources
	if len(res.Info.Status.Resources) > 0 {
		buffer := &bytes.Buffer{}
		re := regexp.MustCompile("  +")

		w := tabwriter.NewWriter(buffer, 0, 0, 2, ' ', tabwriter.TabIndent)
		fmt.Fprintf(w, "RESOURCES:\n%s\n", re.ReplaceAllString(res.Info.Status.Resources, "\t"))
		w.Flush()
		output += buffer.String()
	}

	// Test
	if res.Info.Status.LastTestSuiteRun != nil {
		lastRun := res.Info.Status.LastTestSuiteRun
		output += fmt.Sprintf("TEST SUITE:\n%s\n%s\n\n%s\n",
			fmt.Sprintf("Last Started: %s", timeconv.String(lastRun.StartedAt)),
			fmt.Sprintf("Last Completed: %s", timeconv.String(lastRun.CompletedAt)),
			formatTestResults(lastRun.Results))
	}

	// Notes
	if len(res.Info.Status.Notes) > 0 {
		output += fmt.Sprintf("NOTEST:\n%s\n", res.Info.Status.Notes)
	}

	return output, nil
}

func formatTestResults(results []*release.TestRun) string {
	tbl := uitable.New()
	tbl.MaxColWidth = 50
	tbl.AddRow("TEST", "STATUS", "INFO", "STARTED", "COMPLETED")
	for i := 0; i < len(results); i++ {
		r := results[i]
		n := r.Name
		s := strutil.PadRight(r.Status.String(), 10, ' ')
		i := r.Info
		ts := timeconv.String(r.StartedAt)
		tc := timeconv.String(r.CompletedAt)
		tbl.AddRow(n, s, i, ts, tc)
	}
	return tbl.String()
}
