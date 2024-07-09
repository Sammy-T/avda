<script>
    import EntryItem from './EntryItem.svelte';
    import { closeFile } from '$lib/util';
    import { EventsOn } from '$wails/runtime/runtime';

    let items = [];

    EventsOn('onCodesUpdated', (entryCodes) => items = entryCodes);
</script>

{#if items.length === 0}
    <div id="empty">
        <p>No entries.</p>
        <button on:click={closeFile}>Open Vault</button>
    </div>
{:else}
    <div class="list">
        {#each items as item (item.entry.uuid) }
            <EntryItem {item} />
        {/each}
    </div>
{/if}

<style>
    .list {
        flex-grow: 1;
        padding: calc(var(--pico-spacing) * 1.5);
        overflow-y: auto;
        display: flex;
        align-content: flex-start;
        align-items: flex-start;
        flex-wrap: wrap;
        gap: calc(var(--pico-spacing) * 2);
    }

    #empty {
        height: 100%;
        text-align: center;
        align-self: center;
        display: flex;
        flex-direction: column;
        justify-content: center;
    }
</style>
