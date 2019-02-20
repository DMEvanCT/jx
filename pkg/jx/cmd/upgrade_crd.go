package cmd

import (
	"io"

	"github.com/jenkins-x/jx/pkg/jx/cmd/templates"
	"github.com/jenkins-x/jx/pkg/kube"
	"github.com/jenkins-x/jx/pkg/log"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1/terminal"
)

var (
	upgradeCRDsLong = templates.LongDesc(`
		Upgrades the Jenkins X Custom Resource Definitions in the Kubernetes Cluster
`)

	upgradeCRDsExample = templates.Examples(`
		# Upgrades the Custom Resource Definitions 
		jx upgrade crd
	`)
)

// UpgradeCRDsOptions the options for the upgrade CRDs command
type UpgradeCRDsOptions struct {
	UpgradeOptions
}

// NewCmdUpgradeCRDs defines the command
func NewCmdUpgradeCRDs(f Factory, in terminal.FileReader, out terminal.FileWriter, errOut io.Writer) *cobra.Command {
	options := &UpgradeCRDsOptions{
		UpgradeOptions: UpgradeOptions{
			CommonOptions: CommonOptions{
				Factory: f,
				In:      in,
				Out:     out,
				Err:     errOut,
			},
		},
	}

	cmd := &cobra.Command{
		Use:     "crd",
		Short:   "Upgrades the Jenkins X Custom Resource Definitions in the Kubernetes Cluster",
		Long:    upgradeCRDsLong,
		Example: upgradeCRDsExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			CheckErr(err)
		},
	}
	return cmd
}

// Run implements the command
func (o *UpgradeCRDsOptions) Run() error {
	apisClient, err := o.ApiExtensionsClient()
	if err != nil {
		return errors.Wrap(err, "failed to create the API extensions client")
	}
	err = kube.RegisterAllCRDs(apisClient)
	if err != nil {
		return errors.Wrap(err, "failed to register all CRDs")
	}
	log.Info("Jenkins X CRDs upgraded with success\n")
	return nil
}