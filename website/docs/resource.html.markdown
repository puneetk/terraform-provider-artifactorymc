---
layout: "null"
page_title: "Artifactory MC Resource"
sidebar_current: "docs-artifactorymc-resource"
description: |-
  A resource that does nothing.
---

# Artifactory MC Resource

The `artifactorymc_resource` resource implements the standard resource lifecycle but
takes no further action.

The `triggers` argument allows specifying an arbitrary set of values that,
when changed, will cause the resource to be replaced.

## Example Usage

The primary use-case for the artifactorymc resource is as a do-nothing container for
arbitrary actions taken by a provisioner, as follows:

```hcl
resource "aws_instance" "cluster" {
  count = 3

  # ...
}

resource "artifactorymc_resource" "cluster" {
  # Changes to any instance of the cluster requires re-provisioning
  triggers = {
    cluster_instance_ids = "${join(",", aws_instance.cluster.*.id)}"
  }

  # Bootstrap script can run on any instance of the cluster
  # So we just choose the first in this case
  connection {
    host = "${element(aws_instance.cluster.*.public_ip, 0)}"
  }

  provisioner "local-exec" {
    # Bootstrap script called with private_ip of each node in the clutser
    command = "bootstrap-cluster.sh ${join(" ", aws_instance.cluster.*.private_ip)}"
  }
}
```

In this example, three EC2 instances are created and then a
`artifactorymc_resource` instance is used to gather data about all three and execute
a single action that affects them all. Due to the `triggers` map, the
`artifactorymc_resource` will be replaced each time the instance ids change, and thus
the `remote-exec` provisioner will be re-run.


## Argument Reference

The following arguments are supported:

* `triggers` - (Optional) A map of arbitrary strings that, when changed, will
  force the artifactorymc resource to be replaced, re-running any associated
provisioners.

## Attributes Reference

The following attributes are exported:

* `id` - An arbitrary value that changes each time the resource is replaced.
  Can be used to cause other resources to be updated or replaced in response
  to `artifactorymc_resource` changes.
