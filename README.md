# terraform-aws-panos-firewall

[![terraform validate](https://github.com/mrichardson03/terraform-aws-panos-firewall/workflows/terraform%20validate/badge.svg)](https://github.com/mrichardson03/terraform-aws-panos-firewall/actions?query=workflow%3A%22terraform+validate%22)

Terraform Module: PAN-OS firewall connecting two AWS subnets.

This Terraform module creates a PAN-OS firewall between a public and a private subnet in an
AWS VPC.  The configuration is based off of the
[AWS Deployment Guide - Single VPC Model](https://www.paloaltonetworks.com/apps/pan/public/downloadResource?pagePath=/content/pan/en_US/resources/guides/aws-deployment-guide-single-resource)
reference architecture.

## Usage

Include in a Terraform plan (see [mrichardson03/terraform-aws-panos-bootstrap](https://github.com/mrichardson03/terraform-aws-panos-bootstrap) for easy bootstrapping):

```terraform
module "firewall" {
  source  = "github.com/mrichardson03/terraform-aws-panos-firewall?ref=v1.0.0"

  vpc_id   = module.vpc.vpc_id
  key_name = var.key_name

  mgmt_subnet_id = module.vpc.mgmt_a_id
  mgmt_sg_id     = module.vpc.mgmt_sg_id
  mgmt_ip        = "10.1.9.21"

  eth1_subnet_id = module.vpc.public_a_id
  eth1_sg_id     = module.vpc.public_sg_id
  eth1_ip        = "10.1.10.10"

  eth2_subnet_id = module.vpc.web_a_id
  eth2_sg_id     = module.vpc.internal_sg_id
  eth2_ip        = "10.1.1.10"

  iam_instance_profile = module.bootstrap.instance_profile_name
  bootstrap_bucket     = module.bootstrap.bucket_name
}
```

### Required Inputs

`vpc_id`: VPC ID to create firewall instance in.

`key_name`: Key pair name to provision instances with.

`mgmt_subnet_id`: Subnet ID for firewall management interface.

`mgmt_ip`: Internal IP address for firewall management interface.

`mgmt_sg_id`: Security group ID for firewall management interface.

`eth1_subnet_id`: Subnet ID for firewall ethernet1/1 interface.

`eth1_ip`: Internal IP address for firewall ethernet1/1 interface.

`eth1_sg_id`: Security group ID for firewall ethernet1/1 interface.

`eth2_subnet_id`: Subnet ID for firewall ethernet1/2 interface.

`eth2_ip`: Internal IP address for firewall ethernet1/2 interface.

`eth2_sg_id`: Security group ID for firewall ethernet1/2 interface.

### Optional Inputs

`ami`: Firewall AMI in specified region.  Default is 9.0.3.xfr BYOL in us-east-1.

`instance_type`: Instance type for firewall.

`iam_instance_profile`: IAM Instance Profile used to bootstrap firewall.

`bootstrap_bucket`: S3 bucket containing bootstrap configuration.

`tags`: A map of tags to add to all resources.

### Outputs

`instance_id`: Instance ID of created firewall.

`mgmt_public_ip`: Public IP address of firewall management interface.

`mgmt_interface_id`: Interface ID of created firewall management interface.

`eth1_public_ip`: Public IP address of firewall ethernet1/1 interface.

`eth1_interface_id`: Interface ID of created firewall ethernet1/1 interface.

`eth2_interface_id`: Interface ID of created firewall ethernet1/2 interface.

## Changelog

**v0.9.0** - Initial private release.

## Support Policy

The code and templates in the repo are released under an as-is, best effort,
support policy. These scripts should be seen as community supported and
Palo Alto Networks will contribute our expertise as and when possible.
We do not provide technical support or help in using or troubleshooting the
components of the project through our normal support options such as
Palo Alto Networks support teams, or ASC (Authorized Support Centers)
partners and backline support options. The underlying product used
(the VM-Series firewall) by the scripts or templates are still supported,
but the support is only for the product functionality and not for help in
deploying or using the template or script itself. Unless explicitly tagged,
all projects or work posted in our GitHub repository
(at https://github.com/PaloAltoNetworks) or sites other than our official
Downloads page on https://support.paloaltonetworks.com are provided under
the best effort policy.