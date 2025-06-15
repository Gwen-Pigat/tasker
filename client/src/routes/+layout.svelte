<script>
    import { fade } from 'svelte/transition';
    import { error } from '$lib'

    let { children } = $props()

    $effect(() => {
        if($error !== ""){
            setTimeout(function(){
                error.set("")
            },6000)
        }
    })
    
</script>

{#if $error !== ""}
    <div class="error" transition:fade={{duration:500}}>{$error}</div>
{/if}
{@render children?.()}

<style>
    .error{
        position: absolute;
        right: 0;
        top: 0;
        font-size: 16px;
        padding: 0.5rem;
        border-radius: 0;
        color: var(--white);
        background-color: var(--red);
        z-index: 10000;
    }
</style>