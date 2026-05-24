<script lang="ts">
    import { fetchAPI } from "$lib/_core";
    import { notes } from "$lib";

    let isCreating = $state(false);
    let title = $state("");
    let content = $state("");
    let errorMsg = $state("");

    async function handleAddNote(e: Event) {
        e.preventDefault();

        if (!title.trim() || !content.trim()) {
            errorMsg = "Title and content are required";
            return;
        }

        isCreating = true;
        errorMsg = "";

        const payload = {
            title: title.trim(),
            content: content.trim(),
        };

        const response = await fetchAPI("/notes", "POST", payload);

        isCreating = false;

        if (response.error) {
            errorMsg = response.error;
            return;
        }

        notes.update((currentNotes) => {
            return [response, ...currentNotes];
        });

        title = "";
        content = "";
    }
</script>

<div class="glass-card p-6 mb-8 relative overflow-hidden group">
    <div
        class="absolute inset-0 bg-gradient-to-br from-indigo-500/5 to-purple-500/5 opacity-0 group-hover:opacity-100 transition-opacity duration-500"
    ></div>

    <form onsubmit={handleAddNote} class="relative flex flex-col gap-4">
        {#if errorMsg}
            <div
                class="bg-red-500/10 border border-red-500/20 text-red-400 p-3 rounded-lg text-sm"
            >
                {errorMsg}
            </div>
        {/if}

        <input
            type="text"
            bind:value={title}
            placeholder="Note title..."
            class="w-full bg-slate-900/50 border border-white/5 rounded-xl px-4 py-3 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/50 focus:border-indigo-500/50 transition-all text-lg font-semibold"
            disabled={isCreating}
        />

        <textarea
            bind:value={content}
            placeholder="Write your note here..."
            rows="4"
            class="w-full bg-slate-900/50 border border-white/5 rounded-xl px-4 py-3 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/50 focus:border-indigo-500/50 transition-all resize-none"
            disabled={isCreating}
        ></textarea>

        <div class="flex justify-end">
            <button
                type="submit"
                disabled={isCreating || !title.trim() || !content.trim()}
                class="bg-indigo-500 hover:bg-indigo-600 disabled:bg-slate-800 disabled:text-slate-500 disabled:cursor-not-allowed text-white font-bold py-2.5 px-6 rounded-xl transition-all shadow-lg shadow-indigo-500/25 disabled:shadow-none flex items-center gap-2"
            >
                {#if isCreating}
                    <svg
                        class="animate-spin h-5 w-5"
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
                        class="h-5 w-5"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z"
                            clip-rule="evenodd"
                        />
                    </svg>
                    Save Note
                {/if}
            </button>
        </div>
    </form>
</div>
