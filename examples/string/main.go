package main

import (
	"fmt"
	"os"

	"github.com/Songbird-Project/scsv"
)

func main() {
	scsvString := `#@valuePriority,false
#@strictMode,false

# Core packages
|core,

# Extra packages
|extra,bat
,eza
,zoxide
,fzf
,fish
,ripgrep
,btop

# Multilib packages
|multilib,

# AUR packages
|aur,
	`

	scsvMap, err := scsv.ParseString(scsvString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for repo, pkgs := range scsvMap {
		if len(pkgs) <= 1 {
			continue
		}

		fmt.Printf("%s:\n", repo)

		for _, pkgColumns := range pkgs {
			for _, pkg := range pkgColumns {
				fmt.Printf(" - %s\n", pkg)
			}
		}
	}
}
