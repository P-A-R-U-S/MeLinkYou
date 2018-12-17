FROM alpine

LABEL maintainer="Valentyn_Ponomarenko@melinkyou.com>"

# install google cloud tools
RUN go get cloud.google.com/go/datastore
RUN go get golang.org/x/net/context
# install glide
RUN go get github.com/Masterminds/glide
# install gin
RUN go get github.com/codegangsta/gin

# create a working directory
WORKDIR /go/src/MeLinkYou/Services/Registration
# add source code

# add glide.yaml and glide.lock
# as you can see that glide.yaml and glide.lock are being added separately (instead of doing a ADD . .),
#   resulting in separate layers. By separating out package management to a separate layer,
#   Docker will cache the layer and will only rebuild it if the corresponding files change,
#                           i.e. when a new package is added or an existing one is removed.
#   Hence, glide install won’t be executed for every source code change.
ADD glide.yaml glide.yaml
ADD glide.lock glide.lock

# install packages
RUN glide install

COPY . .

EXPOSE 8080

# run main.go
#RUN export DATASTORE_PROJECT_ID=<project-id>
CMD ["go", "run", "Service.go"]