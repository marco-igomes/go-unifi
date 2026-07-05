package unifi

import "context"

// ConsoleApp is one installed UniFi OS application as reported by
// /api/system -> apps.controllers[]. Name is the app slug ("network",
// "protect", "access", ...); Version is its human-readable version.
type ConsoleApp struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	VersionRaw   string `json:"versionRaw"`
	InstallState string `json:"installState"`
	State        string `json:"state"`
}

// ConsoleSystem is the subset of the UniFi OS /api/system response the SDK
// exposes: the console firmware/uCore versions plus every installed
// application. Unlike Sysinfo (which reads the Network app's own
// /stat/sysinfo and therefore only knows the Network version), this covers all
// UniFi OS apps (Protect, Access, Talk, ...), so callers can track them.
type ConsoleSystem struct {
	Name         string `json:"name"`
	UcoreVersion string `json:"ucore_version"`
	Hardware     struct {
		Name            string `json:"name"`
		Shortname       string `json:"shortname"`
		FirmwareVersion string `json:"firmwareVersion"`
	} `json:"hardware"`
	Apps struct {
		Controllers []ConsoleApp `json:"controllers"`
	} `json:"apps"`
}

// ConsoleSystem fetches the UniFi OS /api/system info: console firmware and the
// list of installed applications with their versions. The leading slash targets
// the console root, bypassing the "/proxy/network" API path.
func (c *ApiClient) ConsoleSystem(ctx context.Context) (*ConsoleSystem, error) {
	var respBody ConsoleSystem
	if err := c.do(ctx, "GET", "/api/system", nil, &respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}
