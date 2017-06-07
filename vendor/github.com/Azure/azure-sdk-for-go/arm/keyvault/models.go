package keyvault

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/satori/uuid"
	"net/http"
)

// CertificatePermissions enumerates the values for certificate permissions.
type CertificatePermissions string

const (
	// All specifies the all state for certificate permissions.
	All CertificatePermissions = "all"
	// Create specifies the create state for certificate permissions.
	Create CertificatePermissions = "create"
	// Delete specifies the delete state for certificate permissions.
	Delete CertificatePermissions = "delete"
	// Deleteissuers specifies the deleteissuers state for certificate
	// permissions.
	Deleteissuers CertificatePermissions = "deleteissuers"
	// Get specifies the get state for certificate permissions.
	Get CertificatePermissions = "get"
	// Getissuers specifies the getissuers state for certificate permissions.
	Getissuers CertificatePermissions = "getissuers"
	// Import specifies the import state for certificate permissions.
	Import CertificatePermissions = "import"
	// List specifies the list state for certificate permissions.
	List CertificatePermissions = "list"
	// Listissuers specifies the listissuers state for certificate permissions.
	Listissuers CertificatePermissions = "listissuers"
	// Managecontacts specifies the managecontacts state for certificate
	// permissions.
	Managecontacts CertificatePermissions = "managecontacts"
	// Manageissuers specifies the manageissuers state for certificate
	// permissions.
	Manageissuers CertificatePermissions = "manageissuers"
	// Setissuers specifies the setissuers state for certificate permissions.
	Setissuers CertificatePermissions = "setissuers"
	// Update specifies the update state for certificate permissions.
	Update CertificatePermissions = "update"
)

// KeyPermissions enumerates the values for key permissions.
type KeyPermissions string

const (
	// KeyPermissionsAll specifies the key permissions all state for key
	// permissions.
	KeyPermissionsAll KeyPermissions = "all"
	// KeyPermissionsBackup specifies the key permissions backup state for key
	// permissions.
	KeyPermissionsBackup KeyPermissions = "backup"
	// KeyPermissionsCreate specifies the key permissions create state for key
	// permissions.
	KeyPermissionsCreate KeyPermissions = "create"
	// KeyPermissionsDecrypt specifies the key permissions decrypt state for
	// key permissions.
	KeyPermissionsDecrypt KeyPermissions = "decrypt"
	// KeyPermissionsDelete specifies the key permissions delete state for key
	// permissions.
	KeyPermissionsDelete KeyPermissions = "delete"
	// KeyPermissionsEncrypt specifies the key permissions encrypt state for
	// key permissions.
	KeyPermissionsEncrypt KeyPermissions = "encrypt"
	// KeyPermissionsGet specifies the key permissions get state for key
	// permissions.
	KeyPermissionsGet KeyPermissions = "get"
	// KeyPermissionsImport specifies the key permissions import state for key
	// permissions.
	KeyPermissionsImport KeyPermissions = "import"
	// KeyPermissionsList specifies the key permissions list state for key
	// permissions.
	KeyPermissionsList KeyPermissions = "list"
	// KeyPermissionsRestore specifies the key permissions restore state for
	// key permissions.
	KeyPermissionsRestore KeyPermissions = "restore"
	// KeyPermissionsSign specifies the key permissions sign state for key
	// permissions.
	KeyPermissionsSign KeyPermissions = "sign"
	// KeyPermissionsUnwrapKey specifies the key permissions unwrap key state
	// for key permissions.
	KeyPermissionsUnwrapKey KeyPermissions = "unwrapKey"
	// KeyPermissionsUpdate specifies the key permissions update state for key
	// permissions.
	KeyPermissionsUpdate KeyPermissions = "update"
	// KeyPermissionsVerify specifies the key permissions verify state for key
	// permissions.
	KeyPermissionsVerify KeyPermissions = "verify"
	// KeyPermissionsWrapKey specifies the key permissions wrap key state for
	// key permissions.
	KeyPermissionsWrapKey KeyPermissions = "wrapKey"
)

// SecretPermissions enumerates the values for secret permissions.
type SecretPermissions string

const (
	// SecretPermissionsAll specifies the secret permissions all state for
	// secret permissions.
	SecretPermissionsAll SecretPermissions = "all"
	// SecretPermissionsDelete specifies the secret permissions delete state
	// for secret permissions.
	SecretPermissionsDelete SecretPermissions = "delete"
	// SecretPermissionsGet specifies the secret permissions get state for
	// secret permissions.
	SecretPermissionsGet SecretPermissions = "get"
	// SecretPermissionsList specifies the secret permissions list state for
	// secret permissions.
	SecretPermissionsList SecretPermissions = "list"
	// SecretPermissionsSet specifies the secret permissions set state for
	// secret permissions.
	SecretPermissionsSet SecretPermissions = "set"
)

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Premium specifies the premium state for sku name.
	Premium SkuName = "premium"
	// Standard specifies the standard state for sku name.
	Standard SkuName = "standard"
)

// AccessPolicyEntry is an identity that have access to the key vault. All
// identities in the array must use the same tenant ID as the key vault's
// tenant ID.
type AccessPolicyEntry struct {
	TenantID      *uuid.UUID   `json:"tenantId,omitempty"`
	ObjectID      *string      `json:"objectId,omitempty"`
	ApplicationID *uuid.UUID   `json:"applicationId,omitempty"`
	Permissions   *Permissions `json:"permissions,omitempty"`
}

// Permissions is permissions the identity has for keys, secrets and
// certificates.
type Permissions struct {
	Keys         *[]KeyPermissions         `json:"keys,omitempty"`
	Secrets      *[]SecretPermissions      `json:"secrets,omitempty"`
	Certificates *[]CertificatePermissions `json:"certificates,omitempty"`
}

// Resource is key Vault resource
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// ResourceListResult is list of vault resources.
type ResourceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Resource `json:"value,omitempty"`
	NextLink          *string     `json:"nextLink,omitempty"`
}

// ResourceListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ResourceListResult) ResourceListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Sku is sKU details
type Sku struct {
	Family *string `json:"family,omitempty"`
	Name   SkuName `json:"name,omitempty"`
}

// Vault is resource information with extended details.
type Vault struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Properties        *VaultProperties    `json:"properties,omitempty"`
}

// VaultCreateOrUpdateParameters is parameters for creating or updating a vault
type VaultCreateOrUpdateParameters struct {
	Location   *string             `json:"location,omitempty"`
	Tags       *map[string]*string `json:"tags,omitempty"`
	Properties *VaultProperties    `json:"properties,omitempty"`
}

// VaultListResult is list of vaults
type VaultListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Vault `json:"value,omitempty"`
	NextLink          *string  `json:"nextLink,omitempty"`
}

// VaultListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client VaultListResult) VaultListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// VaultProperties is properties of the vault
type VaultProperties struct {
	VaultURI                     *string              `json:"vaultUri,omitempty"`
	TenantID                     *uuid.UUID           `json:"tenantId,omitempty"`
	Sku                          *Sku                 `json:"sku,omitempty"`
	AccessPolicies               *[]AccessPolicyEntry `json:"accessPolicies,omitempty"`
	EnabledForDeployment         *bool                `json:"enabledForDeployment,omitempty"`
	EnabledForDiskEncryption     *bool                `json:"enabledForDiskEncryption,omitempty"`
	EnabledForTemplateDeployment *bool                `json:"enabledForTemplateDeployment,omitempty"`
}
