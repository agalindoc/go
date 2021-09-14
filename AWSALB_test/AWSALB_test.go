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

func TestawsALB(t *testing.T) {
	t.Parallel()
	awsRegion := aws.GetRandomStableRegion(t, nil, nil)
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
	TerraformDir: "../awsALB",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	instanceID0 := terraform.Output(t, terraformOptions, "instance_id0")
	instanceID1 := terraform.Output(t, terraformOptions, "instance_id1")

	expectedName0 := fmt.Sprintf("aws-0")
        expectedOwner0 := fmt.Sprintf("InfraTeam-0")
       	expectedName1 := fmt.Sprintf("aws-1")
        expectedOwner1 := fmt.Sprintf("InfraTeam-1")
        

	// Look up the tags for the given Instance ID
	instanceTags := aws.GetTagsForEc2Instance(t, awsRegion, instanceID0)

	// Verify that our expected name tag is one of the tags
	nameTag, containsNameTag := instanceTags["Name"]
	assert.True(t, containsNameTag)
	assert.Equal(t, expectedName0, nameTag)
	
	ownerTag, containsOwnerTag := instanceTags["Owner"]
	assert.True(t, containsOwnerTag)
	assert.Equal(t, expectedOwner0, ownerTag)
	
	// Look up the tags for the given Instance ID
	instanceTags1 := aws.GetTagsForEc2Instance(t, awsRegion, instanceID1)

	// Verify that our expected name tag is one of the tags
	nameTag1, containsNameTag1 := instanceTags1["Name"]
	assert.True(t, containsNameTag1)
	assert.Equal(t, expectedName1, nameTag1)
	
	ownerTag1, containsOwnerTag1 := instanceTags1["Owner"]
	assert.True(t, containsOwnerTag1)
	assert.Equal(t, expectedOwner1, ownerTag1)
	
}
