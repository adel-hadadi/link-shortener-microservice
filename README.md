# Link Shortener Microservice

A Golang-based link shortener microservice utilizing Clean Architecture principles and the CQRS pattern. This microservice comprises three components: Gateway, Link, and Report services. The API Gateway handles requests, the Link Service manages URL shortening and retrieval, and the Report Service generates daily Excel reports on link usage, storing them in MinIO.

This project leverages oapi-codegen to generate type-safe Go code from OpenAPI specifications, ensuring efficient and consistent API interactions. It also incorporates Docker and Docker Compose for containerization, allowing seamless deployment and management of all services. Communication between services is handled via gRPC, enhancing performance and scalability.

## Project Structure

- **Gateway Service**: API gateway that handles incoming HTTP requests and forwards them to the appropriate service using gRPC protocol.
- **Link Service**: Manages the creation and retrieval and generate of shortened links, and store the links in a `PostgreSQL` database.
- **Report Service**: Store clicks data in a `MongoDB` database and Generates daily Excel reports for link clicks and uploads these reports to `MinIO`. The reports file available for download base on time.

## Features

- **CQRS Pattern**: Segregates write (command) and read (query) responsibilities, enhancing scalability and maintainability.
- **MinIO Integration**: Stores daily reports on MinIO, a high-performance object storage system.
- **Clean Architecture**: Ensures separation of concerns, enabling easier testing and future extension.
- **gRPC Communication**: Services communicate using gRPC, ensuring efficient, high-performance interactions between components.
- **Daily Reporting**: Schedules report generation using [gocron](https://github.com/go-co-op/gocron) to generate daily click data in Excel format using goland [excelize](https://github.com/qax-os/excelize) package.

## Tech Stack

- **Golang**: Core language used for the microservices.
- **gRPC**: Communication protocol between services.
- **MinIO**: Object storage solution for storing reports.
- **MongoDB**: Database for storing click data.
- **PostgreSQL**: Database for storing shortened links information.
- **Docker & Docker Compose**: Used to containerize each service and manage multi-container environments.
- **oapi-codegen**: Generates Go code from OpenAPI specifications for type-safe API definitions and request handling.

## Setup and Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/adel-hadadi/link-shortener
   cd link-shortener
   ```

2. **Configure Environment Variables**:
   In each of the microservices a `.env.example` file exist. Create a copy of this file and rename it to `.env`.

3. **Build and Run**:
   To run all microservices, you only need a single make command. This command performs the following steps:
   1. Compiles the binary for each microservice using the Go compiler.
   2. Stops any running Docker containers for a clean start.
   3. Launches all microservices as Docker Compose containers.

```bash
make up-build
```

> Makefile:
> For other available `Make` commands can see the `makefile`.

## Endpoints

This project uses **oapi-codegen** to generate Go code from an OpenAPI specification, ensuring type-safe interactions with the API. The generated code handles route definitions and request/response struct definitions, reducing boilerplate code and enhancing API consistency.

```yaml
openapi: "3.0.0"
info:
  title: Api Gateway
  description: TODO
  version: 1.0.0
servers:
  - url: "http://{hostname}/api"
    variables:
      hostname:
        default: localhost
security:
  - bearerAuth: []
paths:
  /{shortURL}:
    get:
      operationId: redirectToURL
      parameters:
        - in: path
          name: shortURL
          required: true
          schema:
            type: string
      responses:
        "301":
          description: Redirect to original URL
          headers:
            Location:
              description: URL to which the request is redirected
              schema:
                type: string
                format: url
          content: {}
        "404":
          description: Link not found
          content:
            application/json:
              schema:
                type: object
                properties:
                  error:
                    type: string
                    example: "Link not found"

  /links:
    post:
      operationId: createLink
      requestBody:
        description: TODO
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/CreateLink"
      responses:
        "201":
          description: todo
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Link"

  /reports/download/{filename}:
    get:
      description: Stream an Excel file
      operationId: downloadExcelFile
      parameters:
        - name: filename
          in: path
          required: true
          description: the name of the Excel file
          schema:
            type: string
      responses:
        "200":
          description: Successfully streams the Excel file
          content:
            application/vnd.openxmlformats-officedocument.spreadsheetml.sheet:
              schema:
                type: string
                format: binary
                example: 2024-11-09
          headers:
            Content-Disposition:
              description: Indicates the file download attachment with the original filename
              schema:
                type: string

components:
  schemas:
    Link:
      type: object
      required: [original_link, short_link]
      properties:
        original_link:
          type: string
        short_link:
          type: string
    CreateLink:
      type: object
      required: [original_link]
      properties:
        original_link:
          type: string
```
