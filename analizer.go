package gowired

//Analizer has the funcionality to scan and create blueprints and schemas of
// objects that can be use to know how a object its compose
type Analizer struct {
	scanner        Scanner
	objectsSchemas map[string]*ObjectSchema
}

func BuildAnalizer() *Analizer {
	return &Analizer{
		objectsSchemas: make(map[string]*ObjectSchema),
	}
}

//Analize and its depndencies on a recursive maner it will store
// each object schema in a map so for every object it will traverse
// the fields as a tree and create individual @ObjectSchema
func (a Analizer) Analize(component interface{}) *ObjectSchema {

	//Output channel for schemas
	ch := make(chan *ObjectSchema)

	// First schema its the object targt
	var object ObjectSchema
	go func() {
		// When this function ends we need to close the channel so the
		// subcriber (the for listining for schemas) will end and we can
		// return the function
		defer close(ch)
		a.scanner.ScanDeep(component, &object, ch)
	}()

	// Here we listing to schemas and write to the @ObjectSchema map
	//
	for schema := range ch {
		a.objectsSchemas[schema.ID] = schema
	}

	return &object
}

// GenerateBlueprint take a schema and make a simplifier version with only ID of the
// schema and the schema to use, this will allow futher on to change what schema to use
// for example to a interface where we dont have a concrete type, we can specify
// " for this interface us this schema"
// Its also done on a recursive way throughout the tree
func (a Analizer) GenerateBlueprint(schema *ObjectSchema) *Blueprint {
	var out Blueprint
	out.ID = schema.ID
	out.SchemaID = schema.ID
	out.Childs = make([]*Blueprint, 0)

	// Start recursion
	for i, schemaDep := range schema.FieldsMap {
		a.generateBlueprintChildsTree(schemaDep, &out, i)
	}

	return &out
}

//generateBlueprintChildsTree
func (a Analizer) generateBlueprintChildsTree(schema *ObjectSchema, parent *Blueprint, index int) {
	bp := &Blueprint{

		ID:       schema.ID,
		SchemaID: schema.ID,
		Index:    index,
		Childs:   make([]*Blueprint, 0),
	}

	parent.Childs = append(parent.Childs, bp)
	for i, schemaDep := range schema.FieldsMap {
		a.generateBlueprintChildsTree(schemaDep, bp, i)
	}
}

// FindSchema find a schema of a unknow type. creatin a tempSchma of the
// object passed
func (a Analizer) FindSchema(obj interface{}) *ObjectSchema {
	// Scan object to get a name and ID
	var tempSchema ObjectSchema
	a.scanner.Scan(obj, &tempSchema)

	schemaStored := a.objectsSchemas[tempSchema.ID]
	if schemaStored == nil {
		panic("no found")
	}
	return schemaStored
}

// FindSchemaByID directly search the schema b its key on the map
func (a Analizer) FindSchemaByID(schemaID string) (schema *ObjectSchema, err error) {

	schema = a.objectsSchemas[schemaID]
	if schema == nil {
		return nil, NewSchemaNotFound(schemaID)
	}
	return schema, nil
}
