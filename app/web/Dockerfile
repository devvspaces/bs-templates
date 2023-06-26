# dev stage
FROM node:16.15-alpine3.16 AS dev

ENV NODE_ENV development

# Needed for API URL
ARG API_URL
ENV API_URL=$API_URL

# setup workdir (next commands will run under this context, but also the running container will inherit this - if not overriden at runtime)
WORKDIR /usr/src/app

# Needed for VS Code Remote Development
RUN apk add musl libgcc libstdc++ git curl wget bash ca-certificates gcompat openssh-client

# Build App
# - copy package.json and yarn.lock (optimise image build - copy only these files and run install, if these change rarely we can reuse this image layer too)
COPY package*.json ./

# - install packages (optimise image build - run just install after copying package.json, we hope these will change rarer as the code, so we can reuse this image layer too)
RUN npm ci --legacy-peer-deps

# Copy the rest of the code
COPY . .

COPY entrypoint.sh /usr/share/entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/usr/share/entrypoint.sh"]

CMD ["npm", "start"]



# prod (and builder) stages
FROM dev AS builder

ENV NODE_ENV production

# Build the app
RUN npm run build



# prod stage
FROM nginx:1.19.7-alpine AS prod

# Copy the built app to the Nginx HTML directory
COPY --from=builder /usr/src/app/dist/web /usr/share/nginx/html
COPY --from=builder /usr/share/entrypoint.sh /usr/share/entrypoint.sh
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 8080

ENTRYPOINT ["/usr/share/entrypoint.sh"]

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]

