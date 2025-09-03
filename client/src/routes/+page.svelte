<script lang="ts">
    import { error, user } from '$lib'
    import Tasks from '$lib/components/Tasks.svelte'
    import User from '$lib/components/User.svelte'
    import Loader from '$lib/components/Loader.svelte'
    import { onMount } from 'svelte'

    let { data,form } = $props()
    user.set(data.user)

    let view: string = $state("login")

    onMount(() => {
        if(form !== null
        ){
            if(form.error){
                error.set(form.error)
            }
            if(form.view){
                view = form.view
            }
        }
    })

</script>

<Loader />
<img class="logo" src="/images/logo.svg" alt="Tasker Logo" />
{#if !$user}
    <User {view} />
{:else}
    <Tasks />
{/if}