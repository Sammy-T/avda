<script>
    import { onMount } from 'svelte';
    import { selectedGroupUuid, groupsMap } from '$lib/stores';
    import { GetGroups } from '$wails/go/main/App';
    import { getGroupColor } from '$lib/util';

    export let items = [];
    
    onMount(async () => {
        const backendGroups = await GetGroups();
        
        const mapping = new Map();
        mapping.set(null, { uuid: null, name: "All", color: null });  
        backendGroups?.forEach(group => {
            mapping.set(group.uuid, {
                uuid: group.uuid,
                name: group.name,
                color: getGroupColor(group.uuid)
            });
        });
        
        groupsMap.set(mapping);
    });
    
    function selectGroup(groupUuid) {
        selectedGroupUuid.set(groupUuid);
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
        border-bottom: 1px solid var(--pico-muted-border-color);
        background-color: var(--pico-background-color);
        padding-bottom: calc(var(--pico-spacing) * 0.75);
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
