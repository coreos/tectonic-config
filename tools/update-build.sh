#!/usr/bin/env bash
#
# This script updates all BUILD.bazel rules for Go packages. It's just a thin
# wrapper around gazelle.

set -euo pipefail

# Go to the root of the repo
cd "$(git rev-parse --show-cdup)"

bazel run //:gazelle
