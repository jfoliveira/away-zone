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
  }

  identity {
    type = "SystemAssigned"
  }

  tags = {
    environment = var.environment
  }
}
