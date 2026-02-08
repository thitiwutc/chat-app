<script>
	import { onMount } from "svelte";
	const socket = new WebSocket("http://localhost:3000/ws");
	let msg = $state("");
	$inspect(msg);

	onMount(() => {
		// Connection opened
		socket.addEventListener("open", (event) => {
			socket.send("Hello Server!");
		});

		// Listen for messages
		socket.addEventListener("message", (event) => {
			console.log("Message from server ", event.data);
		});
	});

	function sendMsg() {
		socket.send(msg);
		msg = "";
	}
</script>

<h1>Welcome to SvelteKit</h1>
<input type="text" bind:value={msg} />
<button type="button" onclick={sendMsg}>Send</button>
