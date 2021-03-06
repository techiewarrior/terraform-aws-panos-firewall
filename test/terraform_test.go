package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraform(t *testing.T) {
	t.Parallel()

	// Pick a random, stable (older than one year) AWS region to test in.
	awsRegion := aws.GetRandomStableRegion(t, nil, nil)

	// Pick a random PAN-OS license type and version.
	rand.Seed(time.Now().UnixNano())
	licenses := []string{"byol", "bundle1", "bundle2"}
	versions := []string{"9.1", "10.0"}
	create_eips := []bool{true, false}

	license := licenses[rand.Intn(len(licenses))]
	version := versions[rand.Intn(len(versions))]
	create_eip := create_eips[rand.Intn(len(create_eips))]

	// Create temporary SSH key in that region to test with.
	keyPairName := fmt.Sprintf("terratest-ssh-%s", random.UniqueId())
	keyPair := aws.CreateAndImportEC2KeyPair(t, awsRegion, keyPairName)

	terraformOptions := &terraform.Options{
		// Use the Terraform plans in this directory
		TerraformDir: ".",

		// Add variables to Terraform run (like they were specified on the Terraform CLI via -var).
		Vars: map[string]interface{}{
			"aws_region":    awsRegion,
			"key_name":      keyPairName,
			"panos_license": license,
			"panos_version": version,
			"create_mgmt_eip": create_eip,
			"create_eth1_eip": create_eip,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created,
	// and then clean up our key pair.
	defer func() {
		terraform.Destroy(t, terraformOptions)
		aws.DeleteEC2KeyPair(t, keyPair)
	}()

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)
}
