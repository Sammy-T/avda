<script>
    import EntryItem from './EntryItem.svelte';
    import { items, search } from '$lib/stores';
    import { closeFile } from '$lib/util';
    import { EventsOn } from '$wails/runtime/runtime';
    import { getContext } from 'svelte';

    /** @type {import('svelte/store').Writable<boolean>} */
    const displaySearch = getContext('displaySearch');

    let filteredItems = $derived($items.filter((item) => {
        const searchRE = new RegExp($search, 'i');

        const { issuer, name } = item.entry;

        return searchRE.test(issuer) || searchRE.test(name);
    }));

    /**
     * @param event {KeyboardEvent}
     */
    function onKeyEvent(event) {
        if(filteredItems.length === 0 || !event.ctrlKey) return;

        switch(event.key) {
            case 'f':
                $displaySearch = !$displaySearch;
                break;
            
            case 'w':
                closeFile();
                break;
        }
    }

    EventsOn('onCodesUpdated', (entryCodes) => $items = entryCodes);
</script>

<svelte:document on:keyup={onKeyEvent} />

{#if filteredItems.length === 0}
    <div id="empty">
        <p>No entries.</p>
        <button onclick={closeFile}>Open Vault</button>
    </div>
{:else}
    <div class="list">
        {#each filteredItems as item (item.entry.uuid) }
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
