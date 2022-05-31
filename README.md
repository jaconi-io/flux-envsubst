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
