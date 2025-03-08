# Week 1 - Language Portal Backend (Go) - Development and Testing

## Overview

This document outlines the ongoing development of the Go backend server for a language learning portal and the approach taken to test its API endpoints. The backend is being built using the **Gin framework** and utilises **SQLite3** for database connectivity \. My testing strategy involves using **RSpec**, a testing framework for Ruby, chosen for its **developer-friendliness and readability**. The goal is to implement all required API endpoints according to the backend technical specifications file.

## Development Summary

This week, the focus was on setting up the Go backend server \. Key tasks completed included:

*   **Server Setup**: A basic Go server structure was created using the Gin framework, along with the main server file (`cmd/server/main.go`), initial database connection setup using SQLite3, and configuration of the router with all required endpoints \.
*   **API Endpoints Implementation**: An initial set of **18 API endpoints** were implemented across five core API groups: Dashboard, Study Activities, Words, Groups, and Study Sessions \. 
*   **Handler Implementation**: Handlers were created for each API group to manage request processing (`DashboardHandler`, `ActivityHandler`, `WordHandler`, `GroupHandler`, `StudyHandler`) \.
*   **Service Layer**: A service layer was implemented to house the business logic for each API group (`ActivityService`, `WordService`, `GroupService`, `StudyService`) \. Initially, these services returned mock data \.
*   **Testing Setup**: **Comprehensive RSpec tests** were set up for all implemented endpoints, including the creation of test files and implementation of test cases to validate response structure and required fields \.

The project structure was initially organised as follows \:

```
lang-portal/backend_go/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── dashboard.go
│   │   │   ├── activity.go
│   │   │   ├── groups.go
│   │   │   ├── study.go
│   │   │   └── words.go
│   │   └── router.go
│   ├── models/
│   ├── repository/
│   └── service/
└── db/
    ├── migrations/
    └── seeds/
```

The structure was later refined, considering web development patterns and Go-specific conventions \.

## Testing with Ruby Specs

Following the initial implementation, a comprehensive suite of RSpec tests was developed to ensure the API endpoints function as expected \. This involved:

*   Considering using Go's built-in testing but opting for **RubySpec due to its developer-friendliness and readability** \
*   Setting up a Ruby environment, which required potentially installing `rvm`, Ruby, and `bundler` \. Challenges were faced with ensuring the correct Ruby environment and bundler setup \
*   Writing test specifications in Ruby within the `api_tests/spec/api/` directory, focusing on ensuring endpoints **return the expected JSON or response** (good or bad) \. The importance of **readable test code** is emphasized.
*   Addressing challenges such as issues with loading `spec_helper.rb` \ and ensuring correct file paths, sometimes requiring the use of `require_relative` \. Issues with uninitialised constants in the `spec_helper` were also encountered \.
*   Iteratively running and debugging the tests \, often leading to attempts to fix both the test code and the backend Go code to match expectations \. The process involved stepping through each endpoint \ and manually reviewing the JSON responses \.
*   Implementing a test database 
*   Managing the Go server lifecycle during testing, including manually stopping and restarting the server \. 
*   Troubleshooting failures in the specs, which ranged from incorrect response codes and data structures to issues with database state and missing data \. This often involved focusing on individual spec files for more targeted debugging \. The need for **comprehensive RSpec tests for all endpoints** is emphasised \.

# Challenges Faced (Development & Testing)

Both during the initial development and the subsequent testing, several challenges were encountered:

*   **Setting up the correct response format for each endpoint**: Ensuring that the API returns data in the expected JSON structure and with the correct data types proved to be an ongoing effort \. This involved defining JSON schemas before backend implementation and cross-referencing front-end technical specs \. Issues with generated JSON being incorrect were noted \.
*   **Handling pagination in list endpoints**: While the need for pagination was identified as a development challenge \, testing revealed issues with the pagination structure in API responses \. A consistent approach to pagination responses, using 'items', was adopted \.
*   **Managing proper error responses**: Verifying that the API returns appropriate HTTP status codes (e.g., 200) for various scenarios was a key aspect of testing \.
*   **Ensuring consistent API structure across all endpoints**: Maintaining a uniform design across the API was a goal, and testing helped to identify any inconsistencies \. The front-end design significantly influenced the backend API structure \.
*   **Data inconsistencies in the test environment**: Running tests against a development database led to issues with data availability and state \, necessitating the creation and use of a test database and the implementation of seed data loading for both the main and test databases \.
*   **Managing the Go server lifecycle during testing**: Starting, stopping, and ensuring the server was running with the correct configuration for the test environment required manual intervention at times \. 
*   **Issues with the test setup and execution**: Problems with the Ruby environment configuration, RSpec command usage (e.g., `bundle exec rspec`), and loading test files (`spec_helper.rb`) were encountered and resolved iteratively \.

## Summary So Far

The initial development of the Language Portal backend laid the groundwork for a comprehensive API. The adoption of RSpec for testing allowed for the creation of readable and maintainable test specifications. While the testing process involved iterative debugging and troubleshooting, significant progress was made in verifying the functionality of the implemented endpoints. Continuous testing throughout the remaining development phases, along with the implementation of remaining features and proper task automation using Mage, will be essential to ensure the reliability and quality of the Language Portal backend \. The technical specifications were considered a crucial reference point throughout the development and testing process \. Code reviews using tools like Git Smart were also part of the workflow \.

## OPEA Mega Service Build

This entry documents my initial attempt to build a basic "Mega service" using the OPEA framework. The goal was to integrate foundational code from the Open Platform for Enterprise AI (OPEA) into my app, address any import errors, and understand the basic structure.

[OPEA Documentation](https://opea-project.github.io/latest/index.html)

---

### Step 1: Project Setup
I started by creating a new folder within my OPEA Comps directory named `mega-service`. Inside this folder, I created a Python file named `app.py` to house the service code.

---

### Step 2: Adding Initial Code
Using OPEA's technical documentation, I copied the foundational code for a "Mega service" into `app.py`:

```python
from typing import Dict, Any
import os
from comps.core.microservice import MicroService

class ExampleService(MicroService):
    @property
    def service_type(self) -> str:
        return "example"

    async def handle_request(self, data: Dict[str, Any]) -> Dict[str, Any]:
        return {"message": "Hello from the example service!"}

app = ExampleService()
```

---

### Step 3: Resolving Import Errors

The initial code raised unresolved import errors, particularly for `comps`. While `os` is a standard library, `comps` is part of OPEA. According to the documentation, I needed to install `opa-comps`:

```bash
pip install -r requirements.txt
```

I assumed a `requirements.txt` file already included `opa-comps`. Although the installation seemed to work, import errors persisted.

---

### Step 4: Investigating ServiceType

The `service_type` property wasn’t recognized. Searching the OPEA Comps codebase revealed that `ServiceType` is part of `comps.core.mega.constants`. I updated the import:

```python
from comps.core.mega.constants import ServiceType

class ExampleService(MicroService):
    @property
    def service_type(self) -> str:
        return ServiceType.EXAMPLE  # Assuming an EXAMPLE enum exists
```

---

### Step 5: Running the Service

Running the service:

```bash
python app.py
```

resulted in a `cannot import microservice from comps` error, suggesting Python couldn’t locate the `comps` package.

---

### Step 6: Learning from OPA Examples

I explored OPA’s example projects, particularly the `chat_qa` example. Its Dockerfile pointed to a Python script, `chat_qa.py`, that included a `start` function to run the service.

---

### Step 7: Implementing a Start Function

Following the example, I added a `start` method:

```python
if __name__ == "__main__":
    example_service = ExampleService()
    example_service.start()
```

---

### Step 8: Fixing Missing Attributes

Running the app again threw an error about a missing `endpoint`. I added it to the `ExampleService` class:

```python
class ExampleService(MicroService):
    def __init__(self):
        super().__init__()
        self.endpoint = "example_service"
```

---

### Step 9: Basic Web Server with FastAPI

To keep the service running and accepting requests, I integrated FastAPI:

```python
from fastapi import FastAPI
from typing import Dict, Any
import uvicorn

app = FastAPI()

example_service = ExampleService()

@app.post(example_service.endpoint)
async def handle(data: Dict[str, Any]):
    return await example_service.handle_request(data)

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
```

Installed dependencies:

```bash
pip install fastapi uvicorn
```

---

### Step 10: Testing the Service

To test, I used `curl`:

```bash
curl http://localhost:8000/example_service -X POST -H "Content-Type: application/json" -d '{"prompt": "Hello"}'
```

Initially, this returned a `connection refused` error since no LLM service was running.

---

### Step 11: Integrating LLM (ollama)

I launched an ollama instance on port 9000 to simulate a local LLM server:

```bash
llama serve --model llama --port 9000
```

I then updated the `handle_request` method to interact with ollama:

```python
    async def handle_request(self, data: Dict[str, Any]) -> Dict[str, Any]:
        response = await self.service_orchestrator.invoke_remote_service(
            "llm_service",
            {"prompt": data.get("prompt", "")}
        )
        return {"response": {"choices": [{"message": {"content": response.get('content', '')}}]}}
```

---

### Step 12: Disabling Telemetry

Errors from Open Telemetry suggested the Mega service was trying to send data to a default telemetry port. I disabled telemetry by adding:

```python
os.environ["OTEL_EXPORTER_OTLP_ENDPOINT"] = ""
os.environ["OTEL_TRACES_EXPORTER"] = "none"
os.environ["OTEL_METRICS_EXPORTER"] = "none"
```

---

### Step 13: Debugging Responses

After further testing, I noticed empty responses (`"content": null`). This suggested a mismatch between ollama’s output and the Mega service’s expectations.

---

### Step 14: Aligning Request/Response Formats

To ensure compatibility, I revisited the `chat_qa.py` example's `handle_request` implementation. It highlighted the need for matching request/response formats between the Mega service and ollama.

---

### Conclusion

This initial build involved:
- Setting up a Mega service using OPEA.
- Resolving import errors by properly installing and importing `opa-comps`.
- Implementing a `start` function and an endpoint.
- Integrating FastAPI for request handling.
- Connecting the service to a local LLM (ollama).
- Troubleshooting issues with telemetry and response formats.

The next step is to dive deeper into OPEA’s `chat_qa.py` example, ensuring my `handle_request` method properly processes LLM responses.

# Vocabulary Importer Development

This entry outlines the development of an AI-powered vocabulary importer. The goal was to build an internal tool to quickly populate a language learning app with words and word groups.

## UI Prototyping with v0

I began by using **v0** by Vercel to generate the frontend, leveraging its ability to scaffold Next.js components rapidly. My prompt specified the need for:

- A text field for thematic categories
- A button to trigger vocabulary generation via an API

v0 produced a basic UI with an input field and button, but it became clear that while v0 excelled at frontend generation, backend integration—especially with Language Models (LLMs)—would require additional work.

## Transition to Online IDEs: CodeSandbox and Replit

To build the backend and integrate the LLM, we explored online IDEs with both frontend and server-side capabilities.

### CodeSandbox Attempt

Our first stop was **CodeSandbox**. The plan was to:

1. Download the v0-generated frontend as a ZIP file.
2. Import it into a Next.js App Router project on CodeSandbox.

However, we faced configuration mismatches and unresolved platform issues that prevented the code from running smoothly.

### Success with Replit

We then switched to **Replit**, creating a new Next.js App Router project. The following steps were taken to integrate the v0-generated code:

- **File Import**: Dragged and dropped `components`, `hooks`, and `lib` folders into Replit.
- **Dependencies**: Replaced Repet's `package.json` with v0's version, adding packages like `ai`, `@radix-ui/react-*`, and `zod`.
- **Styling**: Updated `tailwind.config.ts` with v0's configuration.
- **Installation**: Ran `npm install` to install dependencies.

## LLM Integration with GroqCloud

We chose **GroqCloud** for its fast, serverless API and generous free tier.

### API Key Management

To secure the GroqCloud API key:

- Used Repet's **Secrets** feature to store `GROQ_API_KEY`.
- Accessed it in code via `process.env.GROQ_API_KEY`.

### Error Handling

We faced two main issues during setup:

1. **"Missing API Key" error**: Resolved by verifying the key in Replit's Secrets and restarting the dev server.
2. **"Model Not Found" error**: Fixed by consulting Groq's API console and updating the `model` parameter to `gemma-7b-it`.

## Fine-Tuning JSON Output with Prompt Engineering

The LLM initially returned JSON wrapped in backticks and "text" fields, causing syntax errors. To fix this, we refined the prompt:

```typescript
const prompt = `Generate at least five vocabulary items related to the topic: "${topic}".
The output should be a JSON array of objects, each structured as:
{ "word": "...", "parts": [...] }
Ensure the output is valid JSON and nothing else.`;
```

This adjustment successfully generated clean JSON.

## Enhancing Vocabulary Structure

To ensure accurate segmentation  pronunciations, I used **few-shot learning** by:

- Providing **bad vs. good examples** of the desired structure.
- Demonstrating correct character breakdowns and transliteration outputs.

After iterative prompt tuning, the LLM began generating precise vocabulary structures.

## Final Output and Reflections

The final importer produced JSON outputs with arabic word breakdowns, easily copyable via a UI button.

Key takeaways:

- **Rapid prototyping** with LLMs is powerful, but prompt engineering is crucial.
- **GroqCloud** offered cost-effective, fast inference for AI tasks.
- Iterative refinement boosted accuracy and consistency.
- Fine-tuning smaller models may further enhance efficiency and reduce costs.

This journey reinforced the importance of blending AI capabilities with thoughtful engineering, laying a strong foundation for future language learning innovations.

