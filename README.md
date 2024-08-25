Golang Basics : a Function that Uses Slices, Maps, and Defer

a function that manages a small inventory of books. The function should:

Accept a slice of book titles (strings) as input.
Use a map to count how many times each book title appears in the slice.
Use a defer statement to print a message once the counting process is complete.
Return the map containing the count of each book title.
Requirements
The function should be named CountBooks.
The function should take one parameter: books []string.
The function should return a map[string]int where the keys are book titles and the values are the number of times each title appears in the input slice.
Use defer to print the message "Counting complete!" after the map has been populated.
