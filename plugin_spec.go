package operation

import "encoding/json"

// Declaration is the JSON manifest shape for an operation exposed by an SDK,
// runtime, or external provider. It is intentionally runtime-neutral.
type Declaration struct {
	Name           string          `json:"name"`
	Description    string          `json:"description,omitempty"`
	Input          json.RawMessage `json:"input_schema,omitempty"`
	Output         json.RawMessage `json:"output_schema,omitempty"`
	ReadOnly       bool            `json:"read_only,omitempty"`
	Compact        bool            `json:"compact,omitempty"`
	SecretPurposes []string        `json:"secret_purposes,omitempty"`
	Effects        []Effect        `json:"effects,omitempty"`
	Risk           RiskLevel       `json:"risk,omitempty"`
	Idempotency    Idempotency     `json:"idempotency,omitempty"`
	Access         []Access        `json:"access,omitempty"`
	AuthScopes     []string        `json:"auth_scopes,omitempty"`
	Render         *RenderSpec     `json:"render,omitempty"`
}

const (
	// EffectRead and EffectWrite are compact manifest effect declarations.
	// The richer read_external/write_external values remain available for
	// semantic policy claims.
	EffectRead        Effect = "read"
	EffectWrite       Effect = "write"
	EffectBrowser     Effect = "browser"
	EffectLocalSystem Effect = "local_system"
)

// Access describes one external capability or privilege an operation requires.
type Access string

const (
	AccessNone        Access = "none"
	AccessAuth        Access = "auth"
	AccessSecret      Access = "secret"
	AccessNetwork     Access = "network"
	AccessProvider    Access = "provider"
	AccessProcess     Access = "process"
	AccessBrowser     Access = "browser"
	AccessFilesystem  Access = "filesystem"
	AccessLocalSystem Access = "local_system"
)

// RenderSpec declares preferred render formats for operation output.
type RenderSpec struct {
	Preferred string   `json:"preferred,omitempty"`
	Formats   []string `json:"formats,omitempty"`
}
