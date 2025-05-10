curl \
    --header "X-Vault-Token: dev-only-token" \
    --request POST \
    --data '{ "data": {"pass": "my-password", "username":"my-username"} }' \
    http://hashicorp-vault:8200/v1/secret/data/my-project
