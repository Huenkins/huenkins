package globalcontext

import (
	"errors"
	"sync"
)

/*

	How to use:

	print.Printf(
		`String is: {{.StringData}}, Int is: {{.IntData}}, Bool is: {{.BoolData}}`,
		print.Add("StringData", "stringData", "IntData", 9999, "BoolData", true)
	)

*/

// Username with password
// Docker Host Certificate Authentication
// SSH Username with private key
// Sauce Labs
// Secret file
// Secret text
// Vault App Role Credential
// Vault GCP Credential
// Vault Github Token Credential
// Vault Token Credential
// Vault Token File Credential
// Certificate

const (
	NewCredentialDefaultID = "new_credential_default_id"
)

type CredentialType struct {
	ID       string `json:"key"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	PluginID string `json:"plugin_id"`
}

type Credential struct {
	TypeID      string   `json:"type"`
	IsGlobal    bool     `json:"is_global"`
	ScopeIDs    []string `json:"scope_ids"`
	Description string   `json:"description"`

	Saved map[string][]byte `json:"save"`
	View  map[string]string `json:"view"`
}

type AllCredentials struct {
	// safeness :)
	mu sync.RWMutex `json:"-"`

	// Credential type => Credentials Type Info
	CredentialTypes map[string]*CredentialType `json:"credential_types"`

	// Credential ID => Credentials Info
	Credentials map[string]*Credential `json:"credentials"`

	// ID Credential => => Credentials Type
	Types map[string]string `json:"view"`
}

func (cr *AllCredentials) Save(globalContext *GlobalContext, typeID, ID string, data map[string][]byte) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	if typeID == "" && ID == "" {
		return errors.New("No valid data for check type of credential.")
	}

	var credential *Credential
	var credentialType *CredentialType
	var find bool
	if ID != "" {
		credential, find = cr.Credentials[ID]
		if !find {
			return errors.New("Unknown credential.")
		}

		credentialType, find = cr.CredentialTypes[credential.TypeID]
		if !find {
			return errors.New("Unknown credential type.")
		}
	} else {
		credentialType, find = cr.CredentialTypes[typeID]
		if !find {
			return errors.New("Unknown credential type.")
		}
		credential = &Credential{
			TypeID:      typeID,
			ScopeIDs:    []string{},
			Description: "",
			Saved:       map[string][]byte{},
			View:        map[string]string{},
		}
	}

	plugin, err := globalContext.GetPlugin(credentialType.PluginID)
	if err != nil {
		return err
	}

	viewData, saveData, err := plugin.CheckCredential(globalContext, data)
	if err != nil {
		return err
	}

	credential.Saved = saveData
	credential.View = viewData

	return nil
}
