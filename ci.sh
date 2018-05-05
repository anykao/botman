#!/bin/bash

set -euo pipefail
IFS=$'\n\t'


GIT_COMMIT_MESSAGE="Automatic stats update"
GIT_EMAIL="you.got@me.com"
GIT_NAME="anykao"
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

git checkout "${GIT_BRANCH}"
go run main.go


if git diff --quiet data/; then
    echo "No changes to commit."
else
    git status
    git add *.md
    git -c "commit.gpgsign=false" \
        -c "user.name=${GIT_NAME}" \
        -c "user.email=${GIT_EMAIL}" \
        commit -m "${GIT_COMMIT_MESSAGE}"
    git push "https://x-token:${GITHUB_TOKEN}@github.com/${GIT_REPO}" "${GIT_BRANCH}"
fi