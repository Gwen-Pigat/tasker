<script lang="ts">
    import { error, user } from '$lib'
    import Tasks from '$lib/components/Tasks.svelte'
    import User from '$lib/components/User.svelte'
    import Loader from '$lib/components/Loader.svelte'
    import { onMount } from 'svelte'

    let { data,form } = $props()
    user.set(data.user)

    onMount(() => {
        if(form !== null
        && form.error
        ){
            error.set(form.error)
        }
    })

</script>

<Loader />
<img class="logo" src="/images/logo.svg" alt="Tasker Logo" />
{#if !$user}
    <User />
{:else}
    <Tasks />
{/if}