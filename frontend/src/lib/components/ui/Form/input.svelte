<script lang="ts">
	import { twMerge } from "tailwind-merge";
	let props = $props()

	console.log(props);
	
</script>


<input
	type={props.type}
	placeholder={props.placeholder ?? ''}
	pattern={props.pattern }
	minlength={props.min ?? "3"}
	maxlength={props.max ?? ""}
	title={props.title ?? ""}
	class={twMerge("border-none",props.class)}
/>
