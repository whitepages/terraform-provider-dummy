# Terraform dummy provider [![Build Status](https://secure.travis-ci.org/whitepages/terraform-provider-dummy.png)](http://travis-ci.org/whitepages/terraform-provider-dummy)

`terraform-provider-dummy` is a dummy Terraform provider. It does not
require any configuration, and does not interact with any external
providers.

## Resources

### dns

Provides a dummy DNS resource.

* `host` - (Required) The address of the resource
* `ip_address` - (Calculated) The IP address for the host

### Example

```
resource "dummy_dns" "host" {
	host = "example.com"
}
```

```sh
$ terraform show
dummy_dns.host:
  id = example.ex
  host = example.ex
  ip_address = 127.0.0.2
```

### server

Provides a dummy server resource.

* `address` - (Required) The address of the resource

### Example

```
resource "dummy_server" "example" {
	address = "host${count.index}.example.com"
	count = 3
}
```

```sh
$ terraform plan
+ dummy_server.example.0
    address: "" => "host0.example.com"

+ dummy_server.example.1
    address: "" => "host1.example.com"

+ dummy_server.example.2
    address: "" => "host2.example.com"
```
