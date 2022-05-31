package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/jaconi-io/flux-envsubst/envsubst"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sigs.k8s.io/kustomize/api/resource"
	"sigs.k8s.io/yaml"

	// Load environment variables from .env files.
	_ "github.com/joho/godotenv/autoload"
)

var rootCmd = &cobra.Command{
	Use:   "flux-envsubst",
	Short: "envsubst for Flux",
	RunE: func(cmd *cobra.Command, args []string) error {
		stdin, err := ioutil.ReadAll(cmd.InOrStdin())
		if err != nil {
			return err
		}

		in, err := envsubst.SplitYAML(stdin)
		if err != nil {
			return err
		}

		for _, i := range in {
			res := &resource.Resource{}
			err = yaml.UnmarshalStrict(i, res)
			if err != nil {
				return err
			}

			// For secrets: filter encrypted ones.
			if res.GetKind() == "Secret" {
				sops := res.Field("sops")
				if sops != nil {
					fmt.Fprintf(cmd.OutOrStderr(), "skipping sops encrypted secret %s/%s\n", res.GetNamespace(), res.GetName())
					continue
				}
			}

			res, err = envsubst.SubstituteVariables(cmd.Context(), res)
			if err != nil {
				return err
			}

			out, err := yaml.Marshal(res)
			if err != nil {
				return err
			}

			_, err = cmd.OutOrStdout().Write(out)
			if err != nil {
				return err
			}

			cmd.OutOrStdout().Write([]byte("---\n"))
		}

		return nil
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(viper.AutomaticEnv)
}