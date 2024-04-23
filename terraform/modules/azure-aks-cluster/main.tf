resource "azurerm_resource_group" "group" {
  name     = local.resource_group_name
  location = var.location
}

resource "azurerm_kubernetes_cluster" "cluster" {
  resource_group_name = azurerm_resource_group.group.name
  name                = var.name
  dns_prefix          = var.name
  location            = azurerm_resource_group.group.location
  kubernetes_version  = var.kubernetes_version

  default_node_pool {
    name                = "default"
    type                = "VirtualMachineScaleSets"
    enable_auto_scaling = false
    node_count          = var.default_node_pool_vm_count
    vm_size             = var.default_node_pool_vm_size
    os_disk_size_gb     = var.default_node_pool_vm_disk_size


    # TO DO: taint all AwayZone's clusters default node pool to prevent workload to be scheduled on it.
    # `default_node_pool` will still be present here, as at least 1 is required by AKS:
    #
    # only_critical_addons_enabled = true
    #
    # WARNING:
    #   - Changing the value will force cluster to be recreated
    #
    # Make sure the taint is set here only for new clusters. For existing clusters, please taint nodes
    # in the default pool externally, without changing the terraform state
    # e.g.:
    #   kubectl taint nodes <node name> CriticalAddonsOnly=true:NoSchedule
    #
    # TO DO: implement pulumi programs to manage the actual workload node pools, to make it easier to
    # increase/decrease node count, increase disk size, add taints, etc, that are dynamic and
    # recurring operations that don't fit so well in the terraform way of life.
  }

  identity {
    type = "SystemAssigned"
  }

  tags = {
    environment = var.environment
  }
}
