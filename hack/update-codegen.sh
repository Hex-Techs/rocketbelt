#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

API_BASE="${SCRIPT_ROOT}"/api/rocketbelt/v1alpha1

mkdir -p ${API_BASE}
cp ${SCRIPT_ROOT}/api/v1alpha1/* ${API_BASE}

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
 bash "${CODEGEN_PKG}"/generate-groups.sh "client,lister,informer" \
    github.com/hex-techs/rocketbelt/generated github.com/hex-techs/rocketbelt/api \
    "rocketbelt:v1alpha1" \
    --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
    --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.."

 if [ -d "${API_BASE}" ]; then
    rm -rf ${SCRIPT_ROOT}/api/rocketbelt
 fi

 find generated -type f -name "*.go" | xargs sed -i".out" -e "s#github.com/hex-techs/rocketbelt/api/rocketbelt/v1alpha1#github.com/hex-techs/rocketbelt/api/v1alpha1#g"
 find generated -type f -name "*go.out" | xargs rm -rf
