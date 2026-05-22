<script lang="ts">
    let { variant = "tasks" }: { variant: "dashboard" | "tasks" | "commonTasks" } = $props();
</script>

{#if variant === "dashboard"}
    <div class="space-y-12">
        <!-- 3 Stats Cards Skeleton -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            {#each Array(3) as _}
                <div class="glass-card p-8 flex flex-col items-center text-center animate-pulse">
                    <div class="w-12 h-12 bg-slate-700/30 rounded-xl mb-4"></div>
                    <div class="w-16 h-8 bg-slate-700/40 rounded-lg mb-2"></div>
                    <div class="w-28 h-4 bg-slate-700/20 rounded-lg"></div>
                </div>
            {/each}
        </div>

        <!-- Productivity Summary Card Skeleton -->
        <div class="glass-card p-8 animate-pulse">
            <div class="w-48 h-6 bg-slate-700/40 rounded-lg mb-6"></div>
            <div class="space-y-4">
                {#each Array(2) as _}
                    <div class="flex justify-between items-center p-4 bg-white/5 rounded-lg border border-white/5">
                        <div class="w-32 h-5 bg-slate-700/20 rounded-lg"></div>
                        <div class="w-12 h-5 bg-slate-700/30 rounded-lg"></div>
                    </div>
                {/each}
            </div>
        </div>
    </div>
{:else if variant === "tasks"}
    <div class="space-y-6">
        {#each Array(3) as _}
            <div class="glass-card p-6 mb-6 animate-pulse">
                <div class="flex flex-col gap-4">
                    <div class="flex justify-between items-start gap-4">
                        <div class="space-y-2 flex-grow">
                            <div class="w-2/3 h-6 bg-slate-700/40 rounded-lg"></div>
                            <div class="w-1/3 h-4 bg-slate-700/20 rounded-lg"></div>
                        </div>
                    </div>
                    <div class="mt-2 pt-6 border-t border-white/5 flex gap-2">
                        <div class="w-20 h-8 bg-slate-700/30 rounded-lg"></div>
                        <div class="w-20 h-8 bg-slate-700/30 rounded-lg"></div>
                        <div class="w-20 h-8 bg-slate-700/20 rounded-lg ml-auto"></div>
                    </div>
                </div>
            </div>
        {/each}
    </div>
{:else if variant === "commonTasks"}
    <div class="space-y-3">
        {#each Array(4) as _}
            <div class="flex items-center justify-between p-3 bg-white/5 rounded-lg border border-white/5 animate-pulse">
                <div class="w-1/2 h-5 bg-slate-700/30 rounded-lg"></div>
                <div class="flex gap-1">
                    <div class="w-7 h-7 bg-slate-700/30 rounded-md"></div>
                    <div class="w-7 h-7 bg-slate-700/30 rounded-md"></div>
                </div>
            </div>
        {/each}
    </div>
{/if}
