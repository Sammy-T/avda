<script>
    import { selectedGroupUuid, groupsMap } from '$lib/stores.svelte';
    
    function selectGroup(groupUuid) {
        $selectedGroupUuid = groupUuid;
    }
</script>

<div class="group-filter">
    {#each Array.from($groupsMap.entries()) as [uuid, group]}
        <button 
            class="group-button" 
            class:active={$selectedGroupUuid === group.uuid}
            onclick={() => selectGroup(group.uuid)}
            style={group.uuid ? `--group-color: ${group.color}` : ''}
        >
            {group.name}
        </button>
    {/each}
</div>

<style>
    .group-filter {
        padding: calc(var(--pico-spacing) * 0.5) calc(var(--pico-spacing) * 1.5);
        border-bottom: 1px solid #ffffffb0;
        background-color: #b5c6e2;
        padding-bottom: calc(var(--pico-spacing) * 0.75);
        position: sticky;
        top: 0;
        z-index: 1;
    }
    
    .group-button {
        background: transparent;
        border: none;
        padding: calc(var(--pico-spacing) * 0.4) calc(var(--pico-spacing) * 0.8);
        margin-right: calc(var(--pico-spacing) * 0.5);
        font-size: 0.9rem;
        transition: all 0.2s ease;
        color: var(--pico-muted-color);
        min-width: 3rem;
        text-align: center;
    }
    
    .group-button.active {
        background: var(--group-color, var(--pico-primary));
        color: var(--pico-primary-inverse);
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }
    
    .group-button.active:not([style*="--group-color"]) {
        background: var(--pico-primary);
    }
</style>
