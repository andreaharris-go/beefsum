# 🥩 Beef Summary API

A simple Go application that provides a JSON API endpoint to count meat words from a bacon text generator. This project fetches paragraphs from [baconipsum.com](https://baconipsum.com), parses the text, and returns a count of each meat type found.

---

## 🚀 Features

- HTTP API endpoint at `/beef/summary`
- Automatically fetches meat-related text from [baconipsum.com](https://baconipsum.com)
- Parses and counts known meat words
- Returns result in JSON format
- Includes unit tests, integration tests, and performance benchmark

---

## 📦 Installation

Clone the repository and run the application:

```bash
git clone https://github.com/your-user/beef-summary-api.git
cd beef-summary-api
go run main.go
```
By default, the server runs at http://localhost:8080.

---

## 🔗 API Usage

### Endpoint
```swagger codegen
GET /beef/summary
```

### Sample Response
```json
{
  "beef": {
    "brisket": 2,
    "ham": 3,
    "bacon": 1,
    "pork": 4
  }
}

```

---
## 🧪 Testing
Run unit, integration, and benchmark tests:
```shell
go test -v
go test -bench=.
```
Test Coverage 
- ✅ Unit test for meat counting 
- ✅ Mocked integration test for `/beef/summary `
- ✅ Error handling test for text fetch failure 
- ✅ Performance benchmark of `CountMeats`
---
## 🧠 Implementation Details
### Key Components
- `main.go`: starts the HTTP server
- `handler.go`: defines `/beef/summary` logic
- `beef_counter.go`: core text-parsing and meat counting
- `beef_counter_test.go`: includes benchmark test
- `handler_test.go`: includes mocked HTTP handler tests

Example Known Meat Words
- bacon, ham, brisket, pork, corned, jowl, beef, boudin, etc.

 You can customize the known meat list in `beef_counter.go`.


---
## 📎 Notes
- The app uses a function variable `fetchBaconTextFunc` for easy mocking in tests.
- If the fetch fails (e.g. no internet), the server responds with HTTP 500.

---
## 📜 License

MIT License

---
## 🙋‍♂️ Author

Developed by [Duckdev84](https://linktr.ee/duckdev84)
