<script lang="ts">
    import { onMount } from "svelte";
    const json = (response: Response) => response.json();
    onMount(() => {
        fetch("http://localhost:3000/api/spotify").then(json);
    })
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
