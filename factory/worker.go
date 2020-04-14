package factory

import "reflect"

type Worker struct {
	providerLine *providerLine
}

// CreateObjectByName create a object you can pass a name a object or anythin
func (f *Worker) BuildObjectT(schema *ObjectSchema) (obj reflect.Value, err error) {

	provider, err := f.providerLine.FindProviderBySchemaID(schema.ID)

	// In case that there its not a provider of this object yet.
	if provider == nil {
		obj = reflect.New(schema.Type)
		provider := &Provider{
			SchemaID:      schema.ID,
			SchemaToUseID: schema.ID,
			Instance:      obj.Interface(),
		}
		f.providerLine.AddProvider(provider)
		f.setDependencies(obj, schema)

		return obj, nil
	}

	// In case that the object dosent exist
	if provider.Instance == nil {
		obj := reflect.New(schema.Type)
		provider.Instance = obj.Interface()
		f.setDependencies(obj, schema)
		return obj, nil
	}

	return reflect.ValueOf(provider.Instance), nil
}

func (f *Worker) setDependencies(prtVal reflect.Value, schema *ObjectSchema) {
	//indiect the value of the Ptr to be able to work fields
	val := reflect.Indirect(prtVal)

	//For each dependencie on the @Blueprint it will get the blueprint dependency
	//build and object and assing it to the correspondent field.
	for index, dep := range schema.FieldsMap {
		if dep == nil {
			continue
		}
		//On the analize process some space are left on the array of dependencies
		//this validate temporally that the index it have a valid @FieldDep value.
		if dep.Kind == reflect.String {
			continue
		}
		prtValDev, err := f.BuildObjectT(dep)
		if err != nil {
			panic(err)

		}

		val.Field(index).Set(reflect.Indirect(prtValDev))
		f.setDependencies(prtValDev, dep)
	}
}
