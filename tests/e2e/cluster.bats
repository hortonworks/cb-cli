#!/usr/bin/env bash

load ../utils/commands
load ../utils/resources

DELAY=$(($SECONDS+2100))
REINSTALL_TEMPFILE="reinstall-template.json"
STATUS_TEMPFILE="status_reason"
UTILS_TEMPFILE="clitestutil"

@test "SETUP: Cleanup stuck OpenStack ["${OS_CLUSTER_NAME}"] cluster" {
  run remove-stuck-cluster "${OS_CLUSTER_NAME}"

  echo "$output" >&2

  [ $status -eq 0 ]
}

@test "SETUP: Cleanup stuck OpenStack ["${OS_CREDENTIAL_NAME}"cluster] credential" {
  run remove-stuck-credential "${OS_CREDENTIAL_NAME}cluster"

  echo "$output" >&2

  [ $status -eq 0 ]
}

@test "SETUP: Create new ["${OS_CREDENTIAL_NAME}"cluster] OpenStack V2 credential" {
  OUTPUT=$(create-credential-openstack-v2 --name "${OS_CREDENTIAL_NAME}cluster" $OS_ARGS_V2 2>&1 | tail -n 2 | head -n 1)

  echo "${OUTPUT}" >&2

  [[ "${OUTPUT}" == *"credential created: ${OS_CREDENTIAL_NAME}cluster"* ]]
  [[ "${OUTPUT}" != *"error"* ]]
}

@test "Create new ["${OS_CLUSTER_NAME}"] OpenStack cluster" {
  OUTPUT=$(create-cluster --name "${OS_CLUSTER_NAME}" --cli-input-json $OS_INPUT_JSON_FILE 2>&1 | tail -n 2 | head -n 1)

  echo "${OUTPUT}" >&2

  [[ "${OUTPUT}" == *"stack created: ${OS_CLUSTER_NAME}"* ]]
  [[ "${OUTPUT}" != *"error"* ]]
}

@test "Wait for ["${OS_CLUSTER_NAME}"] cluster is created" {
  run wait-cluster-status $DELAY "${OS_CLUSTER_NAME}" "AVAILABLE"

  echo "$output" >&2

  [ $status -eq 0 ]
  [ "$output" = "true" ]
}

@test "Change Ambari password for ["${OS_CLUSTER_NAME}"] cluster" {
  run is-cluster-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is NOT created yet!"
  fi

  OUTPUT=$(change-ambari-password --name "${OS_CLUSTER_NAME}" --old-password admin --new-password 4321 --ambari-user admin 2>&1 | awk '{printf "%s",$0} END {print ""}' | grep -o '{.*}' | jq ' . |  [to_entries[].key] == ["oldPassword","password","userName"]')

  [[ "${OUTPUT}" ==  true ]]
}

@test "["${OS_CLUSTER_NAME}"] cluster should be listed" {
  run is-cluster-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is NOT created yet!"
  fi

  for OUTPUT in $(list-clusters | jq ' .[] | [to_entries[].key] == ["Name","Description","CloudPlatform","StackStatus","ClusterStatus"]');
  do
    [[ "$OUTPUT" == "true" ]]
  done
}

@test "["${OS_CLUSTER_NAME}"] cluster should be described" {
  run is-cluster-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is NOT created yet!"
  fi

  OUTPUT=$(describe-cluster --name "${OS_CLUSTER_NAME}" | jq '. "name" == "${OS_CLUSTER_NAME}"')

  [[ "$OUTPUT" == "true" ]]
}

@test "["${OS_CLUSTER_NAME}"] cluster can be stop" {
  run is-cluster-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is NOT created yet!"
  fi

  OUTPUT=$(stop-cluster --name "${OS_CLUSTER_NAME}")

  echo $OUTPUT >&2

  [[ "${OUTPUT}" != *"error"* ]]
}

@test "["${OS_CLUSTER_NAME}"] cluster should be stopped" {
  run is-cluster-status "${OS_CLUSTER_NAME}" "STOP_IN_PROGRESS"
  if [[ "$output" != "true" ]]; then
    skip "Cluster Stop has not been requested!"
  fi

  run wait-cluster-status $DELAY "${OS_CLUSTER_NAME}" "STOPPED"

  echo "$output" >&2

  [ $status -eq 0 ]
  [ "$output" = "true" ]
}

@test "["${OS_CLUSTER_NAME}"] cluster can be start" {
  run is-cluster-status "${OS_CLUSTER_NAME}" "STOPPED"
  if [[ "$output" != "true" ]]; then
    skip "Cluster has not been stopped"
  fi

  OUTPUT=$(start-cluster --name "${OS_CLUSTER_NAME}")

  echo $OUTPUT >&2

  [[ "${OUTPUT}" != *"error"* ]]
}

@test "["${OS_CLUSTER_NAME}"] cluster should be started" {
  run is-cluster-status "${OS_CLUSTER_NAME}" "START_IN_PROGRESS"
  if [[ "$output" != "true" ]]; then
    skip "Cluster Start has not been requested"
  fi

  run wait-cluster-status $DELAY "${OS_CLUSTER_NAME}" "AVAILABLE"

  echo "$output" >&2

  [ $status -eq 0 ]
  [ "$output" = "true" ]
}

@test "["${OS_CLUSTER_NAME}"] cluster can be upscale" {
  run is-cluster-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is NOT created yet!"
  fi

  CHECK_RESULT=$(describe-cluster --name "${OS_CLUSTER_NAME}" | jq ' ."instanceGroups" | .[] | select(."group"=="compute") | . "nodeCount"')
  INSTANCE_COUNT_DESIRED=$(($CHECK_RESULT + 1))

  echo $INSTANCE_COUNT_DESIRED > "${UTILS_TEMPFILE}"

  run scale-cluster --name "${OS_CLUSTER_NAME}" --group-name compute --desired-node-count $INSTANCE_COUNT_DESIRED
  [ $status -eq 0 ]
}

@test "["${OS_CLUSTER_NAME}"] cluster upscale should be started" {
  run is-cluster-status "${OS_CLUSTER_NAME}" "UPDATE_IN_PROGRESS"
  if [[ "$output" != "true" ]]; then
    skip "Cluster upscale has not been requested!"
  fi

  run wait-cluster-status $DELAY "${OS_CLUSTER_NAME}" "AVAILABLE"
  echo "$output" >&2

  [ $status -eq 0 ]
  [ "$output" = "true" ]
}

@test "["${OS_CLUSTER_NAME}"] cluster should be upscaled" {
  run is-cluster-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster upscale has not been done!"
  fi

  INSTANCE_COUNT_DESIRED=`cat "${UTILS_TEMPFILE}"`
  echo $INSTANCE_COUNT_DESIRED

  INSTANCE_COUNT_CURRENT=$(describe-cluster --name "${OS_CLUSTER_NAME}" | jq ' ."instanceGroups" | .[] | select(."group"=="compute") | . "nodeCount" ')
  echo $INSTANCE_COUNT_CURRENT

  STATUS_REASON=`cat "${STATUS_TEMPFILE}"`
  echo $STATUS_REASON

  [[ $INSTANCE_COUNT_DESIRED -eq $INSTANCE_COUNT_CURRENT ]]
}

@test "["${OS_CLUSTER_NAME}"] cluster can be downscale" {
  run is-cluster-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is NOT created yet!"
  fi

  CHECK_RESULT=$( describe-cluster --name "${OS_CLUSTER_NAME}" | jq ' ."instanceGroups" | .[] | select(."group"=="compute") | . "nodeCount" ')
  INSTANCE_COUNT_DESIRED=$(($CHECK_RESULT - 1))

  echo $INSTANCE_COUNT_DESIRED > "${UTILS_TEMPFILE}"

  run scale-cluster --name "${OS_CLUSTER_NAME}" --group-name compute --desired-node-count $INSTANCE_COUNT_DESIRED
  [ $status -eq 0 ]
}

@test "["${OS_CLUSTER_NAME}"] cluster downscale should be started" {
  run is-cluster-status "${OS_CLUSTER_NAME}" "UPDATE_IN_PROGRESS"
  if [[ "$output" != "true" ]]; then
    skip "Cluster downscale has not been requested!"
  fi

  run wait-cluster-status $DELAY "${OS_CLUSTER_NAME}" "AVAILABLE"
  echo "$output" >&2

  [ $status -eq 0 ]
  [ "$output" = "true" ]
}

@test "["${OS_CLUSTER_NAME}"] cluster should be downscaled" {
  run is-cluster-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster downscale has not been done!"
  fi

  INSTANCE_COUNT_DESIRED=`cat "${UTILS_TEMPFILE}"`
  echo $INSTANCE_COUNT_DESIRED

  INSTANCE_COUNT_CURRENT=$(describe-cluster --name "${OS_CLUSTER_NAME}" | jq ' ."instanceGroups" | .[] | select(."group"=="compute") | . "nodeCount" ')
  echo $INSTANCE_COUNT_CURRENT

  STATUS_REASON=`cat "${STATUS_TEMPFILE}"`
  echo $STATUS_REASON

  [[ $INSTANCE_COUNT_DESIRED = $INSTANCE_COUNT_CURRENT ]]
}

@test "Generate reinstall template" {
  run is-cluster-present "${OS_CLUSTER_NAME}"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is not present!"
  fi

  OUTPUT=$(generate-reinstall-template --name "${OS_CLUSTER_NAME}" --blueprint-name "${BLUEPRINT_NAME}" > "${REINSTALL_TEMPFILE}")

  echo "${OUTPUT}" >&2

  [[ -f "${REINSTALL_TEMPFILE}" ]]
}

@test "TEARDOWN: Delete ["${OS_CLUSTER_NAME}"] OpenStack cluster" {
  OUTPUT=$(delete-cluster-wait "${OS_CLUSTER_NAME}" 2>&1 | tail -n 4 | head -n 1)

  echo "${OUTPUT}" >&2

  [[ "${OUTPUT}" == *"[waitForClusterToFinishImpl] stack status: DELETE_COMPLETED, cluster status"* ]]
  [[ "${OUTPUT}" != *"error"* ]]
}

@test "TEARDOWN: Wait for "${OS_CLUSTER_NAME}" cluster is terminated" {
  run is-cluster-status "${OS_CLUSTER_NAME}" "DELETE_IN_PROGRESS"
  if [[ "$output" != "true" ]]; then
    skip "Cluster has already been terminated!"
  fi

  run wait-cluster-delete $DELAY "${OS_CLUSTER_NAME}"

  echo "$output" >&2

  [ $status -eq 0 ]
  [ "$output" != "true" ]
}

@test "TEARDOWN: Delete ["${OS_CREDENTIAL_NAME}"cluster] OpenStack credential" {
  OUTPUT=$(delete-credential "${OS_CREDENTIAL_NAME}cluster" 2>&1 | tail -n 2 | head -n 1)

  echo "${OUTPUT}" >&2

  [[ "${OUTPUT}" == *"[DeleteCredential] credential deleted, name: ${OS_CREDENTIAL_NAME}cluster"* ]]
  [[ "${OUTPUT}" != *"error"* ]]
}

@test "TEARDOWN: Delete status temp file" {
  OUTPUT=$(rm -f "${STATUS_TEMPFILE}")
  echo "${OUTPUT}" >&2
}

@test "TEARDOWN: Delete util temp file" {
  OUTPUT=$(rm -f "${UTILS_TEMPFILE}")
  echo "${OUTPUT}" >&2
}

@test "TEARDOWN: Delete reinstall temp file" {
  OUTPUT=$(rm -f "${REINSTALL_TEMPFILE}")
  echo "${OUTPUT}" >&2
}