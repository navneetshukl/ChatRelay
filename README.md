
# document.txt file has all the explanation about this project.


1. Clone the repository.
2. Create .secrets folder and create .env.localhost file.
3. Define the environment varibale. For example consider .env.example.
   SLACK_APP_TOKEN=""
   SLACK_BOT_TOKEN=""
   MOCK_SERVER_BASE_URL=""
   MOCK_SERVER_PORT=""

   You can Get the SLACK_APP_TOKEN and SLACK_BOT_TOKEN from the steps which is defined inside document.txt. For keeping thing simple make MOCK_SERVER_BASE_URL=http://localhost:8080 and MOCK_SERVER_PORT=":8080".
   
4. After cloning and defining the secrets run the below command.
    For starting mock server run : go run cmd/http/main.go
    For starting the main application: go run cmd/main.go

5.  You can also run the application using docker.
    For building the mock server docker image run: docker build -t http_app -f Dockerfile.http .
    For running the mock server docker image run: docker run -d -p 8080:8080 http_app
    For building the main app docker image run: docker build -t main_app -f Dockerfile.main .
    For running the main app docker image run: docker run main_app

6. Alternatively you can run docker-compose up --build.
7. There  is some error comming in communicating between the docker images which I am trying to fix.So I suggest to run the whole application locally or run the mock server in docker image and main app locally.


     
