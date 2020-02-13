#!/usr/bin/env bash

: ${GIT_VERSION:=latest}
: ${BASE_URL:=https://127.0.0.1}

compose-init() {
    declare desc="Installing Docker-Compose 1.21.2 if it is no present on machine"

    local required_compose=1.21.2
    local compose_version=$(docker-compose --version 2>&1 | grep -E -o "([0-9]+\.)+[0-9]+")
    local compose_binary=$(which docker-compose)

    if [[ -z "${compose_version}" ]]; then
      echo $TERM
      echo "$(tty -s && tput setaf 3)* Docker-Compose required, installing with ${required_compose} version...$(tty -s && tput sgr 0)"
      sudo curl -L "https://github.com/docker/compose/releases/download/${required_compose}/docker-compose-$(uname -s)-$(uname -m)"  -o /usr/local/bin/docker-compose
      sudo mv /usr/local/bin/docker-compose /usr/bin/docker-compose
      sudo chmod +x /usr/bin/docker-compose
    else
      echo $TERM
      echo "$(tty -s && tput setaf 3)* Already installed Compose version: ${compose_version}"
      echo "* Already installed Compose binary: ${compose_binary}$(tty -s && tput sgr 0)"
    fi
}

mock-image-tag() {
    declare desc="Set Cloudbreak Mock Docker image tag"

    echo "GitHub First Parent Tag is: ${GIT_VERSION}"
    echo "CircleCI Branch is: ${CIRCLE_BRANCH}"
    echo "CircleCI Tag is: ${CIRCLE_TAG}"

    if [[ "${GIT_VERSION}" != *"-rc."* ]]; then
        export GIT_VERSION=latest
    fi
}

mock-start-logs() {
    declare desc="Gather Cloudbreak Mock start logs"

    mkdir -pv test_log

    if [[ $(id -u jenkins 2>/dev/null || echo $?) -gt 1 ]]; then
        sudo chown -R jenkins .
        sudo docker logs cbreak_cloudbreak_1 > test_log/cloudbreak_start.log
    else
        docker logs cbreak_cloudbreak_1 > test_log/cloudbreak_start.log
    fi
}

mock-start() {
    declare desc="Start Cloudbreak Mock"

    docker-compose -f tmp/docker-compose.yml -p cbreak up -d
    sleep 30s
}

main() {
    compose-init
    mock-image-tag
    mock-start
    mock-start-logs
}

main "$@"