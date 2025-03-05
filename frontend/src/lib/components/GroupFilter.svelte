<script>
    import { items, selectedGroup } from '$lib/stores';
    import { getGroupColor } from '$lib/util';

    let groups = $derived(getUniqueGroups($items));

    function getUniqueGroups(items) {
        const groupMap = new Map();
        
        const allGroup = { uuid: null, name: "All" };
        groupMap.set(null, allGroup);
        
        items.forEach(item => {
            if (item.groups && item.groups.length > 0) {
                item.groups.forEach(group => {
                    if (!groupMap.has(group.uuid)) {
                        const name = group.name || `Group ${group.uuid.substring(0, 8)}`;
                        
                        groupMap.set(group.uuid, { uuid: group.uuid, name });
                    }
                });
            }
        });
        
        // Convert the Map to an array and sort by name
        return Array.from(groupMap.values()).sort((a, b) => {
            if (a.uuid === null) return -1; // "All" always first
            if (b.uuid === null) return 1;
            return a.name.localeCompare(b.name);
        });
    }

    function selectGroup(groupId) {
        $selectedGroup = groupId;
    }
</script>

<div class="group-filter">
    {#each groups as group}
        <button 
            class="group-button" 
            class:active={$selectedGroup === group.uuid}
            onclick={() => selectGroup(group.uuid)}
            style={group.uuid ? `--group-color: ${getGroupColor(group.uuid)}` : ''}
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
        overflow: visible;
        position: relative;
        padding-bottom: calc(var(--pico-spacing) * 0.75);
        
    }
    
    .group-button {
        background: transparent;
        border: none;
        border-radius: calc(var(--pico-border-radius) * 1.5);
        padding: calc(var(--pico-spacing) * 0.4) calc(var(--pico-spacing) * 0.8);
        margin-right: calc(var(--pico-spacing) * 0.5);
        font-size: 0.9rem;
        cursor: pointer;
        white-space: nowrap;
        transition: all 0.2s ease;
        color: var(--pico-muted-color);
        position: relative;
        min-width: 3rem;
        text-align: center;
        z-index: 1;
        gap: calc(var(--pico-spacing) * 1.75);
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
