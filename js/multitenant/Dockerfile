FROM oven/bun:alpine

# create a new user app and assign group app
# will cause some issue it the volumes are mounted and the files in host need to be edited
# RUN addgroup app && adduser -S -G app app
# USER app

WORKDIR /app

COPY package.json bun.lockb ./

# change the access rights of the app user
# USER app

RUN bun --bun install

COPY . .

# USER root
# RUN chown -R app:app .
# USER root

EXPOSE 3000

# CMD [ "sh" ] # for testing
CMD ["bun", "run", "dev"]
