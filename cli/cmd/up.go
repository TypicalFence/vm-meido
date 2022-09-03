package cmd

import (
	"log"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	vm_meido "zaun.moe/vm-meido"
)

func init() {
	rootCmd.AddCommand(upCmd)
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "provisions the desired vm",
	Run: func(cmd *cobra.Command, args []string) {
		setMeidoFile()
		state, _ := vm_meido.LoadState()

		if state != nil && state.VmStatus != vm_meido.VM_STATUS_GONE {
			print(state.Name)
			print(state.IpAddress)
			os.Exit(0)
		}

		name, err := Provider.ProvisionVm(vm_meido.VmSettings{Size: MeidoFile.Size, Distro: MeidoFile.Distro, Region: Config.Region})

		if err != nil {
			log.Fatal(err)
			os.Exit(1)

		}
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()

		for {
			time.Sleep(15 * time.Second)

			status, err := Provider.GetVmStatus(name)

			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			if status == vm_meido.VM_STATUS_RUNNING {
				break
			}
		}

		s.Stop()
		println(name)
		state, err = Provider.GetVmState(name)
		cobra.CheckErr(err)

		println(state.IpAddress)

		vm_meido.SaveState(*state)
	},
}
