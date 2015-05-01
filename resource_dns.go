package main

import (
	"fmt"
	"net"

	"github.com/whitepages/terraform-provider-dummy/Godeps/_workspace/src/github.com/hashicorp/terraform/helper/schema"
)

func resourceDNS() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSCreate,
		Read:   resourceDNSRead,
		Update: resourceDNSUpdate,
		Delete: resourceDNSDelete,

		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDNSCreate(d *schema.ResourceData, m interface{}) error {
	err := resourceDNSRead(d, m)
	if err != nil {
		return err
	}

	d.SetId(d.Get("host").(string))
	return nil
}

func resourceDNSRead(d *schema.ResourceData, m interface{}) error {
	ip, err := getFirstIP(d.Get("host").(string))
	if err != nil {
		return err
	}

	d.Set("ip_address", ip)

	return nil
}

func resourceDNSUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDNSDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func getFirstIP(host string) (string, error) {
	adds, err := net.LookupIP(host)
	if err != nil {
		return "", err
	}

	if len(adds) != 1 {
		return "", fmt.Errorf("Zero or multiple IP addresses found for %s", host)
	}

	return adds[0].String(), nil
}
