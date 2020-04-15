package factory

//Factory its the one who handles the creation of other objects///
type Factory struct {
	analizer     *Analizer
	scanner      *Scanner
	worker       *Worker
	providerLine *providerLine
	// probably remove later
	objectsCreated map[string]interface{}
	objectSchema   map[string]*ObjectSchema
}

//CreateFactory create a factory this its the constructor function
//it initialize the bleprint map
func CreateFactory() *Factory {

	pl := newProviderLine()

	factory := &Factory{
		scanner:      &Scanner{},
		analizer:     BuildAnalizer(),
		providerLine: pl,
		worker:       &Worker{providerLine: pl},
		objectSchema: make(map[string]*ObjectSchema),
	}
	return factory
}

// GenerateObjectSchema generate a tree schema of the object and its dependencies
func (f *Factory) GenerateObjectSchema(component interface{}) string {
	object := f.analizer.Analize(component)
	//store object schema
	return object.ID
}

// CreateObjectByName create a object you can pass a name a object or anythin
func (f *Factory) CreateObjectByName(name interface{}) (obj interface{}) {
	targetSchema := f.analizer.FindSchema(name)

	created, err := f.worker.BuildObjectT(targetSchema)
	if err != nil {
		panic(err)
	}

	return created.Interface()
}

//  RegisterProvider way to store a provider
func (f *Factory) RegisterProviderSchema(target interface{}, provider interface{}) {
	schemaProvider := f.analizer.Analize(provider)
	targetSchema := f.analizer.FindSchema(target)

	newProvider := &Provider{SchemaID: targetSchema.ID, SchemaToUseID: schemaProvider.ID}

	f.providerLine.AddProvider(newProvider)
}

//  RegisterProviderInstance way to store a provider
func (f *Factory) RegisterProviderInstance(target interface{}, provider interface{}) {
	schemaProvider := f.analizer.Analize(provider)
	targetSchema := f.analizer.FindSchema(target)

	//By default it will be a singleton
	newProvider := &Provider{
		SchemaID:      targetSchema.ID,
		SchemaToUseID: schemaProvider.ID,
		ItsSingleton:  true,
		Instance:      provider,
	}

	f.providerLine.AddProvider(newProvider)
}
