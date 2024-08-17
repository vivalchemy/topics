<main>
  This is the <h1>First</h1>
  app
  <h3>Run this inside the app1 directory</h3>
  <pre>
    docker buildx build --tag app1 .
    docker run -it --rm -p 3000:5173 --volume
    .:/app app1
  </pre>
</main>

