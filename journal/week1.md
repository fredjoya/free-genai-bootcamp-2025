# Week 1 - Language Portal Backend (Go) - Development and Testing

## Overview

This document outlines the ongoing development of the Go backend server for a language learning portal and the approach taken to test its API endpoints. The backend is being built using the **Gin framework** \ and utilises **SQLite3** for database connectivity \. My testing strategy involves using **RSpec**, a testing framework for Ruby, chosen for its **developer-friendliness and readability**. The goal is to implement all required API endpoints according to the backend technical specifications file.

## Development Summary

This week, the focus was on setting up the Go backend server \. Key tasks completed included:

*   **Server Setup**: A basic Go server structure was created using the Gin framework, along with the main server file (`cmd/server/main.go`), initial database connection setup using SQLite3, and configuration of the router with all required endpoints \.
*   **API Endpoints Implementation**: An initial set of **18 API endpoints** were implemented across five core API groups: Dashboard, Study Activities, Words, Groups, and Study Sessions \. 
*   **Handler Implementation**: Handlers were created for each API group to manage request processing (`DashboardHandler`, `ActivityHandler`, `WordHandler`, `GroupHandler`, `StudyHandler`) \.
*   **Service Layer**: A service layer was implemented to house the business logic for each API group (`ActivityService`, `WordService`, `GroupService`, `StudyService`) \. Initially, these services returned mock data \.
*   **Testing Setup**: **Comprehensive RSpec tests** were set up for all implemented endpoints, including the creation of test files and implementation of test cases to validate response structure and required fields \.

The project structure was initially organised as follows \:

```
lang-portal/backend_go/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── dashboard.go
│   │   │   ├── activity.go
│   │   │   ├── groups.go
│   │   │   ├── study.go
│   │   │   └── words.go
│   │   └── router.go
│   ├── models/
│   ├── repository/
│   └── service/
└── db/
    ├── migrations/
    └── seeds/
```

The structure was later refined, considering web development patterns and Go-specific conventions \.

## Testing with Ruby Specs

Following the initial implementation, a comprehensive suite of RSpec tests was developed to ensure the API endpoints function as expected \. This involved:

*   Considering using Go's built-in testing but opting for **RubySpec due to its developer-friendliness and readability** \
*   Setting up a Ruby environment, which required potentially installing `rvm`, Ruby, and `bundler` \. Challenges were faced with ensuring the correct Ruby environment and bundler setup \
*   Writing test specifications in Ruby within the `api_tests/spec/api/` directory, focusing on ensuring endpoints **return the expected JSON or response** (good or bad) \. The importance of **readable test code** is emphasized.
*   Addressing challenges such as issues with loading `spec_helper.rb` \ and ensuring correct file paths, sometimes requiring the use of `require_relative` \. Issues with uninitialised constants in the `spec_helper` were also encountered \.
*   Iteratively running and debugging the tests \, often leading to attempts to fix both the test code and the backend Go code to match expectations \. The process involved stepping through each endpoint \ and manually reviewing the JSON responses \.
*   Implementing a test database 
*   Managing the Go server lifecycle during testing, including manually stopping and restarting the server \. 
*   Troubleshooting failures in the specs, which ranged from incorrect response codes and data structures to issues with database state and missing data \. This often involved focusing on individual spec files for more targeted debugging \. The need for **comprehensive RSpec tests for all endpoints** is emphasised \.

## Challenges Faced (Development & Testing)

Both during the initial development and the subsequent testing, several challenges were encountered:

*   **Setting up the correct response format for each endpoint**: Ensuring that the API returns data in the expected JSON structure and with the correct data types proved to be an ongoing effort \. This involved defining JSON schemas before backend implementation and cross-referencing front-end technical specs \. Issues with generated JSON being incorrect were noted \.
*   **Handling pagination in list endpoints**: While the need for pagination was identified as a development challenge \, testing revealed issues with the pagination structure in API responses \. A consistent approach to pagination responses, using 'items', was adopted \.
*   **Managing proper error responses**: Verifying that the API returns appropriate HTTP status codes (e.g., 200) for various scenarios was a key aspect of testing \.
*   **Ensuring consistent API structure across all endpoints**: Maintaining a uniform design across the API was a goal, and testing helped to identify any inconsistencies \. The front-end design significantly influenced the backend API structure \.
*   **Data inconsistencies in the test environment**: Running tests against a development database led to issues with data availability and state \, necessitating the creation and use of a test database and the implementation of seed data loading for both the main and test databases \.
*   **Managing the Go server lifecycle during testing**: Starting, stopping, and ensuring the server was running with the correct configuration for the test environment required manual intervention at times \. 
*   **Issues with the test setup and execution**: Problems with the Ruby environment configuration, RSpec command usage (e.g., `bundle exec rspec`), and loading test files (`spec_helper.rb`) were encountered and resolved iteratively \.

## Conclusion

The initial development of the Language Portal backend laid the groundwork for a comprehensive API. The adoption of RSpec for testing allowed for the creation of readable and maintainable test specifications. While the testing process involved iterative debugging and troubleshooting, significant progress was made in verifying the functionality of the implemented endpoints. Continuous testing throughout the remaining development phases, along with the implementation of remaining features and proper task automation using Mage, will be essential to ensure the reliability and quality of the Language Portal backend \. The technical specifications were considered a crucial reference point throughout the development and testing process \. Code reviews using tools like Git Smart were also part of the workflow \.



