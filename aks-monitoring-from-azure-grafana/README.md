```bash
az aks enable-addons \
  --addons monitoring \
  --resource-group rg-aks-southindia \
  --name aks-private-india \
  --workspace-resource-id "$LOG_ID"
```

```bash
az aks update \
  --resource-group rg-aks-southindia \
  --name aks-private-india \
  --enable-azure-monitor-metrics \
  --azure-monitor-workspace-resource-id "$PROM_ID"
```

```bash
$ az aks show -g rg-aks-southindia -n aks-private-india --query "addonProfiles.monitoring.enabled"


$ az aks show -g rg-aks-southindia -n aks-private-india --query "azureMonitorProfile.metrics.enabled"
true
```

#### Azure Managed Grafana
```bash
az grafana create   --name go-app-grafana   --resource-group rg-aks-southindia   --location centralindia
```

#### Provide access of AKS Prometheus for Azure Managed Grafana
```bash
az aks update \
  --resource-group rg-aks-southindia \
  --name aks-private-india \
  --enable-azure-monitor-metrics \
  --azure-monitor-workspace-resource-id "$PROM_ID" \
  --grafana-resource-id $(az grafana show --name go-app-grafana -g rg-aks-southindia --query id -o tsv)
```

#### Accessing Grafana Dashboard
```bash
az grafana show --name go-app-grafana -g rg-aks-southindia --query "properties.endpoint" -o tsv
```
