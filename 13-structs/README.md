<h1>Structs</h1>

You can see more about structs [here](https://www.golangprograms.com/go-language/struct.html#:~:text=Structs%20are%20the%20only%20way,data%20structures%20around%20the%20system.)

A struct (short for "structure") is a collection of data fields with declared data types. Golang has the ability to declare and create own data types by combining one or more types, including both built-in and user-defined types. Each data field in a struct is declared with a known type, which could be a built-in type or another user-defined type.
Structs are the only way to create concrete user-defined types in Golang. Struct types are declared by composing a fixed set of unique fields. Structs can improve modularity and allow to create and pass complex data structures around the system. You can also consider Structs as a template for creating a data record, like an employee record or an e-commerce product.
The declaration starts with the keyword type, then a name for the new struct, and finally the keyword struct. Within the curly brackets, a series of data fields are specified with a name and a type.