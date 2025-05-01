<script lang="ts">
    import { onMount } from 'svelte';
    import { fetchAPI } from '$lib/_core';
    import { user } from '$lib'
    import Tasks from '$lib/components/Tasks.svelte';
    import User from '$lib/components/User.svelte';
    import Loader from '$lib/components/Loader.svelte';

    let data = $props()
    console.log(data.user)
    if(data.user !== undefined){
        user.set(data.user) 
    }

    async function loadUser(){
        const data = await fetchAPI("/user","GET")
        userConnect = true
        if(data.error){
            return
        }
        user.set(data)
    }

    let userConnect:boolean = $state(false)

    onMount(() => {
        loadUser()
    })

</script>

<Loader />
<img class="logo" src="/images/logo.svg" alt="Tasker Logo" />
{#if !userConnect}
    <progress></progress>
{:else}
    {#if !$user.id}
        <User />
    {:else}
        <Tasks />
    {/if}
{/if}