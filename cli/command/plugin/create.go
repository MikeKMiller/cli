package plugin

import (
	"github.com/maliceio/cli/cli"
	"github.com/maliceio/cli/cli/command"
	"github.com/spf13/cobra"
)

// // validateTag checks if the given repoName can be resolved.
// func validateTag(rawRepo string) error {
// 	_, err := reference.ParseNormalizedNamed(rawRepo)
//
// 	return err
// }
//
// // validateConfig ensures that a valid config.json is available in the given path
// func validateConfig(path string) error {
// 	dt, err := os.Open(filepath.Join(path, "config.json"))
// 	if err != nil {
// 		return err
// 	}
//
// 	m := types.PluginConfig{}
// 	err = json.NewDecoder(dt).Decode(&m)
// 	dt.Close()
//
// 	return err
// }
//
// // validateContextDir validates the given dir and returns abs path on success.
// func validateContextDir(contextDir string) (string, error) {
// 	absContextDir, err := filepath.Abs(contextDir)
// 	if err != nil {
// 		return "", err
// 	}
// 	stat, err := os.Lstat(absContextDir)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	if !stat.IsDir() {
// 		return "", errors.Errorf("context must be a directory")
// 	}
//
// 	return absContextDir, nil
// }

type pluginCreateOptions struct {
	repoName string
	context  string
	compress bool
}

func newCreateCommand(maliceCli *command.MaliceCli) *cobra.Command {
	options := pluginCreateOptions{}

	cmd := &cobra.Command{
		Use:   "create [OPTIONS] PLUGIN PLUGIN-DATA-DIR",
		Short: "Create a plugin from a rootfs and configuration. Plugin data directory must contain config.json and rootfs directory.",
		Args:  cli.RequiresMinArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.repoName = args[0]
			options.context = args[1]
			return runCreate(maliceCli, options)
		},
	}

	flags := cmd.Flags()

	flags.BoolVar(&options.compress, "compress", false, "Compress the context using gzip")

	return cmd
}

func runCreate(maliceCli *command.MaliceCli, options pluginCreateOptions) error {
	// var (
	// 	createCtx io.ReadCloser
	// 	err       error
	// )
	//
	// if err := validateTag(options.repoName); err != nil {
	// 	return err
	// }
	//
	// absContextDir, err := validateContextDir(options.context)
	// if err != nil {
	// 	return err
	// }
	//
	// if err := validateConfig(options.context); err != nil {
	// 	return err
	// }
	//
	// compression := archive.Uncompressed
	// if options.compress {
	// 	logrus.Debugf("compression enabled")
	// 	compression = archive.Gzip
	// }
	//
	// createCtx, err = archive.TarWithOptions(absContextDir, &archive.TarOptions{
	// 	Compression: compression,
	// })
	//
	// if err != nil {
	// 	return err
	// }
	//
	// ctx := context.Background()
	//
	// createOptions := types.PluginCreateOptions{RepoName: options.repoName}
	// if err = maliceCli.Client().PluginCreate(ctx, createCtx, createOptions); err != nil {
	// 	return err
	// }
	// fmt.Fprintln(maliceCli.Out(), options.repoName)
	return nil
}