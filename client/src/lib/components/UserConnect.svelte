<script lang="ts">
    import { onMount } from "svelte";

    interface Props {
        title: string;
        view: "login" | "register";
        username: string;
        password: string;
    }
    let { title, view = $bindable(), username, password }: Props = $props()

    let input:undefined|HTMLInputElement = $state()

    $effect(() => {
        if(view === "login"){
            title = "Connect"
        } else if(view === "register"){
            title = "Register"
        }
        triggerFocus()
    })

    function triggerFocus()
    {
        input?.focus()
    }

    onMount(() => {
        triggerFocus()
    })

</script>

<div class="glass-card p-8 w-full max-w-md mx-auto animate-in">
    <h1 class="text-3xl font-bold text-center mb-8 bg-gradient-to-r from-indigo-400 to-purple-400 bg-clip-text text-transparent">
        {title}
    </h1>
    
    <form id="setConnect" method="POST" action="?/login" class="space-y-6">
        <div>
            <label for="username" class="block text-sm font-medium text-slate-400 mb-2">Username</label>
            <input id="username" type="text" placeholder="Enter your username" name="username" bind:value={username} bind:this={input} required class="input-field" />
        </div>
        
        <div>
            <label for="password" class="block text-sm font-medium text-slate-400 mb-2">Password</label>
            <input id="password" type="password" placeholder="Enter your password" name="password" bind:value={password} required class="input-field" />
        </div>

        <button type="submit" disabled={username === "" || password === ""} class="btn-primary w-full mt-4">
            Connect to Tasker
        </button>

        <div class="mt-8 text-center">
            <p class="text-slate-400 text-sm">
                Don't have an account? 
                <button type="button" class="text-indigo-400 hover:text-indigo-300 font-semibold underline transition-colors ml-1" onclick={() => view = "register"}>
                    Register here
                </button>
            </p>
        </div>
    </form>
</div>

<style>
</style>