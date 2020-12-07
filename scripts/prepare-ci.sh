#!/usr/bin/env bash

set -euo pipefail

cat << __EOS__ > azure-pipelines.yml
trigger:
- master

resources:
- repo: self

jobs:
- job: Build
  pool:
    vmImage: 'ubuntu-latest'
  steps:
    - template: ci/build-and-push.yml
      parameters:
        problems:
__EOS__

for path in $(ls */Dockerfile); do
    problem_name=$(dirname $path)
    echo "          - '$problem_name'" | tee -a azure-pipelines.yml
done

echo "" | tee -a azure-pipelines.yml

