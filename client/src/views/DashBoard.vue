<script setup lang="ts">
import router from '@/router'
import axios from 'axios'
import { onBeforeMount, ref, reactive } from 'vue'
let user = ref({})
let link = reactive({
    plateForm: '',
    link: '',
})
let passed = ref(false)

let password = reactive({
    oldPassword: '',
    newPassword: '',
})
onBeforeMount(async () => {
    try {
        let { data } = await axios.get('http://localhost:42069/User/login/', {
            headers: {
                'X-Jwt': localStorage.getItem('mhm'),
            },
        })
        user.value = data
    } catch (error) {
        console.log(error)
    }
})

async function addOneLink() {
    if (!user.value.links) {
        user.value.links = {}
    }
    user.value.links[link.plateForm] = link.link

    try {
        let { data } = await axios.put(
            'http://localhost:42069/User/links',
            //@ts-ignore
            user.value.links,
            {
                headers: {
                    'X-Jwt': localStorage.getItem('mhm'),
                },
            }
        )
    } catch (error) {
        console.log(error)
    }
}

async function removeLinks(key: any) {
    // @ts-ignore
    delete user.value.links[key]
    //@ts-ignore
    try {
        let { data } = await axios.put(
            'http://localhost:42069/User/links',
            //@ts-ignore
            user.value.links,
            {
                headers: {
                    'X-Jwt': localStorage.getItem('mhm'),
                },
            }
        )
    } catch (error) {
        console.log(error)
    }
}

async function changePassword() {
    try {
        let { data } = await axios.post(
            'http://localhost:42069/User/changePassword',
            password,
            {
                headers: {
                    'X-Jwt': localStorage.getItem('mhm'),
                },
            }
        )

        localStorage.removeItem('mhm')
        window.location = ''
    } catch (error) {
        console.log(error)
        passed.value = true
    }
}
let disabled = ref(false)
async function deactivate() {
    let { deactivate } = user.value
    deactivate = !deactivate
    console.log(deactivate)

    try {
        let { data } = await axios.post(
            'http://localhost:42069/User/active',
            {
                deactivate,
            },
            {
                headers: {
                    'X-Jwt': localStorage.getItem('mhm'),
                },
            }
        )
    } catch (error) {
        console.log('there was an err')
    }
}

function logOut() {
    localStorage.removeItem('mhm')
    router.push({ path: '/login' })
}

async function sendEmail() {
    try {
        let { data } = await axios.get(`http://localhost:42069/User/verify/`, {
            headers: {
                'X-Jwt': localStorage.getItem('mhm'),
            },
        })
        console.log(data)
    } catch (error) {
        console.log(error)
    }
}
</script>
<template>
    <div class="content">
        <div>userName:{{ user.userName }}</div>
        <div>userEmail:{{ user.userEmail }}</div>
        <div>deactivate:{{ user.deactivate }}</div>

        <div>
            <ul>
                <li v-for="(link, plateform) in user.links" :key="plateform">
                    {{ plateform }}:{{ link }}
                    <button @click.prevent="removeLinks(plateform)">
                        delete Link
                    </button>
                </li>
            </ul>
        </div>

        <!--        <div>userImg:{{ user.userImg }}</div>
        <div>verified:{{ user.verified }}</div>-->
    </div>
    <form @submit.prevent="addOneLink">
        <input
            v-model="link.plateForm"
            placeholder="plateForm"
            required
        /><br />
        <input
            v-model="link.link"
            placeholder="Link"
            required
            type="url"
        /><br />
        <button>submit</button><br />
    </form>

    <div>
        change content

        <div>
            change Password
            <form @submit.prevent="changePassword">
                <input
                    placeholder="old-password"
                    required
                    v-model="password.oldPassword"
                />
                <br />
                <input
                    placeholder="new-password"
                    required
                    v-model="password.newPassword"
                />
                <br />
                <button>Change Password</button>
                <br />
                <div v-if="passed">password has not been changed</div>
            </form>
            <div>
Account deactivated
                <input
                    type="checkbox"
                    @click="deactivate"
                    :disabled="disabled"
                    v-model="user.deactivate"
                />
            </div>
        </div>
    </div>
    <div>
        requist a token to Veryfied
        <button @click.prevent="sendEmail">get Veryfied</button>
    </div>
    <button @click="logOut">logOut</button>
</template>
<style lang="scss" scoped></style>
