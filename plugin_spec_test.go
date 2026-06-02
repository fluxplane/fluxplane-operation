package operation

import (
	"encoding/json"
	"testing"
)

func TestDeclarationJSON(t *testing.T) {
	spec := Declaration{
		Name:        "test.write",
		Description: "Write something",
		ReadOnly:    false,
		Effects:     []Effect{EffectWrite, EffectNetwork},
		Risk:        RiskHigh,
		Idempotency: IdempotencyNonIdempotent,
		Access:      []Access{AccessNetwork},
		Render:      &RenderSpec{Preferred: "json", Formats: []string{"json"}},
	}
	raw, err := json.Marshal(spec)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded Declaration
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if decoded.Name != spec.Name || decoded.Risk != RiskHigh || decoded.Access[0] != AccessNetwork {
		t.Fatalf("decoded = %#v", decoded)
	}
}
