# reserv-go
A Project to explore Clean Architecture principles and strict separation of concerns.

## Live API
https://reservgo.onrender.com

--- 

## Responsabilities per layer

| Layer              | What it knows           | What it does NOT know |
| ---                | ---                     | --- |
| handlers (cmd/api) | HTTP, gin, JSON         | business logic/rules |
| usecases           | full flow, modules      | HTTP, DB |
| modules/service    | it's own domain rules   | Other module concerns |
| modules/repository | DB interface            | business logic/rules |
| infrastructure     | Stripe, Supabase, etc   | internal logic |

---

## Layers

### `cmd/api/` — Entry point HTTP

Everything here is related to the server concern and it is the only layer that can instanciate dependencies and be connected through wiring (DI).
Like the pool conection to the database, services, usecases, handlers, gin, among others.

- `main.go` — Entry point.
- `routes.go` — Endpoints for the application.
- `handlers/*` — HTTP handlers for the endpoints.
  Recieves http requests, gets parameters, calls module's services, serializes the reply. Knows nothing about bussiness logic.
  
  **Why does the handlers live at `cmd/` instead of `internal/`?**
  
  In Go, `internal/` is a special folder that has a compiler restriction: no package outside this Go module can import anything inside this folder. 
  Within this module, all of the packages can import `internal/` without issue. The restriction only applies to external consumers. 
  This isn't the technical reason for not putting handlers there. My real reason is `internal/` is for reusable business logic across multiple project binaries. Handlers are HTTP adapters specific to the binary API, they don't make sense in a worker, a CLI, or any other entry point. That's why it belongs in `cmd/api/`, alongside the router and the entry point that uses it.

---

### `configs/` — App configuration

- `env.go` — Loads environmental variables.

### `infrastructure/` — External Adapters

Everything here is related to external services or utilities overall; Databases, Email-ing, Other APIS, etc.
This layer implemments interfaces of `internal/modules/`. 

- `supabase/supabase.go` — Connection to Database.
- `supabase/repos/*.go` — Data Acces Layer for internal/modules.


### `internal/` — Business Layer

Everything here is related to Bussiness Layer only. This layer must not know nothing about HTTP, Supabase or providers.
The only concern is to know business rules/logic and provide interfaces.

#### `internal/modules/` — Domain Modules

Each module is atomic and independient. No module caanot import another module. 
If two or more modules needs coordination, that is a `usecase` concern.

- `modules/*/repository.go` — Data Acces Layer interface (contract).
  Declares what functions must implemment any function for such module. DAL Equivalency for C#
  Contrains no implemmentation, only the interface.

- `modules/*/service.go` — Business Logic for the module
  Declares the Business Logic for such module. Gets it's Database calls through the repository interface in this same folder.

- `modules/*/models.go` — Data Transfer Objects of the module
  Models for the module. Implemments structs that represents entities of this module.
  This models can be shared for `infrastructure/`, Infrastructure can import the module and the interface for implemmentations, 
  but this module cannot import anything from infrastructure, that is explained previously.


#### `internal/usecases/` — Orquestation

- One file per 'workflow' 
(idk how to call this, it is just a set of instructions that require two or more modules to reach a goal)
    
    Coordinates multiple modules to accomplish a job. This are the only ones that can import more than one module.
    This files should be created only when more than two modules need eachother or have secondary effects chained.
    Example: 
      Payment module returns a confirmation, then i save that data in the database, then whatever the database returns, 
      i can generate a QR code or something.

### `internal/shared/` — Utilies
- Code that any module can use
    Envelopes, Json, Pagination, Validations, among others.
