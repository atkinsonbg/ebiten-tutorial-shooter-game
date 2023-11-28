```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "rds:FailoverGlobalCluster",
                "rds:DescribeGlobalClusters"
            ],
            "Resource": [
                "arn:aws:rds:region:account-id:global-cluster:global-cluster-id",
                "arn:aws:rds:region:account-id:db:db-instance-id"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "route53:ChangeResourceRecordSets",
                "route53:ListResourceRecordSets"
            ],
            "Resource": [
                "arn:aws:route53:::hostedzone/zone-id"
            ]
        }
    ]
}
```
