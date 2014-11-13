# Terraform dummy provider

`terraform-provider-dummy` is a dummy Terraform provider. It does not
require any configuration, and does not interact with any external
providers.

## Example

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

## Resources

### server

Provides a dummy server resource.

* `address` - (Required) The address of the resource
