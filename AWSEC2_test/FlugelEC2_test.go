package test

import (
	"fmt"
	"testing"
    
		"github.com/gruntwork-io/terratest/modules/aws"
	//"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	//test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestawsEC2(t *testing.T) {
	t.Parallel()
	awsRegion := aws.GetRandomStableRegion(t, nil, nil)
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
	TerraformDir: "../awsEC2",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	instanceID := terraform.Output(t, terraformOptions, "instance_id")

	expectedName := fmt.Sprintf("aws")
        expectedOwner := fmt.Sprintf("InfraTeam")

	// Look up the tags for the given Instance ID
	instanceTags := aws.GetTagsForEc2Instance(t, awsRegion, instanceID)

	// Verify that our expected name tag is one of the tags
	nameTag, containsNameTag := instanceTags["Name"]
	assert.True(t, containsNameTag)
	assert.Equal(t, expectedName, nameTag)
	
	ownerTag, containsOwnerTag := instanceTags["Owner"]
	assert.True(t, containsOwnerTag)
	assert.Equal(t, expectedOwner, ownerTag)
}
