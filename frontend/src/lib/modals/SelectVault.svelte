<script>
    import FileInput from '$lib/components/select-vault/FileInput.svelte';
    import avdaIc from '$assets/images/avda.svg?raw';
    import infoIc from '$assets/images/info-circle.svg?raw';
    import { OpenVault } from '$wails/go/main/App';
    import { openExtUrl } from '$lib/util.svelte';
    import { releaseUrl, vaultPath } from '$lib/stores.svelte';
    import { onMount } from 'svelte';

    let respError = $state(false);

    /** @type {String} */
    let selected = $state();

    /** @type {HTMLInputElement} */
    let pwdInput;

    function onSelect() {
        pwdInput.focus();
        respError = false;
    }

    function onVisibilityChange() {
        if(selected && document.visibilityState === 'visible') pwdInput.focus();
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

    onMount(() => {
        if(selected) pwdInput.focus();

        document.addEventListener('visibilitychange', onVisibilityChange);

        return () => document.removeEventListener('visibilitychange', onVisibilityChange);
    });
</script>

<dialog open>
    <article>
        <h1>{@html avdaIc} avda</h1>

        <h5>Select vault file</h5>

        <form onsubmit={onSubmit}>
            <FileInput onselect={onSelect} bind:selected />

            <input type="password" name="password" placeholder="Password" bind:this={pwdInput} />
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

    form > button {
        width: 100%;
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
