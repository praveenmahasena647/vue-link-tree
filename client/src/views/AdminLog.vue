<script setup lang="ts">
import router from '@/router';
import axios from 'axios'
import { reactive } from 'vue'
let cred = reactive({
    key: '',
    password: '',
})
async function LoggedIn() {
    try {
        let { data } = await axios.post('http://localhost:42069/Admin/', cred)
        localStorage.setItem('AdminMhm',data)
        router.push({'path':'/cuminMouth'})
    } catch (error) {
        console.log(error)
    }
}
</script>
<template>
    <h1>fucking get Login</h1>
    <form @submit.prevent="LoggedIn">
        <input placeholder="key" v-model="cred.key" /><br />
        <input placeholder="password" v-model="cred.password" /><br />
        <button>getLoggedIn</button>
    </form>
</template>
<style lang="scss" scoped></style>
