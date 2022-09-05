package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	vm_meido "zaun.moe/vm-meido"
)

func init() {
	rootCmd.AddCommand(sshCmd)
}

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "connect to vm via ssh",
	Run: func(cmd *cobra.Command, args []string) {
		setMeidoFile()
		state, err := vm_meido.LoadState()
		cobra.CheckErr(err)
		if state.VmStatus == vm_meido.VM_STATUS_GONE {
			println("vm gone")
			os.Exit(1)
		}
		state, err = Provider.GetVmState(state.Name)
		cobra.CheckErr(err)

		if state.VmStatus == vm_meido.VM_STATUS_RUNNING {
			cmd := exec.Command("ssh", "root@"+state.IpAddress)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			_ = cmd.Run()
		}
	},
}
