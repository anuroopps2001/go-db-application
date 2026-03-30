
```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
```

```bash
helm repo ls
helm search repo prometheus-community/kube-prometheus-stack --versions
```

```bash
kubectl create ns monitoring
```

```bash
helm install monitoring prometheus-community/kube-prometheus-stack -n monitoring
```


* Onece the installation is done, port-forward both the prometheus and grafana to access from our laptops

* Add the prometheus as datasource and test the connectivity from the grafana URL

* Create a service monitor based on the application service name, exposed port, labels and metrics exposed path and deploy the servive monitor instance in `monitoring` namespace


```bash
azureuser@client-management-vm:~/go-db-application/k8s/observability/prometheus-grafana-alertting-with-helm$ kubectl get prometheus -n monitoring -o yaml | grep -A 3 serviceMonitorSelector
    serviceMonitorSelector:
      matchLabels:
        release: monitoring
    shards: 1
```

If you see something like release: kube-prometheus-stack, your ServiceMonitor must have that exact label.



### Writing prometheus queries

```bash
rate(http_requests_total[1m])
```

Sucsesss i,e status_code=200:
```bash
rate(http_requests_total{status="201"}[1m])
round(increase(http_requests_total{status="404"}[10m]))
```

client errors:
```bash
rate(http_requests_total{status="409"}[1m])
rate(http_requests_total{status=~"4.."}[1m])
```

server errors: any status_code starting with 5
```bash
rate(http_requests_total{status=~"5.."}[1m])
```


```bash
histogram_quantile(0.95, sum by (le, job, namespace) (rate(http_request_duration_seconds_bucket{method="GET", path="/ready"}[5m])))
```

the number of requests by path:
```bash
sum by (path) (increase(http_requests_total[1m]))
```
