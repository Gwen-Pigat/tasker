<script lang="ts">
    import { onMount } from "svelte";
    import UserConnect from "./UserConnect.svelte";
    import UserRegister from "./UserRegister.svelte";
    import { url } from "$lib";

    let username: string = $state("");
    let password: string = $state("");
    let title: string = $state("");
    let input: undefined | HTMLInputElement = $state();

    let { view }: { view: "login" | "register" } = $props();

    $effect(() => {
        if (view === "login") {
            title = "Connect";
        } else if (view === "register") {
            title = "Register";
        }
        triggerFocus();
    });

    function triggerFocus() {
        input?.focus();
    }

    onMount(() => {
        triggerFocus();
    });
</script>

<svelte:head>
    <title>Tasker | {title}</title>
</svelte:head>

<div
    class="min-h-screen flex flex-col items-center justify-center p-6 bg-slate-950/20"
>
    <div class="mb-12 animate-in">
        <img
            src="/images/logo.svg"
            alt="Tasker Logo"
            class="w-100 drop-shadow-[0_0_15px_rgba(99,102,241,0.3)]"
        />
    </div>

    {#if view === "login"}
        <UserConnect bind:view {username} {password} {title} />
    {:else if view === "register"}
        <UserRegister bind:view {username} {password} {title} />
    {/if}

    <div
        class="mt-12 text-slate-500 text-xs font-medium uppercase tracking-widest animate-in [animation-delay:200ms]"
    >
        Secure & Optimized Task Management
    </div>
</div>
