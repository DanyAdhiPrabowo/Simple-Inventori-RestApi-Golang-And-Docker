<h1 align="center">Simple Inventori RestApi with Golang And Docker</h1>

<div align="center">
  <a href="#">
    <img src="https://img.shields.io/badge/Golang-1.12.7-blue.svg?style=flat-square" alt="npm">
  </a>
  <a href="#">
    <img src="https://img.shields.io/badge/MySQL-.-orange.svg?style=flat-square" alt="mysql">
  </a>
  <a href="#">
    <img src="https://img.shields.io/badge/Docker-19.03.1-9cf.svg?style=flat-square" alt="docker">
  </a>
</div>


## Introduction
This is a my final task for online course at the arkademy. This is a simple inventory RestApi made with the Golang programming language and using the MySQL database.

## Prerequiste
- Golang - Download and Install [Golang](https://golang.org/)
- RemoteMysql - Create account RemoteMysql  [here](https://remotemysql.com/)
- Docker - Download and install Docker if you want try this api via docker

## Installation On Your Local Computer
- git clone https://github.com/DanyAdhiPrabowo/Simple-Inventori-RestApi-Golang-And-Docker
- create database on [Remotemysql](https://remotemysql.com/) it's free
- setup config database in [here](https://github.com/DanyAdhiPrabowo/Simple-Inventori-RestApi-Golang-And-Docker/blob/master/src/config/config.go)
- Open file main.go on terminal
- Write "go run main.go"

## Try This API on Docker
- Download and install Docker for [windows]('https://docs.docker.com/docker-for-windows/install/'), and this for [Mac]('https://docs.docker.com/docker-for-mac/install/')
- Reference install docker click [here]('https://runnable.com/docker/getting-started/')
- After install Doker, just copy this code "docker run -p 4149:60 danyadhi/tugas-golang-course-app:0.0.1" on your terminal and press enter
- Open your Postman and try this endpoint "localhost:4149/api/product"

## Endpoint
#### Product
- /api/product                        (method GET, for get all data product)
- /api/product/id                     (method GET, for get specific data product)
- /api/product/search/"name_product"  (method GET, for search data product by name product)
- /api/product                        (method POST, for create new data product)
- /api/product/id                     (method PUT, for update data product)
- /api/product/id                     (method Delete, for delete data product)

#### Category
- /api/category                        (method GET, for get all data category)
- /api/category/id                     (method GET, for get specific data category)
- /api/category/search/"name_product"  (method GET, for search data product by name category)
- /api/category                        (method POST, for create new data category)
- /api/category/id                     (method PUT, for update data category)
- /api/category/id                     (method Delete, for delete data category)

## Link Docker hub
https://hub.docker.com/r/danyadhi/tugas-golang-course-app

## Built With

* [Golang](https://golang.org/)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [RemoteMysql](https://remotemysql.com/)
* [Docker](https://www.docker.com)


<br />
<br />

## Contact
danyadhi4149@gmail.com
