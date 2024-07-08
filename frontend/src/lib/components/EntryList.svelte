<script>
    import EntryItem from './EntryItem.svelte';

    let entries = [];
</script>

<div class="list">
    {#if entries.length === 0}
        <div id="empty">
            <p>No entries.</p>
            <button>Open Vault</button>
        </div>
    {:else}
        {#each entries as entry (entry.uuid) }
            <EntryItem />
        {/each}
    {/if}
</div>

<style>
    .list {
        flex-grow: 1;
        padding: calc(var(--pico-spacing) * 1.5);
        overflow-y: auto;
        display: flex;
        flex-wrap: wrap;
        align-items: flex-start;
        gap: calc(var(--pico-spacing) * 2);
    }

    #empty {
        width: 100%;
        text-align: center;
        align-self: center;
    }
</style>
