# set the base image. Since we're running 
# a Python application a Python base image is used
FROM  golang:1.16-alpine
# set a key-value label for the Docker image
LABEL AppType="simple Go website with 3 endpoints"
# copy files from the host to the container filesystem. 
# For example, all the files in the current directory
# to the  `/app` directory in the container
COPY go.mod ./
COPY go.sum ./
COPY . go/src/app
#  defines the working directory within the container
WORKDIR go/src/app
# run commands within the container. 
# For example, invoke a pip command 
EXPOSE 6111
#RUN pip install -r requirements.txt
# provide a command to run on container start. 
# For example, start the `app.py` application.
CMD [ "go", "run","main.go"]