terraform {
  required_providers {
    keycloak = {
      source  = "mrparkers/keycloak"
      version = "~> 5.0"
    }
  }
}

provider "keycloak" {
  client_id  = "admin-cli"
  realm      = "master"
  url        = "http://localhost:8080"
  username   = "admin"
  password   = "admin"
}

# Create a new realm
resource "keycloak_realm" "my_realm" {
  realm   = "my_realm"
  enabled = true
}

# Create a client for the Go API
resource "keycloak_openid_client" "go_api_client" {
  realm_id                     = keycloak_realm.my_realm.id
  client_id                    = "go_api_client"
  enabled                      = true
  access_type                  = "CONFIDENTIAL"
  standard_flow_enabled        = true
  direct_access_grants_enabled = true
  client_secret                = "go-api-secret"
}

# Create roles
resource "keycloak_role" "api_user_role" {
  realm_id = keycloak_realm.my_realm.id
  name     = "api_user"
}

# Create a test user
resource "keycloak_user" "test_user" {
  realm_id    = keycloak_realm.my_realm.id
  username    = "testuser"
  enabled     = true
  email       = "testuser@example.com"
  first_name  = "Test"
  last_name   = "User"
  initial_password {
    value     = "Password123"
    temporary = false
  }
}

# Assign the role to the user
resource "keycloak_user_roles" "test_user_roles" {
  realm_id = keycloak_realm.my_realm.id
  user_id  = keycloak_user.test_user.id
  roles    = [
    keycloak_role.api_user_role.id
  ]
}
