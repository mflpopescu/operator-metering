#!/bin/bash -e

ROOT_DIR=$(dirname "${BASH_SOURCE}")/..
source "${ROOT_DIR}/hack/common.sh"

# can also be specified as an argument
METERING_CR_FILE="${1:-$METERING_CR_FILE}"

if [ "$CREATE_NAMESPACE" == "true" ]; then
    echo "Creating namespace ${METERING_NAMESPACE}"
    kubectl create namespace "${METERING_NAMESPACE}" || true
elif ! kubectl get namespace "${METERING_NAMESPACE}" 2> /dev/null; then
    echo "Namespace '${METERING_NAMESPACE}' does not exist, please create it before starting"
    exit 1
fi

msg "Installing Custom Resource Definitions"
kube-install \
    "$MANIFESTS_DIR/custom-resource-definitions"

if [ "$SKIP_METERING_OPERATOR_DEPLOYMENT" == "true" ]; then
    echo "\$SKIP_METERING_OPERATOR_DEPLOYMENT=true, not creating metering-operator"
else
    msg "Installing metering-operator service account and RBAC resources"
    kube-install \
        "$INSTALLER_MANIFESTS_DIR/metering-operator-service-account.yaml" \
        "$INSTALLER_MANIFESTS_DIR/metering-operator-role.yaml" \
        "$INSTALLER_MANIFESTS_DIR/metering-operator-rolebinding.yaml"

    msg "Installing metering-operator"
    kube-install \
        "$INSTALLER_MANIFESTS_DIR/metering-operator-deployment.yaml"
fi

msg "Installing Metering Resource"
kube-install \
    "$METERING_CR_FILE"
