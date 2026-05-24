<script lang="ts">
    import { onMount } from "svelte";
    import { dashboardStats } from "$lib";
    import { fetchAPI } from "$lib/_core";
    import Skeleton from "./Skeleton.svelte";
    import { fade } from "svelte/transition";
    import Note from "./Note.svelte";

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
        <h2 class="text-3xl font-extrabold text-white mb-2">
            Performance Dashboard
        </h2>
        <p class="text-slate-400">
            Your productivity insights for the last 30 days.
        </p>
    </header>

    {#if isFetching}
        <Skeleton variant="dashboard" />
    {:else if stats}
        <div
            class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-12"
            transition:fade={{ duration: 500 }}
        >
            <!-- Stats Cards -->
            <div class="glass-card p-8 flex flex-col items-center text-center">
                <div
                    class="w-12 h-12 bg-indigo-500/10 rounded-xl flex items-center justify-center text-indigo-400 mb-4"
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="w-6 h-6"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-11a1 1 0 10-2 0v2H7a1 1 0 100 2h2v2a1 1 0 102 0v-2h2a1 1 0 100-2h-2V7z"
                            clip-rule="evenodd"
                        />
                    </svg>
                </div>
                <span class="text-4xl font-black text-white mb-1"
                    >{stats.tasksAdded}</span
                >
                <span
                    class="text-xs font-bold text-slate-500 uppercase tracking-widest"
                    >Tasks Created</span
                >
            </div>

            <div
                class="glass-card p-8 flex flex-col items-center text-center border-emerald-500/10"
            >
                <div
                    class="w-12 h-12 bg-emerald-500/10 rounded-xl flex items-center justify-center text-emerald-400 mb-4"
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="w-6 h-6"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                            clip-rule="evenodd"
                        />
                    </svg>
                </div>
                <span class="text-4xl font-black text-white mb-1"
                    >{stats.tasksDone}</span
                >
                <span
                    class="text-xs font-bold text-slate-500 uppercase tracking-widest"
                    >Tasks Completed</span
                >
            </div>

            <div
                class="glass-card p-8 flex flex-col items-center text-center border-amber-500/10"
            >
                <div
                    class="w-12 h-12 bg-amber-500/10 rounded-xl flex items-center justify-center text-amber-400 mb-4"
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="w-6 h-6"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z"
                            clip-rule="evenodd"
                        />
                    </svg>
                </div>
                <span class="text-4xl font-black text-white mb-1"
                    >{Math.round(stats.avgDuration)}m</span
                >
                <span
                    class="text-xs font-bold text-slate-500 uppercase tracking-widest"
                    >Avg. Completion Time</span
                >
            </div>
        </div>

        <div class="glass-card p-8 mb-12">
            <h3 class="text-xl font-bold text-white mb-6">
                Productivity Summary
            </h3>
            <div class="space-y-4">
                <div
                    class="flex justify-between items-center p-4 bg-white/5 rounded-lg border border-white/5"
                >
                    <span class="text-slate-300 font-medium"
                        >Completion Rate</span
                    >
                    <span class="text-emerald-400 font-bold">
                        {stats.tasksAdded > 0
                            ? Math.round(
                                  (stats.tasksDone / stats.tasksAdded) * 100,
                              )
                            : 0}%
                    </span>
                </div>
                <div
                    class="flex justify-between items-center p-4 bg-white/5 rounded-lg border border-white/5"
                >
                    <span class="text-slate-300 font-medium"
                        >Efficiency Score</span
                    >
                    <span class="text-indigo-400 font-bold">
                        {stats.avgDuration > 0
                            ? Math.round(100 - stats.avgDuration / 60)
                            : 100}
                    </span>
                </div>
            </div>
        </div>

        <div class="glass-card p-8">
            <h3
                class="text-xl font-bold text-white mb-6 flex items-center gap-2"
            >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-5 w-5 text-indigo-400"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                >
                    <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
                    />
                </svg>
                Last Note Added
            </h3>

            {#if Object.entries(stats.lastNote).length}
                <div
                    class="relative overflow-hidden bg-gradient-to-br from-indigo-500/10 to-purple-500/10 rounded-2xl border border-indigo-500/20 p-6 group"
                >
                    <div
                        class="absolute inset-0 bg-gradient-to-br from-indigo-500/20 to-purple-500/20 opacity-0 group-hover:opacity-100 transition-opacity duration-500 pointer-events-none"
                    ></div>

                    <h4
                        class="text-xl font-bold text-white mb-3 relative z-10 break-words"
                    >
                        {stats.lastNote.title}
                    </h4>
                    <p
                        class="text-slate-300 whitespace-pre-wrap leading-relaxed relative z-10 break-words"
                    >
                        {stats.lastNote.content}
                    </p>

                    <div
                        class="mt-5 pt-4 border-t border-indigo-500/20 flex justify-between items-center relative z-10"
                    >
                        <div class="flex flex-col gap-0.5">
                            <span class="text-xs text-slate-500 font-medium">
                                Added on {new Date(
                                    stats.lastNote.dateAdd.replace(" ", "T") +
                                        "Z",
                                ).toLocaleDateString(undefined, {
                                    month: "short",
                                    day: "numeric",
                                    year: "numeric",
                                    hour: "2-digit",
                                    minute: "2-digit",
                                })}
                            </span>
                            {#if stats.lastNote.dateUpdate}
                                <span
                                    class="text-xs text-slate-500/80 font-medium"
                                >
                                    Updated on {new Date(
                                        stats.lastNote.dateUpdate.replace(
                                            " ",
                                            "T",
                                        ) + "Z",
                                    ).toLocaleDateString(undefined, {
                                        month: "short",
                                        day: "numeric",
                                        year: "numeric",
                                        hour: "2-digit",
                                        minute: "2-digit",
                                    })}
                                </span>
                            {/if}
                        </div>
                    </div>
                </div>
            {:else}
                <div
                    class="p-6 bg-white/5 rounded-2xl border border-white/5 text-center"
                >
                    <p class="text-slate-500 italic">No notes added.</p>
                </div>
            {/if}
        </div>
    {/if}
</div>
