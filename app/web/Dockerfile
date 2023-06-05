# dev stage
FROM node:15-alpine AS dev
WORKDIR /usr/src/app

ARG BASE_API_URL
ENV BASE_API_URL $BASE_API_URL

# Copy package.json and package-lock.json
COPY package*.json ./

# - install packages (optimise image build - run just install after copying package.json, we hope these will change rarer as the code, so we can reuse this image layer too)
RUN npm ci

# Copy the rest of the code
COPY . .

EXPOSE 80

CMD ["npm", "start"]

# prod (and builder) stages
FROM dev AS builder

RUN npm run build

# prod stage
FROM nginx:1.19.7-alpine AS prod

COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /usr/src/app/dist/web /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]


