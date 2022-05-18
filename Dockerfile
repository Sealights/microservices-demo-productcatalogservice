FROM golang:1.17.5 as development
ENV GO111MODULE=on

# Add a work directory
WORKDIR /app
COPY productcatalogservice ./
COPY products.json ./

# Expose ports
ENV GOTRACEBACK=single

EXPOSE 3550
# Start app
ENTRYPOINT ["./productcatalogservice"]

