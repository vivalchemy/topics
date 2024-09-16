<script lang="ts">
  import { Button } from "$lib/components/ui/button/index.js";
  import * as Card from "$lib/components/ui/card/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import { Label } from "$lib/components/ui/label/index.js";
  //import { insertUser } from "$lib/drizzle/functions"; // require nodejs Buffer module so it can only be run on server side sed life
  //
  let name: any = $state("");
  let username: any = $state("");
  let password: any = $state("");

  //function onsubmit() {
  //  if (name && username && password) insertUser(name, username, password);
  //}
</script>

<Card.Root class="mx-auto max-w-sm">
  <Card.Header>
    <Card.Title class="text-xl">Sign Up</Card.Title>
    <Card.Description
      >Enter your information to create an account</Card.Description
    >
  </Card.Header>
  <Card.Content>
    <div class="grid gap-4">
      <div class="grid gap-2">
        <Label for="name">Name</Label>
        <Input id="name" placeholder="Max" required bind:value={name} />
      </div>
      <div class="grid gap-2">
        <Label for="username">Username</Label>
        <Input
          id="username"
          type="username"
          placeholder="m@example.com"
          required
          bind:value={username}
        />
      </div>
      <div class="grid gap-2">
        <Label for="password">Password</Label>
        <Input id="password" type="password" required bind:value={password} />
      </div>
      <Button type="submit" class="w-full" onclick={onsubmit}
        >Create an account</Button
      >
    </div>
  </Card.Content>
</Card.Root>
