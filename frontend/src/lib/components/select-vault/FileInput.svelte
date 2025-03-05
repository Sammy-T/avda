<script>
    import chevronDownIc from '$assets/images/chevron-down.svg?raw';
    import { SelectVault } from '$wails/go/main/App';
    import { OnFileDrop, OnFileDropOff } from '$wails/runtime/runtime';
    import { onMount } from 'svelte';

    /**
     * @typedef {Object} Props
     * @property {String} [placeholder]
     * @property {Function} onselect
     */

    /** @type {Props} */
    let { placeholder = 'No file chosen', onselect } = $props();

    /** @type {String} */
    let selected = $state();

    let dragover = $state(false);

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
        // Set the dropped file as the selected file.
        OnFileDrop((x, y, paths) => selected = paths.at(0), true);

        return () => {
            OnFileDropOff();
        }
    });
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="file-input" class:dragover 
    ondragover={onDragOver} ondragleave={onDrop} ondragend={onDrop} ondrop={onDrop}>

    <fieldset id="file-select">
        <button type="button" class="secondary" onclick={select}>
            Choose File
        </button>

        <details class="dropdown">
            <!-- svelte-ignore a11y_no_redundant_roles -->
            <summary role="button" class="secondary">{@html chevronDownIc}</summary>

            <ul>
                <li><a href="##">Recent File 1</a></li>
                <li><a href="##">Recent File 2</a></li>
                <li><a href="##">Recent File 3</a></li>
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
    }

    .dragover {
        outline: 2px solid rgba(255, 255, 255, 0.845);
    }
</style>
