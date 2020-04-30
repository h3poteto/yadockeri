package helm

import (
	"github.com/pkg/errors"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/strvals"
)

type Deploy struct {
	config    *action.Configuration
	settings  *cli.EnvSettings
	DryRun    bool
	StackName string
}

func debug(format string, v ...interface{}) {
}

func New(stack string, dryRun bool) (*Deploy, error) {
	actionConfig := new(action.Configuration)
	settings := cli.New()
	helmDriver := "memory"
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), helmDriver, debug); err != nil {
		return nil, err
	}

	c := &Deploy{
		config:    actionConfig,
		settings:  settings,
		DryRun:    dryRun,
		StackName: stack,
	}
	return c, nil
}

func yamlVals(values []string) (map[string]interface{}, error) {
	base := map[string]interface{}{}

	for _, value := range values {
		if err := strvals.ParseInto(value, base); err != nil {
			return nil, err
		}
	}
	return base, nil
}

func isChartInstallable(ch *chart.Chart) (bool, error) {
	switch ch.Metadata.Type {
	case "", "application":
		return true, nil
	}
	return false, errors.Errorf("%s charts are not installable", ch.Metadata.Type)
}

// NewRelease create a new helm release using specified helm chart.
// It is overrided with specified values and update image tag with revision.
func (d *Deploy) NewRelease(chartPath, namespace string, overrides []string) (*release.Release, error) {
	chartRequested, err := loader.Load(chartPath)
	if err != nil {
		return nil, err
	}

	validInstallableChart, err := isChartInstallable(chartRequested)
	if !validInstallableChart {
		return nil, err
	}

	if namespace == "" {
		namespace = d.settings.Namespace()
	}

	rawValues, err := yamlVals(overrides)
	if err != nil {
		return nil, err
	}

	client := action.NewInstall(d.config)
	client.Namespace = namespace

	return client.Run(chartRequested, rawValues)
}

// UpdateRelease updates a exist helm release using specified helm chart.
// It is overrided with specified values and update image tag with revision.
func (d *Deploy) UpdateRelease(releaseName, chartPath string, overrides []string) (*release.Release, error) {
	chartRequested, err := loader.Load(chartPath)
	if err != nil {
		return nil, err
	}

	rawValues, err := yamlVals(overrides)
	if err != nil {
		return nil, err
	}

	client := action.NewUpgrade(d.config)
	client.Namespace = d.settings.Namespace()

	rel, err := client.Run(releaseName, chartRequested, rawValues)
	if err != nil {
		return nil, err
	}
	return rel, nil
}

func (d *Deploy) PrintRelease(rel *release.Release) (string, error) {
	if rel == nil {
		return "", nil
	}

	return PrintStatus(rel)
}

func (d *Deploy) Delete(releaseName string) (*release.UninstallReleaseResponse, error) {
	client := action.NewUninstall(d.config)
	return client.Run(releaseName)
}
