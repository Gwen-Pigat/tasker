<script lang="ts">
    import { error, user } from "$lib";
    import Tasks from "$lib/components/Tasks.svelte";
    import User from "$lib/components/User.svelte";
    import Loader from "$lib/components/Loader.svelte";

    let { data, form } = $props();
    user.set(data.user);

    let view: "login" | "register" = $state("login");

    $effect(() => {
        if (form !== null) {
            if (form.error) {
                error.set(form.error);
            }
            if (form.view) {
                view = form.view as "login" | "register";
            }
        }
    });
</script>

<Loader />
{#if !$user}
    <User {view} />
{:else}
    <Tasks />
{/if}
