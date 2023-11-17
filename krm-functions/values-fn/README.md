# values-fn

## Overview

<!--mdtogo:Short-->

This function is a simple values.yaml modifier.

<!--mdtogo:Long-->

## Usage
```
kpt fn source ./package/ | go run main.go
```


Example of wanted resource:

```
kind: ConfigMap
metadata:
  name: policy-acm-values
  namespace: acm
apiVersion: v1
data:
  values.yaml: |+
    replicaCount: 3
    image:
      repository: nginx
      pullPolicy: IfNotPresent
      tag: ""
```