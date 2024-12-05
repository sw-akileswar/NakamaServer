**Start:**
The recommended workflow is to use Docker and the compose file to build and run the game server, database resources and tensorflow-serving (AI model server).

docker-compose up --build nakama

**Recompile / Run:**
When the containers have been started as shown above you can replace just the game server custom code and recompile it with the -d option.

docker-compose up -d --build nakama

**Stop:**
To stop all running containers you can use the Docker compose sub-command.

docker-compose down

You can wipe the database and workspace with docker-compose down -v to remove the disk volumes.
