<script lang="ts">
    import "../app.css";
    import { fade } from "svelte/transition";
    import { error, url } from "$lib";

    let { children, data } = $props();

    url.set(data.apiURL);

    $effect(() => {
        if ($error !== "") {
            setTimeout(function () {
                error.set("");
            }, 6000);
        }
    });
</script>

{#if $error !== ""}
    <div class="fixed top-4 right-4 bg-red-500 text-white px-6 py-3 rounded-xl shadow-2xl z-[10000] flex items-center gap-2" transition:fade={{ duration: 500 }}>
        <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
        </svg>
        {$error}
    </div>
{/if}
{@render children?.()}

<style>
</style>
