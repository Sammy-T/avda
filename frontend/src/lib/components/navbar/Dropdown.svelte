<script>
    import { onDestroy } from 'svelte';

    /**
     * @callback onselect
     * @param {String} selected
     */

    /**
     * @typedef {Object} DropdownItem
     * @property {String} label
     * @property {String} value
     */

    /**
     * @typedef {Object} Props
     * @property {String} name
     * @property {DropdownItem[]} items
     * @property {String} selected
     * @property {import('svelte').Snippet} children
     * @property {onselect} onselect
     */

    /** @type {Props} */
    let { name, items, selected, children, onselect } = $props();

    let listVisible = $state(false);

    /** @type {HTMLDivElement} */
    let dropdownContainer;

    $effect(() => {
        if(listVisible) {
            window.addEventListener('click', handleWindowClick);
        } else {
            window.removeEventListener('click', handleWindowClick);
        }
    });

    /**
     * @param {Event} event
     */
    function handleWindowClick(event) {
        // @ts-ignore
        if(dropdownContainer.contains(event.target)) return; // Ignore clicks on this component

        listVisible = false; // The user clicked away so hide the list
    }

    /**
     * @param {Event} event
     */
    function toggleList(event) {
        event.preventDefault();

        // @ts-ignore
        event.currentTarget.blur(); // Remove focus from the anchor element so the tooltip doesn't remain
        
        listVisible = !listVisible;
    }

    /**
     * @param {Event} event
     */
    function select(event) {
        listVisible = false;

        // @ts-ignore
        onselect(event.currentTarget.value);
    }

    onDestroy(() => {
        window.removeEventListener('click', handleWindowClick);
    });
</script>

{#snippet dropdownItem(/** @type {DropdownItem} */ item)}
    <li>
        <label>
            <input type="radio" {name} value={item.value} checked={item.value === selected} onclick={select} />
            {item.label}
        </label>
    </li>
{/snippet}

<div class="dropdown-container" bind:this={dropdownContainer}>
    <a href="##" class="contrast" data-tooltip="Sort" data-placement="bottom" onclick={toggleList}>
        {@render children()}
    </a>
    
    <ul class="dropdown-list" data-theme="light" style:display={listVisible ? 'flex' : 'none'}>
        {#each items as item}
            {@render dropdownItem(item)}
        {/each}
    </ul>
</div>

<style>
    .dropdown-container {
        position: relative;
    }

    .dropdown-list {
        width: max-content;
        margin: 0;
        position: absolute;
        right: 0;
        top: calc(var(--pico-spacing) * 2);
        background-color: var(--pico-background-color);
        border-radius: var(--pico-border-radius);
        box-shadow: var(--pico-dropdown-box-shadow);
        flex-direction: column;
    }

    li, label {
        width: 100%;
    }

    li {
        padding: 0;
    }

    label {
        margin: 0;
        padding: calc(var(--pico-nav-element-spacing-vertical) * 0.65) var(--pico-nav-element-spacing-horizontal);
        border-radius: var(--pico-border-radius);
    }

    label:hover {
        background-color: var(--pico-dropdown-hover-background-color);
    }

    a:hover {
        color: var(--pico-primary-background);
    }
</style>
