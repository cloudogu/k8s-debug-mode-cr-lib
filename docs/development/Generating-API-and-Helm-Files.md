

kubebuilder create api --group k8s.cloudogu.com --version v1 --kind DebugMode
kustomize build config/default > output2.yaml | helmify
kustomize build config/crd > output.yaml | helmify