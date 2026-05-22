<script lang="ts">
    import { onMount } from "svelte";
    import { commonTasks, tasks } from "$lib";
    import { fetchAPI } from "$lib/_core";
    import Skeleton from "./Skeleton.svelte";

    let isFetching = $state(true);
    let newTitle = $state("");
    let isSubmitting = $state(false);

    async function loadCommonTasks() {
        isFetching = true;
        const response = await fetchAPI("/common-tasks", "GET");
        isFetching = false;
        if (!response.error) {
            commonTasks.set(response);
        }
    }

    async function addCommonTask(e: Event) {
        e.preventDefault();
        if (!newTitle) return;
        isSubmitting = true;
        const response = await fetchAPI("/common-tasks", "POST", { title: newTitle });
        isSubmitting = false;
        if (!response.error) {
            commonTasks.set(response);
            newTitle = "";
        }
    }

    async function validateCommonTask(id: number) {
        const response = await fetchAPI(`/common-tasks/${id}/validate`, "POST");
        if (!response.error) {
            // Refresh today's tasks
            const tasksResponse = await fetchAPI("/tasks?date=" + new Date().toISOString().split('T')[0], "GET");
            if (!tasksResponse.error) {
                tasks.set(tasksResponse);
            }
        }
    }

    async function deleteCommonTask(id: number) {
        const response = await fetchAPI(`/common-tasks/${id}`, "DELETE");
        if (!response.error) {
            commonTasks.set(response);
        }
    }

    onMount(() => {
        loadCommonTasks();
    });
</script>

<div class="glass-card p-6">
    <div class="flex items-center gap-3 mb-6">
        <div class="w-10 h-10 bg-purple-500/10 rounded-lg flex items-center justify-center text-purple-400">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 20 20" fill="currentColor">
                <path d="M7 3a1 1 0 000 2h6a1 1 0 100-2H7zM4 7a1 1 0 011-1h10a1 1 0 110 2H5a1 1 0 01-1-1zM2 11a2 2 0 012-2h12a2 2 0 012 2v4a2 2 0 01-2 2H4a2 2 0 01-2-2v-4z" />
            </svg>
        </div>
        <h3 class="text-xl font-bold text-white">Common Tasks</h3>
    </div>

    <form onsubmit={addCommonTask} class="flex gap-2 mb-6">
        <input 
            type="text" 
            bind:value={newTitle} 
            placeholder="New common task..." 
            class="input-field flex-grow text-xs" 
            required 
        />
        <button
            type="submit"
            disabled={isSubmitting || !newTitle}
            class="bg-purple-500 hover:bg-purple-600 text-white p-2.5 rounded-lg transition-all disabled:opacity-50"
            aria-label="Add common task"
            title="Add common task"
        >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z" clip-rule="evenodd" />
            </svg>
        </button>
    </form>

    {#if isFetching}
        <Skeleton variant="commonTasks" />
    {:else}
        <div class="space-y-3">
            {#each $commonTasks as item (item.id)}
                <div class="flex items-center justify-between p-3 bg-white/5 rounded-lg border border-white/5 group hover:border-purple-500/30 transition-all">
                    <span class="text-slate-300 text-sm font-medium">{item.title}</span>
                    <div class="flex items-center gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
                        <button 
                            onclick={() => validateCommonTask(item.id)}
                            class="p-1.5 text-emerald-400 hover:bg-emerald-500/10 rounded-md transition-colors"
                            title="Add to today"
                        >
                            <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" viewBox="0 0 20 20" fill="currentColor">
                                <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                            </svg>
                        </button>
                        <button 
                            onclick={() => deleteCommonTask(item.id)}
                            class="p-1.5 text-red-400 hover:bg-red-500/10 rounded-md transition-colors"
                            title="Delete"
                        >
                            <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" viewBox="0 0 20 20" fill="currentColor">
                                <path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" />
                            </svg>
                        </button>
                    </div>
                </div>
            {:else}
                <p class="text-center text-slate-500 text-xs py-4 italic">No common tasks set.</p>
            {/each}
        </div>
    {/if}
</div>
