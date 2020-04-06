package gowired

import (
	"reflect"
	"sync"

	"github.com/go-wired/models"
)

//Scanner its a wrappr for reflect standar pacakge
type Scanner struct{}

func (s Scanner) FindName(obj interface{}) string {
	kind := reflect.TypeOf(obj).Kind()

	switch kind {
	case reflect.String:
		name := obj.(string)
		return name
	case reflect.Struct:
		name := reflect.TypeOf(obj).Name()
		return name
	default:
		return ""
	}
}

func (c Scanner) ScanDeep(obj interface{}, out *models.ObjectSchema, wg *sync.WaitGroup, ch chan *models.ObjectSchema) {
	c.Scan(obj, out)

	if ch != nil {
		ch <- out
	}

	defer wg.Done()

	for i := 0; i < out.Type.NumField(); i++ {
		dKind := out.Type.Field(i).Type.Kind()
		var depout models.ObjectSchema

		if dKind == reflect.Struct {
			wg.Add(1)
			go func(index int) {
				c.ScanDeep(out.Type.Field(index).Type, &depout, wg, ch)
				mutex.Lock()
				out.FieldsMap[index] = &depout
				mutex.Unlock()
			}(i)

		} else if dKind == reflect.Interface {
			c.Scan(out.Type.Field(i).Type, &depout)

			if ch != nil {
				ch <- &depout
			}
			mutex.Lock()
			out.FieldsMap[i] = &depout
			mutex.Unlock()
		}

	}
}

func (s Scanner) Scan(obj interface{}, out *models.ObjectSchema) {
	switch obj.(type) {
	case reflect.Type:
		s.buildObject(obj.(reflect.Type), out)
	default:

		oType := reflect.TypeOf(obj)
		oKind := oType.Kind()

		switch oKind {
		case reflect.String:
			out.ID = obj.(string)
			out.Name = obj.(string)
		case reflect.Struct:
			s.buildObject(oType, out)
		case reflect.Interface:
			s.buildObject(oType, out)
		case reflect.Ptr:
			s.buildObject(oType, out)
		default:
			panic("Indefinido")
		}
	}
}

func (s Scanner) buildObject(oType reflect.Type, out *models.ObjectSchema) {
	out.FieldsMap = make(map[int]*models.ObjectSchema)

	out.Name = oType.Name()
	out.Type = oType
	out.Package = oType.PkgPath()
	out.ID = out.Package + "." + out.Name
}
