# NLW Journey Golang

This repository contains the code created during the NLW Journey event by Rocketseat.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Routes](#routes)
- [TODO](#todo)

## Installation

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/mwives/nlw-journey-golang.git
cd nlw-journey-golang
```

Build the Docker containers:

```bash
docker-compose up --build
```

## Usage

To start the application, run:

```bash
docker-compose up
```

## Routes

### Trips

- **Confirm a trip and send e-mail invitations**  
  `GET /trips/{tripId}/confirm`

- **Create a new trip**  
  `POST /trips`

- **Get a trip details**  
  `GET /trips/{tripId}`

- **Update a trip**  
  `PUT /trips/{tripId}`

### Participants

- **Confirms a participant on a trip**  
  `PATCH /participants/{participantId}/confirm`

- **Invite someone to the trip**  
  `POST /trips/{tripId}/invites`

- **Get a trip participants**  
  `GET /trips/{tripId}/participants`

### Activities

- **Create a trip activity**  
  `POST /trips/{tripId}/activities`

- **Get a trip activities**  
  `GET /trips/{tripId}/activities`

### Links

- **Create a trip link**  
  `POST /trips/{tripId}/links`

- **Get a trip links**  
  `GET /trips/{tripId}/links`

## TODO

- Finish the endpoints.
- Add user authentication.
- Add HTML templates for emails.
