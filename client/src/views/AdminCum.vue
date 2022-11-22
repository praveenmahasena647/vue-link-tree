<script setup lang="ts">
import axios from 'axios'
import { ref } from 'vue'
let deleteUserName = ref<string>('')
let deactivateUserName = ref<string>('')
async function deleteUser() {
    try {
        let { data } = await axios.post(
            'http://localhost:42069/Admin/deleteUser',
            { username: deleteUserName.value },
            {
                headers: {
                    'X-Admin': localStorage.getItem('adminCum'),
                },
            }
        )
        console.log(data)
    } catch (error) {
        console.log(error)
    }
}

async function deactivateUser() {
    try {
        let { data } = await axios.post(
            'http://localhost:42069/Admin/deactivateUser',
            { username: deactivateUserName.value },
            {
                headers: {
                    'X-Admin': localStorage.getItem('adminCum'),
                },
            }
        )
        console.log(data)
    } catch (error) {
        console.log(error)
    }
}
</script>
<template>
    <div>
        DeleteUser
        <form @submit.prevent="deleteUser">
            <input v-model="deleteUserName" placeholder="deleteUser" />
            <button>deleteEm</button>
        </form>
    </div>
    <div>
        DeactivateEm
        <form @submit.prevent="deactivateUser">
            <input v-model="deactivateUserName" placeholder="deactivateEm" />

            <button>deactivateEm</button>
        </form>
    </div>
</template>
<style lang="scss" scoped></style>
