![GoFiery logo](https://github.com/Emeto/gofiery/blob/main/docs/gofiery.png?raw=true)

[![Go Reference](https://pkg.go.dev/badge/github.com/Emeto/gofiery.svg)](https://pkg.go.dev/github.com/Emeto/gofiery)

# GoFiery

GoFiery is a pure go implementation of the Fiery API. It provides a client and easy access to all the API endpoints to help you write your own application.

**This is an in-progress, non-tested project and is therefore not ready for production use**

## What is Fiery API?
The Fiery API is an API available on Fiery controllers,
machines that are used to manage production-grade printers and its printing jobs.
The API helps write applications that talk to a Fiery controller
and provide remote access to the Fiery server.

## Roadmap

- [x] Versions
- [x] Authentication
  - [x] Login
  - [x] Logout
- [ ] Status
- [ ] Info
  - [ ] Info detail
- [ ] Licenses
  - [ ] License info
  - [ ] Activate license
  - [ ] License details
  - [ ] Deactivate license
  - [ ] Get license request file
- [x] Consumables
- [x] Features
- [ ] Printable pages
  - [ ] Lists printable system pages
  - [ ] Request output of printable system page
- [ ] Server
  - [ ] Start the Fiery server
  - [ ] Stop the Fiery server
  - [ ] Perform actions on the Fiery server
  - [ ] Server status
- [ ] Devices
  - [ ] Device info
- [ ] Localize
  - [ ] Retrieve translated user-displayable strings from the Fiery server
- [ ] Jobs
  - [ ] Get jobs
  - [ ] Upload job
  - [ ] Get job details
  - [ ] Delete job
  - [ ] Retrieve PDL content of a job
  - [ ] Get jobs by state
  - [ ] Perform actions on a job
  - [ ] Retrieve job preview
- [ ] Accounting
  - [ ] List accounting info from printed jobs
  - [ ] Retrieve job log entry settings
- [ ] Presets
  - [ ] List presets and attributes
  - [ ] Get preset attribute settings
- [ ] Queues
  - [ ] List printer queues
  - [ ] Get attributes of a printer queue
- [ ] Properties
  - [ ] List all job properties
  - [ ] Retrieve property scopes
  - [ ] Retrieve all properties in scope
  - [ ] Retrieve all properties of particular group
  - [ ] Retrive particular property
  - [ ] Retrieve particular property constraints
  - [ ] Perform constraint check for particular property
- [ ] Paper Catalog
  - [ ] List paper catalog
  - [ ] Retrieve paper catalog entry
