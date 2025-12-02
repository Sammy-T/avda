<script>
    import chevronDownIc from '$assets/images/chevron-down.svg?raw';
    import { STORAGE_KEY_RECENT_FILES } from '$lib/util.svelte';
    import { SelectVault } from '$wails/go/main/App';
    import { OnFileDrop, OnFileDropOff } from '$wails/runtime/runtime';
    import { onMount } from 'svelte';

    /**
     * @typedef {Object} Props
     * @property {String} [placeholder]
     * @property {String} [selected]
     * @property {Function} onselect
     */

    /** @type {Props} */
    let { placeholder = 'No file chosen', selected = $bindable(), onselect } = $props();

    /** @type {String[]} */
    let recentFiles = $state([]);

    let dragover = $state(false);

    /** @type {HTMLDetailsElement}*/
    let recentDropdown;

    function updateRecentFiles() {
        recentFiles = [selected, ...recentFiles.filter((f) => f !== selected)].slice(0, 3);
        
        localStorage.setItem(STORAGE_KEY_RECENT_FILES, JSON.stringify(recentFiles));
    }

    /**
     * @param {Event} event
     */
    function selectRecent(event) {
        event.preventDefault();

        // @ts-ignore
        selected = event.target.dataset.file;

        updateRecentFiles();

        recentDropdown.removeAttribute('open');

        onselect();
    }

    /**
     * @param {Event} event
     */
    async function select(event) {
        event.preventDefault();

        onselect();

        const resp = await SelectVault();
        
        if(resp.status !== 'success' || resp.data === '') {
            console.warn('No file selected.', resp);
            return;
        }

        selected = resp.data;

        updateRecentFiles();
    }

    /**
     * @param {Event} _event
     */
    function onDragOver(_event) {
        dragover = true;
    }

    /**
     * A handler for when the drop is completed or cancelled.
     * @param {Event} _event
     */
    function onDrop(_event) {
        dragover = false;
    }

    onMount(() => {
        recentFiles = JSON.parse(localStorage.getItem(STORAGE_KEY_RECENT_FILES)) ?? [];

        if(recentFiles.length > 0) selected = recentFiles.at(0);

        // Set the dropped file as the selected file.
        OnFileDrop((x, y, paths) => selected = paths.at(0), true);

        return () => {
            OnFileDropOff();
        }
    });
</script>

{#snippet recent(/** @type {String} */ file)}
    <li>
        <a href="##" data-file={file} onclick={selectRecent}>
            {file.split(/[\\/]/).at(-1)}
        </a>
    </li>
{/snippet}

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="file-input" class:dragover 
    ondragover={onDragOver} ondragleave={onDrop} ondragend={onDrop} ondrop={onDrop}>

    <fieldset id="file-select">
        <button type="button" class="secondary" onclick={select}>
            Choose File
        </button>

        <details class="dropdown" bind:this={recentDropdown}>
            <!-- svelte-ignore a11y_no_redundant_roles -->
            <summary role="button" class="secondary">{@html chevronDownIc}</summary>

            <ul>
                {#if recentFiles.length === 0}
                    <li class="empty">No recent files</li>
                {/if}

                {#each recentFiles as file}
                    {@render recent(file)}
                {/each}
            </ul>
        </details>
    </fieldset>

    <p>{selected?.split(/[\\/]/).at(-1) ?? placeholder}</p>

    <input type="hidden" name="filepath" bind:value={selected} />
</div>

<style>
    .file-input {
        margin-bottom: calc(var(--pico-spacing) * 1);
        border-radius: calc(var(--pico-spacing) * 0.1);
        display: flex;
        align-items: center;
        gap: calc(var(--pico-spacing) * 1);
        --wails-drop-target: drop;

        & > p {
            margin: 0;
        }
    }

    #file-select {
        width: unset;
        display: flex;
        margin: 0;

        & * {
            margin: 0;
        }

        & details li {
            padding: 0;
        }

        & button {
            border-top-right-radius: 0;
            border-bottom-right-radius: 0;
            padding: calc(var(--pico-form-element-spacing-vertical) * 0.5) calc(var(--pico-form-element-spacing-horizontal) * 0.5);
            text-wrap: nowrap;
        }

        & summary {
            border-top-left-radius: 0;
            border-bottom-left-radius: 0;
            padding: calc(var(--pico-form-element-spacing-vertical) * 0.5) calc(var(--pico-form-element-spacing-horizontal) * 0.25);
        }

        & summary::after {
            display: none;
        }

        & .empty {
            padding: calc(var(--pico-form-element-spacing-vertical) * 0.5) calc(var(--pico-form-element-spacing-horizontal) * 0.25);
        }
    }

    .dragover {
        outline: 2px solid rgba(255, 255, 255, 0.845);
    }
</style>
