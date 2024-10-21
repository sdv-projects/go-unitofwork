# go-unitofwork
> [!WARNING]
> **THE PROJECT IS UNDER DEVELOPMENT!!!**

This project implements a **Unit of Work** pattern in Go, providing a convenient way to encapsulate a set of database operations within a single transaction when the **Repository** pattern is used. 

The Unit of Work allows to register repositories and perform operations on them, ensuring that all actions are executed within the same transaction.

## Installation
To install this package, use the go get command:

```sh
go get github.com/sdv-projects/go-unitofwork
```
