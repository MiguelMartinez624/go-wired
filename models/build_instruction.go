package models

//BuildInstruction indicate how should a object/component be build
// by th factory, what component use for a field, mustly use it on fields that
// wich type its a interface type
type BuildInstruction struct {
	//Name of this configuration
	Name string
	//In who should this instructions by apply if thre its not target
	//it will apply to all objects
	Target interface{}
	//A list of injection definitions
	Instruction []Instruction

	InstructionMap map[string]*Instruction
}

//Instruction definition of what component use on a field
type Instruction struct {
	//If you witch to specify the exact field
	Field interface{}
	//Field type
	FieldType interface{}
	Use       interface{}
}

type Provider struct {
	SchemaID      string
	ToUseSchemaID string
}
