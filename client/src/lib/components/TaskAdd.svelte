<script lang="ts">
    import { tasks } from "$lib";
    import { fetchAPI } from "$lib/_core";

    let form: HTMLFormElement | undefined = $state()
    let title: string = $state("")
    let isSubmit: boolean = $state(false)

    async function submitTask(e: Event): Promise<void> {
        e.preventDefault()
        if (!form) return
        
        isSubmit = true
        const response = await fetchAPI(
            "/tasks", 
            "POST", 
            new FormData(form)
        )
        isSubmit = false
        
        if (response.error) {
            console.error(response.error)
            return
        }
        title = ""
        tasks.set(response)
    }

</script>

<div class="glass-card p-6 mb-8 animate-in">
    <form id="setTask" onsubmit={submitTask} bind:this={form} class="flex flex-col sm:flex-row gap-4 items-center">
        <input type="text" name="title" placeholder="What needs to be done?" bind:value={title} required class="input-field flex-grow" />
        <button type="submit" disabled={title === "" || isSubmit} class="btn-primary whitespace-nowrap w-full sm:w-auto">
            Add Task
        </button>
    </form>
</div>

<style>
</style>