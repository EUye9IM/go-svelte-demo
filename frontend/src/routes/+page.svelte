<script lang="ts">
	let text = "";
	function submit(event: SubmitEvent) {
		const input = document.querySelector("#name") as HTMLInputElement;
		if (input.value == "") {
			text = "";
			return;
		}
		fetch("/api/demo", {
			method: "POST",
			headers:{'content-type': 'application/json'},
			body: JSON.stringify({ name: input.value }),
		})
			.then((res) => res.json())
			.then((data) => (text = data.msg))
			.catch((err) => console.log(err));
	}
</script>

<h1>Welcome to SvelteKit</h1>
<p>
	Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
</p>

<form on:submit={submit}>
	<input id="name" autocomplete="off" />
	<input type="submit" />
</form>

<p>{text}</p>
