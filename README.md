# Sample app: Task Service
A simple Golang web application for testing purposes.

## API
| Endpoint | Type   | Usage             |
| -------- | ------ | ----------------- |
| `/tasks` | `GET`  | Get all tasks     |
| `/tasks` | `POST` | Create a new task |



## How to start
- Clone the repository and `cd` inside it:
  ``` bash
  git clone https://github.com/thai-nm/sample-task-service.git
  cd sample-task-service
  ```

- Build Golang code:
  ``` bash
  go build
  ```
- Start server:
  ``` bash
  ./sample-task-service
  ```
  
## Using Docker
You can run this application using Docker:

  ```bash
  docker run -d -p 8080:8080 --name task-service nmthai/sample-task-service
  ```

Or build it yourself:

  ```bash
  docker build -t task-service .
  docker run -d -p 8080:8080 --name task-service task-service
  ```
  
## API call examples

- `POST` request:
  ``` bash
  curl --header "Content-Type: application/json" \
  --data '{"Name":"Project A", "Description":"A simple description"}' \
  -X POST http://localhost:8080/tasks
  ```

- `GET` request:
  ``` bash
  curl http://localhost:8080/tasks
  ```