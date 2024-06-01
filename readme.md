# Clean Architecture Diagram

       +-------------------+
       |   Presentation    |       (port: http, grpc)
       |-------------------|
       |   Handlers, DTOs  |       (dto)
       +--------^----------+
                |
       +--------|----------+
       |    Application    |       (service)
       |-------------------|
       |  Business Logic   |
       +--------^----------+
                |
       +--------|----------+
       |      Domain       |       (model)
       |-------------------|
       |   Entities, Repos |
       +--------^----------+
                |
       +--------|----------+
       |   Infrastructure  |       (repository)
       |-------------------|
       | Database, APIs    |
       +-------------------+

## Explanation

## Presentation Layer:

### port/http, port/grpc:

This layer contains the entry points to the application, such as HTTP or gRPC handlers.
It includes the logic for handling incoming requests and sending responses.

### dto:

Data Transfer Objects (DTOs) are used to define the structure of the data that is sent and received by the handlers.
They help in transferring data between the client and the server.

## Application Layer:

### service:

This layer contains the business logic of the application.
It uses the repositories to perform operations and applies business rules.
It orchestrates the flow of data between the Presentation Layer and the Domain Layer.

## Domain Layer:

### model:

This layer contains the core business entities and domain models.
It defines the essential data and behavior of the application without depending on any external frameworks.
Entities here are pure Go structs that represent the core objects.

## Infrastructure Layer:

### repository:

This layer contains the implementation details for data access.
It interacts with the database or other external systems to fetch and store data.
It uses the domain entities to perform CRUD operations.

# Sequence of Writing Code

<!-- <li>internal/address/model/address.go.</li>
<li>internal/address/dto/address.go.</li>
<li>internal/address/repository/address.go.</li>
<li>internal/address/service/address.go.</li>
<li>internal/address/port/http/routes.go  ==>> Contains the HTTP routes.</li>
<li>internal/address/port/http/handlers.go  ==>> Contains the HTTP handlers and routes.</li>
<li>internal/address/port/grpc/server.go  ==>> Initialization, Registration, Configuration for Server and Starting the Server. </li>
<li>internal/address/port/grpc/handlers.go  ==>> Contains the Service Implementation and  Business Logic</li>
<li>proto/address/address.proto  ==>> create proto file and generate pb and grpc_pb for grpc server</li>
<li>internal/server/http/server.go  ==>> Add addressHttp to maproutes</li>
<li>internal/server/grpc/server.go  ==>>  Add addressGrpc to run server</li>
<li>cmd/api/main.go  ==>>  Add addressModel, grpcServer and httpServer to run server</li> -->

<table>
  <thead>
    <tr>
      <th>File Path</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>internal/address/model/address.go</td>
      <td>Defines the data structures for address entities.</td>
    </tr>
    <tr>
      <td>internal/address/dto/address.go</td>
      <td>Contains Data Transfer Objects (DTOs) for address operations.</td>
    </tr>
    <tr>
      <td>internal/address/repository/address.go</td>
      <td>Provides the repository interface and implementations for address persistence.</td>
    </tr>
    <tr>
      <td>internal/address/service/address.go</td>
      <td>Implements the business logic for address-related operations.</td>
    </tr>
    <tr>
      <td>internal/address/port/http/routes.go</td>
      <td>Defines the HTTP routes for address endpoints.</td>
    </tr>
    <tr>
      <td>internal/address/port/http/handlers.go</td>
      <td>Implements the HTTP handlers for address-related requests.</td>
    </tr>
    <tr>
      <td>internal/address/port/grpc/server.go</td>
      <td>Handles initialization, registration, configuration, and starting of the gRPC server for address services.</td>
    </tr>
    <tr>
      <td>internal/address/port/grpc/handlers.go</td>
      <td>Implements the gRPC service methods and business logic for address services.</td>
    </tr>
    <tr>
      <td>proto/address/address.proto</td>
      <td>Defines the Protocol Buffers (proto) schema for the address service and generates the corresponding gRPC code.</td>
    </tr>
    <tr>
      <td>internal/server/http/server.go</td>
      <td>Integrates address HTTP routes into the main HTTP server.</td>
    </tr>
    <tr>
      <td>internal/server/grpc/server.go</td>
      <td>Integrates address gRPC services into the main gRPC server.</td>
    </tr>
    <tr>
      <td>cmd/api/main.go</td>
      <td>Bootstraps and runs the application, including address models, gRPC server, and HTTP server.</td>
    </tr>
  </tbody>
</table>

## 1- Define Domain Models (model):

<li>Create the core business entities.</li>
<li>Example: model/address.go.</li>

```go
package model
type Address struct {
ID string
UserID string
Name string
City string
Street string
Lat string
Long string
}
```

## 2- Define Data Transfer Objects (dto):

<li>Create structs to define the shape of data sent and received.</li>
<li>Example: dto/address.go.</li>

```go
package dto

type AddressDTO struct {
    UserID  string `json:"userid"`
    Name    string `json:"name"`
    City    string `json:"city"`
    Street  string `json:"street"`
    Lat     string `json:"lat"`
    Long    string `json:"long"`
}
```

## 3- Implement Repository Interfaces (repository):

<li>Create interfaces and their implementations for data access.</li>
<li>Example: repository/address.go.</li>

```go
package repository

import (
    "main/model"
    "gorm.io/gorm"
)

type AddressRepository struct {
    db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
    return &AddressRepository{db: db}
}

func (r *AddressRepository) Save(address model.Address) error {
    return r.db.Create(&address).Error
}

func (r *AddressRepository) FindByID(id string) (model.Address, error) {
    var address model.Address
    err := r.db.First(&address, "id = ?", id).Error
    return address, err
}
```

## 4- Implement Business Logic (service):

<li>Implement the core logic that uses the repositories.</li>
<li>Example: service/address.go.</li>

```go
package service

import (
    "main/dto"
    "main/model"
    "main/repository"
    "github.com/google/uuid"
)

type AddressService struct {
    addressRepo repository.AddressRepository
}

func NewAddressService(addressRepo repository.AddressRepository) *AddressService {
    return &AddressService{addressRepo: addressRepo}
}

func (s *AddressService) CreateAddress(addressDTO dto.AddressDTO) error {
    address := model.Address{
        ID:     uuid.New().String(),
        UserID: addressDTO.UserID,
        Name:   addressDTO.Name,
        City:   addressDTO.City,
        Street: addressDTO.Street,
        Lat:    addressDTO.Lat,
        Long:   addressDTO.Long,
    }
    return s.addressRepo.Save(address)
}

```

## 5- Implement Handlers and Routes (port/http):

<li>Create handlers to process incoming requests and route them to the appropriate service methods.</li>

<li>Example: port/http/routes.go.</li>

```go
package http

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "main/service"
    "main/dto"
)

type AddressHandler struct {
    addressService service.AddressService
}

func NewAddressHandler(addressService service.AddressService) *AddressHandler {
    return &AddressHandler{addressService: addressService}
}

func (h *AddressHandler) CreateAddress(c *gin.Context) {
    var dto dto.AddressDTO
    if err := c.ShouldBindJSON(&dto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := h.addressService.CreateAddress(dto)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "success"})
}
```

<li>Example: port/http/handlers.go.</li>

```go
package http

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "main/service"
    "main/dto"
)

type AddressHandler struct {
    addressService service.AddressService
}

func NewAddressHandler(addressService service.AddressService) *AddressHandler {
    return &AddressHandler{addressService: addressService}
}

func (h *AddressHandler) CreateAddress(c *gin.Context) {
    var dto dto.AddressDTO
    if err := c.ShouldBindJSON(&dto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := h.addressService.CreateAddress(dto)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "success"})
}
```

## 7- Implement Handlers and Routes (port/grpc):

<li>Create handlers to process incoming requests and route them to the appropriate service methods.</li>

<li>Example: port/grpc/server.go.</li>

```go
package http

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "main/service"
    "main/dto"
)

type AddressHandler struct {
    addressService service.AddressService
}

func NewAddressHandler(addressService service.AddressService) *AddressHandler {
    return &AddressHandler{addressService: addressService}
}

func (h *AddressHandler) CreateAddress(c *gin.Context) {
    var dto dto.AddressDTO
    if err := c.ShouldBindJSON(&dto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := h.addressService.CreateAddress(dto)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "success"})
}
```

<li>Example: port/grpc/handlers.go.</li>

```go
package http

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "main/service"
    "main/dto"
)

type AddressHandler struct {
    addressService service.AddressService
}

func NewAddressHandler(addressService service.AddressService) *AddressHandler {
    return &AddressHandler{addressService: addressService}
}

func (h *AddressHandler) CreateAddress(c *gin.Context) {
    var dto dto.AddressDTO
    if err := c.ShouldBindJSON(&dto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := h.addressService.CreateAddress(dto)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "success"})
}
```
