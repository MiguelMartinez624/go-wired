
Dependency injection tool for go programs (golang).

DI handles the life cycle of the objects in your application. It creates them when they are needed, resolves their dependencies and closes them properly when they are no longer used.

If you do not know if DI could help improving your application, learn more about dependency injection and dependency injection containers:

- [What is a dependency injection container and why use one ?](https://www.sarulabs.com/post/2/2018-06-12/what-is-a-dependency-injection-container-and-why-use-one.html)

There is also an [Examples](#examples) section at the end of the documentation.

DI is focused on performance. It does not rely on reflection.


# Table of contents


- [Basic usage](#basic-usage)
	* [Object Blueprint](#object-blueprint)
	* [Object Creation](#object-creation)

- [Errors](#errors)



# Basic usage

## Object Blueprint

A Blueprint contains a the data necesary to build and object its extracted using the 'reflect' library, and you can add one just by passing and object to be analize it.


The blueprint can be added to a Factory with the `AddBlueprint` method:

```go
  //Initialize the factory
    factory := gowired.CreateFactory()

    //Register this component with the name 'ComponentTwo'
	factory.AddBlueprint(true, ComponentTwo{}, "ComponentTwo")
```

The object you pass to this method will be analized generating a blueprint that will be stored, this process
will look for dependences field and will create blueprint for any field of type `struct`, this will be inyected
later on with a default zero value of the struct  


## Object Creation

Once the blueprints have been added to a Factory, the Factory can start creating objects.

```go
    // you can pass and object and will return a instnace of the same kind with all its
    //dependencies
    componentOne := factory.CreateObjectByName(ComponentTwo{}).(*ComponentTwo)

    //or pass a string
	componentOne := factory.CreateObjectByName("ComponentTwo").(*ComponentTwo)

```

The `CreateObjectByName` method returns an `interface{}`. You need to cast the interface before using the object.





## Errors

`BlueprintNotFound` : when you try to Create and object that its not been registered on the factory

