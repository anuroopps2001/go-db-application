```bash
azureuser@client-management-vm:~$ kubectl get svc -n go-db-app-dev
NAME                              TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)    AGE
go-db-service-sample-version-01   ClusterIP   10.0.5.201    <none>        8080/TCP   17h
go-db-service-sample-version-02   ClusterIP   10.0.56.223   <none>        8080/TCP   17h
postgres-db-service               ClusterIP   10.0.72.85    <none>        5432/TCP   17h
azureuser@client-management-vm:~$ export PGPASSWORD='SecurePassword123!'
azureuser@client-management-vm:~$ pg_dump -h postgres-db-service.go-db-app-dev.svc.cluster.local -U postgres -d mydb | psql -h pg-flex-india.postgres.databa
se.azure.com -U psqladmin -d k8s_migration
pg_dump: error: could not translate host name "postgres-db-service.go-db-app-dev.svc.cluster.local" to address: Temporary failure in name resolution
psql: error: connection to server at "pg-flex-india.postgres.database.azure.com" (10.2.1.4), port 5432 failed: FATAL:  database "k8s_migration" does not exist
azureuser@client-management-vm:~$ psql -h pg-flex-india.postgres.database.azure.com -U psqladmin -d postgres -c "CREATE DATABASE k8s_migration;"
CREATE DATABASE
azureuser@client-management-vm:~$
```


```bash
azureuser@client-management-vm:~$ pg_dump -h localhost -U postgres -d mydb | psql -h pg-flex-india.postgres.database.azure.com -U psqladmin -d k8s_migration
SET
SET
SET
ERROR:  unrecognized configuration parameter "transaction_timeout"
SET
SET
 set_config
------------

(1 row)

SET
SET
SET
SET
SET
SET
CREATE TABLE
ERROR:  role "postgres" does not exist
CREATE SEQUENCE
ERROR:  role "postgres" does not exist
ALTER SEQUENCE
ALTER TABLE
COPY 1   --> This line says that data migration is successful
 setval
--------
      1
(1 row)

ALTER TABLE
ALTER TABLE
azureuser@client-management-vm:~$ psql -h pg-flex-india.postgres.database.azure.com -U psqladmin -d k8s_migration -c "\dt"
         List of relations
 Schema | Name  | Type  |   Owner
--------+-------+-------+-----------
 public | users | table | psqladmin
(1 row)

azureuser@client-management-vm:~$ psql -h pg-flex-india.postgres.database.azure.com -U psqladmin -d k8s_migration -c "SELECT * FROM users;"
 id |  name   |        email        | age
----+---------+---------------------+-----
  1 | Anuroop | anuroop@example.com |   0
(1 row)

azureuser@client-management-vm:~$
```

#### Once Migration is done, Create the secrets in the Azure keyVault 
Create the Azure KeyVault secrets with keyNames required for the application and the values accordingly.


#### Providing Permissions on AKS to access Azure KeyVault secrets
```bash
azureuser@client-management-vm:~$ AKS_OIDC_URL=$(az aks show -n aks-private-india -g rg-aks-southindia --query "oidcIssuerProfile.issuerUrl" -otsv)
echo "Your Cluster OIDC URL is: $AKS_OIDC_URL"
Your Cluster OIDC URL is: https://southindia.oic.prod-aks.azure.com/3241b922-8d4a-4339-9b21-5009e9c776b8/69567099-f5b2-4d6e-b635-fb3adddba920/


azureuser@client-management-vm:~$ az identity federated-credential list \
  --identity-name aks-kv-identity \
  --resource-group rg-aks-southindia
[
  {
    "audiences": [
      "api://AzureADTokenExchange"
    ],
    "id": "/subscriptions/3c744587-46e1-4a41-b95f-3bca3fd5e622/resourcegroups/rg-aks-southindia/providers/Microsoft.ManagedIdentity/userAssignedIdentities/aks-kv-identity/federatedIdentityCredentials/kv-federation",
    "issuer": "https://southindia.oic.prod-aks.azure.com/3241b922-8d4a-4339-9b21-5009e9c776b8/69567099-f5b2-4d6e-b635-fb3adddba920/",
    "name": "kv-federation",
    "resourceGroup": "rg-aks-southindia",
    "subject": "system:serviceaccount:default:go-kv-sa",
    "type": "Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"
  }
]
```

* Create federated credential for specific namespace with specific serviceAccount Name
```bash
az identity federated-credential create \
  --name kv-federation \
  --identity-name aks-kv-identity \
  --resource-group rg-aks-southindia \
  --issuer "https://southindia.oic.prod-aks.azure.com/3241b922-8d4a-4339-9b21-5009e9c776b8/69567099-f5b2-4d6e-b635-fb3adddba920/" \
  --subject "system:serviceaccount:go-db-app-dev:go-app-sa"
```

* Create namespace and service account if doesn't exists
```bash
# Create the namespace if it's missing
kubectl create namespace go-db-app-dev --dry-run=client -o yaml | kubectl apply -f -


# Create the ServiceAccount
kubectl create serviceaccount go-app-sa -n go-db-app-dev --dry-run=client -o yaml | kubectl apply -f -
```

#### Annotate it with the Client ID of aks-kv-identity
#### Get the Client ID of the ManagedIdentity on which federatedCredential was created:
```bash
CLIENT_ID=$(az identity show -n aks-kv-identity -g rg-aks-southindia --query clientId -otsv)
```

Annotate ServiceAccount:
```bash
kubectl annotate serviceaccount go-app-sa -n go-db-app-dev \
  azure.workload.identity/client-id=$CLIENT_ID --overwrite
```

Verification:
```bash
azureuser@client-management-vm:~$ az identity federated-credential list   --identity-name aks-kv-identity   --resource-group rg-aks-southindia
[
  {
    "audiences": [
      "api://AzureADTokenExchange"
    ],
    "id": "/subscriptions/3c744587-46e1-4a41-b95f-3bca3fd5e622/resourcegroups/rg-aks-southindia/providers/Microsoft.ManagedIdentity/userAssignedIdentities/aks-kv-identity/federatedIdentityCredentials/kv-federation",
    "issuer": "https://southindia.oic.prod-aks.azure.com/3241b922-8d4a-4339-9b21-5009e9c776b8/69567099-f5b2-4d6e-b635-fb3adddba920/",
    "name": "kv-federation",
    "resourceGroup": "rg-aks-southindia",
    "subject": "system:serviceaccount:go-db-app-dev:go-app-sa", ## This makes, federated credential and It;s managedIdentity trust this specific cluster, Namespace and ServiceAccount
    "type": "Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"
  }
]
azureuser@client-management-vm:~$
```


* Get the PrincipalID of the ManagedIdentity and provide access at AzureKeyVault
```bash
PRINCIPAL_ID=$(az identity show -n aks-kv-identity -g rg-aks-southindia --query principalId -otsv)
```
* Assign a role on AKV, for PrincipalID:
```bash
azureuser@client-management-vm:~$ az role assignment create --assignee $PRINCIPAL_ID \
  --role "Key Vault Secrets User" \
  --scope "/subscriptions/3c744587-46e1-4a41-b95f-3bca3fd5e622/resourceGroups/rg-aks-southindia/providers/Microsoft.KeyVault/vaults/kv-aks-gitops-prod"
{
  "condition": null,
  "conditionVersion": null,
  "createdBy": null,
  "createdOn": "2026-04-08T05:24:03.168159+00:00",
  "delegatedManagedIdentityResourceId": null,
  "description": null,
  "id": "/subscriptions/3c744587-46e1-4a41-b95f-3bca3fd5e622/resourceGroups/rg-aks-southindia/providers/Microsoft.KeyVault/vaults/kv-aks-gitops-prod/providers/Microsoft.Authorization/roleAssignments/b4629db4-6b20-4c44-bd0a-7a147414b1e9",
  "name": "b4629db4-6b20-4c44-bd0a-7a147414b1e9",
  "principalId": "febc1b80-d826-4c3f-9387-ae50515911f2",
  "principalType": "ServicePrincipal",
  "resourceGroup": "rg-aks-southindia",
  "roleDefinitionId": "/subscriptions/3c744587-46e1-4a41-b95f-3bca3fd5e622/providers/Microsoft.Authorization/roleDefinitions/4633458b-17de-408a-b874-0445c86b69e6",
  "scope": "/subscriptions/3c744587-46e1-4a41-b95f-3bca3fd5e622/resourceGroups/rg-aks-southindia/providers/Microsoft.KeyVault/vaults/kv-aks-gitops-prod",
  "type": "Microsoft.Authorization/roleAssignments",
  "updatedBy": "9a901f05-62a3-4e02-92bc-aaa80d9d7707",
  "updatedOn": "2026-04-08T05:24:03.799166+00:00"
}
azureuser@client-management-vm:~$
```


* **Create SecretProviderClass**

* **Mount the SecretProviderClass** as a volume inside the application deployments.