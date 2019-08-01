package helm

import (
	"fmt"

	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/kube"
	"k8s.io/helm/pkg/proto/hapi/release"
	"k8s.io/helm/pkg/proto/hapi/services"
)

type Deploy struct {
	client      *helm.Client
	DryRun      bool
	kubeContext string
	kubeConfig  string
	StackName   string
}

func New(stack, kubeContext, kubeConfig string) (*Deploy, error) {
	cli, err := NewClient(kubeContext, kubeConfig)
	if err != nil {
		return nil, err
	}
	c := &Deploy{
		client:      cli,
		kubeContext: kubeContext,
		kubeConfig:  kubeConfig,
		DryRun:      false,
		StackName:   stack,
	}
	return c, nil
}

func (d *Deploy) Version() (string, error) {
	version, err := d.client.GetVersion()
	if err != nil {
		return "", err
	}
	return version.Version.GetSemVer(), nil
}

func (d *Deploy) NewRelease(chartPath, namespace, revision string, overrides map[string]interface{}) (*release.Release, error) {
	chartRequested, err := chartutil.Load(chartPath)
	if err != nil {
		return nil, err
	}

	if namespace == "" {
		n, _, err := kube.GetConfig(d.kubeContext, d.kubeConfig).Namespace()
		if err != nil {
			return nil, err
		}
		namespace = n
	}

	values, err := d.overrideValues(revision, overrides)
	if err != nil {
		return nil, err
	}

	res, err := d.client.InstallReleaseFromChart(
		chartRequested,
		namespace,
		helm.ValueOverrides(values),
		helm.ReleaseName(d.StackName),
		helm.InstallDryRun(d.DryRun),
		helm.InstallReuseName(false),
		helm.InstallDisableHooks(false),
		helm.InstallDisableCRDHook(false),
		helm.InstallSubNotes(false),
		helm.InstallTimeout(300),
		helm.InstallWait(false),
		helm.InstallDescription(""),
	)
	if err != nil {
		return nil, err
	}
	release := res.GetRelease()
	if release == nil {
		return nil, nil
	}
	return release, nil
}

func (d *Deploy) UpdateRelease(releaseName, chartPath, revision string, overrides map[string]interface{}) (*release.Release, error) {
	chartRequested, err := chartutil.Load(chartPath)
	if err != nil {
		return nil, err
	}

	values, err := d.overrideValues(revision, overrides)
	if err != nil {
		return nil, err
	}

	res, err := d.client.UpdateReleaseFromChart(
		releaseName,
		chartRequested,
		helm.UpdateValueOverrides(values),
		helm.UpgradeDryRun(d.DryRun),
		helm.UpgradeRecreate(false),
		helm.UpgradeForce(false),
		helm.UpgradeDisableHooks(false),
		helm.UpgradeTimeout(300),
		helm.ResetValues(false),
		helm.ReuseValues(false),
		helm.UpgradeSubNotes(false),
		helm.UpgradeWait(false),
		helm.UpgradeDescription(""),
		helm.UpgradeCleanupOnFail(false),
	)
	if err != nil {
		return nil, err
	}
	release := res.GetRelease()
	if release == nil {
		return nil, nil
	}
	return release, nil
}

func (d *Deploy) overrideValues(revision string, overrides map[string]interface{}) ([]byte, error) {
	base := map[string]interface{}{
		"image": map[string]interface{}{
			"tag": revision,
		},
	}
	err := mergo.Merge(&base, overrides)
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(base)
}

func (d *Deploy) PrintRelease(rel *release.Release) (string, error) {
	if rel == nil {
		return "", nil
	}

	output := fmt.Sprintf("NAME:    %s\n", rel.Name)
	if !d.DryRun {
		status, err := d.client.ReleaseStatus(rel.Name)
		if err != nil {
			return "", err
		}
		res, err := PrintStatus(status)
		if err != nil {
			return "", err
		}
		output += res
	}
	return output, nil
}

func (d *Deploy) Delete(releaseName string) (*services.UninstallReleaseResponse, error) {
	opts := []helm.DeleteOption{
		helm.DeleteDryRun(d.DryRun),
		helm.DeleteDisableHooks(false),
		helm.DeletePurge(true),
		helm.DeleteTimeout(300),
		helm.DeleteDescription(""),
	}
	return d.client.DeleteRelease(releaseName, opts...)
}
