# REST API Documentation


## User Endpoints

### Create User
- **URL**: `/user`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "username": "name",
    "password": "Password123"
  }
  ```
- **Response**:
  - **Status**: `201 Created`
  - **Body**:
    ```json
    {
      "state": {
          "status": "Success"
      },
      "data": {
          "userId": 3
      }
    }
    ```

### Get User by User ID
- **URL**: `/user`
- **Method**: `GET`
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
      "state": {
          "status": "Success"
      },
      "data": {
          "user": {
              "userId": 3,
              "username": "name",
              "createdAt": "2025-03-08T18:22:46.628854Z"
          }
      }
    }
    ```

### Update User Password
- **URL**: `/user/password`
- **Method**: `PUT`
- **Request Body**:
  ```json
  {
    "password": "Newpassword123"
  }
  ```
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
        "state": {
            "status": "Success"
        }
    }
    ```

### Delete User
- **URL**: `/user/:userId`
- **Method**: `DELETE`
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
        "state": {
            "status": "Success"
        }
    }
    ```

## Error Responses
- **Status**: `400 Bad Request`
  - **Body**:
    ```json
    {
      "status": "error",
      "message": "failed to bind request"
    }
    ```
- **Status**: `404 Not Found`
  - **Body**:
    ```json
    {
      "status": "error",
      "message": "resource not found"
    }
    ```
- **Status**: `500 Internal Server Error`
  - **Body**:
    ```json
    {
      "status": "error",
      "message": "unexpected server error"
    }

## Task Endpoints

### Create Task
- **URL**: `/tasks`
- **Method**: `POST`
- **Request Body**:
  ```json
    {
        "taskContent": "Hello World!"
    }
  ```
- **Response**:
  - **Status**: `201 Created`
  - **Body**:
    ```json
    {
        "state": {
            "status": "Success"
        },
        "data": {
            "taskId": 1
        }
    }
    ```

### Get Tasks by User ID
- **URL**: `/tasks`
- **Method**: `GET`
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
        "state": {
            "status": "Success"
        },
        "data": {
            "tasks": [
                {
                    "taskId": 1,
                    "userId": 1,
                    "taskContent": "Hello World!",
                    "createdAt": "2025-03-08T18:28:31.800531+05:00"
                }
            ]
        }
    }
    ```

### Get Task by Task ID
- **URL**: `/tasks/:taskId`
- **Method**: `GET`
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
        "state": {
            "status": "Success"
        },
        "data": {
            "task": {
                "taskId": 1,
                "userId": 1,
                "taskContent": "Hello World!",
                "createdAt": "2025-03-08T18:28:31.800531+05:00"
            }
        }
    }
    ```

### Update Task
- **URL**: `/tasks`
- **Method**: `PUT`
- **Request Body**:
  ```json
  {
      "taskContent": "Hello"
  }
  ```
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
        "state": {
            "status": "Success"
        }
    }
    ```

### Delete Task
- **URL**: `/tasks`
- **Method**: `DELETE`
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
        "state": {
            "status": "Success"
        }
    }
    ```