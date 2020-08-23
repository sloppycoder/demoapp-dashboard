package api

import (
	"bytes"
	"encoding/json"
)

type Customer struct {
	CustomerId string `json:"customer_id,omitempty"`
	Name       string `json:"name,omitempty"`
	LoginName  string `json:"login_name,omitempty"`
}

type CasaAccountStatus int32

const (
	CasaAccountACTIVE  CasaAccountStatus = 0
	CasaAccountBLOCKED CasaAccountStatus = 1
	CasaAccountDORMANT CasaAccountStatus = 2
)

// Enum value maps for CasaAccountStatus.
var (
	CasaAccountStatusName = map[CasaAccountStatus]string{
		CasaAccountACTIVE:  "ACTIVE",
		CasaAccountBLOCKED: "BLOCKED",
		CasaAccountDORMANT: "DORMANT",
	}
	CasaAccountStatusValue = map[string]CasaAccountStatus{
		"ACTIVE":  CasaAccountACTIVE,
		"BLOCKED": CasaAccountBLOCKED,
		"DORMANT": CasaAccountDORMANT,
	}
)

func (s CasaAccountStatus) String() string {
	return CasaAccountStatusName[s]
}

// MarshalJSON marshals the enum as a quoted json string
func (s CasaAccountStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(CasaAccountStatusName[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmarshals a quoted json string to the enum value
func (s *CasaAccountStatus) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value
	*s = CasaAccountStatusValue[j]
	return nil
}

type CasaAccount struct {
	AccountId string            `json:"account_id,omitempty"`
	Nickname  string            `json:"nickname,omitempty"`
	ProdCode  string            `json:"prod_code,omitempty"`
	ProdName  string            `json:"prod_name,omitempty"`
	Currency  string            `json:"currency,omitempty"`
	Status    CasaAccountStatus `json:"status,omitempty"`
	Balances  []*Balance        `json:"balances,omitempty"`
}

type BalanceType int32

const (
	BalanceCURRENT   BalanceType = 0
	BalanceAVAILABLE BalanceType = 1
)

// Enum value maps for BalanceType.
var (
	BalanceTypeName = map[BalanceType]string{
		BalanceCURRENT:   "CURRENT",
		BalanceAVAILABLE: "AVAILABLE",
	}
	BalanceTypeValue = map[string]BalanceType{
		"CURRENT":   BalanceCURRENT,
		"AVAILABLE": BalanceAVAILABLE,
	}
)

func (s BalanceType) String() string {
	return BalanceTypeName[s]
}

// MarshalJSON marshals the enum as a quoted json string
func (s BalanceType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(BalanceTypeName[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmarshals a quoted json string to the enum value
func (s *BalanceType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value
	*s = BalanceTypeValue[j]
	return nil
}

type Balance struct {
	Amount     float64     `json:"amount,omitempty"`
	Type       BalanceType `json:"type,omitempty"` // balance type
	CreditFlag bool        `json:"credit_flag,omitempty"`
}

type Dashboard struct {
	Customer *Customer      `json:"customer,omitempty"`
	Casa     []*CasaAccount `json:"casa,omitempty"`
}
