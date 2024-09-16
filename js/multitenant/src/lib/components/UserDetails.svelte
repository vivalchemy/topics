<script lang="ts">
  //import { getUserAndOrganizations } from "$lib/drizzle/functions";
  import { getContext } from "svelte";

  const user: { username: string; password: string } = getContext("user");
  //const userDetails = getUserAndOrganizations(user.username, user.password);
</script>

<h1 class="text-3xl">User Details</h1>

<!-- <pre class="text-white">{JSON.stringify(userDetails, null, 2)}</pre> -->
<!-- This won't work cause postgres only work on server side we can't call drizzle related functions on client side -->
