package main

/*
	context package in go provides a way to carry deadlines, cancellation signals and other
	request-scoped value across API boundaries and between processes.
*/

func main() {
	//contextTimeout()
	//cancelContext()
	//valueContext()
	//httpRequest()
	handleErrorGroupWithContext()

}

/*
	good practice
	- Pass context.Context as the first parameter in functions.
	- Avoid storing contexts in structs.
	- Use context.WithTimeout and context.WithDeadline to control execution time.
	- Minimize usage of context.WithValue; prefer structs for complex data.
	- Always handle context cancellation by checking ctx.Done() or ctx.Err().
	- Shouldn't be used for general-purpose storage.
	- Use for storing request-scoped data
*/

/*
	in Go, the HTTP request context is automatically canceled
	when the request is complete, regardless of success, failure, or cancellation.
*/
