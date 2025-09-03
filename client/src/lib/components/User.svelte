<script lang="ts">
    import { onMount } from "svelte";
    import UserConnect from "./UserConnect.svelte";
    import UserRegister from "./UserRegister.svelte";


    let username:string = $state("")

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
    <UserConnect {view} {username} {title} />
{:else if view === "register"}
    <UserRegister {view} {username} {title} />
{/if}