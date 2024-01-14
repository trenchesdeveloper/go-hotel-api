# Go Hotel Reservation System

This is a hotel reservation system built with Go. It provides APIs for managing users, hotels, and bookings.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16 or higher
- MongoDB

### Installing

1. Clone the repository:
```sh
git clone https://github.com/trenchesdeveloper/go-hotel.git

2. Navigate to the project directory:
```sh
cd go-hotel
```

3. Install dependencies:
```sh
go mod download
```

4. Create a `.env` file in the root directory of the project and add the following environment variables:
```sh
PORT=
JWT_SECRET=
JWT_ISSUER=
JWT_AUDIENCE=
DBNAME=
DBURI=
TESTDBNAME=
```
```

5. Run the project:
```sh
make run
```
```
6. Run tests:
```sh
make test
```
```

## API Endpoints

### Users
- `GET /users` - Get all users
- `GET /users/:id` - Get a user by ID
- `POST /users` - Create a new user
- `PUT /users/:id` - Update a user by ID
- `DELETE /users/:id` - Delete a user by ID

### Hotel
- `GET /hotels` - Get all hotels
- `GET /hotels/:id` - Get a hotel by ID
- `POST /hotels` - Create a new hotel
- `PUT /hotels/:id` - Update a hotel by ID

### Booking
- `GET /bookings` - Get all bookings
- `GET /bookings/:id` - Get a booking by ID
- `POST /bookings` - Create a new booking
- `PUT /bookings/:id` - Update a booking by ID
- `DELETE /bookings/:id` - Delete a booking by ID

## Built With
- [Go](https://golang.org/) - Programming language
- [MongoDB](https://www.mongodb.com/) - Database
- [MongoDB Go Driver]
- [Go Fiber](https://gofiber.io/) - Web framework