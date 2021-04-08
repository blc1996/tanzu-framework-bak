# Tanzu CLI

## Summary
SIG CLI/API's objective is to support building CLI and API experiences within Tanzu, built on a shared framework to enable a consistent, unified Tanzu product experience for our users. We have designed the CLI around a pluggable model to support a broader set of the Tanzu portfolio’s products as they are ready to adopt it.

For more detail, please see the [Tanzu CLI purpose doc](https://docs.google.com/document/d/1X34ZNkPG_kEMSySpFjAQsmX2Xn1dXTksbVxXUgUk-QM/edit?usp=sharing).

## Overview
The CLI is based on a plugin architecture. This architecture enables teams to build, own, and release their own piece of functionality as well as enable external partners to integrate with the system.

For more detail, please see the [Tanzu CLI architecture doc](https://docs.google.com/document/d/1qCarTtSUxJzYJweiHsOQhObTc2L4f9smAXAlIheMFfI/edit#).
Additionally, an architectual diagram is available on [confluence](https://confluence.eng.vmware.com/display/TKG/Tanzu+CLI+Architecture+and+components).

---

## Installation

### Dev
### MacOS
```shell
curl -o tanzu https://storage.googleapis.com/tanzu-cli/artifacts/core/latest/tanzu-core-darwin_amd64 && \
    mv tanzu /usr/local/bin/tanzu && \
    chmod +x /usr/local/bin/tanzu
```
### Linux
#### i386
```shell
curl -o tanzu https://storage.googleapis.com/tanzu-cli/artifacts/core/latest/tanzu-core-linux_386 && \
    mv tanzu /usr/local/bin/tanzu && \
    chmod +x /usr/local/bin/tanzu
```
#### AMD64
```shell
curl -o tanzu https://storage.googleapis.com/tanzu-cli/artifacts/core/latest/tanzu-core-linux_amd64 && \
    mv tanzu /usr/local/bin/tanzu && \
    chmod +x /usr/local/bin/tanzu
```

### Windows
Windows executable can be found at https://storage.googleapis.com/tanzu-cli/artifacts/core/latest/tanzu-core-windows_amd64.exe

### Release Versions
TKG 1.3 Docs [link](https://docs.vmware.com/en/VMware-Tanzu-Kubernetes-Grid/1.1/vmware-tanzu-kubernetes-grid-11/GUID-install-tkg-set-up-tkg.html#unpack-cli)

---

## Usage
```
tanzu [command]

Available command groups:

  Manage
    clustergroup            A group of Kubernetes clusters

  Run
    cluster                 Kubernetes cluster operations
    kubernetes-release      Kubernetes release operations
    management-cluster      Commands for creating and managing TKG clusters

  System
    completion              Output shell completion code
    config                  Configuration for the CLI
    init                    Initialize the CLI
    login                   Login to the platform
    plugin                  Manage CLI plugins
    update                  Update the CLI
    version                 Version information


Flags:
  -h, --help   help for tanzu

Use "tanzu [command] --help" for more information about a command.
```
## Documentation
Tanzu CLI and Kickstart documentation on [TKG Confluence](https://confluence.eng.vmware.com/display/TKG/Tanzu+CLI+and+Kickstart+UI+Documentation).

## Contribution

👍🎉 First off, thanks for taking the time to contribute! 🎉👍

The first step to contributing is to come say hello at [SIG meetings](https://confluence.eng.vmware.com/display/CNA/SIG+CLI-API). You can also join the [slack channel](https://vmware.slack.com/archives/C016NB72DQW) or take a peek at our [issues](https://github.com/vmware-tanzu-private/core/issues). 

The [plugin author guide](https://confluence.eng.vmware.com/pages/viewpage.action?spaceKey=CNA&title=End-to-End+CLI+Plugin+Creation+Process) provides details on the getting the process started, but in short:
* The core framework and components are written in go
* Every plugin requires a README that explains its usage.
* Core plugins are required to conform to [design guidance](https://github.com/vmware-tanzu-private/core/blob/main/docs/cli/style_guide.md) to ensure a consistent user experience.
* The plugin will live in an approved repo to enable review and testing.
* Every CLI plugin requires a [nested test executable](https://github.com/vmware-tanzu-private/core/blob/main/docs/cli/plugin_implementation_guide.md#tests).

Non-core/3rd party plugins are welcome and encouraged, and can be added by users via `tanzu plugin repo add NAME`.  
Note that non-core plugins will not be included in plugin distributions distributed by VMware and commands will be namespaced differently in help.

For more detail, please see the [Tanzu CLI Implementation Guide](/docs/cli/plugin_implementation_guide.md)

## License
TODO
