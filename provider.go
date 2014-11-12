package main

import (
	"github.com/svend/terraform-provider-dummy/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/schema"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"dummy_server": resourceServer(),
		},
	}
}
