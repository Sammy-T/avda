<dialog open>
    <article>
        <h1>avdap</h1>

        <h5>Select vault file</h5>
        <form>
            <div class="file-input">
                <button class="secondary">Choose File</button>
                No file chosen
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
