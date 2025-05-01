<script>

    import { fade } from 'svelte/transition';
    import { error, user } from '$lib'

    $effect(() => {
        if($error !== ""){
            setTimeout(function(){
                error.set("")
            },6000)
        }
    })

    let data = $props()
    console.log(data)
    if(data.form && !data.form.error){
        user.set(data.form) 
    }
    
</script>

{#if $error !== ""}
    <div class="error" transition:fade={{duration:500}}>{$error}</div>
{/if}

<slot />


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