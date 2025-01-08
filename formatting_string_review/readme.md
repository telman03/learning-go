## Assignment

Our users frequently try to run custom analytics queries on their message deliverability metrics, and end up writing a bad query that tries to divide a number by zero. It's become such a problem that we need to make a new type of error for division by zero.

Update the code so that the divideError type implements the error interface. Its Error() method should just return a string formatted in the following way:

`can not divide DIVIDEND by zero`
Where DIVIDEND is the actual dividend of the divideError.
Use the %v verb to format the dividend as a float.