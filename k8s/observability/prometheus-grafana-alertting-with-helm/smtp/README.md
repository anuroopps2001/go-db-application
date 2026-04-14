### 1. Create 16 char app password from your Gmail account

### 2. Create secret in k8s using the app password
```bash
kubectl create secret generic grafana-smtp-password   --from-literal=password='xsXXXXXXXXXXXhw'   -n monitoring
```


### 3. Create a values.yaml
```bash
grafana:
  enabled: true
  # Add this to ensure the pod actually restarts when secrets change
  envValueFrom:
    GF_SMTP_PASSWORD:
      secretKeyRef:
        name: grafana-smtp-password
        key: password

  grafana.ini:
    smtp:
      enabled: true
      host: smtp.gmail.com:587
      user: anuroopps2001@gmail.com
      from_address: anuroopps2001@gmail.com
      from_name: Grafana Alerts
      # Tell Grafana to use the Env Var we defined above
      password: "$__env{GF_SMTP_PASSWORD}"
```

### 4. Update helm chart with the above values.yaml
```bash
helm upgrade prometheus prometheus-community/kube-prometheus-stack   --namespace monitoring   -f smtp-val
ues.yaml
```

### 5. Verify from the grafana pod ENV Vars
```bash
azureuser@vm-circleci-runner:~/secure-deployments$ kubectl exec -it prometheus-grafana-6986c98d78-zpjzr -n monitoring -- env | grep GF_
GF_PATHS_CONFIG=/etc/grafana/grafana.ini
GF_PATHS_HOME=/usr/share/grafana
GF_SECURITY_ADMIN_PASSWORD=1rkrXXXXXXXXXXXXXXXXXXXXXvow
GF_PATHS_DATA=/var/lib/grafana/
GF_PATHS_LOGS=/var/log/grafana
GF_PATHS_PLUGINS=/var/lib/grafana/plugins
GF_PATHS_PROVISIONING=/etc/grafana/provisioning
GF_UNIFIED_STORAGE_INDEX_PATH=/var/lib/grafana-search/bleve
GF_SMTP_PASSWORD=xXXXXXXXXXXXXXXXXhw  # This is the password var required
GF_SECURITY_ADMIN_USER=admin
azureuser@vm-circleci-runner:~/secure-deployments$
```
### 6. Create the contact point inside the Grafana and Test