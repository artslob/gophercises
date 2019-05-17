# Exercise 15 - "Recover Middleware w/ Source Code"

View on [gophercises.com](https://gophercises.com/exercises/recover_chroma).

## Description
In the [recover](https://gophercises.com/exercises/recover) exercise we learned how to create some HTTP middleware
that recovers from any panics in our application and renders a stack trace if we are in a local development environment.
In this exercise we will be taking that code a step further; we will be adding in the ability to navigate to any source
file in the panic stack trace in order to make it easier to debug issues when they arise in a development environment.

Given the web server and the recovery middleware in `main.go`, add the following to the application:

#### 1. An HTTP handler that will render source files in the browser
#### 2. Add syntax highlighting to the source file rendering
#### 3. Parse the stack trace & creating links
#### 4. Add line highlighting
