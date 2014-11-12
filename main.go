package main

import (
	"github.com/svend/terraform-provider-dummy/Godeps/_workspace/src/github.com/hashicorp/terraform/plugin"
	"github.com/svend/terraform-provider-dummy/Godeps/_workspace/src/github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return Provider()
		},
	})
}
