<dialog open>
    <article>
        <h1>avdap</h1>

        <h5>Select vault file</h5>
        <form>
            <input type="file" />

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

    form > button {
        width: 100%;
    }
</style>
