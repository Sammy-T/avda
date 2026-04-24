<script>
    import closeIc from '$assets/images/x.svg?raw';
    import { showSettings } from '$lib/stores.svelte';
    import { MAX_FONT, MIN_FONT, STORAGE_KEY_AUTO_CLOSE, STORAGE_KEY_AUTO_CLOSE_TIME, STORAGE_KEY_UI_SCALE, STORAGE_KEY_THEME } from '$lib/util.svelte';

    let theme = $state(localStorage.getItem(STORAGE_KEY_THEME) || '');
    let fontScale = $state(Number(localStorage.getItem(STORAGE_KEY_UI_SCALE)) || getCurrentFontScale());
    let autoClose = $state(localStorage.getItem(STORAGE_KEY_AUTO_CLOSE) === 'true');
    let autoCloseTime = $state(Number(localStorage.getItem(STORAGE_KEY_AUTO_CLOSE_TIME)) || 5);

    /** @type {HTMLFormElement} */
    let form;

    $effect(() => {
        // Pico uses a percent-based root font size
        const root = document.documentElement;
        root.style.setProperty('--pico-font-size', `${fontScale}%`);
    });

    function getCurrentFontScale() {
        const root = document.documentElement;
        const style = getComputedStyle(root);
        const fontSize = Number(style.getPropertyValue('--pico-font-size').replace('%', ''));

        return Math.round(fontSize);
    }

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
                <select id="theme" name="theme" bind:value={theme} oninput={onThemeSelect}>
                    <option value="">System</option>
                    <option value="light">Light</option>
                    <option value="dark">Dark</option>
                </select>

                <label for="scale">UI Scale</label>
                <fieldset class="row">
                    <input id="scale" name="scale" type="number" min={MIN_FONT} max={MAX_FONT} bind:value={fontScale} />
                    <input type="range" min={MIN_FONT} max={MAX_FONT} bind:value={fontScale} />
                </fieldset>

                <label class="controlled">
                    Auto-close vault
                    <input type="checkbox" role="switch" bind:checked={autoClose} />
                    <input name="auto-close" type="hidden" value={autoClose} />
                </label>

                <fieldset class="row">
                    <input id="auto-close-time" name="auto-close-time" type="number" min="1" max="15" bind:value={autoCloseTime}
                        disabled={!autoClose} />
                    <label for="auto-close-time">minutes</label>
                </fieldset>
            </div>
        </form>
    </article>
</dialog>

<style>
    article {
        width: min(80dvw, 600px);
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

    form {
        padding: 0 1rem;
    }

    .fields {
        display: grid;
        grid-template-columns: auto 1fr;
        align-items: center;
        gap: calc(var(--pico-spacing) * 2);

        & fieldset, & input {
            margin: 0;
        }

        & .controlled {
            display: flex;
            align-items: center;
            gap: 0.5rem;
            flex-wrap: wrap;
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

            & input[type="number"] {
                width: 5rem;
            }

            & :last-child {
                flex-grow: 1;
            }
        }
    }
</style>
