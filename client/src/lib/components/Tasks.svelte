<script lang="ts">
    import { tasks, user, type Task as TaskType } from "$lib";
    import Task from "./Task.svelte";
    import TaskAdd from "./TaskAdd.svelte";
    import { fetchAPI } from "$lib/_core";
    import { onMount } from "svelte";
    let isFetching: boolean = $state(true);

    async function loadTasks() {
        tasks.set([]);
        isFetching = true;
        const dataFetch = await fetchAPI("/tasks", "GET");
        isFetching = false;
        if (dataFetch.error) {
            return;
        }
        tasks.set(dataFetch);
    }

    const activeTasks = $derived($tasks.filter((t: TaskType) => !t.isDeleted));

    onMount(() => {
        loadTasks();
    });
</script>

<div class="max-w-4xl mx-auto px-4 py-12 animate-in">
    <div
        class="flex flex-col sm:flex-row justify-between items-center gap-6 mb-12"
    >
        <div class="space-y-1 text-center sm:text-left">
            <h1 class="text-4xl font-extrabold tracking-tight text-white">
                Hello, <span class="text-indigo-400">{$user?.username}</span>
            </h1>
            <p class="text-slate-400 font-medium">Ready to conquer your day?</p>
        </div>
        <form method="POST" action="?/logout">
            <button
                type="submit"
                class="bg-slate-800 hover:bg-slate-700 text-slate-300 text-xs font-bold py-2.5 px-6 rounded-xl border border-white/5 transition-all flex items-center gap-2"
            >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="w-4 h-4"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                >
                    <path
                        fill-rule="evenodd"
                        d="M3 3a1 1 0 00-1 1v12a1 1 0 102 0V4a1 1 0 00-1-1zm10.293 9.293a1 1 0 001.414 1.414l3-3a1 1 0 000-1.414l-3-3a1 1 0 10-1.414 1.414L14.586 9H7a1 1 0 100 2h7.586l-1.293 1.293z"
                        clip-rule="evenodd"
                    />
                </svg>
                Logout
            </button>
        </form>
    </div>

    <TaskAdd />

    <div class="space-y-4">
        {#if activeTasks.length > 0}
            {#each activeTasks as task (task.id)}
                <Task {task} />
            {/each}
        {:else if isFetching}
            <div class="flex justify-center py-20">
                <div
                    class="w-12 h-12 border-4 border-indigo-500/20 border-t-indigo-500 rounded-full animate-spin"
                ></div>
            </div>
        {:else}
            <div class="glass-card p-12 text-center animate-in">
                <div
                    class="w-20 h-20 bg-indigo-500/10 rounded-full flex items-center justify-center mx-auto mb-6 text-indigo-400"
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="w-10 h-10"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                    >
                        <path d="M9 2a1 1 0 000 2h2a1 1 0 100-2H9z" />
                        <path
                            fill-rule="evenodd"
                            d="M4 5a2 2 0 012-2 3 3 0 003 3h2a3 3 0 003-3 2 2 0 012 2v11a2 2 0 01-2 2H6a2 2 0 01-2-2V5zm3 4a1 1 0 000 2h.01a1 1 0 100-2H7zm3 0a1 1 0 000 2h3a1 1 0 100-2h-3zm-3 4a1 1 0 100 2h.01a1 1 0 100-2H7zm3 0a1 1 0 100 2h3a1 1 0 100-2h-3z"
                            clip-rule="evenodd"
                        />
                    </svg>
                </div>
                <h3 class="text-xl font-bold text-white mb-2">No tasks yet</h3>
                <p class="text-slate-400 max-w-xs mx-auto">
                    Start by adding your first task above. Success favors the
                    organized!
                </p>
            </div>
        {/if}
    </div>
</div>

<style>
</style>
