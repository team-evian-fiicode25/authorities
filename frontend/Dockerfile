FROM node:22 AS build

WORKDIR /build

COPY package.json .
COPY yarn.lock .

RUN yarn --frozen-lockfiles

COPY . .

ARG PUBLIC_BACKEND_URL="/api"

RUN yarn run build

FROM node:22

WORKDIR /app

COPY --from=build /build/build .

ENTRYPOINT [ "node", "." ]
