<script lang="ts">
    import { enhance } from "$app/forms";
    import { goto } from "$app/navigation";

    const { form } = $props();
</script>

<form
    method="POST"
    class="bg-base-200 mx-auto my-5 box-border flex w-full max-w-md flex-col items-center p-8"
    use:enhance={() => {
        return (e) => {
            e.update();
            if (e.result.type != "error" && e.result.type != "failure") {
                goto("/home");
            }
        };
    }}
>
    <h1 class="mb-2 text-xl font-bold">Log In</h1>
    <fieldset class="fieldset w-full">
        <legend class="fieldset-legend">Username</legend>
        <input type="text" class="input w-full" name="username" value={form?.username ?? ""} />
    </fieldset>
    <fieldset class="fieldset w-full">
        <legend class="fieldset-legend">Password</legend>
        <input type="password" class="input w-full" name="password" />
    </fieldset>
    <button type="submit" class="btn btn-primary mt-4 w-full">Submit</button>
    {#if form?.error}
        {#if form.error == "unknown"}
            <p class="text-error mt-2 w-full text-right">Something went wrong</p>
        {/if}
        {#if form.error == "credentials"}
            <p class="text-error mt-2 w-full text-right">Wrong credentials</p>
        {/if}
    {/if}
</form>
