<script lang="ts" setup>
import { reactive, watch, ref } from 'vue'
import { RouterLink } from 'vue-router'
import axios from 'axios'
import Img from '../components/Img.vue'
let newUser = reactive({
    userName: '',
    userEmail: '',
    userPassword: '',
    userPassword1: '',
    userImg: '',
})
let exist = reactive({
    userName: false,
    userEmail: false,
})
let done = ref(false)
async function createUser() {
    try {
        //@ts-ignore
        let { data } = await axios.post('http://localhost:42069/User/', newUser)
        done.value = true
    } catch (error) {
        console.log('error')
        console.log(error)
    }
}
watch(
    () => {
        return newUser.userName
    },
    async () => {
        cheakOut('X-userName')
    }
)

watch(
    () => {
        return newUser.userEmail
    },
    async () => {
        cheakOut('X-userEmail')
    }
)

async function cheakOut(userDe: string) {
    // I can import this Func TODO: refactorrrrrrrr
    if (userDe == 'X-userName') {
        try {
            let { data } = await axios.get('http://localhost:42069/User/c/', {
                headers: {
                    'X-UserName': newUser.userName,
                },
            })
            if (data) {
                exist.userName = true
                return
            }
            exist.userName = false
        } catch (error) {
            console.log(error)
        }
        return
    }

    try {
        let { data } = await axios.get('http://localhost:42069/User/c/', {
            headers: {
                'X-UserEmail': newUser.userEmail,
            },
        })
        if (data) {
            exist.userEmail = true
            return
        }
        exist.userEmail = !true
    } catch (error) {
        console.log(error)
    }
}
</script>
<template>
    <form @submit.prevent="createUser">
        <div>
            <input
                type="text"
                v-model.lazy="newUser.userName"
                placeholder="UserName"
            />
            <span v-if="exist.userName"> userName already there </span>
        </div>
        <div>
            <input
                type="text"
                v-model.lazy="newUser.userEmail"
                placeholder="EmailId"
            />
            <span v-if="exist.userEmail"> useEmail already there </span>
        </div>
        <div>
            <input
                type="password"
                v-model="newUser.userPassword"
                placeholder="password"
            />
        </div>
        <div>
            <input
                type="password"
                v-model="newUser.userPassword1"
                placeholder="password"
            />
        </div>
        <div>
            <Img v-model="newUser.userImg" />
        </div>
        <button>Submit it bitch</button>
    </form>
    <router-link to="login" v-show="done">Login Now</router-link>
</template>
<style scoped lang="scss"></style>
