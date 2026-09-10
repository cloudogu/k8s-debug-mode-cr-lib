## Die Debug-Mode CRD generieren

Folgende Befehle ausführen: 

```bash
kustomize build config/default > output2.yaml | helmify
kustomize build config/crd > output.yaml | helmify
```
