<script lang="ts" setup>
import axios from 'axios'
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
let router = useRouter()
let person = reactive({
    userPassword: '',
    userEmail: '',
})
async function login() {
    try {
        let { data } = await axios.post(
            'http://localhost:42069/User/login',
            person
        )
        localStorage.setItem('mhm', data)
        router.push({ path: '/DashBoard' })
    } catch (error) {
        console.log(error)
    }
}
</script>
<template>

    <div class="form">
        <form @submit.prevent="login">
            <input v-model="person.userEmail" placeholder="EmailId" /><br />
            <input v-model="person.userPassword" placeholder="password" /><br />
            <button>logIn</button>
        </form>
    </div>
</template>
<style lang="scss" scoped></style>
