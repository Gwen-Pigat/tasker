<script lang="ts">
    import { error, user } from "$lib";
    import Tasks from "$lib/components/Tasks.svelte";
    import User from "$lib/components/User.svelte";
    import Dashboard from "$lib/components/Dashboard.svelte";
    import Loader from "$lib/components/Loader.svelte";

    let { data, form } = $props();
    user.set(data.user);

    let view: "login" | "register" = $state("login");
    let currentTab: "dashboard" | "tasks" = $state("dashboard");

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
                                onclick={() => currentTab = 'dashboard'}
                                class="px-4 py-1.5 rounded-lg text-sm font-bold transition-all {currentTab === 'dashboard' ? 'bg-indigo-500 text-white' : 'text-slate-400 hover:bg-white/5'}"
                            >
                                Dashboard
                            </button>
                            <button 
                                onclick={() => currentTab = 'tasks'}
                                class="px-4 py-1.5 rounded-lg text-sm font-bold transition-all {currentTab === 'tasks' ? 'bg-indigo-500 text-white' : 'text-slate-400 hover:bg-white/5'}"
                            >
                                Tasks
                            </button>
                        </div>
                    </div>
                    <div class="space-y-8">
                        <form method="POST" action="?/logout">
                            <button type="submit" class="bg-slate-800 hover:bg-red-500/20 hover:text-red-400 text-slate-300 text-xs font-bold py-3 px-6 rounded-xl border border-white/5 transition-all flex items-center justify-center gap-2 w-full">
                                <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" viewBox="0 0 20 20" fill="currentColor">
                                    <path fill-rule="evenodd" d="M3 3a1 1 0 00-1 1v12a1 1 0 102 0V4a1 1 0 00-1-1zm10.293 9.293a1 1 0 001.414 1.414l3-3a1 1 0 000-1.414l-3-3a1 1 0 10-1.414 1.414L14.586 9H7a1 1 0 100 2h7.586l-1.293 1.293z" clip-rule="evenodd" />
                                </svg>
                                Logout Account
                            </button>
                        </form>
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
