package gowired

import (
	"reflect"
)

//Factory its the one who handles the creation of other objects///
type Factory struct {
	blueprints     *BlueprintMap
	analizer       *Analizer
	scanner        *Scanner
	objectsCreated map[string]interface{}
	objectSchema   map[string]*ObjectSchema
	providers      map[string]*Provider
}

//CreateFactory create a factory this its the constructor function
//it initialize the bleprint map
func CreateFactory() *Factory {
	factory := &Factory{
		blueprints:     NewBlueprintMap(),
		scanner:        &Scanner{},
		analizer:       BuildAnalizer(),
		objectsCreated: make(map[string]interface{}),
		objectSchema:   make(map[string]*ObjectSchema),
		providers:      make(map[string]*Provider),
	}
	return factory
}

// GenerateObjectSchema generate a tree schema of the object and its dependencies
func (f *Factory) GenerateObjectSchema(component interface{}) string {
	object := f.analizer.Analize(component)
	//store object schema
	f.objectSchema[object.ID] = object
	return object.ID
}

// CreateBlueprint create a Object Schema with a schema ID
func (f *Factory) CreateBlueprint(schemaID string) (ID string, err error) {

	schema := f.objectSchema[schemaID]
	if schema == nil {
		return "", NewSchemaNotFound(schemaID)
	}

	blueprint := f.analizer.GenerateBlueprint(schema)
	f.blueprints.AddBlueprint(blueprint)
	return blueprint.ID, nil
}

// CreateObjectByName create a object you can pass a name a object or anythin
func (f *Factory) CreateObjectByName(name interface{}) (obj interface{}) {
	targetSchema := f.analizer.FindSchema(name)
	blueprint, err := f.blueprints.FindBlueprint(targetSchema.ID)
	if err != nil {
		panic(err)
	}

	//if its a singleton we check the builded object history map
	if value, exist := f.objectsCreated[blueprint.Name]; blueprint.ItsSingleton && exist {
		return value
	}

	//here we have the core object now we need to create its dependencies
	prtVal, err := f.BuildObject(blueprint)
	if err != nil {
		panic(err)

	}

	//Set all dependencies of the object.
	if blueprint.Childs != nil {
		f.setDependencies(prtVal, blueprint)
	}

	//if its a singleton we check that there its not other reference of this object and
	// we stored this one for future use
	if _, exist := f.objectsCreated[blueprint.Name]; blueprint.ItsSingleton && !exist {
		f.objectsCreated[blueprint.Name] = prtVal.Interface()
	}

	return prtVal.Interface()
}

// setDependencies iterates though all @FieldDep and using the @Blueprints stored in
// the map it will procced to create and assign the value field, this its done on a
// recursive manner so it guarantee the childs tree dependencies are fullfilled too
func (f *Factory) setDependencies(prtVal reflect.Value, blueprint *Blueprint) {
	//indiect the value of the Ptr to be able to work fields
	val := reflect.Indirect(prtVal)

	//For each dependencie on the @Blueprint it will get the blueprint dependency
	//build and object and assing it to the correspondent field.
	for _, dep := range blueprint.Childs {
		if dep == nil {
			continue
		}
		//On the analize process some space are left on the array of dependencies
		//this validate temporally that the index it have a valid @FieldDep value.
		if dep.SchemaID == "" {
			continue
		}
		prtValDev, err := f.BuildObject(dep)
		if err != nil {
			panic(err)

		}

		val.Field(dep.Index).Set(reflect.Indirect(prtValDev))
		f.setDependencies(prtValDev, dep)
	}
}

//BuildObject build and object using a blueprint
func (f *Factory) BuildObject(blueprint *Blueprint) (obj reflect.Value, err error) {
	var schemaToUse string

	// Check if there its a provider to the object, if there its one
	// for that schema it will have priority over de native (the one on the type)
	if provider, ok := f.providers[blueprint.SchemaID]; ok {
		schemaToUse = provider.SchemaToUseID
	} else {
		schemaToUse = blueprint.SchemaID
	}

	//Find the schema to be used.
	objectSchema, err := f.analizer.FindSchemaByID(schemaToUse)
	if err != nil {
		panic(err)
	}
	value := reflect.New(objectSchema.Type)
	return value, nil
}

//  RegisterProvider way to store a provider
func (f *Factory) RegisterProvider(target interface{}, provider interface{}) {
	schemaProvider := f.analizer.Analize(provider)
	targetSchema := f.analizer.FindSchema(target)

	newProvider := &Provider{SchemaID: targetSchema.ID, SchemaToUseID: schemaProvider.ID}

	f.providers[newProvider.SchemaID] = newProvider

}

func (f *Factory) RegisterObject(objct interface{}) {
	schemaID := f.GenerateObjectSchema(objct)
	_, err := f.CreateBlueprint(schemaID)
	if err != nil {
		panic(err)
	}

}
