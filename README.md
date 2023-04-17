# E-Commerce.API

This is a basic e-commerce api were you can create a user, user can add products (does not implement admin authorisations) so therefore the products are added by the users, add items to a users cart and finally checkout (i.e buy the product)

## Installation And Running

### Installation

To install the api, all you need to do is just clone the repository using `git clone https://github.com/jamesconfy/e-commerce-api.git`.

Do a `go mod tidy` to download all the required packages or by doing `make gotidy`. A [makefile](Makefile) was provided that contains all the commands relative to the repository.

**DATABASE**

You will need to set up a mysql database as it was the database of choice in this code base (either locally or remotely). Get the username, password, host, port and database name as this will be needed before running the application.

You then make a copy of the [app-sample.env](app-sample.env) into an [app.env](app.env) file and provide the required parameters.

- **`DATABASE`** is the database name in the format `username:password@tcp(host_network:host_port(usually 3306))/database_name(e_commerce_api can be used)?parseTime=true`. Either **_`DEVELOPMENT_DATABASE`_** or **_`PRODUCTION_DATABASE`_**
- **`MODE`** is the application mode could either be production or development, this is needed to know how to configure the application as there are different settings for different mode.
- **`ADDR`** is the port the application is running on, provided default is `8080` but can be changed to anything you prefer.
- **`SECRET_KEY_TOKEN`** is a secret key that is used for encrypting the jwt service. Could be anything you like it to be, but you can generate a unique secret key using [Random Key Generator](https://acte.ltd/utils/randomkeygen).

After the database has been set up, the next thing you want to do is migrate the database schema to the database as for how to do that, you can use the `make migrateup` command.

**NB:** You need to provide the database url you recently created by using the format `mysql://username:password@tcp(host_network:host_port(usually 3306))/database_name(e_commerce_api can be used)`. Or you can just run this command to do this instead

```terminal
migrate -path db/migration -database "mysql://username:password@tcp(host_network:host_port(usually 3306))/database_name(e_commerce_api can be used)" -verbose up
```

### Running

You can start the local server by using the makefile command `make run`. This build a new e_commerce_api executable binary on the root directory and starts the server. If mode is set to development, gin mode is set to debugging and everything about the server is logged to the stdout. If **ADDR** in the [app.env](app.env) was not changed from the one in the [app-sample.env](app-sample.env), you can access the application on [localhost](localhost:8080/api/v1). This leads you to the homepage. You can access the openai documentation on [swagger](localhost:8080/api/v1/swagger/index.html) and you should see something resembling this ![image](/assets/swagger.png)

### Testing

Three different package tests was done and you can see in the [tests](/tests/) folder. We have the [repo_test](), [service_test]() and [handler_test](). These tests are automated tests, i.e it creates a custom mysql database for you, using mysql:8.0.32 docker image (you need to have docker desktop running for it to work). It tests each package and destroys the docker mysql image after it is done, so you don't need to worry about your database integrity.

- **repo_test:** This test deals on everything about the repo (user, product, cart and cart_item). Run the command `make test_repo` to run this test.
- **service_test:** This test deals on everything about the service aspect (user, product, cart and cart_item). Run the command `make test_service` to run this test.
- **handler_test:** This is an integration test that test the whole application as one. It first of all starts a local server that it uses to test the application. to run this test, do a `make test_handler`.

### Others

You can read more about the [makefile](Makefile) on what other commands you can carryout.

### Live

Project is live on [live](https://e-commerce-api.fly.dev/api/v1/swagger/index.html)
