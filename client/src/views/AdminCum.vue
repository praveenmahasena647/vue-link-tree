<script setup lang="ts">
import axios from 'axios'
import { reactive } from 'vue'
let deleteBool = reactive({
    deactivate: false,
    deleteEm: false,
    deleteUserName: '',
    deactivateUserName: '',
})

async function deleteUser() {
    try {
        let { data } = await axios.post(
            'http://localhost:42069/Admin/deleteUser',
            deleteBool.deleteUserName,
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
            deleteBool.deactivateUserName,
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
            <input
                v-model="deleteBool.deleteUserName"
                placeholder="deleteUser"
            />
            <button>deleteEm</button>
        </form>
    </div>
    <div>
        DeactivateEm
        <form @submit.prevent="deactivateUser">
            <input
                v-model="deleteBool.deactivateUserName"
                placeholder="deactivateEm"
            />

            <button>deactivateEm</button>
        </form>
    </div>
</template>
<style lang="scss" scoped></style>
