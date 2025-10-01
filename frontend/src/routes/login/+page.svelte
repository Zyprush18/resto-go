<script lang="ts">
	import * as Cards from "$lib/components/ui/Card";
	import * as Form from "$lib/components/ui/Form";
    import { Mail,KeyRound } from "@lucide/svelte";

</script>

<div class="h-screen flex items-center justify-center">
    <Cards.Card class="shadow-lg bg-gray-50 p-4">
        <h1 class="text-center text-4xl font-bold mb-4">Login</h1>
        <Cards.CardBody >
            <form action="" method="post">
                <Form.FormLabel class="mb-3 rounded-lg bg-transparent">
                    <Mail />
                    <Form.FormInput type="email" min="3" pattern="[A-Za-z][A-Za-z0-9\-]*" placeholder="Enter your username" />
                </Form.FormLabel>
                <Form.FormLabel class="mb-3 rounded-lg bg-transparent">
                    <KeyRound />
                    <Form.FormInput type="password" min="6" pattern="(?=.*\d)(?=.*[a-z])(?=.*[A-Z])" placeholder="Enter your password" />
                </Form.FormLabel>
                <div class="text-right mb-4">
                    <span >Belum punya akun? <a href="/register" class="text-blue-400 hover:text-blue-500 underline">register</a></span>
                </div>

                <button type="submit" class="bg-green-400 hover:bg-green-500 w-full p-2 rounded-lg font-bold text-white">Login</button>
            </form>
        </Cards.CardBody>
    </Cards.Card>
</div>
