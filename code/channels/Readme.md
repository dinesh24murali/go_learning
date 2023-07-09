# Channels

- Channels are used to communicate between child routines
- In go we can create a child routine by calling a function with the keywork `go` before it

- First, run the `without_routines.go` file to see what we are trying to build. We make api calls to differnet URLs to see if they are working. If they are working we print success, if there is an error we print the error. In this implementation
the `http` call we are making gets blocked until the api call ends. The idea of using routines and channels is to make these API calls concurrently/parallely
- The main function itself is a go routine
- In the `not_working_routines.go` file we have implmented child routines, but it didn't print because the main routine doesn't know that the child routines haven't finished. Channels are a mechanism through which we can communicate between these routines so that the main routine ends only after all the child routine have finished
- Receiving a message from a channel is blocking