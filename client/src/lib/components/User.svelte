<script lang="ts">
    import { onMount } from "svelte";
    import UserConnect from "./UserConnect.svelte";
    import UserRegister from "./UserRegister.svelte";
    import { url } from "$lib";


    let username:string = $state("")
    let password:string = $state("")
    let title:string = $state("")
    let input:undefined|HTMLInputElement = $state()

    let { view } = $props()

    $effect(() => {
        if(view === "login"){
            title = "Connect"
        } else if(view === "register"){
            title = "Register"
        }
        triggerFocus()
    })

    function triggerFocus()
    {
        input?.focus()
    }

    onMount(() => {
       triggerFocus()
    })

</script>

<svelte:head>
    <title>Tasker | {title}</title>
</svelte:head>

{#if view === "login"}
    <UserConnect bind:view {username} {password} {title} />
{:else if view === "register"}
    <UserRegister bind:view {username} {password} {title} />
{/if}