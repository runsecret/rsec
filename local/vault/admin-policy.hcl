# vault-admin-policy.hcl
# This policy grants administrative capabilities to manage all aspects of Vault
# while following the principle of least privilege where possible.

# Allow management of the token store
path "auth/token/*" {
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

# Allow management of all auth methods
path "auth/*" {
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

# Allow management of all secret engines
path "sys/mounts/*" {
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

# Allow management and usage of database secrets engine
path "database/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}

# Allow access to database credentials
path "database/creds/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}

# Allow database configuration
path "database/config/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}

# Allow database roles management
path "database/roles/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}

# Allow management of all policies
path "sys/policies/*" {
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}
path "sys/policy/*" {
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

# Allow listing and viewing all secrets
path "secret/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}

# Allow management of KV secrets engine (v1 and v2)
path "kv/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}
path "+/data/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}
path "+/metadata/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}

# Allow management of system configuration
path "sys/config/*" {
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

# Allow management of system health and status
path "sys/health" {
  capabilities = ["read", "sudo"]
}
path "sys/seal" {
  capabilities = ["read", "update", "sudo"]
}
path "sys/seal-status" {
  capabilities = ["read"]
}
path "sys/unseal" {
  capabilities = ["read", "update"]
}

# Allow management of audit devices
path "sys/audit*" {
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

# Allow management of leases
path "sys/leases/*" {
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

# Allow management of storage
path "sys/storage/*" {
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

# Allow management of capabilities and permissions
path "sys/capabilities*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}

# Allow management of audit logs
path "sys/audit-hash/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}

# Allow management of plugins
path "sys/plugins*" {
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

# Allow management of replication
path "sys/replication/*" {
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

# Allow tools functionality
path "sys/tools/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}

# Allow rotation of encryption keys
path "sys/rotate" {
  capabilities = ["update", "sudo"]
}

# Allow management of system initialization
path "sys/init" {
  capabilities = ["read", "update"]
}

# Allow root token generation (should be used very carefully)
path "sys/generate-root/*" {
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

# Allow license management
path "sys/license" {
  capabilities = ["read", "update"]
}

# Allow reading metrics
path "sys/metrics" {
  capabilities = ["read"]
}
