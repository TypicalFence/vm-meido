package cmd

import (
	"github.com/spf13/cobra"
	vm_meido "zaun.moe/vm-meido"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show vm status",
	Run: func(cmd *cobra.Command, args []string) {
	setMeidoFile()
		state, err := vm_meido.LoadState()
		cobra.CheckErr(err)
		state, err = Provider.GetVmState(state.Name)
		cobra.CheckErr(err)
		err = vm_meido.SaveState(*state)
		cobra.CheckErr(err)
		println(state.IpAddress)
		println(state.VmStatus)
		println(state.Name)
		println(state.Provider)
	},
}

