<script>
    import avdaIc from '$assets/images/avda.svg?raw';
    import infoIc from '$assets/images/info-circle.svg?raw';
    import { OpenVault, SelectVault } from '$wails/go/main/App';
    import { OnFileDrop, OnFileDropOff } from '$wails/runtime/runtime';
    import { openExtUrl } from '$lib/util';
    import { releaseUrl, vaultPath } from '$lib/stores';
    import { onMount } from 'svelte';

    let dragover = $state(false);
    let selected = $state();
    let respError = $state(false);

    /**
     * @param {Event} event
     */
    async function select(event) {
        event.preventDefault();

        respError = false;

        const resp = await SelectVault();
        
        if(resp.status !== 'success' || resp.data === '') {
            console.warn('No file selected.', resp);
            return;
        }

        selected = resp.data;
    }

    /**
     * Passes the filepath and password to the Go backend
     * to attempt to read the vault file.
     * @param {SubmitEvent} event
     */
    async function onSubmit(event) {
        event.preventDefault();

        respError = false;

        // @ts-ignore
        const formData = new FormData(event.target);

        const filepath = formData.get('filepath');
        const password = formData.get('password');

        if(filepath === '') return;

        const resp = await OpenVault(filepath.toString(), password.toString());

        if(resp.status !== 'success') {
            console.error(resp.message);
            respError = true;
            return;
        }

        $vaultPath = filepath;
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

<dialog open>
    <article>
        <h1>{@html avdaIc} avda</h1>

        <h5>Select vault file</h5>

        <form onsubmit={onSubmit}>
            <!-- svelte-ignore a11y_no_static_element_interactions -->
            <div class="file-input" class:dragover 
                ondragover={onDragOver} ondragleave={onDrop} ondragend={onDrop} ondrop={onDrop}>

                <button type="button" class="secondary" onclick={select}>
                    Choose File
                </button>

                <p>{selected?.split(/[\\/]/).at(-1) ?? 'No file chosen'}</p>

                <input type="hidden" name="filepath" bind:value={selected} />
            </div>

            <input type="password" name="password" placeholder="Password" />
            <small>Leave blank for unencrypted vault files.</small>

            {#if respError}
                <small class="warning">Unable to open vault.</small>
            {/if}

            <button>Unlock</button>

            {#if $releaseUrl}
                <div class="info">
                    <small><a href={$releaseUrl} onclick={openExtUrl}>{@html infoIc} Update Available</a></small>
                </div>
            {/if}
        </form>
    </article>
</dialog>

<style>
    article {
        max-width: 500px;
        padding: calc(var(--pico-spacing) * 2);
        background-image: linear-gradient(135deg, #d7e3f4ff, #8f9bc9ff);
    }

    h1 {
        text-align: center;
        margin-bottom: calc(var(--pico-typography-spacing-vertical) * 2.5);
        color: var(--pico-primary);
    }

    h1, h5 {
        font-weight: 400;
    }

    .file-input {
        margin-bottom: calc(var(--pico-spacing) * 1);
        border-radius: calc(var(--pico-spacing) * 0.1);
        display: flex;
        align-items: center;
        gap: calc(var(--pico-spacing) * 1);
        --wails-drop-target: drop;

        & > button {
            padding: calc(var(--pico-form-element-spacing-vertical) * 0.5) calc(var(--pico-form-element-spacing-horizontal) * 0.5);
            margin: 0;
            text-wrap: nowrap;
        }

        & > p {
            margin: 0;
        }
    }

    form > button {
        width: 100%;
    }

    .dragover {
        outline: 2px solid rgba(255, 255, 255, 0.845);
    }

    .info {
        display: flex;
        justify-content: center;
        padding-top: calc(var(--pico-spacing) * 0.5);

        & * {
            text-decoration: none;
        }
    }
</style>
