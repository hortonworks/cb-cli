#!/bin/bash
# -e  Exit immediately if a command exits with a non-zero status.
# -x  Print commands and their arguments as they are executed.

: ${CB_VERSION:=latest}
: ${CB_TARGET_BRANCH:=master}
: ${BASE_URL:=https://localhost}
: ${STOP_MOCK:=false}

determine-versions() {
    declare desc="Determine the MOCK version based on Swagger availability"

    echo "Checking latest available MOCK for"${CB_VERSION}
    if [[ -n "$GITHUB_TOKEN" ]]; then
        access_token="access_token=$GITHUB_TOKEN"
    fi
    BRANCH_VERSION=$(echo $CB_VERSION | cut -d'-' -f1)
    versions=($(curl https://api.github.com/repos/hortonworks/cloudbreak-deployer/git/refs/tags\?$access_token | jq -rc ".[] | select(.ref | contains(\"$BRANCH_VERSION\")) | .ref"))

    if [[ -z $versions ]]; then
        echo "Can not fetch cbd versions from github!"
        exit 1
    fi

    for ref_ver in "${versions[@]}"
    do
        ver=${ref_ver#"refs/tags/"}
        if [[ $ver < $CB_VERSION || $ver = $CB_VERSION ]]; then
            highest_built_version=$ver
        fi
    done

    echo "Latest available mock: $highest_built_version"

    if [[ -z $highest_built_version ]]; then
        echo "Can not determine highest built version!"
        exit 1
    fi
    CB_VERSION=$highest_built_version
}

download-swagger-jsons() {
    declare desc="Download micro services Swagger JSONs from S3 if necessary"

    if [[ -z $CLOUDBREAK_ADDRESS ]]; then
        echo "$(tput setaf 3)CLOUDBREAK_ADDRESS environment variable is not defined."
        echo "So the swagger-${CB_VERSION}.json is going to be downloaded from S3.$(tput sgr 0)"

        curl -k https://s3-eu-central-1.amazonaws.com/cloudbreak-swagger/swagger-"$CB_VERSION".json -o swagger-cloudbreak.json
        curl -k https://s3-us-east-2.amazonaws.com/environment-swagger/swagger-"$CB_VERSION".json -o swagger-environment.json
        curl -k https://s3-us-east-2.amazonaws.com/datalake-swagger/swagger-"$CB_VERSION".json -o swagger-datalake.json
        curl -k https://s3-us-east-2.amazonaws.com/redbeams-swagger/swagger-"$CB_VERSION".json -o swagger-redbeams.json
        curl -k https://s3-us-east-2.amazonaws.com/freeipa-swagger/swagger-"$CB_VERSION".json -o swagger-freeipa.json
        curl -k https://s3-us-east-2.amazonaws.com/autoscale-swagger/swagger-"$CB_VERSION".json -o swagger-autoscale.json
    else
        echo "$(tput setaf 2)CLOUDBREAK_ADDRESS environment variable is defined."
        echo "So the ${CLOUDBREAK_ADDRESS} swagger.json is supposed to already been downloaded to 'integration-tests' folder$(tput sgr 0)"
    fi
}

get-cbd() {
    declare desc="Downloading CBD binary"
    local os=$(uname -s)
    local latest_tag=$(curl "http://release.infra.cloudera.com/hwre-api/listbuilds?stack=CB&release=${CB_VERSION}&type=dev" | jq -r '.latest_build_version')
    local url=$(echo "https://public-repo-1.hortonworks.com/HDP/cloudbreak/cloudbreak-deployer_${latest_tag}_${os}_x86_64.tgz")
    local dest=.

    if curl -Ls $url | tar -xz -C ${dest}; then
        echo "CBD has been downloaded with version: $CB_VERSION"
    else
        echo "$(tput setaf 3) Building ${CB_TARGET_BRANCH} CBD for ${os}$(tput sgr 0)"

        docker volume rm cbd-source 2>/dev/null || :
        docker volume create --name cbd-source 1>/dev/null
        docker run --rm --user="0" --entrypoint init.sh -v cbd-source:/var/workspace jpco/git:1.0 https://github.com/hortonworks/cloudbreak-deployer.git ${CB_TARGET_BRANCH} cloudbreak-deployer
        docker run --user="0" -e GOPATH=/usr -v cbd-source:/usr/src/github.com/hortonworks -w /usr/src/github.com/hortonworks/cloudbreak-deployer golang:1.12 make bindata
        docker run --rm --user="0" -e GOPATH=/usr -v cbd-source:/usr/src/github.com/hortonworks -w /usr/src/github.com/hortonworks/cloudbreak-deployer golang:1.12 make build
        docker run --rm --user="0" -v cbd-source:/var/workspace jpco/git:1.0 cat /var/workspace/cloudbreak-deployer/build/${os}/cbd > cbd
        docker volume rm cbd-source 1>/dev/null || :

        chmod +x cbd
    fi
    mkdir etc
    echo "TEST_LICENSE" > etc/license.txt

    echo "CBD installed into "$(pwd)
    echo "CBD has been found with version: $(./cbd --version)"
}

mock-start() {
    declare desc="Start Cloudbreak Mock"

    mv docker-compose.yml docker-compose-mocks.yml
    echo "Starting CBD with minimal set:"
    cat <<EOF > Profile
export PUBLIC_IP=$(hostname -i)
export CB_LOCAL_DEV_LIST=cloudbreak,periscope,datalake,freeipa,redbeams,environment,uluwatu,cluster-proxy,core-gateway
export VAULT_AUTO_UNSEAL=true
export COMMON_DB_VOL=mock-common
export DOCKER_NETWORK_NAME=mock-cbreak-network
export CB_ENABLEDPLATFORMS=AZURE,OPENSTACK,AWS,GCP,YARN,MOCK
export ENVIRONMENT_ENABLEDPLATFORMS=AZURE,OPENSTACK,AWS,GCP,YARN,MOCK
export UMS_HOST=''
EOF
    ./cbd generate
    ./cbd start
    ./.deps/bin/docker-compose -f docker-compose-mocks.yml -p cbreak up -d
    echo "Mock APIs version is: " $(./cbd version)

    sleep 30s
}

mock-start-logs() {
    declare desc="Gather Cloudbreak Mock start logs"

    mkdir -pv test-logs

    if [[ $(id -u jenkins 2>/dev/null || echo $?) -gt 1 ]]; then
        sudo chown -R jenkins .
        sudo docker logs cbreak_cloudbreak_1 > test-logs/cloudbreak_start.log
        sudo docker logs cbreak_environment_1 > test-logs/environment_start.log
        sudo docker logs cbreak_datalake_1 > test-logs/datalake_start.log
        sudo docker logs cbreak_redbeams_1 > test-logs/redbeams_start.log
    else
        docker logs cbreak_cloudbreak_1 > test-logs/cloudbreak_start.log
        docker logs cbreak_environment_1 > test-logs/environment_start.log
        docker logs cbreak_datalake_1 > test-logs/datalake_start.log
        docker logs cbreak_redbeams_1 > test-logs/redbeams_start.log
    fi
}

mock-stop() {
  declare desc="Stop Cloudbreak Mock"

	./.deps/bin/docker-compose -f docker-compose.yml -f docker-compose-mocks.yml -p cbreak down --volumes
	./cbd kill
}

main() {
    cd tmp
    if [[ "$STOP_MOCK" == "true" ]]; then
        mock-stop
        cd ..
    else
        determine-versions
        download-swagger-jsons
        get-cbd
        mock-start
        cd ..
        mock-start-logs
    fi
}

main "$@"