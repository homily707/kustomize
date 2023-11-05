apiVersion: v1
data:
  metricsAggregator: |-
    {
      "enableMetricAggregation": "false",
      "enablePrometheusScraping" : "false"
    }
kind: ConfigMap
metadata:
  name: inferenceservice-config
  namespace: kserve
---
apiVersion: serving.kserve.io/v1alpha1
kind: ClusterStorageContainer
metadata:
  name: default
  namespace: kserve
spec:
  container:
    image: kserve/storage-initializer:latest
    name: storage-initializer
    resources:
      limits:
        cpu: "1"
        memory: 1Gi
      requests:
        cpu: 100m
        memory: 100Mi
  supportedUriFormats:
  - prefix: gs://
  - prefix: s3://
  - prefix: hdfs://
  - prefix: webhdfs://
  - regex: https://(.+?).blob.core.windows.net/(.+)
  - regex: https://(.+?).file.core.windows.net/(.+)
  - regex: https?://(.+)/(.+)
