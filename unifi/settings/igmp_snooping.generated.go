// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package settings

import (
	"encoding/json"
	"fmt"
)

type IgmpSnooping struct {
	BaseSetting

	Enabled                              bool     `json:"enabled"`
	FailoverQuerier                      string   `json:"failover_querier,omitempty"`
	FastleaveForNetworkIDs               []string `json:"fastleave_for_network_ids,omitempty"`
	FloodKnownProtocols                  bool     `json:"flood_known_protocols"`
	FloodUnknownMulticastForNetworkIDs   []string `json:"flood_unknown_multicast_for_network_ids,omitempty"`
	ForwardUnknownMcastRouterPorts       bool     `json:"forward_unknown_mcast_router_ports"`
	NetworkIDs                           []string `json:"network_ids,omitempty"`
	PrimaryQuerier                       string   `json:"primary_querier,omitempty"`
	QuerierAddresses                     []string `json:"querier_addresses,omitempty"`
	QuerierMode                          string   `json:"querier_mode,omitempty"`              // OFF|ON|AUTO
	QuerierSubscriptionMode              string   `json:"querier_subscription_mode,omitempty"` // ALL|CUSTOM
	QuerierSwitches                      []string `json:"querier_switches,omitempty"`
	SubscriptionMode                     string   `json:"subscription_mode,omitempty"` // ALL|CUSTOM
	Switches                             []string `json:"switches,omitempty"`
}

func (dst *IgmpSnooping) UnmarshalJSON(b []byte) error {
	type Alias IgmpSnooping
	aux := &struct{ *Alias }{Alias: (*Alias)(dst)}

	if err := json.Unmarshal(b, &dst.BaseSetting); err != nil {
		return fmt.Errorf("unable to unmarshal base setting: %w", err)
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	return nil
}
