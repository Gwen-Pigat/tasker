<script lang="ts">
    import { notes, user, type Note as NoteType } from "$lib";
    import Note from "./Note.svelte";
    import NoteAdd from "./NoteAdd.svelte";
    import { fetchAPI } from "$lib/_core";
    import { onMount } from "svelte";
    import Skeleton from "./Skeleton.svelte";

    let isFetching: boolean = $state(true);

    async function loadNotes() {
        notes.set([]);
        isFetching = true;
        const dataFetch = await fetchAPI("/notes", "GET");
        isFetching = false;
        if (dataFetch.error) {
            return;
        }
        notes.set(dataFetch);
    }

    onMount(() => {
        loadNotes();
    });
</script>

<div class="max-w-6xl mx-auto px-4 py-12 animate-in">
    <div class="flex flex-col gap-8">
        <div
            class="flex flex-col sm:flex-row justify-between items-center gap-6 mb-4"
        >
            <div class="space-y-1 text-center sm:text-left">
                <h1 class="text-4xl font-extrabold tracking-tight text-white">
                    Your <span class="text-indigo-400">Notes</span>
                </h1>
                <p class="text-slate-400 font-medium">
                    Jot down your thoughts and ideas.
                </p>
            </div>
        </div>

        <NoteAdd />

        <div class="space-y-4">
            {#if $notes.length > 0}
                <div
                    class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6"
                >
                    {#each $notes as note (note.id)}
                        <Note {note} />
                    {/each}
                </div>
            {:else if isFetching}
                <Skeleton variant="notes" />
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
                            <path
                                fill-rule="evenodd"
                                d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4zm2 6a1 1 0 011-1h6a1 1 0 110 2H7a1 1 0 01-1-1zm1 3a1 1 0 100 2h6a1 1 0 100-2H7z"
                                clip-rule="evenodd"
                            />
                        </svg>
                    </div>
                    <h3 class="text-xl font-bold text-white mb-2">
                        No notes yet
                    </h3>
                    <p class="text-slate-400 max-w-xs mx-auto">
                        Start by creating your first note above.
                    </p>
                </div>
            {/if}
        </div>
    </div>
</div>
