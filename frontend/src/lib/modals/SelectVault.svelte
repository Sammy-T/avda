<script>
    import { OpenVault, SelectVault } from '$wails/go/main/App';
    import { vaultPath } from '$lib/stores';

    let selected;

    async function select() {
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
     * @param event {SubmitEvent}
     */
    async function onSubmit(event) {
        // @ts-ignore
        const formData = new FormData(event.target);

        const filepath = formData.get('filepath');
        const password = formData.get('password');

        if(filepath === '') return;

        const resp = await OpenVault(filepath.toString(), password.toString());

        if(resp.status !== 'success') {
            console.error(resp.message);
            return;
        }

        $vaultPath = filepath;
    }
</script>

<dialog open>
    <article>
        <h1>avda</h1>

        <h5>Select vault file</h5>

        <form on:submit|preventDefault={onSubmit}>
            <div class="file-input">
                <button class="secondary" on:click|preventDefault={select}>
                    Choose File
                </button>

                {selected?.split(/[\\/]/).at(-1) ?? 'No file chosen'}

                <input type="hidden" name="filepath" bind:value={selected} />
            </div>

            <input type="password" name="password" placeholder="Password" />
            <small>Leave blank for unencrypted vault files.</small>

            <button>Unlock</button>
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
    }

    .file-input > button {
        padding: calc(var(--pico-form-element-spacing-vertical) * 0.5) calc(var(--pico-form-element-spacing-horizontal) * 0.5);
        margin-right: calc(var(--pico-spacing) * 0.5);
    }

    form > button {
        width: 100%;
    }
</style>
