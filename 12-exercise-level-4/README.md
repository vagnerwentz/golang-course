<h1>🥷🏻 Ninja Level 4 🥷🏻</h1>

<h2>Hands-on exercise #1</h2>

* Using a COMPOSITE LITERAL:
  * Create an array which holds 5 values of type int
  * Assign values to each index position
* Range over the array and print the values out
* Using format printing
  * Print out the type of the array

<h2>Hands-on exercise #2</h2>

* Using a COMPOSITE LITERAL:
   * Create a slice of type int
   * Assign 10 values
* Range over the slice and print the values out
* Using format printing
  * Print out the type of the array

<h2>Hands-on exercise #3</h2>

Using the code from the previous example, use slicing to create the following new slices
which are then printed:
* [42 43 44 45 46]
* [47 48 49 50 51]
* [44 45 46 47 48]
* [43 44 45 46 47]

<h2>Hands-on exercise #4</h2>

Follow these steps:
1. Start with this slice
   1. si := []int{42...51}
2. Append to that slice this value
   1. 52
3. Print out the slice
4. In one statement append to that slice these values
   1. 53, 54, 55
5. Print out the slice
6. Append to the slice this slice
   1. osi := []int{56...60}
7. Print out the slice

<h2>Hands-on exercise #5</h2>

To delete from a slice, we use append along with slicing.
For this hands-on exercise, follow these steps:
1. Start with this slice
   1. si := []int{42...51}
2. Use append and slicing to get the values here which you should then print
   1. [42, 43, 44, 48, 49, 50, 51]

<h2>Hands-on exercise #6</h2>

Create a slice to store the names of all the states in the
United States of America. Use **make** and **append** to do this.
Goal: do not have the array that underlies the slice created more than
once. What is the length of your slice? What is the capacity? Print out all of the values, along
with their index position, without using the range clause. Here is a list of the 50 states:

` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`

<h2>Hands-on exercise #7</h2>

Create a slice of a slice of string. Store the following data in the
multi-dimensional slice:
* Vagner, Bruna, Antonio
* Wentz, Eloisa, Roberval

Range over the records, then range over the data in each record.

<h2>Hands-on exercise #8</h2>

Create a map with a key of type string which is a person's "last_first" name, and
a value of type []string which stores their favorite things. Store three records in your map.
Print out all of the values, along with their index position in the slice.

"wentz_vagner", "Go", "Terere", "Bruna Eloisa"
"eloisa_bruna", "Design", "Chá", "Vagner Wentz"
"nina_wentz_schvingel", "Passear", "Biofresh", "Papai e Mamãe"

<h2>Hands-on exercise #9</h2>

Using the code from previous example, add a record to your map. Now print the map
out using the "range" loop

<h2>Hands-on exercise #10</h2>

Using the code from previous example, delete a record to your map. Now print the map
out using the "range" loop