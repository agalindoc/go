package test

import (
	"fmt"
	"testing"
    
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAWSS3Bucket(t *testing.T) {
	t.Parallel()
	awsRegion :=  fmt.Sprintf("us-east-2")
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
	TerraformDir: "../AWSS3Bucket",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	bucketID := terraform.Output(t, terraformOptions, "bucket_id")

	expectedBucket := fmt.Sprintf("TFStateLogs/")

	// Look up the tags for the given bucket ID
	bucketObjectID := aws.GetS3BucketLoggingTargetPrefix(t, awsRegion, bucketID)

	// Verify that our expected name tag is one of the tags
	assert.Equal(t, expectedBucket,  bucketObjectID)
	
}
