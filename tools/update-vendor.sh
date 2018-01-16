#!/usr/bin/env bash
#
# This script updates Go vendoring using glide and then fixes BUILD.bazel
# rules.

set -euo pipefail

# Go to the root of the repo
cd "$(git rev-parse --show-cdup)"

# Run glide.
glide up -v

# Sometimes vendor BUILD files cause problems, so delete them for now.
# Gazelle will re-create the ones we need.
# e.g. https://github.com/kubernetes/kubernetes/issues/50975
find vendor -name 'BUILD*' -delete

# Re-run Gazelle.
./tools/update-build.sh
