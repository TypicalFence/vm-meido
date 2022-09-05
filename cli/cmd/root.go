package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	vm_meido "zaun.moe/vm-meido"
	"zaun.moe/vm-meido/providers"
	. "zaun.moe/vm-meido/providers/digitalocean"
)

var Config *vm_meido.Config
var MeidoFile *vm_meido.MeidoFile
var Provider vm_meido.Provider
var ProviderName providers.ProviderName

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vm-meido",
	Short: "meido that provisions a vm somewhere",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	config, err := vm_meido.LoadConfig()
	cobra.CheckErr(err)
	Config = config

	switch Config.Provider {
	case providers.DIGITALOCEAN:
		provider := CreateDigitaloceanProvider(*config)
		Provider = provider
		ProviderName = providers.DIGITALOCEAN
	default:
		println("Provider: " + config.Provider)
		cobra.CheckErr(errors.New("Provider does not exist"))
	}
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
