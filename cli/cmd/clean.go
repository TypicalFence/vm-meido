package cmd

import (
	"os"
	"path"

	"github.com/spf13/cobra"
	vm_meido "zaun.moe/vm-meido"
)

func init() {
	rootCmd.AddCommand(cleanCmd)
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "cleans the folder of state files",
	Run: func(cmd *cobra.Command, args []string) {
		setMeidoFile()
		state, err := vm_meido.LoadState()
		cobra.CheckErr(err)
		if state.VmStatus == vm_meido.VM_STATUS_GONE {
			cwd, err := os.Getwd()
			if err != nil {
				println(err)
				os.Exit(1)
			}

			meidoDir := path.Join(cwd, ".vm-meido")

			os.RemoveAll(meidoDir)
		} else {
			println("please destroy the vm first.")
			os.Exit(1)
		}
	},
}
