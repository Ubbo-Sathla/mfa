package main

import (
	"github.com/Ubbo-Sathla/mfa/pkg/mfa"
	"github.com/olekukonko/tablewriter"

	"os"
)

func main() {
	err := mfa.LoadConfig()
	if err != nil {
		return
	}
	c := mfa.GetConfig()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Issuer", "Account", "Code"})

	for _, j := range c {
		j.Load()
		table.Append([]string{j.Name, j.Issuer, j.AccountName, j.GenerateCode()})
	}
	table.Render()
}
