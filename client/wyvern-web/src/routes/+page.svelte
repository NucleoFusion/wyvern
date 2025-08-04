<script>
  import Button from "$lib/components/ui/button/button.svelte";
  import * as Card from "$lib/components/ui/card/index";
  import { Github } from "@lucide/svelte";
</script>

<div class="w-[100vw] h-[100vh] grid justify-center items-center">
  <Card.Root class="">
    <Card.Header>
      <Card.Title class="text-center text-3xl">Wyvern</Card.Title>
    </Card.Header>
    <Card.Content class="h-[15vh] w-[25vw] grid items-center justify-center">
      <Button
        onclick={() => {
          window.location.href = "http://localhost:3000/auth/github";
        }}
        class="text-md p-[10px 5px]">Login with Github <Github /></Button
      >
    </Card.Content>
  </Card.Root>
</div>
