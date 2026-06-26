<script lang="ts">
    import { user } from "$lib";

    let { currentTab = $bindable() }: { currentTab: "dashboard" | "tasks" | "notes" } =
        $props();

    function handleLogout() {
        localStorage.removeItem("user");
        user.set(null);
    }
</script>

<svelte:head>
    {#if currentTab === "dashboard"}
        <title>Tasker | Dashboard</title>
    {:else if currentTab === "tasks"}
        <title>Tasker | Tasks list</title>
    {:else if currentTab === "notes"}
        <title>Tasker | Notes</title>
    {/if}
</svelte:head>

<nav
    class="sticky top-0 z-50 bg-slate-900/80 backdrop-blur-xl border-b border-white/5"
>
    <div class="max-w-6xl mx-auto px-4">
        <div class="flex items-center justify-between h-16 gap-4">
            <div class="flex items-center gap-3 sm:gap-8 min-w-0">
                <img
                    src="/images/logo.svg"
                    alt="Logo"
                    class="w-20 sm:w-24 flex-shrink-0"
                />
                <div class="flex gap-1">
                    <button
                        onclick={() => (currentTab = "dashboard")}
                        class="px-2.5 sm:px-4 py-1.5 rounded-lg text-xs sm:text-sm font-bold transition-all {currentTab ===
                        'dashboard'
                            ? 'bg-indigo-500 text-white shadow-lg shadow-indigo-500/20'
                            : 'text-slate-400 hover:bg-white/5'}"
                    >
                        Dashboard
                    </button>
                    <button
                        onclick={() => (currentTab = "tasks")}
                        class="px-2.5 sm:px-4 py-1.5 rounded-lg text-xs sm:text-sm font-bold transition-all {currentTab ===
                        'tasks'
                            ? 'bg-indigo-500 text-white shadow-lg shadow-indigo-500/20'
                            : 'text-slate-400 hover:bg-white/5'}"
                    >
                        Tasks
                    </button>
                    <button
                        onclick={() => (currentTab = "notes")}
                        class="px-2.5 sm:px-4 py-1.5 rounded-lg text-xs sm:text-sm font-bold transition-all {currentTab ===
                        'notes'
                            ? 'bg-indigo-500 text-white shadow-lg shadow-indigo-500/20'
                            : 'text-slate-400 hover:bg-white/5'}"
                    >
                        Notes
                    </button>
                </div>
            </div>
            <div class="flex-shrink-0">
                <button
                    onclick={handleLogout}
                    class="bg-slate-800 hover:bg-red-500/20 hover:text-red-400 text-slate-300 text-xs font-bold p-2.5 sm:py-2.5 sm:px-4 rounded-lg sm:rounded-xl border border-white/5 transition-all flex items-center justify-center gap-2"
                    title="Logout Account"
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="w-4 h-4 flex-shrink-0"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M3 3a1 1 0 00-1 1v12a1 1 0 102 0V4a1 1 0 00-1-1zm10.293 9.293a1 1 0 001.414 1.414l3-3a1 1 0 000-1.414l-3-3a1 1 0 10-1.414 1.414L14.586 9H7a1 1 0 100 2h7.586l-1.293 1.293z"
                            clip-rule="evenodd"
                        />
                    </svg>
                    <span class="hidden sm:inline">Logout</span>
                </button>
            </div>
        </div>
    </div>
</nav>
