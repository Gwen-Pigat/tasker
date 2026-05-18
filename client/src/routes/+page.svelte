<script lang="ts">
    import { error, user } from "$lib";
    import Tasks from "$lib/components/Tasks.svelte";
    import User from "$lib/components/User.svelte";
    import Dashboard from "$lib/components/Dashboard.svelte";
    import Loader from "$lib/components/Loader.svelte";

    let { data, form } = $props();
    user.set(data.user);

    let view: "login" | "register" = $state("login");
    let currentTab: "tasks" | "dashboard" = $state("tasks");

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
    <div class="min-h-screen bg-slate-950/20">
        <!-- Navigation Header -->
        <nav class="sticky top-0 z-50 bg-slate-900/80 backdrop-blur-xl border-b border-white/5">
            <div class="max-w-6xl mx-auto px-4">
                <div class="flex items-center justify-between h-16">
                    <div class="flex items-center gap-8">
                        <img src="/images/logo.svg" alt="Logo" class="w-24" />
                        <div class="flex gap-1">
                            <button 
                                onclick={() => currentTab = 'tasks'}
                                class="px-4 py-1.5 rounded-lg text-sm font-bold transition-all {currentTab === 'tasks' ? 'bg-indigo-500 text-white' : 'text-slate-400 hover:bg-white/5'}"
                            >
                                Tasks
                            </button>
                            <button 
                                onclick={() => currentTab = 'dashboard'}
                                class="px-4 py-1.5 rounded-lg text-sm font-bold transition-all {currentTab === 'dashboard' ? 'bg-indigo-500 text-white' : 'text-slate-400 hover:bg-white/5'}"
                            >
                                Dashboard
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </nav>

        {#if currentTab === 'tasks'}
            <Tasks />
        {:else}
            <Dashboard />
        {/if}
    </div>
{/if}
