<script setup>
import axios from 'axios'
import { useRoute } from 'vue-router'
import { onBeforeMount, reactive } from 'vue'
let routes = useRoute()
let person = reactive({
    data: '',
})

onBeforeMount(async () => {
    try {
        let { data } = await axios.get(
            `http://localhost:42069/User/${routes.params.id}`
        )
        person.data = data
        console.log(person.data)
    } catch (error) {
        console.log(error)
    }
})
</script>
<template>
    <div>
        <img
            :src="person.data.userImg"
            :alt="person.data.userName"
            width="100"
        />
    </div>
    <div class="personName">UserName:{{ person.data.userName }}</div>
    <div class="for-links">
        <ul>
            <li v-for="(links, index) in person.data.links" :key="index">
                {{ index }}:{{ links }}
            </li>
        </ul>
    </div>
</template>
<style lang="scss" scoped></style>
