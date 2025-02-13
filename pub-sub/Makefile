# Variables
KAFDROP_PORT=9000

# Targets
.PHONY: all up down logs kafdrop

# Start the Kafka and Kafdrop containers
up:
	@make down
	@docker-compose up -d
	@docker-compose logs -f

# Stop and remove the containers
down:
	@docker-compose down

# View logs from the Kafka container
logs:
	@docker-compose logs -f kafka

# Open Kafdrop in the browser
kafdrop:
	@echo "Opening Kafdrop in your default browser..."
	@open http://localhost:$(KAFDROP_PORT)

start-sub:
	@ go run subscriber/subscriber.go

start-pub:
	@ go run publisher/publisher.go
