# go-gin-api

GoLang + Gin + MongoDB Stack

Web Application stack built with the following technologies:

- GoLang for developing the backend
- Gin for API routing
- MongoDB for the Database (NoSQL)

## Local Setup

Pre-requisite:

- [GoLang](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)
- [Nodejs](https://nodejs.org/)

## <a name="building"></a> Building

Developers can easily build GoGAM using NPM and gulp.

- [Builds - Under the Hood](docs/guides/BUILD.md)

First install or update your local project's **npm** tools:

```bash
# First install all the NPM tools:
npm install

# Or update
npm update


```

## Get Go Packages needed

1. Download and install gin-gonic:

   ```sh
   $ go get gopkg.in/gin-gonic/gin.v1
   ```

2. Download and install mgo:

   ```sh
   $ go get gopkg.in/mgo.v2
   ```

3. Download and install mgo & mgo bson:

   ```sh
   $ go get gopkg.in/mgo.v2
   $ go get gopkg.in/mgo.v2/bson
   ```

For more details on how the build process works and additional commands (available for testing and debugging) developers should read the [Build Instructions](docs/guides/BUILD.md).
