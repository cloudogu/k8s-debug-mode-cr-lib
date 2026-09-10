## How to generate Debug-Mode CRD

Execute the following commands: 

```bash
kustomize build config/default > output2.yaml | helmify
kustomize build config/crd > output.yaml | helmify
```