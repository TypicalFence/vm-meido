package cmd

import (
	"github.com/spf13/cobra"
	vm_meido "zaun.moe/vm-meido"
)

func setMeidoFile() {
	meidoFile, err := vm_meido.LoadMeidoFile()
	cobra.CheckErr(err)
	MeidoFile = meidoFile
}
