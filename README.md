# Tectonic Operator Configuration

This repository contains configuration objects used by the [Tectonic
Installer](https://github.com/coreos/tectonic-installer) to configure the
operators that install and update components in a Tectonic cluster.

## Configuration Objects

Configuration objects are defined for the following operators:

* [Kube Addon Operator](config/kube-addon)
* [Kube Core Operator](config/kube-core)
* [Tectonic Networking Operator](config/tectonic-network)
* [Tectonic Utility Operator](config/tectonic-utiltiy)

In addition, some utility functions are defined in the [config
library](config/).

## Build/Test

Build and test objects using Bazel. There are utilities tools
in the [tools](tools/) directory for updating `BUILD` rules
and vendoring.

## Vendoring

This repository uses glide for vendoring. Since this is a library the `vendor/`
directory is not committed to the repository. To install the correct libraries
to the `vendor/` directory for local development use `glide install`.
