<h1>🥷🏻 Ninja Level 5 🥷🏻</h1>

<h2>Hands-on exercise #1</h2>

Create your own type person which will have an underlying type of struct
so that it can store the following data
* firstName
* lastName
* favorite ice cream flavors

Create two values of type person. Print out the values, ranging over the elements in the slice

<h2>Hands-on exercise #2</h2>

Take the code from the previous exercise, then store the values of type person in a map with the key of
last name. Access each value in the map. Print out the values, ranging over the slice.

<h2>Hands-on exercise #3</h2>

* Create a new type: vehicle
  * The underlying type is a struct
  * The fields:
    * doors
    * colors
* Create two new type: truck and sedan
  * The underlying type of each of these new types is a struct
  * Embed the vehicle type in both and sedan
  * Give truck the field fourWheel which will be set to bool
  * Give sedan the field luxury which will be set to bool
* Using the vehicle, truck and sedan structs:
  * Using a composite literal, create a value of type truck and assign values to the fields
  * Using a composite literal, create a value of type sedan and assign values to the fields
* Print out each of these values
* Print out a single field from each of these values