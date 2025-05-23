sleep 2
# Initialize Vault
# Set Token
export VAULT_TOKEN=dev-only-token
export VAULT_ADDR=http://hashicorp-vault:8200

# Initialize secret in KVv2 engine
vault kv put secret/my-secret \
    username="my-username" \
    pass="my-password"

# Setup postgres database secret engine for testing
vault secrets enable database
vault write database/config/my-database \
    plugin_name=postgresql-database-plugin \
    allowed_roles="readonly" \
    connection_url="postgresql://{{username}}:{{password}}@postgres:5432/postgres?sslmode=disable" \
    username="postgres" \
    password="postgres"
vault write database/roles/readonly \
    db_name=my-database \
    creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}';" \
    default_ttl="1h" \
    max_ttl="24h" \
    revocation_statements="REVOKE ALL PRIVILEGES ON ALL TABLES IN SCHEMA public FROM \"{{name}}\";"

# Create Admin policy
vault policy write admin /usr/local/bin/admin-policy.hcl

# Enable Userpass auth method
vault auth enable userpass
vault write auth/userpass/users/test-user \
    password=my-password \
    policies=admin

# Enable AppRole auth method
vault auth enable approle
vault write auth/approle/role/test-role \
    secret_id_ttl=60m \
    token_num_uses=10 \
    token_ttl=20m \
    token_max_ttl=30m \
    policies=admin
