// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package settings

import (
	"encoding/json"
	"fmt"
)

type GlobalNetwork struct {
	BaseSetting

	DefaultSecurityPosture string `json:"default_security_posture,omitempty"` // ALLOW_ALL|BLOCK_ALL
}

func (dst *GlobalNetwork) UnmarshalJSON(b []byte) error {
	type Alias GlobalNetwork
	aux := &struct{ *Alias }{Alias: (*Alias)(dst)}

	if err := json.Unmarshal(b, &dst.BaseSetting); err != nil {
		return fmt.Errorf("unable to unmarshal base setting: %w", err)
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	return nil
}
