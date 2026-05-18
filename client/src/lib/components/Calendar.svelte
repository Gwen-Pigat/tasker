<script lang="ts">
    let { selectedDate = $bindable(), onDateChange } = $props();

    function handleInput(e: Event) {
        const value = (e.target as HTMLInputElement).value;
        selectedDate = value;
        onDateChange(value);
    }
</script>

<div class="relative group">
    <div class="absolute -inset-0.5 bg-gradient-to-r from-indigo-500 to-purple-500 rounded-xl blur opacity-20 group-hover:opacity-40 transition duration-300"></div>
    <div class="relative flex items-center bg-slate-900 rounded-xl border border-white/10 overflow-hidden">
        <div class="pl-4 text-slate-500">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd" />
            </svg>
        </div>
        <input 
            type="date" 
            value={selectedDate}
            oninput={handleInput}
            class="bg-transparent text-white text-sm font-bold py-2.5 px-4 focus:outline-none w-full cursor-pointer [color-scheme:dark]"
        />
    </div>
</div>
