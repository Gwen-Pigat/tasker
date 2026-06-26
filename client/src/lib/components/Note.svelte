<script lang="ts">
    import { notes, type Note as NoteType } from "$lib";
    import { fetchAPI } from "$lib/_core";

    let { note }: { note: NoteType } = $props();

    let isDeleting = $state(false);
    let errorMsg = $state("");

    async function handleDelete() {
        if (!confirm("Are you sure you want to delete this note?")) return;

        isDeleting = true;
        const response = await fetchAPI(`/notes/${note.id}`, "DELETE");
        isDeleting = false;

        if (response.error) {
            errorMsg = response.error;
            return;
        }

        notes.update((currentNotes) => {
            return currentNotes.filter((n) => n.id !== note.id);
        });
    }

    let isEditing = $state(false);
    let isSaving = $state(false);
    let editTitle = $state("");
    let editContent = $state("");

    function startEditing() {
        editTitle = note.title;
        editContent = note.content;
        isEditing = true;
    }

    async function handleSave() {
        if (!editTitle.trim() || !editContent.trim()) {
            errorMsg = "Title and content are required";
            return;
        }

        isSaving = true;
        errorMsg = "";

        const response = await fetchAPI(`/notes/${note.id}`, "PUT", {
            title: editTitle.trim(),
            content: editContent.trim(),
        });

        isSaving = false;

        if (response.error) {
            errorMsg = response.error;
            return;
        }

        notes.update((currentNotes) => {
            return currentNotes.map((n) => (n.id === note.id ? response : n));
        });

        isEditing = false;
    }
</script>

<div
    class="glass-card p-6 flex flex-col h-full group relative overflow-hidden transition-all duration-300 hover:shadow-lg hover:shadow-indigo-500/10"
>
    <div
        class="absolute inset-0 bg-gradient-to-br from-indigo-500/5 to-purple-500/5 opacity-0 group-hover:opacity-100 transition-opacity duration-500 pointer-events-none"
    ></div>

    {#if errorMsg}
        <div
            class="bg-red-500/10 border border-red-500/20 text-red-400 p-2 mb-4 rounded-lg text-xs"
        >
            {errorMsg}
        </div>
    {/if}

    <div class="flex flex-col h-full relative z-10">
        {#if isEditing}
            <div class="flex flex-col gap-3 flex-grow mb-4">
                <input
                    id="note_title"
                    type="text"
                    name="title"
                    class="w-full h-10 px-4 py-2 border border-white/10 bg-white/5 text-white placeholder-slate-500 focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none transition-all rounded-lg font-bold"
                    bind:value={editTitle}
                    disabled={isSaving}
                />
                <textarea
                    id="note_content"
                    name="content"
                    class="w-full flex-grow px-4 py-2 border border-white/10 bg-white/5 text-slate-300 placeholder-slate-500 focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none transition-all rounded-lg resize-none text-sm min-h-[120px]"
                    bind:value={editContent}
                    disabled={isSaving}
                ></textarea>
                <div class="flex justify-end gap-2">
                    <button
                        onclick={() => {
                            isEditing = false;
                            errorMsg = "";
                        }}
                        disabled={isSaving}
                        class="px-3 py-1.5 rounded-lg text-sm font-medium text-slate-400 hover:text-white hover:bg-white/5 transition-colors disabled:opacity-50"
                    >
                        Cancel
                    </button>
                    <button
                        onclick={handleSave}
                        disabled={isSaving ||
                            !editTitle.trim() ||
                            !editContent.trim()}
                        class="px-3 py-1.5 rounded-lg text-sm font-medium bg-indigo-500 text-white hover:bg-indigo-600 transition-colors disabled:opacity-50 disabled:bg-slate-800 flex items-center gap-2"
                        title="Save note"
                    >
                        {#if isSaving}
                            <svg
                                class="animate-spin h-3.5 w-3.5"
                                xmlns="http://www.w3.org/2000/svg"
                                fill="none"
                                viewBox="0 0 24 24"
                            >
                                <circle
                                    class="opacity-25"
                                    cx="12"
                                    cy="12"
                                    r="10"
                                    stroke="currentColor"
                                    stroke-width="4"
                                ></circle>
                                <path
                                    class="opacity-75"
                                    fill="currentColor"
                                    d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                                ></path>
                            </svg>
                            Saving...
                        {:else}
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                class="h-4 w-4"
                                viewBox="0 0 20 20"
                                fill="currentColor"
                            >
                                <path
                                    fill-rule="evenodd"
                                    d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z"
                                    clip-rule="evenodd"
                                />
                            </svg>
                            Save
                        {/if}
                    </button>
                </div>
            </div>
        {:else}
            <div class="flex justify-between items-start gap-3 mb-2">
                <h3
                    class="font-bold text-lg text-white break-words overflow-hidden text-ellipsis line-clamp-2"
                >
                    {note.title}
                </h3>
            </div>

            <p
                class="text-slate-400 text-sm whitespace-pre-wrap break-words flex-grow overflow-hidden text-ellipsis line-clamp-6 mb-4"
            >
                {note.content}
            </p>
        {/if}

        <div
            class="flex justify-between items-center mt-auto pt-4 border-t border-white/5"
        >
            <div class="flex flex-col gap-0.5">
                <span class="text-xs text-slate-500 font-medium">
                    Added on {new Date(
                        note.dateAdd.replace(" ", "T") + "Z",
                    ).toLocaleDateString(undefined, {
                        month: "short",
                        day: "numeric",
                        year: "numeric",
                        hour: "2-digit",
                        minute: "2-digit",
                    })}
                </span>
                {#if note.dateUpdate}
                    <span class="text-xs text-slate-500/80 font-medium">
                        Updated on {new Date(
                            note.dateUpdate.replace(" ", "T") + "Z",
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

            <div
                class="flex items-center gap-1 opacity-100 md:opacity-0 md:group-hover:opacity-100 transition-opacity"
            >
                {#if !isEditing}
                    <button
                        onclick={startEditing}
                        class="p-1.5 rounded-md text-slate-400 hover:text-indigo-400 hover:bg-indigo-400/10 transition-colors"
                        title="Edit note"
                    >
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            class="h-4 w-4"
                            viewBox="0 0 20 20"
                            fill="currentColor"
                        >
                            <path
                                d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z"
                            />
                        </svg>
                    </button>
                    <button
                        onclick={handleDelete}
                        disabled={isDeleting}
                        class="p-1.5 rounded-md text-slate-400 hover:text-red-400 hover:bg-red-400/10 transition-colors disabled:opacity-50"
                        title="Delete note"
                    >
                        {#if isDeleting}
                            <svg
                                class="animate-spin h-4 w-4 text-red-400"
                                xmlns="http://www.w3.org/2000/svg"
                                fill="none"
                                viewBox="0 0 24 24"
                            >
                                <circle
                                    class="opacity-25"
                                    cx="12"
                                    cy="12"
                                    r="10"
                                    stroke="currentColor"
                                    stroke-width="4"
                                ></circle>
                                <path
                                    class="opacity-75"
                                    fill="currentColor"
                                    d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                                ></path>
                            </svg>
                        {:else}
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                class="h-4 w-4"
                                viewBox="0 0 20 20"
                                fill="currentColor"
                            >
                                <path
                                    fill-rule="evenodd"
                                    d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"
                                    clip-rule="evenodd"
                                />
                            </svg>
                        {/if}
                    </button>
                {/if}
            </div>
        </div>
    </div>
</div>
