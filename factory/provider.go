package factory

// Provider its represents what schema to Use on a certain
// schema.
type Provider struct {
	SchemaID      string
	SchemaToUseID string
	ItsSingleton  bool
	Instance      interface{}
}
