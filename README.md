# DTO Pattern Challenge

## Overview
This challenge focuses on implementing the Data Transfer Object (DTO) pattern in a Go application. The goal is to create a clean separation between the internal data model and the data exposed through the API.

## Challenge Requirements

### 1. User Entity Structure
The application manages a User entity with the following fields:
- `id` (UUID)
- `firstName`
- `lastName`
- `email`
- `password`
- `dateOfBirth`
- `createdAt`

### 2. DTO Implementation
Two DTOs must be implemented:

#### UserResponseDTO
- Should never expose sensitive information (password, security details)
- Combines firstName and lastName into a single fullName field
- Includes only relevant fields for client display
- Adds a calculated field 'age' based on dateOfBirth

#### UserCreateDTO
- Handles user creation requests
- Validates input data
- Includes proper type definitions
- Implements field validation

### 3. Evaluation Criteria
The solution will be evaluated based on:

1. **Separation of Concerns**
   - Clear distinction between entity and DTOs
   - Proper data transformation
   - Clean architecture principles

2. **Type Safety**
   - Proper use of Go's type system
   - Clear struct definitions
   - JSON serialization handling

3. **Data Validation**
   - Input validation
   - Field constraints
   - Error handling

4. **Security Considerations**
   - Sensitive data protection
   - Input sanitization
   - Secure data transfer

5. **Code Organization**
   - Clean code structure
   - Proper error handling
   - Documentation

## Implementation Details

### Project Structure
```
internal/
  └── dto/
      ├── user.dto.go    # DTO implementations
      └── user_test.go   # Test cases
```

### Key Features
- Data transformation between entity and DTOs
- Input validation
- Error handling
- Security measures

## Getting Started

1. Clone the repository
2. Install dependencies
3. Set up the database connection
4. Run the tests
5. Start the application

## Testing
Run the test suite:
```bash
go test ./internal/dto/...
```

## Best Practices
- Use proper error handling instead of panics
- Implement input validation
- Protect sensitive data
- Follow clean code principles
- Write comprehensive tests

## Learning Objectives
- Understand the DTO pattern
- Learn proper data validation
- Implement security best practices
- Master error handling in Go
- Practice clean architecture principles 