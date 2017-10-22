package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/puneetk/terraform-provider-artifactorymc/artifactorymc"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: artifactorymc.Provider})
}
