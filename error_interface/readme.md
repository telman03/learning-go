### Assignment
We offer a product that allows businesses to send pairs of messages to couples. It is mostly used by flower shops and movie theaters.

Complete the sendSMSToCouple function. It should send 2 messages, first to the customer and then to the customer's spouse.

Use sendSMS() to send the msgToCustomer. If an error is encountered, return 0 and the error.
Do the same for the msgToSpouse
If both messages are sent successfully, return the total cost of the messages added together.
When you return a non-nil error in Go, it's conventional to return the "zero" values of all other return values.