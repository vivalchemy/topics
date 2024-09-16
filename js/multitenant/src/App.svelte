<script>
  import createUserState from "$lib/components/createUserState.svelte";
  import LoginForm from "$lib/components/LoginForm.svelte";
  import ModeToggle from "$lib/components/ModeToggle.svelte";
  import Signup from "$lib/components/Signup.svelte";
  import Button from "$lib/components/ui/button/button.svelte";
  import { setContext } from "svelte";
  import "./app.css";
  import { ModeWatcher } from "mode-watcher";
  import UserDetails from "$lib/components/UserDetails.svelte";

  let isNewUser = $state(false);
  let msg = $derived(isNewUser ? "Log in" : "Sign In");

  // updated at LoginForm.svelte
  const user = createUserState({ username: "", password: "" });
  setContext("user", user);
</script>

<!-- required by mode watcher -->
<ModeWatcher />
<main
  class="max-w-3xl mx-auto min-h-screen flex flex-col justify-center items-center gap-2"
>
  <div class="flex">
    <h1 class="text-3xl flex-1 text-center">Hello world</h1>
    <ModeToggle />
  </div>

  <Button onclick={() => (isNewUser = !isNewUser)}>{msg}</Button>
  {#if isNewUser}
    <Signup />
  {:else}
    <LoginForm />
  {/if}

  <UserDetails />
</main>
