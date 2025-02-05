FROM golang:1.23-alpine AS build
WORKDIR /app
COPY . .
RUN go mod tidy
EXPOSE 5050

FROM nginx:latest
WORKDIR /usr/share/nginx/html
EXPOSE 80
COPY --from=build /app /usr/share/nginx/html
CMD ["nginx"]

