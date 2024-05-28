# sayakaya-scheduler

![alt](images/systemdesign.png)
Let's assume we already have a backend service, we want to make a scheduler service for birthday promo scenario. In real world maybe we have a notification service that provide for notification features, but for simplify this test project we make notification directly in the scheduler. More advance, we also can use message broker for assign promo to users.

## 1. Flow cart

![alt](images/flowchart.png)

Scheduler will run automatically on 00.01 WIB. It will get verified users that birthday is the day, we create promo for them and then notify them via email/whatsapp.

## 2. ERD

![alt](images/dbdiagram.png)

We will use PostgreSQL. We create `promos` and `promo_eligibilities`, `promo_eligibilities` will safe eligible users for a promo, so one promo can be used by some users.

After get the verified birthday users, we create a `promos` first and then some `promo_eligibilities`. To make sure data concistency we use database transaction, we can afford this since we use RDBMS.

## 3. Code Implementation

For those functions that requested on the assignment, use the links below for direct view.

### 1. [FetchUsers](internal/users/service.go#L21)


### 2. [GeneratePromoCode](internal/promos/service.go#L32)
This function implement database transaction to ensure data concistency.

### 3. [SendNotification](internal/notification/service.go#L31)
We create [SendNotifications](internal/notification/service.go#L18) that encapsule `SendNotification`, we use go routine so we can efficiently use server resource (based on core/thread availability).

# How to Run the project

We use [mailtrap](https://mailtrap.io/) (a free email service) for this project. If you want this project run correctly you need to create mailtrap account first.

## Setup The Database Table and Data
We need to setup the databse. I use [golang-migrate](https://github.com/golang-migrate/migrate) to manage database migration. If you already install it, you can run:

```sh
make up-migration dsn="postgres://user:password@localhost:5432/dbname?sslmode=disable"
```

please edit the dsn to your dsn.

Another option is just copy SQL from [link](migrations/000001_migration.up.sql) and run manually to add table dan some data to your database.

## Setup ENV
example env:
```sh
MAILTRAP_API_TOKEN="182721hwsg873816"
EMAIL_SENDER="admin@demomailtrap.com"
DB_HOST="localhost"
DB_PORT="5432"
DB_DATABASE="sayakaya"
DB_USERNAME="postgres"
DB_PASSWORD="postgres123"
DB_SSLMODE="disable"
```

## Run
```sh
go build cmd/app/main.go
./main
```






