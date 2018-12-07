run:
	@go run .

dstart:
	@docker-compose up -d

log:
	@docker ps --filter name="blog-api" --format "{{.Names}}" | peco | xargs docker logs -f 
