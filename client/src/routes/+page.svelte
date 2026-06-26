<script lang="ts">
    import { error, user } from "$lib";
    import Tasks from "$lib/components/Tasks.svelte";
    import User from "$lib/components/User.svelte";
    import Dashboard from "$lib/components/Dashboard.svelte";
    import Notes from "$lib/components/Notes.svelte";
    import Loader from "$lib/components/Loader.svelte";
    import Navigation from "$lib/components/Navigation.svelte";

    let view: "login" | "register" = $state("login");
    let currentTab: "dashboard" | "tasks" | "notes" = $state("dashboard");
</script>

<Loader />

{#if !$user}
    <User {view} />
{:else}
    <div class="min-h-screen bg-slate-950/20">
        <Navigation bind:currentTab />
        {#if currentTab === "tasks"}
            <Tasks />
        {:else if currentTab === "dashboard"}
            <Dashboard />
        {:else if currentTab === "notes"}
            <Notes />
        {/if}
    </div>
{/if}
