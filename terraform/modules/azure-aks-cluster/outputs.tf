output "resource_group_name" {
  value       = azurerm_resource_group.group.name
  description = "Name of the generated resource group, that contains the resources for this managed kubernetest cluster"
}

output "cluster_id" {
  value       = azurerm_kubernetes_cluster.cluster.id
  description = "ID of the kubernetes cluster"
}

output "cluster_name" {
  value       = azurerm_kubernetes_cluster.cluster.name
  description = "Name of the kubernetes cluster"
}

output "cluster_kube_config" {
  value       = azurerm_kubernetes_cluster.cluster.kube_config_raw
  description = "Raw Kubernetes config to be used by kubectl and other compatible tools."
  sensitive   = true
}
