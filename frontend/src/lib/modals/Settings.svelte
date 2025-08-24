<script>
    import closeIc from '$assets/images/x.svg?raw';
    import { showSettings } from '$lib/stores.svelte';
    import { STORAGE_KEY_AUTO_CLOSE, STORAGE_KEY_AUTO_CLOSE_TIME, STORAGE_KEY_THEME } from '$lib/util.svelte';
    import { onMount } from 'svelte';

    let theme = $state('');
    let autoClose = $state(false);
    let autoCloseTime = $state(5);

    /** @type {HTMLFormElement} */
    let form;

    function onThemeSelect(ev) {
        const body = document.querySelector('body');
        body.dataset.theme = ev.target.value;
    }

    function save() {
        const formData = new FormData(form);

        for(const [key, value] of formData.entries()) {
            localStorage.setItem(`avda:${key}`, value.toString());
        }
    }

    /**
     * @param {Event} event
     */
    function close(event, override = false) {
        if(!override && event.currentTarget !== event.target) return;

        save();

        $showSettings = false;
    }

    onMount(() => {
        theme = localStorage.getItem(STORAGE_KEY_THEME) || '';
        autoClose = localStorage.getItem(STORAGE_KEY_AUTO_CLOSE) === 'true';
        autoCloseTime = Number(localStorage.getItem(STORAGE_KEY_AUTO_CLOSE_TIME)) || 5;
    });
</script>

<dialog open onclick={close}>
    <article>
        <header>            
            <strong>Settings</strong>

            <button class="outline" onclick={(ev) => close(ev, true)}>{@html closeIc}</button>
        </header>

        <form onsubmit={(ev) => ev.preventDefault()} bind:this={form}>
            <div class="fields">
                <label for="theme">Theme</label>
                <select name="theme" bind:value={theme} oninput={onThemeSelect}>
                    <option value="">System</option>
                    <option value="light">Light</option>
                    <option value="dark">Dark</option>
                </select>

                <label>
                    Auto-close vault
                    <input type="checkbox" role="switch" bind:checked={autoClose} />
                    <input name="auto-close" type="hidden" value={autoClose} />
                </label>

                <fieldset class="row">
                    <input name="auto-close-time" type="number" min="1" max="15" bind:value={autoCloseTime}
                        disabled={!autoClose} />
                    <label for="auto-close-time">minutes</label>
                </fieldset>
            </div>
        </form>
    </article>
</dialog>

<style>
    article {
        max-width: 500px;
    }

    header {
        display: flex;
        justify-content: space-between;
        align-items: center;

        & button {
            padding: 0;
            border: none;
        }
    }

    .fields {
        display: grid;
        grid-template-columns: auto 1fr;
        align-items: center;
        gap: calc(var(--pico-spacing) * 2);

        & fieldset, & input {
            margin: 0;
        }

        & .row {
            display: flex;
            align-items: center;
            gap: calc(var(--pico-spacing) * 0.5);

            & * {
                margin: 0;
            }

            & input {
                width: revert;
            }
        }
    }
</style>
