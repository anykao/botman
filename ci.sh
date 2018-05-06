#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'


GIT_COMMIT_MESSAGE="daily update"
GIT_REPO="anykao/botman"
GIT_BRANCH="master"


if [[ -z "${TRAVIS+x}" ]]; then
    echo "Error: this script is meant to be run on Travis CI."
    exit 1
fi


if [[ -z "${GITHUB_TOKEN}" ]]; then
    echo "Error: the \$GITHUB_TOKEN environment variable is not set!"
    exit 1
fi

if [[ -z "${GIT_NAME}" ]]; then
    echo "Error: the \$GIT_NAME environment variable is not set!"
    exit 1
fi

if [[ -z "${GIT_EMAIL}" ]]; then
    echo "Error: the \$GIT_EMAIL environment variable is not set!"
    exit 1
fi

git checkout "${GIT_BRANCH}"
go run main.go

git status
git add *.md
git -c "commit.gpgsign=false" \
    -c "user.name=${GIT_NAME}" \
    -c "user.email=${GIT_EMAIL}" \
    commit -m "${GIT_COMMIT_MESSAGE}"
git push "https://x-token:${GITHUB_TOKEN}@github.com/${GIT_REPO}" "${GIT_BRANCH}"
