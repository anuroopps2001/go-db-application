```bash
az aks scale \
  --resource-group rg-aks-southindia \
  --name aks-private-india \
  --node-count 2
```

OR

```bash
az aks update \
  --resource-group rg-aks-southindia \
  --name aks-private-india \
  --enable-cluster-autoscaler \
  --min-count 1 \
  --max-count 3
```