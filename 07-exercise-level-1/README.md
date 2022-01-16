<h1>🥷🏻 Ninja Level 1 🥷🏻</h1>

<h2>Hands-on exercise #1</h2>

1. Using the short declaration operator, assign these values to variables with the identifiers "x" and "y" and "z".
   1. 42
   2. "James Bond"
   3. true
2. Now print the values stored in those variables using
   1. a single print statement
   2. multiple print statements
   
<h2>Hands-on exercise #2</h2>

1. Use var to declare three variables. The variables should have package level scope. Do not assign values to the variables. Use the following identifiers for the variables and makes sure te variables store values of the following type.
   1. identifier "x" type int
   2. identifier "y" type string
   3. identifier "z" type bool
2. In func main
   1. print out the values for each identifier
   2. the compiler assigned value to the variables. What are these values called?

<h2>Hands-on exercise #3</h2>

Using the code from the previous exercise.
1. At the package level scope, assign the following values to the three variables.
   1. for x assign 42
   2. for y assign "James Bond"
   3. for z assign true
2. In func main
   1. use fmt.Sprintf() to print all of the values to one single string. Assign the returned value of type string using the short declaration operator to a variable with identifier "s"

<h2>Hands-on exercise #4</h2>

1. Create your own type. Have the underlying type be an int
2. Create a variable of your new type with the identifier "x" using the var keyword
3. In func main
   1. print out the value of the variable "x"
   2. print out the type of the variable "x"
   3. assign 42 to the variable using "x" the = operator
   4. print out the value of the variable "x"

<h2>Hands-on exercise #5</h2>

Using the code from the previous exercise.
1. At the package level scope, using the var keyword, create a variable with the identifier "y". The variable should be of the underlying type of your custom type "x".
2. In func main
   1. This should already be done
      1. print out the value of the variable "x"
      2. print out the type of the variable "x"
      3. assign 42 to the variable using "x" the = operator
      4. print out the value of the variable "x"
   2. Now do this
      1. now use conversion to convert the type of the value stored in "x" to te underlying type
         1. then use the short declaration operator to assign that value to y