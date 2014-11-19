package main

import (
	"github.com/svend/terraform-provider-dummy/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/schema"
	"github.com/svend/terraform-provider-dummy/Godeps/_workspace/src/github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"dummy_dns":    resourceDNS(),
			"dummy_server": resourceServer(),
		},
	}
}
