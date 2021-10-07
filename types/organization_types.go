package types

import "time"

type Organization struct {
	// The name of the organization.
	Name string `json:"name,omitempty"`
	// The short name of the organization.
	ShortName string `json:"short_name,omitempty"`
	// The identifier of the organization.
	UUID      string                 `json:"UUID,omitempty"`
	Policies  *OrganizationPolicies  `json:"policies,omitempty"`
	Authority *OrganizationAuthority `json:"authority,omitempty"`
	Time      *OrganizationTime      `json:"time,omitempty"`
}

type OrganizationAuthority struct {
	ActiveRoles []string `json:"active_roles,omitempty"`
}

type OrganizationPolicies struct {
	// The number of CVE IDs the organization is allowed to have in the RESERVED state at one time.
	IdQuota int32 `json:"id_quota,omitempty"`
}

type OrganizationTime struct {
	// The time the organization was created.
	Created time.Time `json:"created,omitempty"`
	// The last time the organization was modified.
	Modified time.Time `json:"modified,omitempty"`
}
