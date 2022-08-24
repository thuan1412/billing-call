# Introduction

A project that uses [gin](https://github.com/gin-gonic/gin) and [go-ent](https://github.com/ent/ent)

# Instructions

## Using docker

Run `docker-compose up` to start the server and database container

## Without docker

Create `.env` file based on the `.env.example` file.

Run `go run main.go` to start the server

# Seed data

The project does not provide API to create users.
You can use `go run cmd/* seedDb` command to seed data for `users` table