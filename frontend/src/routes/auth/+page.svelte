<script lang="ts">
    import { enhance } from '$app/forms';
    import { goto } from '$app/navigation'
    
	const { form } = $props();

</script>

<form method="POST" class="w-full max-w-md flex flex-col items-center mx-auto my-5 p-8 box-border bg-base-200 " use:enhance={() => {
        return (e) => {
            e.update();
            if (e.result.type != "error" && e.result.type != "failure") {
                goto("/home");
            }
        }
    }}>
    <h1 class="text-xl font-bold mb-2">Log In</h1>
    <fieldset class="fieldset w-full">
      <legend class="fieldset-legend">Username</legend>
      <input type="text" class="input w-full" name="username" value={form?.username ?? ""} />
    </fieldset>
    <fieldset class="fieldset w-full">
      <legend class="fieldset-legend">Password</legend>
      <input type="password" class="input w-full" name="password" />
    </fieldset>
    <button type="submit" class="btn btn-primary w-full mt-4">Submit</button>
    {#if form?.error}
        {#if form.error == "unknown"}
            <p class="w-full text-right mt-2 text-error">Something went wrong</p>
        {/if}
        {#if form.error == "credentials"}
            <p class="w-full text-right mt-2 text-error">Wrong credentials</p>
        {/if}
    {/if}
</form>
