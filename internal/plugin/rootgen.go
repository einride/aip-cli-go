package plugin

import "github.com/einride/ctl/internal/codegen"

const (
	programName string = "einridectl2"
	devHost     string = "api-g4oz7jceaa-ew.a.run.app:443"
	prodHost    string = "api-pe3g7ntwkq-ew.a.run.app:443"
)

func GenerateRootFile(f *codegen.File) {
	f.Pf(`
import (
	"context"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"github.com/spf13/pflag"
	"github.com/spf13/cobra"
)

var cfgFile string
const programName string = "%s"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   programName,
	Short: "A brief description of your application",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var CustomConnect func(ctx context.Context) (*grpc.ClientConn, error)

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func RootFlags() *pflag.FlagSet {
	return rootCmd.PersistentFlags()
}

var token string
var address string
var insecure bool
var prod bool

const devHost string = "%s"
const prodHost string = "%s"

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ctl.yaml)")

    rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "identity token")
	rootCmd.PersistentFlags().StringVarP(&address, "address", "a", "", "custom address, e.g localhost:8080")
	rootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "i", false, "make insecure request")
    rootCmd.PersistentFlags().BoolVarP(&prod, "prod", "p", false, "use production environment")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}`, programName, devHost, prodHost)
}
