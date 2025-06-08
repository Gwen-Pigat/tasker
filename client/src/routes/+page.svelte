<script lang="ts">
    import { user } from '$lib'
    import Tasks from '$lib/components/Tasks.svelte';
    import User from '$lib/components/User.svelte';
    import Loader from '$lib/components/Loader.svelte';
    import { onMount } from 'svelte';

    let data = $props()

    onMount(() => {
        user.set(data.data.user)
        console.log(data.form)
    })

</script>

<Loader />
<img class="logo" src="/images/logo.svg" alt="Tasker Logo" />
{#if !$user}
    <User result={data.form} />
{:else}
    <Tasks />
{/if}