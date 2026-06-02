package operation

import (
	"encoding/json"
	"testing"
)

func TestPluginSpecJSON(t *testing.T) {
	spec := PluginSpec{
		Name:        "test.write",
		Description: "Write something",
		ReadOnly:    false,
		Effects:     []PluginEffect{PluginEffectWrite, PluginEffectNetwork},
		Risk:        PluginRiskHigh,
		Idempotency: PluginNonIdempotent,
		Access:      []Access{AccessNetwork},
		Render:      &PluginRenderSpec{Preferred: "json", Formats: []string{"json"}},
	}
	raw, err := json.Marshal(spec)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded PluginSpec
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if decoded.Name != spec.Name || decoded.Risk != PluginRiskHigh || decoded.Access[0] != AccessNetwork {
		t.Fatalf("decoded = %#v", decoded)
	}
}
