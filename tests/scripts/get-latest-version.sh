#!/bin/bash
# -e  Exit immediately if a command exits with a non-zero status.
# -x  Print commands and their arguments as they are executed.

: ${CB_TARGET_BRANCH:=master}

get-latest-version() {
    if [[ -z $CB_VERSION ]]; then
        if [[ $CB_TARGET_BRANCH = master* ]]; then
            export LATEST_VERSION=$(curl "http://release.infra.cloudera.com/hwre-api/getreleaseversion?stack=CB&releaseline=master" | jq -r '.version')
            export CB_VERSION=$(curl "http://release.infra.cloudera.com/hwre-api/listbuilds?stack=CB&release=${LATEST_VERSION}&type=dev" | jq -r '.latest_build_version')
        elif [[ $CB_TARGET_BRANCH = rc-* ]]; then
            export BRANCH_VERSION=$(echo $CB_TARGET_BRANCH | cut -f 2 -d '-')
            export CB_VERSION=$(curl "http://release.infra.cloudera.com/hwre-api/listbuilds?stack=CB&release=${BRANCH_VERSION}.0&type=rc" | jq -r '.latest_build_version')
        elif [[ $CB_TARGET_BRANCH = CB-* ]]; then
            export BRANCH_VERSION=$(echo $CB_TARGET_BRANCH | cut -f 2 -d '-')
            export CB_VERSION=$(curl "http://release.infra.cloudera.com/hwre-api/listbuilds?stack=CB&release=${BRANCH_VERSION}&type=cb" | jq -r '.latest_build_version')
        else
            export BRANCH_VERSION=$(echo $CB_TARGET_BRANCH | cut -f 2 -d '-')
            export CB_VERSION=$(curl "http://release.infra.cloudera.com/hwre-api/listbuilds?stack=CB&release=${BRANCH_VERSION}.0" | jq -r '.latest_build_version')
        fi
    fi
    echo $CB_VERSION
}

main() {
    get-latest-version
}

main "$@"