<script lang="ts">
    import { fetchAPI } from "$lib/_core";
    import { onMount } from "svelte";
    import { fade, slide } from "svelte/transition";
    import { tasks, type Task as TaskType } from "$lib";

    let { task }: { task: TaskType } = $props();
    let isSubmit: boolean = $state(false);
    let edit: boolean = $state(false);

    async function patchTask() {
        isSubmit = true;
        const data = await fetchAPI(`/tasks/${task.id}`, "PATCH");
        isSubmit = false;
        if (data.error) return;

        tasks.update((all) =>
            all.map((t) => {
                if (t.id === task.id) {
                    const isDone = !t.isDone;
                    return {
                        ...t,
                        isDone,
                        dateTo: isDone ? new Date().toISOString() : null,
                    };
                }
                return t;
            }),
        );
    }

    async function removeTask() {
        isSubmit = true;
        const data = await fetchAPI(`/tasks/${task.id}`, "DELETE");
        isSubmit = false;
        if (data.error) return;

        tasks.update((all) =>
            all.map((t) => (t.id === task.id ? { ...t, isDeleted: true } : t)),
        );
    }

    const dateAddFormat = $derived(new Date(task.dateAdd).toLocaleString());

    const dateToFormat = $derived.by(() => {
        if (!task.dateTo) return "";
        return new Date(task.dateTo).toLocaleString();
    });

    let nowTime = $state(Date.now());

    $effect(() => {
        if (task.dateFrom && !task.isDone) {
            const interval = setInterval(() => {
                nowTime = Date.now();
            }, 1000);
            return () => clearInterval(interval);
        }
    });

    const timerLabel = $derived.by(() => {
        if (!task.dateFrom || task.isDone) return "";
        const start = new Date(task.dateFrom).getTime();
        const diffMs = nowTime - start;
        const diffSecs = Math.max(0, Math.floor(diffMs / 1000));
        const diffMinutes = Math.floor(diffSecs / 60);
        const totalHours = Math.floor(diffMinutes / 60);

        const days = Math.floor(totalHours / 24);
        const hours = totalHours % 24;
        const minutes = diffMinutes % 60;
        const seconds = diffSecs % 60;

        const pad = (n: number) => String(n).padStart(2, "0");

        if (days > 0) {
            return `${days}d ${pad(hours)}:${pad(minutes)}:${pad(seconds)}`;
        }
        return `${pad(hours)}:${pad(minutes)}:${pad(seconds)}`;
    });

    const diffLabel = $derived.by(() => {
        if (!task.dateTo) return "";
        const dStart = new Date(task.dateFrom || task.dateAdd);
        const dTo = new Date(task.dateTo);
        const diffMs = dTo.getTime() - dStart.getTime();
        const diffSecs = Math.max(0, Math.floor(diffMs / 1000));
        const diffMinutes = Math.floor(diffSecs / 60);
        const totalHours = Math.floor(diffMinutes / 60);

        const days = Math.floor(totalHours / 24);
        const hours = totalHours % 24;
        const minutes = diffMinutes % 60;

        if (days > 0) return `Done in ${days}d ${hours}h ${minutes}m`;
        if (hours > 0) return `Done in ${hours}h ${minutes}m`;
        if (minutes > 0) return `Done in ${minutes}m`;
        return `Done in ${diffSecs}s`;
    });

    function toLocalISO(dateStr: string | null) {
        if (!dateStr) return "";
        const date = new Date(dateStr);
        const offset = date.getTimezoneOffset() * 60000;
        const localDate = new Date(date.getTime() - offset);
        return localDate.toISOString().slice(0, 16);
    }

    async function startTask() {
        isSubmit = true;
        const nowStr = new Date().toISOString();
        const payload = {
            title: task.title,
            dateAdd: task.dateAdd,
            dateFrom: nowStr,
            dateTo: task.dateTo,
        };
        const data = await fetchAPI(`/tasks/${task.id}`, "PUT", payload);
        isSubmit = false;
        if (data.error) return;

        tasks.update((all) =>
            all.map((t) => (t.id === task.id ? { ...t, dateFrom: nowStr } : t)),
        );
    }

    async function saveTask(e: Event) {
        e.preventDefault();
        isSubmit = true;
        const formData = new FormData(e.target as HTMLFormElement);

        // Convert local datetime back to UTC ISO
        const dateAddInput = formData.get("dateAdd") as string;
        const dateFromInput = formData.get("dateFrom") as string;
        const dateToInput = formData.get("dateTo") as string;

        const payload = {
            title: formData.get("title") as string,
            dateAdd: dateAddInput
                ? new Date(dateAddInput).toISOString()
                : task.dateAdd,
            dateFrom: dateFromInput ? new Date(dateFromInput).toISOString() : null,
            dateTo: dateToInput ? new Date(dateToInput).toISOString() : null,
        };

        const data = await fetchAPI(`/tasks/${task.id}`, "PUT", payload);
        isSubmit = false;
        if (data.error) return;

        edit = false;
        tasks.update((all) =>
            all.map((t) =>
                t.id === task.id
                    ? {
                          ...t,
                          ...payload,
                          isDone: !!payload.dateTo,
                      }
                    : t,
            ),
        );
    }

    onMount(() => {
        // No longer needed to call dateFormat manually
    });
</script>

{#if !task.isDeleted}
    <div
        class="glass-card p-6 mb-6 transition-all duration-300 hover:border-indigo-500/30 group animate-in"
        class:opacity-60={task.isDone}
        transition:slide
    >
        {#if !edit}
            <div class="flex flex-col gap-4">
                <div class="flex justify-between items-start gap-4">
                    <div class="space-y-1">
                        <div class="flex items-center gap-2">
                            <h3
                                class="text-xl font-semibold text-slate-100 transition-all"
                                class:line-through={task.isDone}
                                class:text-slate-400={task.isDone}
                            >
                                {task.title}
                            </h3>
                            {#if task.isCommon}
                                <span
                                    class="px-2 py-0.5 rounded text-[9px] font-black uppercase tracking-tighter bg-purple-500/20 text-purple-400 border border-purple-500/30"
                                >
                                    Common
                                </span>
                            {/if}
                        </div>
                        <div
                            class="flex flex-wrap gap-x-4 gap-y-1.5 text-xs font-medium text-slate-500"
                        >
                            <span class="flex items-center gap-1">
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    class="w-3.5 h-3.5"
                                    viewBox="0 0 20 20"
                                    fill="currentColor"
                                >
                                    <path
                                        fill-rule="evenodd"
                                        d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z"
                                        clip-rule="evenodd"
                                    />
                                </svg>
                                Added {dateAddFormat}
                            </span>
                            {#if task.dateFrom && !task.isDone}
                                <span
                                    class="flex items-center gap-1 text-indigo-400 font-semibold animate-pulse"
                                >
                                    <svg
                                        xmlns="http://www.w3.org/2000/svg"
                                        class="w-3.5 h-3.5 animate-spin-slow"
                                        fill="none"
                                        viewBox="0 0 24 24"
                                        stroke="currentColor"
                                        stroke-width="2.5"
                                    >
                                        <circle cx="12" cy="12" r="10" />
                                        <path d="M12 6v6l4 2" />
                                    </svg>
                                    Running: {timerLabel}
                                </span>
                            {/if}
                            {#if task.isDone}
                                <span
                                    class="flex items-center gap-1 text-emerald-500/80"
                                >
                                    <svg
                                        xmlns="http://www.w3.org/2000/svg"
                                        class="w-3"
                                        viewBox="0 0 20 20"
                                        fill="currentColor"
                                    >
                                        <path
                                            fill-rule="evenodd"
                                            d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                                            clip-rule="evenodd"
                                        />
                                    </svg>
                                    Completed {dateToFormat}
                                </span>
                            {/if}
                        </div>
                    </div>
                    {#if task.dateFrom && !task.isDone}
                        <span
                            class="px-3 py-1 rounded-full text-[10px] font-bold uppercase tracking-wider bg-indigo-500/10 text-indigo-400 border border-indigo-500/20 whitespace-nowrap flex items-center gap-1.5"
                        >
                            <span class="w-1.5 h-1.5 rounded-full bg-indigo-400 animate-ping"></span>
                            {timerLabel}
                        </span>
                    {/if}
                    {#if task.isDone}
                        <span
                            class="px-3 py-1 rounded-full text-[10px] font-bold uppercase tracking-wider bg-emerald-500/10 text-emerald-400 border border-emerald-500/20 whitespace-nowrap"
                        >
                            {diffLabel}
                        </span>
                    {/if}
                </div>
            </div>
        {:else}
            <form onsubmit={saveTask} class="space-y-4">
                <div>
                    <label
                        for="edit-title-{task.id}"
                        class="block text-xs font-bold text-slate-500 uppercase mb-2"
                        >Title</label
                    >
                    <input
                        id="edit-title-{task.id}"
                        type="text"
                        name="title"
                        value={task.title}
                        required
                        class="input-field"
                    />
                </div>

                <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
                    <div>
                        <label
                            for="edit-dateAdd-{task.id}"
                            class="block text-xs font-bold text-slate-500 uppercase mb-2"
                            >Creation Date</label
                        >
                        <input
                            id="edit-dateAdd-{task.id}"
                            type="datetime-local"
                            name="dateAdd"
                            value={toLocalISO(task.dateAdd)}
                            required
                            class="input-field"
                        />
                    </div>
                    <div>
                        <label
                            for="edit-dateFrom-{task.id}"
                            class="block text-xs font-bold text-slate-500 uppercase mb-2"
                            >Start Date</label
                        >
                        <input
                            id="edit-dateFrom-{task.id}"
                            type="datetime-local"
                            name="dateFrom"
                            value={toLocalISO(task.dateFrom)}
                            class="input-field"
                        />
                    </div>
                    <div>
                        <label
                            for="edit-dateTo-{task.id}"
                            class="block text-xs font-bold text-slate-500 uppercase mb-2"
                            >Termination Date</label
                        >
                        <input
                            id="edit-dateTo-{task.id}"
                            type="datetime-local"
                            name="dateTo"
                            value={toLocalISO(task.dateTo)}
                            class="input-field"
                        />
                    </div>
                </div>

                <div class="flex gap-2 pt-2">
                    <button
                        type="submit"
                        class="bg-indigo-500 hover:bg-indigo-600 text-white text-xs font-bold py-2.5 px-6 rounded-lg transition-all"
                        disabled={isSubmit}>Save Changes</button
                    >
                    <button
                        type="button"
                        class="bg-slate-700 hover:bg-slate-600 text-white text-xs font-bold py-2.5 px-6 rounded-lg transition-all"
                        onclick={() => (edit = false)}>Cancel</button
                    >
                </div>
            </form>
        {/if}

        <div class="mt-6 pt-6 border-t border-white/5 flex flex-wrap gap-2">
            {#if !task.isDone}
                {#if !task.dateFrom}
                    <button
                        class="bg-indigo-600 hover:bg-indigo-500 text-white text-[11px] font-bold py-2 px-4 rounded-lg border border-indigo-500/30 transition-all flex items-center gap-2 shadow-lg shadow-indigo-500/10"
                        onclick={startTask}
                        disabled={isSubmit}
                    >
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            class="w-3.5 h-3.5"
                            viewBox="0 0 20 20"
                            fill="currentColor"
                        >
                            <path
                                fill-rule="evenodd"
                                d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z"
                                clip-rule="evenodd"
                            />
                        </svg>
                        Start
                    </button>
                {:else}
                    <button
                        class="bg-emerald-600 hover:bg-emerald-500 text-white text-[11px] font-bold py-2 px-4 rounded-lg border border-emerald-500/30 transition-all flex items-center gap-2 shadow-lg shadow-emerald-500/10"
                        onclick={patchTask}
                        disabled={isSubmit}
                    >
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            class="w-3.5 h-3.5"
                            viewBox="0 0 20 20"
                            fill="currentColor"
                        >
                            <path
                                fill-rule="evenodd"
                                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                                clip-rule="evenodd"
                            />
                        </svg>
                        Finish
                    </button>
                {/if}
            {:else}
                <button
                    class="bg-slate-800 hover:bg-slate-700 text-slate-300 text-[11px] font-bold py-2 px-4 rounded-lg border border-white/5 transition-all flex items-center gap-2"
                    onclick={patchTask}
                    disabled={isSubmit}
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="w-3.5 h-3.5"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                            clip-rule="evenodd"
                        />
                    </svg>
                    Undo
                </button>
            {/if}
            <button
                class="bg-slate-800 hover:bg-slate-700 text-slate-300 text-[11px] font-bold py-2 px-4 rounded-lg border border-white/5 transition-all flex items-center gap-2"
                onclick={() => (edit = !edit)}
            >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="w-3.5 h-3.5"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                >
                    <path
                        d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z"
                    />
                </svg>
                Edit
            </button>
            <button
                class="ml-auto bg-slate-800 hover:bg-red-500/20 hover:text-red-400 text-slate-500 text-[11px] font-bold py-2 px-4 rounded-lg border border-white/5 transition-all flex items-center gap-2"
                onclick={removeTask}
                disabled={isSubmit}
            >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="w-3.5 h-3.5"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                >
                    <path
                        fill-rule="evenodd"
                        d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"
                        clip-rule="evenodd"
                    />
                </svg>
                Delete
            </button>
        </div>
    </div>
{/if}

<style>
    @keyframes spin-slow {
        from { transform: rotate(0deg); }
        to { transform: rotate(360deg); }
    }
    :global(.animate-spin-slow) {
        animation: spin-slow 10s linear infinite;
    }
</style>
