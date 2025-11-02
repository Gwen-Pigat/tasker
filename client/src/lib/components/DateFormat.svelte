<script lang="ts">
    import { onMount } from "svelte";

    let { date } = $props<{
        date: Date
    }>()

    let dateFormat:Date = $state(
        new Date(date)
    )
    let dateFormatStr: string = $state("")
    let diffLabel:string = $state("")

    let dateAdd:any

    function convertNumber(number:number, addTo:boolean = false):string
    {
        if(addTo) number++
        let numberVal = String(number)
        if(number < 10){
            numberVal = "0"+number
        }
        return numberVal
    }

    function setDateFormat(){
        if(date === null){
            return
        }
        dateFormatStr = convertNumber(date.getDate())+"/"+convertNumber(date.getMonth())+"/"+date.getFullYear()+" "+convertNumber(date.getHours())+":"+convertNumber(date.getMinutes())+":"+convertNumber(date.getSeconds())
        const diffMs = date.getTime() - dateAdd.getTime()
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

    onMount(() => {
        setDateFormat()
    })

</script>


<span>{dateFormatStr}</span>