# Stock Reward System

This project implements a stock reward system using the Gin framework in Go. It provides a set of REST APIs to manage and retrieve stock rewards for users.

## Project Structure

```
stock-reward-system
├── src
│   ├── main.go                # Entry point of the application
│   ├── controllers            # Contains the controller logic for handling requests
│   │   └── reward_controller.go
│   ├── routes                 # Defines the API routes
│   │   └── reward_routes.go
│   ├── models                 # Contains the data models
│   │   └── reward.go
│   ├── database               # Handles database connection and setup
│   │   └── db.go
│   ├── middleware             # Contains middleware functions
│   │   └── error_handler.go
│   └── utils                  # Utility functions for edge case handling
│       └── edge_cases.go
├── go.mod                     # Module definition and dependencies
├── go.sum                     # Checksums for module dependencies
└── README.md                  # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd stock-reward-system
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Set up the PostgreSQL database:**
   - Create a database named `assignment`.
   - Update the database connection settings in `src/database/db.go` if necessary.

4. **Run the application:**
   ```
   go run src/main.go
   ```

## API Specifications

### Endpoints

- **POST /reward**
  - Records a user's stock reward.
  - Request Body: `{ "userId": "string", "stockSymbol": "string", "quantity": decimal }`

- **GET /today-stocks/{userId}**
  - Returns today's stock rewards for a user.

- **GET /historical-inr/{userId}**
  - Returns the INR value of a user's stock rewards for past days.

- **GET /stats/{userId}**
  - Returns total shares rewarded today and current INR value of the user's portfolio.

## Edge Case Handling

The project includes utility functions to handle various edge cases:

- **CheckDuplicateReward:** Ensures that duplicate reward events are not recorded.
- **HandleStockEvents:** Manages stock splits, mergers, or delisting scenarios.
- **HandlePriceAPIDowntime:** Manages scenarios where the price API is down.

## Scaling Considerations

- Consider implementing caching for frequently accessed data to improve performance.
- Use a message queue for processing rewards asynchronously if the volume of requests increases significantly.
- Monitor the application and database performance to identify bottlenecks and optimize accordingly.

## License

This project is licensed under the MIT License.