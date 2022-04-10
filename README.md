<!-- GETTING STARTED -->
## Getting Started

### Prerequisites
- Go(lang)
- Docker

### Installation
1. Clone the repo


<!-- USAGE EXAMPLES -->
## Usage
### Start Up
1. Build and run the executable (`go build .`) OR run the `main.go` file (`go run main.go`)
2. Start the DB Docker container using Docker Compose (docker-compose up)

### API Endpoints
- GET `/api/v1/articles` - Get all news articles
- GET `/api/v1/article/{id}` - Get one specific news article
- GET `/api/v1/article-sources` - Get all available news article sources
- GET `/api/v1/article-source/{id}` - Get one specific news article source
- POST `/api/v1/article-source/{id}` - Set one specific news article source as the active one

<!-- FUTURE FEATURES -->
## Future Features
1. Update the news articles when the active news source changes
2. Implement a Redis cache store
3. Add article filtering by category and provider
