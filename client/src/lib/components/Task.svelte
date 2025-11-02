<script lang="ts">
    
    import { fetchAPI } from "$lib/_core";
    import { onMount } from "svelte";

    let { task = $bindable() } = $props()
    let isSubmit:boolean = $state(false)

    async function patchTask(){
        isSubmit = true
        const data = await fetchAPI(`/tasks/${task.id}`, "PATCH")
        isSubmit = false
        if(data.error){
            return
        }
        task = data.result
        dateFormat()
    }

    async function removeTask(){
        isSubmit = true
        const data = await fetchAPI(`/tasks/${task.id}`, "DELETE")
        isSubmit = false
        if(data.error){
            return
        }
        task.isDeleted = true
    }

    let dateTo:Date
    let dateToFormat:string = $state("")
    let diffLabel:string = $state("")

    function dateFormat(){
        if(task.dateTo === null){
            return
        }
        dateTo = new Date(task.dateTo)
        dateToFormat = convertNumber(dateTo.getDate())+"/"+convertNumber(dateTo.getMonth()+1)+"/"+dateTo.getFullYear()+" "+convertNumber(dateTo.getHours())+":"+convertNumber(dateTo.getMinutes())+":"+convertNumber(dateTo.getSeconds())
        const diffMs = dateTo.getTime() - dateAdd.getTime()
        const diffSecs  = Math.floor(diffMs / 1000)
        const diffMinutes = Math.floor(diffSecs / 60)
        const hours = Math.floor(diffMinutes / 60)
        const minutes = diffMinutes % 60
        const seconds = diffSecs % 60

        diffLabel = "Done in "
        if(hours > 0){
            diffLabel += hours+" hour"
            if(hours > 1) diffLabel += "s"
            if(minutes > 0 || seconds > 0) diffLabel +=" et "
        }
        if(minutes > 0){
            diffLabel += minutes+" minute"
            if(minutes > 1) diffLabel += "s"
            if(seconds > 0) diffLabel += " et "
        }
        if(seconds > 0){
            diffLabel += seconds+" second"
            if(seconds > 1) diffLabel += "s"
        }
    }

    let dateAdd:any
    let dateAddFormat:any = $state("")

    let formPUT:undefined|HTMLFormElement = $state()

    async function saveTask():Promise<void>
    {
        isSubmit = true
        const data = await fetchAPI(`/tasks/${task.id}`, "PUT", new FormData(formPUT))
        isSubmit = false
        if(data.error){
            return
        }
        edit = false
        task = data
        dateFormat()
    }

    function convertNumber(number:number, addTo:boolean = false):string
    {
        if(addTo) number++
        let numberVal = String(number)
        if(number < 10){
            numberVal = "0"+number
        }
        return numberVal
    }

    $effect(() => {
        dateAdd = new Date(task.dateAdd)
        dateAddFormat = convertNumber(dateAdd.getDate())+"/"+convertNumber(dateAdd.getMonth(), true)+"/"+dateAdd.getFullYear()+" "+convertNumber(dateAdd.getHours())+":"+convertNumber(dateAdd.getMinutes())+":"+convertNumber(dateAdd.getSeconds())
    })

    onMount(() => {
        dateFormat()
        console.log(task)
    })

    let edit:boolean = $state(false)

</script>

<article>
    {#if !edit}
        <header>
            {task.title} 
            {#if task.isDone}({diffLabel}){/if}
        </header>
        Add on {dateAddFormat}
        {#if task.dateTo !== null}<br />Done the {dateToFormat}{/if}
    {:else}
        <form onsubmit={saveTask} bind:this={formPUT}>
            <header>
                <input type="text" name="title" value={task.title} /> 
            </header>
            Add on <input type="datetime-local" step="1" name="dateAdd" value={task.dateAdd} /><br />
            Finished on <input type="datetime-local" step="1" name="dateTo" value={task.dateTo} />
            <button class="primary" style="width:100%;">Save</button>
        </form>
    {/if}

    <footer role="group">
        <button class="secondary" 
        onclick={patchTask} 
        disabled={isSubmit === true}>
            {#if task.dateTo}Cancel{:else}Finish{/if}
        </button>
        <button class="contrast" onclick={() => edit = !edit}>
            {#if edit}Cancel{:else}Edit{/if}
        </button>
        <button class="outline secondary" 
        onclick={removeTask} 
        disabled={isSubmit === true}>Delete</button>
    </footer>
</article>