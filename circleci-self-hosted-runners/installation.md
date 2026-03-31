### Create a resource class in UI



### Create a servicePrincipal

```bash
az ad sp create-for-rbac --name "circleci-aks-sp"
```

### Add federated identity
```bash
az ad app federated-credential create --id <client ID of your app registration> \
--parameters credential.json
```

### Example credential.json

```bash
{
  "name": "circleci-federated-identity",
  "issuer": "https://oidc.circleci.com/org/<YOUR_CIRCLECI_ORG_ID>",
  "subject": "org/<YOUR_CIRCLECI_ORG_ID>/project/<YOUR_PROJECT_ID>/user/<YOUR_USER_ID>",
  "description": "Federated identity for CircleCI Canary Runner",
  "audiences": [
    "<YOUR_CIRCLECI_ORG_ID>"
  ]
}
```


### Assign role
```bash
az role assignment create \
  --assignee <APP_ID> \
  --role "Azure Kubernetes Service Cluster User Role" \
  --scope /subscriptions/<sub>/resourceGroups/<rg>/providers/Microsoft.ContainerService/managedClusters/<aks>
```