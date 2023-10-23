ARG TAG=1.20

FROM ghcr.io/anglo-korean/go-builder-static:$TAG as build
FROM ghcr.io/anglo-korean/go-scratch:$TAG

COPY --from=build /app/app /app
