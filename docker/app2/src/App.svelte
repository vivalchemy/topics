<main>
  This is the <h1>Second</h1>
  app

  <h3>Run this inside the app2 directory</h3>
  <pre> 
    docker buildx build --tag app2 .
    docker run -it --rm -p 4000:5173 --volume .:/app app2 </pre>
</main>
