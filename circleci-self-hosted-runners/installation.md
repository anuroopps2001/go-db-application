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

OR 

```bash

$ APP_ID=$(az ad app create --display-name "circleci-gitops-aks" --query appId -o tsv)

ANUROOP P S@ANU MINGW64 /d/go_k8s/circleci-self-hosted-runners (feature/circleci-metrics)
$ az ad sp create --id $APP_ID
{
  "@odata.context": "https://graph.microsoft.com/v1.0/$metadata#servicePrincipals/$entity",
  "accountEnabled": true,
  "addIns": [],
  "alternativeNames": [],
  "appDescription": null,
  "appDisplayName": "circleci-gitops-aks",
  "appId": "b6a04818-afd5-4851-bbdc-5f7ebe36567e",
  "appOwnerOrganizationId": "3241b922-8d4a-4339-9b21-5009e9c776b8",
  "appRoleAssignmentRequired": false,
  "appRoles": [],
  "applicationTemplateId": null,
  "createdByAppId": "04b07795-8ddb-461a-bbee-02f9e1bf7b46",
  "createdDateTime": null,
  "deletedDateTime": null,
  "description": null,
  "disabledByMicrosoftStatus": null,
  "displayName": "circleci-gitops-aks",
  "homepage": null,
  "id": "dfc25b42-2234-445f-8103-88b4b7acbc7b",
  "info": {
    "logoUrl": null,
    "marketingUrl": null,
    "privacyStatementUrl": null,
    "supportUrl": null,
    "termsOfServiceUrl": null
  },
  "isDisabled": null,
  "keyCredentials": [],
  "loginUrl": null,
  "logoutUrl": null,
  "notes": null,
  "notificationEmailAddresses": [],
  "oauth2PermissionScopes": [],
  "passwordCredentials": [],
  "preferredSingleSignOnMode": null,
  "preferredTokenSigningKeyThumbprint": null,
  "replyUrls": [],
  "resourceSpecificApplicationPermissions": [],
  "samlSingleSignOnSettings": null,
  "servicePrincipalNames": [
    "b6a04818-afd5-4851-bbdc-5f7ebe36567e"
  ],
  "servicePrincipalType": "Application",
  "signInAudience": "AzureADMyOrg",
  "tags": [],
  "tokenEncryptionKeyId": null,
  "verifiedPublisher": {
    "addedDateTime": null,
    "displayName": null,
    "verifiedPublisherId": null
  }
}

ANUROOP P S@ANU MINGW64 /d/go_k8s/circleci-self-hosted-runners (feature/circleci-metrics)
$ az ad app federated-credential create --id $APP_ID --parameters credential.json 
{
  "@odata.context": "https://graph.microsoft.com/v1.0/$metadata#applications('55e411f1-9e75-458a-91a0-eb6afaff500d')/federatedIdentityCredentials/$entity",
  "audiences": [
    "df943366-0e4d-4301-814a-796da580b6cc"
  ],
  "description": "Trust CircleCI Project for go-db-application",
  "id": "1f31a273-c2a0-466a-98da-e939111c2021",
  "issuer": "https://oidc.circleci.com/org/df943366-0e4d-4301-814a-796da580b6cc",
  "name": "circleci-oidc-trust",
  "subject": "org/df943366-0e4d-4301-814a-796da580b6cc/project/ee3b3e0e-a9b7-4abe-9d7a-7324ba24e786/*"
}

ANUROOP P S@ANU MINGW64 /d/go_k8s/circleci-self-hosted-runners (feature/circleci-metrics)
$ az account show
{
  "environmentName": "AzureCloud",
  "homeTenantId": "3241b922-8d4a-4339-9b21-5009e9c776b8",
  "id": "3c744587-46e1-4a41-b95f-3bca3fd5e622",
  "isDefault": true,
  "managedByTenants": [],
  "name": "Azure subscription 1",
  "state": "Enabled",
  "tenantDefaultDomain": "anuroopps2127outlook.onmicrosoft.com",
  "tenantDisplayName": "Default Directory",
  "tenantId": "3241b922-8d4a-4339-9b21-5009e9c776b8",
  "user": {
    "name": "anuroopps2127@outlook.com",
    "type": "user"
  }
}
```

### Working command
```bash
az ad app federated-credential create \
  --id c378042f-4d4f-4724-918b-9944ebbee106 \
  --parameters '{
    "name": "circleci-oidc-trust",
    "issuer": "https://oidc.circleci.com/org/df943366-0e4d-4301-814a-796da580b6cc",
    "subject": "org/df943366-0e4d-4301-814a-796da580b6cc/project/ee3b3e0e-a9b7-4abe-9d7a-7324ba24e786/user/58bd21dd-34fd-4dfa-be1b-f8909fae9137/vcs-origin/github.com/anuroopps2001/go-db-application/vcs-ref/refs/heads/feature/circleci-metrics",
    "audiences": ["df943366-0e4d-4301-814a-796da580b6cc"]
  }'
```