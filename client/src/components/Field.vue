<script setup lang="ts">
import axios from 'axios';
import { ref } from 'vue'
let props = defineProps(['work'])
let user = ref('')

async function doIt() {
    try {
let {data}=await axios.post('http://localhost:42069/Admin/')
    } catch (error) {
        console.log(error)
    }
}
</script>
<template>
    <form @submit.prevent="doIt">
        <input v-model="user" placeholder="username" />
        <button>getEmOut</button>
    </form>
</template>
<style></style>
