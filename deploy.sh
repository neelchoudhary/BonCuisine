# Install AWS Command Line Interface

# https://aws.amazon.com/cli/

apk add --update python python-dev py-pip
pip install awscli --upgrade

docker pull $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_NAME

# Set AWS config variables used during the AWS get-login command below

export AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID
export AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY

# AWS login
$(aws ecr get-login --no-include-email --region us-east-2)

# Build new docker image on go-lang-server ECR
docker build -t go-lang-server .

# Tag new docker image on go-lang-server ECR as latest
docker tag go-lang-server:latest 729017073046.dkr.ecr.us-east-2.amazonaws.com/go-lang-server:latest

# Push new docker image to ECR
docker push 729017073046.dkr.ecr.us-east-2.amazonaws.com/go-lang-server:latest

# Register New Task Definition
aws ecs register-task-definition --family boncuisine-production-definition --requires-compatibilities FARGATE --cpu 256 --memory 512 --cli-input-json file://boncuisine-task-definition-production.json --region "us-east-2"

# Update service with new task and start task. This should end old task
aws ecs update-service --cluster golang-cluster --service golang-container-prod-service --task-definition boncuisine-production-definition --region "us-east-2"

# Done
# docker tag $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_NAME $AWS_REGISTRY_IMAGE:$CI_ENVIRONMENT_SLUG
# docker push $AWS_REGISTRY_IMAGE:$CI_ENVIRONMENT_SLUG
# aws ecs register-task-definition --family webcaptioner-$CI_ENVIRONMENT_SLUG	 --requires-compatibilities FARGATE --cpu 256 --memory 512 --cli-input-json file://aws/webcaptioner-task-definition-$CI_ENVIRONMENT_SLUG.json --region \$AWS_REGION
# aws ecs update-service --cluster webcaptioner-$CI_ENVIRONMENT_SLUG --service webcaptioner --task-definition webcaptioner-$CI_ENVIRONMENT_SLUG --region \$AWS_REGION
