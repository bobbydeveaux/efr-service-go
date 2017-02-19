# deploy.sh
#! /bin/bash

SHA1=$1

aws s3 cp ./release/efr-service-go.zip "s3://elasticbeanstalk-eu-west-1-696293867939/efr-service-go.zip"

aws elasticbeanstalk create-application-version --region=eu-west-1 --application-name "My First Elastic Beanstalk Application" \
  --version-label $SHA1 --source-bundle S3Bucket="elasticbeanstalk-eu-west-1-696293867939",S3Key="efr-service-go.zip"


# Update Elastic Beanstalk environment to new version
aws elasticbeanstalk update-environment --region=eu-west-1 --environment-name Default-Environment \
    --version-label $SHA1
