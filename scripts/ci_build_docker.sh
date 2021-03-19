#!/bin/bash
set -ex

AWS_ECR_REGISTRY=${AWS_ECR_REGISTRY}
ECR_REPO=${ECR_REPO:-heycar-mgmt-citools}

echo "logging into ecr"

aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin ${AWS_ECR_REGISTRY}

docker_build () {
    docker build -t ${AWS_ECR_REGISTRY}/${ECR_REPO}:$CIRCLE_SHA1 . 
    docker tag ${AWS_ECR_REGISTRY}/${ECR_REPO}:$CIRCLE_SHA1 ${AWS_ECR_REGISTRY}/${ECR_REPO}:semver-latest
}

docker_build_push () {
    docker_build
    docker push ${AWS_ECR_REGISTRY}/${ECR_REPO}:$CIRCLE_SHA1 && \
    docker push ${AWS_ECR_REGISTRY}/${ECR_REPO}:semver-latest
}

(
    "$@"
)