// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package settings

import (
	"encoding/json"
	"fmt"
)

type Ipsec struct {
	BaseSetting

	IKEv2ReauthenticationMethod string `json:"ikev2_reauthentication_method,omitempty"` // make-before-break|break-before-make
}

func (dst *Ipsec) UnmarshalJSON(b []byte) error {
	type Alias Ipsec
	aux := &struct{ *Alias }{Alias: (*Alias)(dst)}

	if err := json.Unmarshal(b, &dst.BaseSetting); err != nil {
		return fmt.Errorf("unable to unmarshal base setting: %w", err)
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	return nil
}
