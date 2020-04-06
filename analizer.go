package gowired

import (
	"sync"

	"github.com/go-wired/errors"
	"github.com/go-wired/models"
)

var (
	mutex sync.Mutex
)

type Analizer struct {
	scanner        Scanner
	objectsSchemas map[string]*models.ObjectSchema
}

func BuildAnalizer() *Analizer {
	return &Analizer{
		objectsSchemas: make(map[string]*models.ObjectSchema),
	}
}

//Analize a object to createa  blueprint wit its dependencies.
func (a Analizer) Analize(component interface{}) *models.ObjectSchema {
	//Start analizin and examine the componente type
	ch := make(chan *models.ObjectSchema)

	var object models.ObjectSchema
	go func() {
		defer close(ch)
		a.scanner.ScanDeep(component, &object, ch)
	}() //channel to store al schemas

	for schema := range ch {
		mutex.Lock()
		a.objectsSchemas[schema.ID] = schema
		mutex.Unlock()

	}

	return &object
}

func (a Analizer) GenerateBlueprint(schema *models.ObjectSchema) *models.Blueprint {
	var out models.Blueprint
	out.ID = schema.ID
	out.SchemaID = schema.ID
	out.Childs = make([]*models.Blueprint, 0)

	for i, schemaDep := range schema.FieldsMap {
		a.generateBlueprintChildsTree(schemaDep, &out, i)
	}

	return &out
}

func (a Analizer) generateBlueprintChildsTree(schema *models.ObjectSchema, parent *models.Blueprint, index int) {
	bp := &models.Blueprint{

		ID:       schema.ID,
		SchemaID: schema.ID,
		Index:    index,
		Childs:   make([]*models.Blueprint, 0),
	}

	parent.Childs = append(parent.Childs, bp)
	for i, schemaDep := range schema.FieldsMap {
		a.generateBlueprintChildsTree(schemaDep, bp, i)
	}
}

func (a Analizer) FindSchema(obj interface{}) *models.ObjectSchema {
	var tempSchema models.ObjectSchema
	a.scanner.Scan(obj, &tempSchema)
	mutex.Lock()
	schemaStored := a.objectsSchemas[tempSchema.ID]
	mutex.Unlock()
	if schemaStored == nil {
		panic("no found")
	}
	return schemaStored
}

func (a Analizer) FindSchemaByID(schemaID string) (schema *models.ObjectSchema, err error) {

	mutex.Lock()
	schema = a.objectsSchemas[schemaID]
	mutex.Unlock()
	if schema == nil {
		return nil, errors.NewSchemaNotFound(schemaID)
	}
	return schema, nil
}
