1. You have an Azure virtual machine that you back up by using Azure Backup.

The backup policy sub type is Standard, and the backup policy has the following configurations:

Backup schedule frequency: Weekly
Retain instant recovery snapshot(s) for: 5 days
Retention of weekly backup point: On Sunday at 8:00 AM for 12 weeks
You discover that Instant Restore is consuming more storage than expected.

You need to reduce the amount of storage consumed by Instant Restore.

What should you do first?

Options:
```bash
Change the backup schedule frequency to Daily.


Change the retention of weekly backup points to 1 week.


Modify the backup policy to reduce the retention of instant recovery snapshots.


Provision an additional blob storage container.
```

2. You have an Azure Resource Manager (ARM) template named deploy.json that is stored in an Azure Blob storage container.

You plan to deploy the template by running the New-AzDeployment cmdlet.

Which parameter should you use to reference the template?

Options:
```bash
-Tag


-Templatefile


-TemplateSpecId


-TemplateUri
```

3. You have an Azure Resource Manager (ARM) template named Template1 that is used to deploy Azure virtual machines.

Template1 contains the following text. 

"resources": [
  {
    "type": "Microsoft.Compute/virtualMachines",
    "apiVersion": "2025-04-01",
    "name": "[parameters('vmName')]",
    "location": "[resourceGroup().location]",
    "properties": {
      &lt;text removed&gt;
    }
  }
]

You need to deploy two Azure virtual machines by using Template1.

What should you add to Template1?


Options:
```bash

a copy element


the API version


the Azure subscription ID


the resource group location
```

4. Your company plans to host an application on four Azure virtual machines.

You need to ensure that at least two virtual machines are available if a single Azure datacenter fails.

Which availability option should you select for the virtual machine?


Options:
```bash
azure azs
azure avaibality sets
vm sets
```

5. You have an Azure virtual machine.

You receive a notification that the virtual machine is going to be affected by an underlying maintenance activity on the physical infrastructure.

You need to move the virtual machine to a different host to avoid a service interruption.

What should you do?

Options:
```bash
Apply an Azure policy.


Apply an Azure tag.


Move the virtual machine to another Azure subscription.


Redeploy the virtual machine.
```

6. You are deploying a virtual machine by using an availability set in the East US Azure region.

You have deployed 18 virtual machines in two fault domains and 10 update domains.

Microsoft performed planned physical hardware maintenance in the East US region.

What is the maximum number of virtual machines that will be unavailable?

Options:
```bash
2


8


9


18
```

7. You plan to deploy an Azure virtual machine.

You are evaluating whether to use an Azure Spot instance.

Which two factors can cause an Azure Spot instance to be evicted? Each correct answer presents a complete solution.

Options:
```bash
the average CPU usages of the instance


the Azure capacity needs


the current price of the instance


the time of day
```

8. You have an Azure subscription that contains an Azure Storage account named vmstorageaccount1.  

You create an Azure container instance named container1.

You need to configure persistent storage for container1.

What should you create in vmstorageaccount1?

Options:
```bash
a blob container


a file share


a queue


a table
```

9. Your development team plans to deploy an Azure container instance. The container needs a persistent storage layer.

Which service should you use?

Options:
```bash
Azure Blob storage


Azure Files


Azure Queue Storage


Azure SQL Database
```

10. Your company has an Azure subscription that is linked to a Microsoft Entra tenant.

You have been asked to limit the access to the Kubernetes API server.

Which two options should you choose? Each correct answer presents a complete solution.

Options:
```bash
API server authorized IP ranges


public cluster


private cluster


Azure tags
```

11. You have an Azure subscription that contains a Docker container image named container1.

You plan to create a new Azure App Service web app named WebApp1.

You need to ensure that you can use container1 for WebApp1.

Which WebApp1 setting should you configure?

Options:
```bash
Continuous deployment


Pricing plan


Publish


Runtime stack
```

12. You have an Azure subscription that contains multiple resource groups and Azure App Service web apps. A resource group named RG1 hosts a web app named appservice1.

The App Service uses an SSL certificate.

You create a resource group named RG2.

You plan to move all the resources in RG1 to RG2.  

Which two actions should you perform? Each correct answer presents part of the solution.

Options:
```bash
Create a new App Service plan in RG2.


Create a new web app in RG2.


Delete the SSL certificate from RG1 and upload it to RG2.


Move all the resources from RG1 to RG2.
```

13. You have a Basic Azure App Service plan that contains a web app.

You need to ensure that the web app can scale automatically when the CPU usage is over 80% for a duration of 15 minutes.

Which two actions should you perform? Each correct answer presents part of the solution.

Options:
```bash
Configure a deployment slot.


Configure a scaling condition to scale based on a metric, and then add the rules.


Configure a scaling condition to scale based on an instance count, and then set the instance count.


Scale out the App Service plan.


Scale up the App Service plan.
```

14. You have an Azure subscription that contains an App Service web app named App1.

You configure App1 with a custom domain name of webapp1.contoso.com.

You need to create a DNS record for App1. The solution must ensure that App1 remains accessible if the IP address changes.

Which type of DNS record should you create?

Options:
```bash
A


CNAME


SOA


SRV


TXT
```

15. You have an Azure subscription that is linked to a Microsoft Entra tenant named contoso.com.

All users in contoso.com are currently able to invite external users to B2B collaboration.

You need to ensure that only members of the Guest Inviter, User Administrator, and Global Administrator roles can invite guest users.

What should you configure?

Options:
```bash
Access reviews


Conditional Access


Cross-tenant access settings


External collaboration settings
```

16. You have an Azure subscription that is linked to a Microsoft Entra tenant named contoso.com.

All users in contoso.com are currently able to invite external users to B2B collaboration.

You need to ensure that only members of the Guest Inviter, User Administrator, and Global Administrator roles can invite guest users.

What should you configure?

Options:
```bash
Access reviews


Conditional Access


Cross-tenant access settings


External collaboration settings
```

17. You have an Azure subscription.

From PowerShell, you run the Get-MgUser cmdlet for a user and receive the following details:

Id: 8755b347-3545-3876-3987-999999999999
DisplayName: Ben Smith
Mail: bsmith@contoso.com
UserPrincipalName: bsmith_contoso.com#EXT#@fabrikam.com
Based upon the output of the cmdlet, which statement accurately describes the user?

Options:
```bash
The user account is disabled.


The user is a guest in the tenant.


The user is assigned an administrative role.


The user is deleted.
```

18. Your Microsoft Entra tenant and on-premises Active Directory domain contain multiple users.

You need to configure self-service password reset (SSPR) functionality. The solution must minimize costs.

Which Microsoft Entra ID edition should you use?


Options:
```bash
Microsoft Entra ID Free


Microsoft Entra ID P1


Microsoft Entra ID P2
```

19. You have a Microsoft Entra tenant that contains a user named User1.

You need to ensure that User1 can invite external users to the tenant. The solution must follow the principle of least privilege.

Which role should you assign to User1?

Options:
```bash
Global Administrator


Groups Administrator


Guest Inviter


Security Administrator
```


20. You have an Azure subscription that contains multiple users and administrators.  

You are creating a new custom role by using the following JSON.  

``

{   "Name": "Custom Role",   "Id": null,   "IsCustom": true,   "Description": "Custom Role description",   "Actions": [     "Microsoft.Compute/*/read",     “Microsoft.Compute/snapshots/write”,     “Microsoft.Compute/snapshots/read”,   ],   "NotActions": [   “Microsoft.Compute/snapshots/delete”   ],   "AssignableScopes": [     "/subscriptions/00000000-0000-0000-0000-000000000000",     "/subscriptions/11111111-1111-1111-1111-111111111111"   ] }

Which two actions can be performed by a user that is assigned the custom role? Each correct answer presents a complete solution.

Select all answers that apply.


Options:
```bash
Create and delete a snapshot.


Create and read a snapshot.


Create virtual machines.


Read all virtual machine settings.
```

21. You have an Azure subscription that contains multiple virtual machines.  

You need to ensure that a user named User1 can view all the resources in a resource group named RG1. You must use the principle of least privilege.

Which role should you assign to User1?

Select only one answer.


Options:
```bash
Billing Reader


Contributor


Reader


Tag Contributor
```

22. You have an Azure subscription and a user named User1.

You need to assign User1 a role that allows the user to create and manage all types of resources in the subscription. The solution must ensure that User1 is not able to assign roles to other users.

Which Azure role should you assign to User1?

Options:

```bash
API Management Service Contributor


Contributor


Owner


Reader
```

23. You have several management groups and Azure subscriptions.

You want to prevent the accidental deletion of resources.

To which three resource types can you apply delete locks? Each correct answer presents a complete solution.

Options:
```bash
management groups


resource groups


storage account data


subscriptions


virtual machines
```

24. You have an Azure subscription.

You plan to create an Azure Policy definition named Policy1.

You need to include remediation information in Policy.

To which definition section should you add remediation information for Policy1?

Select only one answer.


Options:
```bash
metadata


mode


parameters


policyRule
```

25. You have an Azure subscription that contains 10 virtual machines.

You need to ensure that a user named User1 can tag all the virtual machines by using the Azure portal. The solution must follow the principle of least privilege.

What should you do?

Select only one answer.


Options:
```bash
From the Azure portal, create a custom role that has the Microsoft.Compute virtual machines/*/write permission.


From the Azure portal, modify the Access control (IAM) settings of the virtual machines.


From the Azure portal, modify the Policies settings of the Azure subscription.


From the command line, run the az role assignment create command.
```


26. You have an Azure subscription that contains a resource group named RG1.

RG1 contains 10 resources.

You need to prevent the resources from being deleted accidentally. The solution must ensure that RG1 can be deleted if it no longer contains any resources.

What should you do?

Select only one answer.

```bash
From Azure Cloud Shell, run the New-AzureRmResourceGroup cmdlet.


From Azure Cloud Shell, run the Set-AzResourceGroup cmdlet.


From the Azure portal, add a tag on RG1.


From the Azure portal, add a lock on RG1.
```

27. You have an Azure subscription that contains a storage account named storage1.

You need to provide a partner organization with access to storage1. Access to storage1 must automatically expire after 24 hours.

What should you configure?

Select only one answer.

Options:
```bash
a shared access signature (SAS)


an access key


Azure Content Delivery Network (CDN)


lifecycle management
```

28. You have an Azure subscription that contains a storage account named storage1.

You need to ensure that access to storage1 is prevented from the internet.

What should you configure on storage1?

Select only one answer.


Options:
```bash
Access keys


Data protection


Encryption


Networking
```


29. You have an on-premises network.

You have an Azure subscription that contains a virtual network named VNet1. VNet1 is connected to the on-premises network by using ExpressRoute.

You perform the following actions:

Create a storage account named storage1
Associate VNet1 to storage1 and configure network routing to use Microsoft network routing.
You need to ensure that only connections from the on-premises network are allowed to access storage1. The solution must minimize administrative effort.

What should you do?

Select only one answer.

Options:
```bash
Configure the network settings of storage1.


Create a routing table. Add a filter rule to the table.


Create a shared access signature (SAS).


Create an ExpressRoute circuit. Create a filter on the ExpressRoute connection.
```

30. You have an Azure subscription that contains a storage account named storage1. storage1 contains an Azure Files share named share1.

You need to ensure that users can authenticate to share1 by using Microsoft Entra and access the file share by using SMB.

What should you do?

Select only one answer.

Options:
```bash
Configure identity-based access.


Generate a shared access signature (SAS) and a connection string.


Enable public network access.


Regenerate the access keys.
```

31. You have an Azure Storage account.

You need to copy data to the storage account by using the AzCopy tool.

Which two types of data storage are supported by AzCopy? Each correct answer presents a complete solution.

Options:
```bash
blob


file


queue


table
```

32. You have two premium block blob Azure Storage accounts named storage1 and storage2.

You need to configure object replication from storage1 to storage2.

Which three features should be enabled before configuring object replication? Each correct answer presents part of the solution.

Select all answers that apply.


Options:
```bash
blob versioning for storage1


blob versioning for storage2


change feed for storage1


change feed for storage2


point-in-time restore for the containers on storage1


point-in-time restore for the containers on storage2
```

33. You have an Azure subscription.

You plan to create a storage account named storage1.

You need to ensure that storage1 provides POSIX-compliant access control lists (ACLs).

Which option should you configure when creating storage1?

Select only one answer.

Options:
```bash
access tier


hierarchical namespace


SFTP


version-level immutable support
```

34. You have an Azure Storage account named storageaccount1 with a blob container named container1 that stores confidential information.

You need to ensure that content in container1 is not modified or deleted for six months after the last modification date.

What should you configure?

Select only one answer.


Options:
```bash
a custom Azure role


lifecycle management


the change feed


the immutability policy
```

35. You have an Azure subscription that contains a storage account.

You need to recommend a storage solution for storing infrequently accessed data. The solution must meet the following requirements:

The data must be stored for at least 90 days.
The data must be available within seconds.
Storage costs must be minimized.
Which tier should you recommend?

Select only one answer.

Options:
```bash
Cold


Cool


Hot


Premium
```

36. You have an Azure subscription that contains two virtual networks named VNet1 and VNet2.

You need to ensure that the resources on both VNet1 and VNet2 can communicate seamlessly between both networks.

What should you configure from the Azure portal?

Select only one answer.

Options:
```bash
connected devices


firewall


peerings

service endpoint
```

37. You have two Azure subscriptions named Sub1 and Sub2.

Sub1 contains a virtual network named VNet1 and a VPN gateway. Sub2 contains a virtual network named VNet2.

You have an on-premises device named Device1 that runs Windows and has a Point-to-Site (P2S) VPN client installed.

You configure network peering between VNet1 and VNet2.

You need to ensure that Device1 can access VNet2 when a VPN connection is established.

What should you do?

Options:
```bash
Create a private endpoint in Sub2.


Deploy Azure Front Door to Sub2.


Download and reinstall the P2S VPN client on Device1.


Run the New-SelfSignedCertificate cmdlet on Device1.
```

38. You have an Azure subscription that contains two resource groups named RG1 and RG2.

RG1 contains the following resources:

A virtual network named VNet1 located in the East US Azure region
A network security group (NSG) named NSG1 located in the West US Azure region
RG2 contains the following resources:

A virtual network named VNet2 located in the East US Azure region
A virtual network named VNet3 located in the West US Azure region
You need to associate NSG1.

To which subnets can you associate NSG1?

Select only one answer.


Options:
```bash
the subnets of all the virtual networks


the subnets of VNet1 only


the subnets of VNet1 and VNet2


the subnets of VNet3 only
```

39. You have an Azure virtual network that contains four subnets. Each subnet contains 10 virtual machines.

You plan to configure a network security group (NSG) that will allow inbound traffic over TCP port 8080 to two virtual machines on each subnet. The NSG will be associated to each subnet.

You need to recommend a solution to configure the inbound access by using the fewest number of NSG rules possible.

What should you use as the destination in the NSG?

Select only one answer.

Options:
```bash
an application security group


a service tag


the subnets of the virtual machines
```


40. You create several Azure virtual machines that run Windows Server.

You need to connect to the virtual machines without exposing RDP ports over the internet.

Which Azure service should you deploy?

Select only one answer.


Options:
```bash
Azure Bastion


Azure Front Door


Azure Network Watcher


Azure Virtual Desktop
```

41. Your company plans to migrate servers from on-premises to Azure. There will be dev, test, and production virtual machines on a single virtual network.

You need to restrict traffic between the dev, test, and production virtual machines to specific ports.

What should you use?

Select only one answer.

Options
```bash
a network security group (NSG)


an Azure firewall


an Azure load balancer


an Azure virtual network
```

42. You have an Azure subscription that contains a resource group named RG1.

You plan to create and configure a network security group (NSG) named NSG1 for the following types of traffic:

Remote Desktop Management
HTTP
NSG1 will be used on the subnets of multiple virtual networks.

Which two cmdlets should you run? Each correct answer presents part of the solution.

Select all answers that apply.


Options:
```bash
Add-AzLoadBalancerFrontendIpConfig


Add-AzNetworkInterfaceTapConfig


New-AzNetworkSecurityGroup


New-AzNetworkSecurityRuleConfig
```

43. You have an Azure subscription that contains an ASP.NET application. The application is hosted on four Azure virtual machines that run Windows Server.

You have a load balancer named LB1 that load balances requests to the virtual machines.

You need to ensure that site users connect to the same web server for all requests made to the application.

Which two actions should you perform? Each correct answer presents part of the solution.

Select all answers that apply.

Options:
```bash
Configure an inbound NAT rule.


Set Session persistence to Client IP.


Set Session persistence to None.


Set Session persistence to Protocol.
```

44. You have an Azure subscription.

You plan to implement four Azure virtual networks that will be peered. All virtual machines will use a DNS suffix of contoso.com.

You need to configure name resolution for the virtual networks to ensure that all the virtual machines can communicate by using their FQDNs. The solution must minimize administrative effort.

What should you use?

Select only one answer.


Options:
```bash
a DNS server on an Azure virtual machine


an Azure Private DNS zone


an Azure public DNS zone


Azure-provided name resolution
```

45. You have an Azure subscription that contains an Azure DNS zone named contoso.com.

You add a new subdomain named test.contoso.com.

You plan to delegate test.contoso.com to a different DNS server.

How should you configure the domain delegation?

Select only one answer.

Options:
```bash
Add an A record for test.contoso.com.


Add an NS record set named test to the contoso.com zone.


Create the SOA record for test.contoso.com.


Modify the A record for contoso.com.
```