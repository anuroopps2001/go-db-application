```bash
azureuser@vm-circleci-runner:~$ kubectl get crd | grep gateway
authenticationfilters.gateway.nginx.org                     2026-04-13T23:14:56Z
clientsettingspolicies.gateway.nginx.org                    2026-04-13T23:14:56Z
gatewayclasses.gateway.networking.k8s.io                    2026-04-13T10:58:20Z
gateways.gateway.networking.k8s.io                          2026-04-13T10:58:21Z
grpcroutes.gateway.networking.k8s.io                        2026-04-13T10:58:21Z
httproutes.gateway.networking.k8s.io                        2026-04-13T10:58:21Z
nginxgateways.gateway.nginx.org                             2026-04-13T23:14:56Z
nginxproxies.gateway.nginx.org                              2026-04-13T23:14:56Z
observabilitypolicies.gateway.nginx.org                     2026-04-13T23:14:57Z
proxysettingspolicies.gateway.nginx.org                     2026-04-13T23:14:57Z
ratelimitpolicies.gateway.nginx.org                         2026-04-13T23:14:57Z
referencegrants.gateway.networking.k8s.io                   2026-04-13T10:58:23Z
snippetsfilters.gateway.nginx.org                           2026-04-13T23:14:57Z
snippetspolicies.gateway.nginx.org                          2026-04-13T23:14:57Z
upstreamsettingspolicies.gateway.nginx.org                  2026-04-13T23:14:57Z
azureuser@vm-circleci-runner:~$
```

### Install the gateway helm chart with an internal Load Balancer within an AKS subnet
```bash
$ helm install ngf oci://ghcr.io/nginx/charts/nginx-gateway-fabric \
  --create-namespace \
  --namespace nginx-gateway \
  --set nginx.service.type=LoadBalancer \
  --set nginx.service.annotations."service\.beta\.kubernetes\.io/azure-load-balancer-internal"="true" \
  --set nginx.service.annotations."service\.beta\.kubernetes\.io/azure-load-balancer-internal-subnet"="snet-aks"
```

```bash
$ kubectl get gatewayclasses.gateway.networking.k8s.io
NAME    CONTROLLER                                   ACCEPTED   AGE
nginx   gateway.nginx.org/nginx-gateway-controller   True       3m42s

$ kubectl get pods -n nginx-gateway
NAME                                            READY   STATUS      RESTARTS   AGE
ngf-nginx-gateway-fabric-c98866d6f-s2l6t        1/1     Running     0          21s
ngf-nginx-gateway-fabric-cert-generator-d287v   0/1     Completed   0          31s
```
### Create gateway instance
```bash
 kubectl get gateway -A
NAMESPACE       NAME               CLASS   ADDRESS       PROGRAMMED   AGE
nginx-gateway   internal-gateway   nginx   4.213.41.36   True         42s
```

### Role assignment
```bash
NODE_RG=$(az aks show -g rg-spoke-workloads -n aks-private-cluster --query nodeResourceGroup -o tsv)v)
azureuser@vm-circleci-runner:~/secure-deployments$
azureuser@vm-circleci-runner:~/secure-deployments$
azureuser@vm-circleci-runner:~/secure-deployments$
azureuser@vm-circleci-runner:~/secure-deployments$ echo $NODE_RG
MC_rg-spoke-workloads_aks-private-cluster_centralindia
azureuser@vm-circleci-runner:~/secure-deployments$
azureuser@vm-circleci-runner:~/secure-deployments$
azureuser@vm-circleci-runner:~/secure-deployments$
azureuser@vm-circleci-runner:~/secure-deployments$ MSI_ID=$(az identity list -g $NODE_RG --query "[?contains(name, 'agentpool')].principalId" -o tsv)
azureuser@vm-circleci-runner:~/secure-deployments$ echo $MSI_ID
8f2d60b2-cac8-4322-aa16-ca6d3beb808a
azureuser@vm-circleci-runner:~/secure-deployments$ az role assignment create --assignee $MSI_ID \
  --role "Network Contributor" \
  --scope /subscriptions/3c744587-46e1-4a41-b95f-3bca3fd5e622/resourceGroups/rg-spoke-workloads/providers/Microsoft.Network/virtualNetworks/vnet-spoke-workloads
Subscription '3c744587-46e1-4a41-b95f-3bca3fd5e622' not found. Check the spelling and casing and try again.
azureuser@vm-circleci-runner:~/secure-deployments$ az role assignment create --assignee $MSI_ID   --role "Network Contributor"   --scope /subscriptions/5f47c6fe-549b-45c6-ae21-792c56637b80/resourceGroups/rg-spoke-workloads/providers/Microsoft.Network/virtualNetworks/vnet-spoke-aks
{
  "condition": null,
  "conditionVersion": null,
  "createdBy": null,
  "createdOn": "2026-04-13T23:44:38.437183+00:00",
  "delegatedManagedIdentityResourceId": null,
  "description": null,
  "id": "/subscriptions/5f47c6fe-549b-45c6-ae21-792c56637b80/resourceGroups/rg-spoke-workloads/providers/Microsoft.Network/virtualNetworks/vnet-spoke-aks/providers/Microsoft.Authorization/roleAssignments/586e7e5f-2856-426f-8766-bd6897eb7385",
  "name": "586e7e5f-2856-426f-8766-bd6897eb7385",
  "principalId": "8f2d60b2-cac8-4322-aa16-ca6d3beb808a",
  "principalType": "ServicePrincipal",
  "resourceGroup": "rg-spoke-workloads",
  "roleDefinitionId": "/subscriptions/5f47c6fe-549b-45c6-ae21-792c56637b80/providers/Microsoft.Authorization/roleDefinitions/4d97b98b-1d4f-4787-a291-c67834d212e7",
  "scope": "/subscriptions/5f47c6fe-549b-45c6-ae21-792c56637b80/resourceGroups/rg-spoke-workloads/providers/Microsoft.Network/virtualNetworks/vnet-spoke-aks",
  "type": "Microsoft.Authorization/roleAssignments",
  "updatedBy": "e4df0a8f-d074-4bd5-b0c2-bb88d8fcf748",
  "updatedOn": "2026-04-13T23:44:39.138193+00:00"
}
```
