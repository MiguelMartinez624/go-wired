# Factory

Factory its a small and simple package to build objects in a easy way, also allow to handle its dependencies especifing providers for a field Type.


### Installing

A step by step series of examples that tell you how to get a development env running

Say what the step will be

```
go get github.com/miguelmartinez624/go-wired/factory 
```


## Usage
The Factory its really simple to use, you just create a instance of the factory, each factory 
instance asolate its schemas and objects, so you can use it as a container if you want too.

```
  // Create a factory instance
	f := factory.CreateFactory()

```

After creating a factory you must create a schema of the object that you want to create using the 
**GenerateObjectSchema** method

```
	// Component one will be analize and registered
  f.GenerateObjectSchema(ComponentOne{})

```
Now you can create an object of that type by using the **CreateObjectByName** method and remember
to cast the result to the type you are building this its becouse the method return a **interface{}**
in order to use it as the type you are specting you must casted to that type.

```
  // Create the object and () cast
	c := f.CreateObjectByName(factorytest.ComponentOne{}).(*factorytest.ComponentOne)

```
## Authors

* **Miguel Martinez** - *Initial work* - [MiguelMartinez624](https://github.com/MiguelMartinez624)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details


