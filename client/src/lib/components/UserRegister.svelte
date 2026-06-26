<script lang="ts">
    import { onMount } from "svelte";
    import { fetchAPI } from "$lib/_core";
    import { user, error } from "$lib";

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

    async function handleRegister(e: SubmitEvent) {
        e.preventDefault();
        error.set("");
        const formData = new FormData();
        formData.append("username", username);
        formData.append("password", password);

        const result = await fetchAPI("/user", "POST", formData);
        if (result.error) {
            error.set(result.error);
            return;
        }

        localStorage.setItem("user", JSON.stringify(result));
        user.set(result);
    }

</script>

<div class="glass-card p-8 w-full max-w-md mx-auto animate-in">
    <h1 class="text-3xl font-bold text-center mb-8 bg-gradient-to-r from-indigo-400 to-purple-400 bg-clip-text text-transparent">
        {title}
    </h1>
    
    <form id="setRegister" onsubmit={handleRegister} class="space-y-6">
        <div>
            <label for="reg-username" class="block text-sm font-medium text-slate-400 mb-2">Username</label>
            <input id="reg-username" type="text" placeholder="Choose a username" name="username" bind:value={username} bind:this={input} required class="input-field" />
        </div>
        
        <div>
            <label for="reg-password" class="block text-sm font-medium text-slate-400 mb-2">Password</label>
            <input id="reg-password" type="password" placeholder="Create a password (min 8 characters)" name="password" bind:value={password} required class="input-field" />
        </div>

        <button type="submit" disabled={username === "" || password.length < 8} class="btn-primary w-full mt-4">
            Create Account
        </button>

        <div class="mt-8 text-center">
            <p class="text-slate-400 text-sm">
                Already have an account? 
                <button type="button" class="text-indigo-400 hover:text-indigo-300 font-semibold underline transition-colors ml-1" onclick={() => view = "login"}>
                    Login here
                </button>
            </p>
        </div>
    </form>
</div>

<style>
</style>