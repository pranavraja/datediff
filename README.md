
# datediff

A tool to calculate the number of days difference between certain dates.

## Spec

Fuller requirements:

> This is required to be a CLI application implemented in golang using only primitives, built-in language functions and no external npm (sic) dependencies. You should build the application so it accepts input from stdin. Please do guide the user through the steps required to use your application, and do validate your input.
> 
> You need to calculate the distance in whole days between two dates, counting only the days in between those dates, i.e. 01/01/2001 to 03/01/2001 yields “1”. The valid date range is between 01/01/1900 and 31/12/2999, all other dates should be rejected. When testing your solution, use the following sample data to demonstrate your code works:
> 
> a) 2/6/1983 to 22/6/1983 19 days
> b) 4/7/1984 to 25/12/1984 173 days
> c) 1/3/1989 to 3/8/1983 2036 days

# Why?

I have no idea why someone would want to use this, or why using only primitives is a realistic or desirable idea in practice.

# Running the tests

	go test -v

# Running the program

	go run .

You can then type things like: `2/6/1983 to 22/6/1983` into stdin and get results.

Once you're done, press Ctrl+D to close stdin and finish the program.
