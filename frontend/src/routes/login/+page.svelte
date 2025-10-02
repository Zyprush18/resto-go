<script lang="ts">
	import Button from '$lib/components/ui/Button/Button.svelte';
import * as Cards from '$lib/components/ui/Card';
	import * as Form from '$lib/components/ui/Form';
	import { Mail, KeyRound, EyeClosed, Eye } from '@lucide/svelte';
 
    let show = false;
    let ButtonShow : HTMLButtonElement;
    let ButtonHidden : HTMLButtonElement;
    const passwordPattern = "(?=.*\\d)(?=.*[a-z])(?=.*[A-Z])[A-Za-z\\d]{8,}";
    const ShowPw = () => {
        show = true
        // hidden eye closed icon
        ButtonShow.style.display = 'none'
        // show eye icon
        ButtonHidden.style.display = 'block'
    }

    const hiddenPw = () => {
        show = false
        ButtonHidden.style.display = 'none'
        ButtonShow.style.display = 'block'
    }


</script>

<div class="flex h-screen items-center justify-center">
	<Cards.Card class="bg-gray-50 p-4 shadow-lg">
		<h1 class="mb-4 text-center text-4xl font-bold">Login</h1>
		<Cards.CardBody>
			<form action="" method="post">

                <div class="mb-3">
                    <Form.FormLabel class="rounded-lg bg-transparent">
                        <Mail />
                        <Form.FormInput
                            type="email"
                            min="3"
                            ptrn="[A-Za-z]+@gmail\.com"
                            placeholder="Enter your Email"
                            title="Field Must Be Email"
                            required
                        />
                    </Form.FormLabel>
                    <p class="validator-hint hidden">Field Must be Email</p>
                </div>
                <div class="mb-3">
                    <Form.FormLabel class="rounded-lg bg-transparent">
                        <KeyRound />
                        <Form.FormInput
                            id="password"
                            type={show ? 'text' : 'password'}
                            min="8"
                            ptrn={passwordPattern}
                            placeholder="Enter your password"
                            required
                        />

                        <!-- <Button val={EyeClosed}/> -->
                        <button type="button" id="hidden-pw" bind:this={ButtonShow} on:click={ShowPw}><EyeClosed /></button>
                        <button type="button" class="hidden" id="show-pw" bind:this={ButtonHidden} on:click={hiddenPw}><Eye /></button>
                    </Form.FormLabel>
                    <p class="validator-hint hidden">
                        Must be more than 8 characters, including
                        <br />At least one number <br />At least one lowercase letter <br />At least one uppercase
                        letter
                    </p>
                </div>

				<div class="mb-4 text-right">
					<span
						>Belum punya akun? <a
							href="/register"
							class="text-blue-400 underline hover:text-blue-500">register</a
						></span
					>
				</div>

				<button
					type="submit"
					class="w-full rounded-lg bg-green-400 p-2 font-bold text-white hover:bg-green-500"
					>Login</button
				>
			</form>
		</Cards.CardBody>
	</Cards.Card>
</div>
