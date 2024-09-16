<script lang="ts">
  import { Button } from "$lib/components/ui/button/index.js";
  import * as Card from "$lib/components/ui/card/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import { Label } from "$lib/components/ui/label/index.js";
  import { getContext } from "svelte";

  let username: string = $state("");
  let password: string = $state("");

  const user: { username: string; password: string } = getContext("user");

  console.log("login user get", user);

  function onSubmit() {
    if (!username || !password) return;
    user.username = username;
    user.password = password;
    //console.log(
    //  "context user",
    //  user.username,
    //  "context pass",
    //  user.password,
    //);
  }
  //$effect(() => {
  //  console.log("effect -> username", user.username, "password", user.password);
  //});
</script>

<Card.Root class="w-full max-w-sm">
  <Card.Header>
    <Card.Title class="text-2xl">Login</Card.Title>
    <Card.Description
      >Enter your username below to login to your account.</Card.Description
    >
  </Card.Header>
  <Card.Content class="grid gap-4">
    <div class="grid gap-2">
      <Label for="username">Username</Label>
      <Input id="username" type="username" required bind:value={username} />
    </div>
    <div class="grid gap-2">
      <Label for="password">Password</Label>
      <Input id="password" type="password" required bind:value={password} />
    </div>
  </Card.Content>
  <Card.Footer>
    <Button onclick={onSubmit} class="w-full">Log in</Button>
  </Card.Footer>
</Card.Root>
