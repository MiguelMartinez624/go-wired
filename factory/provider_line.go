package factory

type providerLine struct {
	providers map[string]*Provider
}

func newProviderLine() *providerLine {
	return &providerLine{
		providers: make(map[string]*Provider),
	}
}

func (pl *providerLine) AddProvider(provider *Provider) {
	if _, exist := pl.providers[provider.SchemaID]; !exist {
		//Get the type of the element to store in the blueprint
		pl.providers[provider.SchemaID] = provider
	}
}

func (pl *providerLine) FindProviderBySchemaID(ID string) (provider *Provider, err error) {
	if provider, exist := pl.providers[ID]; exist {
		return provider, nil
	}

	return nil, NewProviderNotFound(ID)
}
