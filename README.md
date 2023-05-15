# Slices-in-GO
In this code I've learnt about slices and would share some key points.

First I created a slice fruitList with 3 string values and later on usingf the append method added 2 more values.
Next I created a value of highScores but didn't use the syntax I was using so far,I use make and we first define what type of data you want to have and then we define what size you are going to have.
So I want a slice which is going to hold integers this time and I want the value to be 5.

By default highScores is an array but getting a layer of abstraction and thus we call them slices.
If I go ahead and try to add one more value say the 5th value and run it, the code is gonna crash case we want just the 5 values here and we try to have a value which is out of the bound.

But interestingly you can actually use the append method in the highScores and add values, and this is where GO behaves a little bit different and doesn't gives an error.

The way how it works that "make" is definitely going to give you a slice of integers and you can store 5 values but this is the default allocation of the memory as soon as you use the method append it is going to reallocate the memory and all of the values will be accommodated .
As soon as the new values come in the entire memory allocation happens again and this saves a lot of time and a lot of performance optimization comes in right from this one.

Also you can use the package sort in slices but not in arrays.
