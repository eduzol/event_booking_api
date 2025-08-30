# Event Booking API

This is a simple Go (Gin) REST API for managing events, with a SQLite backend and a lightweight HTML/jQuery client UI.

## Features
- Create and retrieve events via REST API
- Store events in SQLite
- User-friendly HTML UI for event management

## API Usage

### Retrieve all events

```sh
curl -X GET http://localhost:8080/events
```

### Create a new event

```sh
curl -X POST http://localhost:8080/events \
	-H "Content-Type: application/json" \
	-d '{
		"name": "Sample Event",
		"description": "A test event",
		"location": "Test location",
		"dateTime": "2025-01-01T15:30:00.000Z"
	}'
```

## HTML UI

To use the web client, start the server and open:

```
http://localhost:8080/client/app.html
```

You can create and view events using the browser interface.
