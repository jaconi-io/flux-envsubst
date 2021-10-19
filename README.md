# flux-envsubst

```
git clone git@github.com:jaconi-io/flux-envsubst.git
cd flux-envsubst 
go install
```

```
kustomize build . | flux-envsubst | kubectl apply -f -
```
