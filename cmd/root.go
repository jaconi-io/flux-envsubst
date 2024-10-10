package cmd

import (
	"fmt"

	"github.com/jaconi-io/flux-envsubst/v4/envsubst"

	"github.com/spf13/cobra"
	"sigs.k8s.io/kustomize/api/resource"
	"sigs.k8s.io/yaml"

	// Load environment variables from .env files.
	_ "github.com/joho/godotenv/autoload"
)

var rootCmd = &cobra.Command{
	Use:   "flux-envsubst",
	Short: "envsubst for Flux",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := envsubst.SplitYAML(cmd.InOrStdin(), func(b []byte) error {
			res := &resource.Resource{}
			err := yaml.UnmarshalStrict(b, res)
			if err != nil {
				return err
			}

			// For secrets: filter encrypted ones.
			if res.GetKind() == "Secret" {
				sops := res.Field("sops")
				if sops != nil {
					fmt.Fprintf(cmd.OutOrStderr(), "skipping sops encrypted secret %s/%s\n", res.GetNamespace(), res.GetName())
					return nil
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
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
