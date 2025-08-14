## How to generate Debug-Mode CRD

Execute following commands: 

kustomize build config/default > output2.yaml | helmify
kustomize build config/crd > output.yaml | helmify