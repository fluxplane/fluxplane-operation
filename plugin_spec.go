package operation

import (
	"encoding/json"

	datasource "github.com/fluxplane/fluxplane-datasource"
)

// PluginSpec is the legacy plugin-facing operation declaration used by dex-style plugin manifests.
// It is intentionally runtime-neutral and suitable for reuse by plugin SDKs and engines.
type PluginSpec struct {
	Name           string            `json:"name"`
	Description    string            `json:"description,omitempty"`
	Input          json.RawMessage   `json:"input_schema,omitempty"`
	Output         json.RawMessage   `json:"output_schema,omitempty"`
	ReadOnly       bool              `json:"read_only,omitempty"`
	Compact        bool              `json:"compact,omitempty"`
	SecretPurposes []string          `json:"secret_purposes,omitempty"`
	Effects        []PluginEffect    `json:"effects,omitempty"`
	Risk           PluginRisk        `json:"risk,omitempty"`
	Idempotency    PluginIdempotency `json:"idempotency,omitempty"`
	Access         []Access          `json:"access,omitempty"`
	AuthScopes     []string          `json:"auth_scopes,omitempty"`
	Render         *PluginRenderSpec `json:"render,omitempty"`
}

// PluginEffect is the legacy plugin manifest effect vocabulary.
type PluginEffect string

const (
	PluginEffectRead        PluginEffect = "read"
	PluginEffectWrite       PluginEffect = "write"
	PluginEffectNetwork     PluginEffect = "network"
	PluginEffectProcess     PluginEffect = "process"
	PluginEffectBrowser     PluginEffect = "browser"
	PluginEffectFilesystem  PluginEffect = "filesystem"
	PluginEffectLocalSystem PluginEffect = "local_system"
)

// PluginRisk is the legacy plugin manifest risk vocabulary.
type PluginRisk string

const (
	PluginRiskLow         PluginRisk = "low"
	PluginRiskMedium      PluginRisk = "medium"
	PluginRiskHigh        PluginRisk = "high"
	PluginRiskDestructive PluginRisk = "destructive"
)

// PluginIdempotency is the legacy plugin manifest idempotency vocabulary.
type PluginIdempotency string

const (
	PluginIdempotent    PluginIdempotency = "idempotent"
	PluginNonIdempotent PluginIdempotency = "non_idempotent"
	PluginConditional   PluginIdempotency = "conditional"
	PluginUnknown       PluginIdempotency = "unknown"
)

// Access describes one external capability or privilege an operation requires.
type Access = datasource.Access

const (
	AccessNone        = datasource.AccessNone
	AccessAuth        = datasource.AccessAuth
	AccessSecret      = datasource.AccessSecret
	AccessNetwork     = datasource.AccessNetwork
	AccessProvider    = datasource.AccessProvider
	AccessProcess     = datasource.AccessProcess
	AccessBrowser     = datasource.AccessBrowser
	AccessFilesystem  = datasource.AccessFilesystem
	AccessLocalSystem = datasource.AccessLocalSystem
)

// PluginRenderSpec declares preferred render formats for plugin operation output.
type PluginRenderSpec struct {
	Preferred string   `json:"preferred,omitempty"`
	Formats   []string `json:"formats,omitempty"`
}
