### Insall helm
```bash
sudo apt-get install curl gpg apt-transport-https --yes
curl -fsSL https://packages.buildkite.com/helm-linux/helm-debian/gpgkey | gpg --dearmor | sudo tee /usr/share/keyrings/helm.gpg > /dev/null
echo "deb [signed-by=/usr/share/keyrings/helm.gpg] https://packages.buildkite.com/helm-linux/helm-debian/any/ any main" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list
sudo apt-get update
sudo apt-get install helm
```

### Add PLG stack repos
```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
```

```bash
helm repo ls
helm search repo prometheus-community/kube-prometheus-stack --versions
```

### Create namespace and install the chart into `monitoring` namespace
```bash
helm install prometheus prometheus-community/kube-prometheus-stack \
  --namespace monitoring \
  --create-namespace \
  --set grafana.enabled=true
```


### Loki setup to be used with above created grafana instance from prometheus chart 
```bash
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update

# Install Loki (using the simple scalable mode or monolithic for smaller clusters)
helm install loki grafana/loki-stack \
  --namespace monitoring \
  --set loki.persistence.enabled=true \
  --set loki.persistence.size=10Gi \
  --set loki.persistence.storageClass=default \
  --set promtail.enabled=true \
  --set grafana.enabled=false # We use the Grafana from the Prometheus stack created above
```



### Access Grafana 
```bash
kubectl port-forward svc/prometheus-grafana 8080:80 -n monitoring
azureuser@vm-circleci-runner:~$ kubectl get secrets -n monitoring  prometheus-grafana -ojsonpath='{.data.admin-password}' | base64 -d
```
### Add Loki as datasource
```bash
Test the Connection: In Data Sources > Loki, click Save & Test.

If you still get the unexpected IDENTIFIER error, ignore it. 3.  Explore the Logs: Go to the Explore tab (compass icon), select Loki, and use the Label Browser. You should see labels like job, namespace, and pod populated with data from your AKS cluster.
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
