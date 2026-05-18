<script lang="ts">
    import { onMount } from "svelte";
    import { dashboardStats } from "$lib";
    import { fetchAPI } from "$lib/_core";
    import Loader from "./Loader.svelte";

    let isFetching = $state(true);

    async function loadStats() {
        isFetching = true;
        const response = await fetchAPI("/dashboard/stats", "GET");
        isFetching = false;
        if (!response.error) {
            dashboardStats.set(response);
        }
    }

    onMount(() => {
        loadStats();
    });

    const stats = $derived($dashboardStats);
</script>

<div class="max-w-4xl mx-auto px-4 py-8 animate-in">
    <header class="mb-12">
        <h2 class="text-3xl font-extrabold text-white mb-2">Performance Dashboard</h2>
        <p class="text-slate-400">Your productivity insights for the last 30 days.</p>
    </header>

    {#if isFetching}
        <div class="flex justify-center py-20">
            <Loader />
        </div>
    {:else if stats}
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-12">
            <!-- Stats Cards -->
            <div class="glass-card p-8 flex flex-col items-center text-center">
                <div class="w-12 h-12 bg-indigo-500/10 rounded-xl flex items-center justify-center text-indigo-400 mb-4">
                    <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-11a1 1 0 10-2 0v2H7a1 1 0 100 2h2v2a1 1 0 102 0v-2h2a1 1 0 100-2h-2V7z" clip-rule="evenodd" />
                    </svg>
                </div>
                <span class="text-4xl font-black text-white mb-1">{stats.tasksAdded}</span>
                <span class="text-xs font-bold text-slate-500 uppercase tracking-widest">Tasks Created</span>
            </div>

            <div class="glass-card p-8 flex flex-col items-center text-center border-emerald-500/10">
                <div class="w-12 h-12 bg-emerald-500/10 rounded-xl flex items-center justify-center text-emerald-400 mb-4">
                    <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                    </svg>
                </div>
                <span class="text-4xl font-black text-white mb-1">{stats.tasksDone}</span>
                <span class="text-xs font-bold text-slate-500 uppercase tracking-widest">Tasks Completed</span>
            </div>

            <div class="glass-card p-8 flex flex-col items-center text-center border-amber-500/10">
                <div class="w-12 h-12 bg-amber-500/10 rounded-xl flex items-center justify-center text-amber-400 mb-4">
                    <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd" />
                    </svg>
                </div>
                <span class="text-4xl font-black text-white mb-1">{Math.round(stats.avgDuration)}m</span>
                <span class="text-xs font-bold text-slate-500 uppercase tracking-widest">Avg. Completion Time</span>
            </div>
        </div>

        <div class="glass-card p-8">
            <h3 class="text-xl font-bold text-white mb-6">Productivity Summary</h3>
            <div class="space-y-4">
                <div class="flex justify-between items-center p-4 bg-white/5 rounded-lg border border-white/5">
                    <span class="text-slate-300 font-medium">Completion Rate</span>
                    <span class="text-emerald-400 font-bold">
                        {stats.tasksAdded > 0 ? Math.round((stats.tasksDone / stats.tasksAdded) * 100) : 0}%
                    </span>
                </div>
                <div class="flex justify-between items-center p-4 bg-white/5 rounded-lg border border-white/5">
                    <span class="text-slate-300 font-medium">Efficiency Score</span>
                    <span class="text-indigo-400 font-bold">
                        {stats.avgDuration > 0 ? Math.round(100 - (stats.avgDuration / 60)) : 100}
                    </span>
                </div>
            </div>
        </div>
    {/if}
</div>
