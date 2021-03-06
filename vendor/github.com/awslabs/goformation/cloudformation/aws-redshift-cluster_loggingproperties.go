package cloudformation

// AWSRedshiftCluster_LoggingProperties AWS CloudFormation Resource (AWS::Redshift::Cluster.LoggingProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html
type AWSRedshiftCluster_LoggingProperties struct {

	// BucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-bucketname
	BucketName *StringIntrinsic `json:"BucketName,omitempty"`

	// S3KeyPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-s3keyprefix
	S3KeyPrefix *StringIntrinsic `json:"S3KeyPrefix,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRedshiftCluster_LoggingProperties) AWSCloudFormationType() string {
	return "AWS::Redshift::Cluster.LoggingProperties"
}
