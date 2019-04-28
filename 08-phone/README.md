# Exercise 8 - "Phone Number Normalizer"

View on [gophercises.com](https://gophercises.com/exercises/phone).

## Description
This exercise is fairly straight-forward - we are going to be writing a program that will iterate through
a database and normalize all of the phone numbers in the DB. After normalizing all of the data we might
find that there are duplicates, so we will then remove those duplicates keeping just one entry in our
database.


There are many ways to use SQL in Go. Rather than just picking one, I am going to try to cover a few.
If you would like to see any additional libraries covered feel free to reach out and I'll try to add it.
For now, here are the libraries I intend to cover:

- Writing raw SQL and using the [database/sql](https://golang.org/pkg/database/sql/) package in
the standard library
- Using the very popular [sqlx](https://github.com/jmoiron/sqlx) third party package,
which is basically an extension of Go's sql package.
- Using a relatively minimalistic ORM (I will be using [gorm](https://github.com/jinzhu/gorm))
