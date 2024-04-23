variable "location" {
  description = "Azure location (region)"
  type        = string
}

variable "name" {
  description = "The nae of the cluster to be created"
  type        = string
}

variable "default_node_pool_vm_count" {
  description = "AKS cluster desired node count for the default (first) node pool"
  type        = number
  default     = 1
}

variable "default_node_pool_vm_size" {
  description = "Default node pool VM size"
  type        = string
  default     = "Standard_B2s"
}

variable "default_node_pool_vm_disk_size" {
  description = "The size of the OS Disk (GB) which should be used for each agent in the Node Pool"
  type        = number
  default     = 30
}

variable "environment" {
  description = "Type of environment on which this cluster will be used. One of: dev|staging|production"
  type        = string
  default     = "dev"
}

variable "kubernetes_version" {
  description = "Set the kubernetes version used in control plane"
  type        = string
  # For a list of AKS supported kubernetes version please refer to:
  # https://learn.microsoft.com/en-us/azure/aks/supported-kubernetes-versions?tabs=azure-cli#aks-kubernetes-release-calendar
  default = "1.27.9"
}
