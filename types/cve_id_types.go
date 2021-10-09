package types

import (
	"github.com/antihax/optional"
	"time"
)

type CveId struct {
	CveId     string      `json:"cve_id"`
	CveYear   string      `json:"cve_year"`
	OwningCNA string      `json:"owning_cna"`
	State     string      `json:"state"`
	RequestBy RequestedBy `json:"request_by"`
	Reserved  time.Time   `json:"reserved"`
	Time      CveIdTime   `json:"time"`
}

type CveIdNoTime struct {
	CveId   string `json:"cve_id,omitempty"`
	CveYear string `json:"cve_year,omitempty"`
	// The shortname for the organization that owns the CVE-ID.
	OwningCna   string            `json:"owning_cna,omitempty"`
	State       string            `json:"state,omitempty"`
	RequestedBy *CveIdRequestedBy `json:"requested_by,omitempty"`
	// The time the ID was reserved.
	Reserved time.Time `json:"reserved,omitempty"`
}

type CveIdRequestedBy struct {
	// The shortname for the organization of the user that requested the ID.
	Cna string `json:"cna,omitempty"`
	// The username for the account that requested the ID.
	User string `json:"user,omitempty"`
}

type CveIdTime struct {
	// The time the CVE ID was assigned.
	Created time.Time `json:"created,omitempty"`
	// The last time the CVE ID was modified.
	Modified time.Time `json:"modified,omitempty"`
}

type ListCveIdsOpts struct {
	State          optional.String
	CveIdYear      optional.Int32
	TimeReservedLt optional.Time
	TimeReservedGt optional.Time
	TimeModifiedLt optional.Time
	TimeModifiedGt optional.Time
	Page           optional.Int32
}

type ListCveIdsResponse struct {
	// Total CVE records found
	TotalCount int32 `json:"totalCount,omitempty"`
	// Number of CVE records in a page
	ItemsPerPage int32 `json:"itemsPerPage,omitempty"`
	// Total number of pages
	PageCount int32 `json:"pageCount,omitempty"`
	// Current page
	CurrentPage int32 `json:"currentPage,omitempty"`
	// Previous page
	PrevPage int32 `json:"prevPage,omitempty"`
	// Next page
	NextPage int32    `json:"nextPage,omitempty"`
	CveIds   *[]CveId `json:"cve_ids,omitempty"`
}

type RequestedBy struct {
	CNA  string `json:"cna"`
	user string `json:"user"`
}

type ReserveCveIdOpts struct {
	// Valid values are sequential or non-sequential
	BatchType optional.String
}

type ReserveCveIdResponse struct {
	CveIds *[]CveIdNoTime `json:"cve_ids,omitempty"`
}
