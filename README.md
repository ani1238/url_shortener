# url_shortener
Simple Url shortener written in GOlang

## Features

- Shorten long URLs into compact and easy-to-share links.
- Redirect users to the original URL when they access the shortened link.
- Metrics API to track the top domains with the most shortened links.
- Data storage using Redis for fast and efficient access.

### Installation
- Initialize and start Redis server.
- Build and run the Go application:
  ```
  go build
  ./your-app-name
  ```

### How to build the project ? 

1. Build the docker image : ` docker build -t urlshortner .`
2. Run the image in detached mode on port 5000 : `docker run -d -p 8080:8080 urlshortner` 
3. Now can serve the api's on `http://localhost:8080/<api endpoint>`

## Usage

To shorten a URL, make a POST request to the /shorten endpoint with a JSON payload:

```curl -X POST -H "Content-Type: application/json" -d '{"url": "http://example.com"}' http://localhost:8080/shorten```

Access a shortened URL to be redirected to the original URL:

```http://localhost:8080/shortened-url```

Retrieve the top domains with the most shortened links:

```http://localhost:8080/metrics/topdomains```
