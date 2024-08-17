
```sh
docker run -it --rm -p 3000:5173 --volume .:/app app1
```
-it : interative
--rm: remove the container after stopping
-p : map the expose port HOST:CONTAINER
--volume: mount a volume HOST:CONTAINER(sexy way of mounting volumes)
