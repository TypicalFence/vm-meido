package cmd

import (
	"github.com/spf13/cobra"
	vm_meido "zaun.moe/vm-meido"
)

func init() {
	rootCmd.AddCommand(destoryCmd)
}

var destoryCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroys a provisioned vm",
	Run: func(cmd *cobra.Command, args []string) {
		setMeidoFile()
		state, err := vm_meido.LoadState()
		cobra.CheckErr(err)		
		_, err = Provider.DestroyVm(state.Name)
		cobra.CheckErr(err)
		state.VmStatus = vm_meido.VM_STATUS_GONE
		vm_meido.SaveState(*state)
	},
}
