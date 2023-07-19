#!/bin/sh
set -euo pipefail

VERSION=${1#"v"}
echo "VERSION: "$VERSION
if [ -z "$VERSION" ]; then
    echo "Must specify version!"
    exit 1
fi
MODS=($(
    curl -sS https://raw.githubusercontent.com/kubernetes/kubernetes/v${VERSION}/go.mod |
    sed -n 's|.*k8s.io/\(.*\) => ./staging/src/k8s.io/.*|k8s.io/\1|p'
))
echo "MODs": ${MODS[*]}
for MOD in "${MODS[@]}"; do
    MOD_VERSION=${MOD}@kubernetes-${VERSION}
    echo "MOD": $MOD_VERSION
    V=$(
        go mod download -json "${MOD_VERSION}" |
        sed -n 's|.*"Version": "\(.*\)".*|\1|p'
    )
    go mod edit "-replace=${MOD}=${MOD}@${V}"
done
go get -u -v "k8s.io/kubernetes@v${VERSION}"