# REST API Documentation

## Task Endpoints

### Create Task
- **URL**: `/task`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "userId": 1,
    "taskContent": "Task content"
  }
  ```
- **Response**:
  - **Status**: `201 Created`
  - **Body**:
    ```json
    {
      "status": "ok",
      "userId": 1,
      "taskId": 1
    }
    ```

### Get Tasks by User ID
- **URL**: `/tasks/user`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "userId": 1
  }
  ```
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
      "status": "ok",
      "tasks": [
        {
          "id": 1,
          "content": "Task content"
        }
      ]
    }
    ```

### Get Task by Task ID
- **URL**: `/task`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "taskId": 1
  }
  ```
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
      "status": "ok",
      "task": {
        "id": 1,
        "content": "Task content"
      }
    }
    ```

### Update Task
- **URL**: `/task`
- **Method**: `PUT`
- **Request Body**:
  ```json
  {
    "taskId": 1,
    "taskContent": "Updated task content"
  }
  ```
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
      "status": "ok",
      "taskId": 1
    }
    ```

### Delete Task
- **URL**: `/task`
- **Method**: `DELETE`
- **Request Body**:
  ```json
  {
    "taskId": 1
  }
  ```
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
      "status": "ok"
    }
    ```

## User Endpoints

### Create User
- **URL**: `/user`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "username": "user1",
    "password": "password123"
  }
  ```
- **Response**:
  - **Status**: `201 Created`
  - **Body**:
    ```json
    {
      "status": "ok",
      "userId": 1
    }
    ```

### Get User by User ID
- **URL**: `/user`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "userId": 1
  }
  ```
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
      "status": "ok",
      "user": {
        "id": 1,
        "username": "user1"
      }
    }
    ```

### Update User Password
- **URL**: `/user/password`
- **Method**: `PUT`
- **Request Body**:
  ```json
  {
    "userId": 1,
    "password": "newpassword123"
  }
  ```
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
      "status": "ok",
      "userId": 1
    }
    ```

### Delete User
- **URL**: `/user`
- **Method**: `DELETE`
- **Request Body**:
  ```json
  {
    "userId": 1
  }
  ```
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
      "status": "ok"
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