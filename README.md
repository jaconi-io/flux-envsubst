[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white)](https://conventionalcommits.org)
[![semantic-release: angular](https://img.shields.io/badge/semantic--release-angular-e10079?logo=semantic-release)](https://github.com/semantic-release/semantic-release)
[![go report](https://goreportcard.com/badge/github.com/jaconi-io/flux-envsubst)](https://goreportcard.com/report/github.com/jaconi-io/flux-envsubst)
![CI](https://github.com/jaconi-io/flux-envsubst/actions/workflows/ci.yaml/badge.svg)
[![GitHub release](https://img.shields.io/github/release/jaconi-io/flux-envsubst.svg)](https://github.com/jaconi-io/flux-envsubst/releases/)

# flux-envsubst

[Flux](https://fluxcd.io) recommends using Drone's [envsubst](https://github.com/drone/envsubst) to replicate post-build
substitutions locally. However, envsubst lacks support for

* `.env` files
* the `kustomize.toolkit.fluxcd.io/substitute: disabled` annotation / label
* [SOPS](https://github.com/mozilla/sops)

## Installation

```
go install github.com/jaconi-io/flux-envsubst@latest
```

## Usage

```
kustomize build . | flux-envsubst | kubectl apply -f -
```
