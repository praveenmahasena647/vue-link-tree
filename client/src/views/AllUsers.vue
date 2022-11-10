<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { reactive, onBeforeMount, ref } from 'vue'
import { getAllUsers } from '../hooks/getAll'
let people = reactive([])
let backgroundClass: any[] = ['some-flowers', 'some-spars', 'some-views']

let selectBg = ref(null)
onBeforeMount(async () => {
    // @ts-ignore
    let result = await getAllUsers()
    //@ts-ignore
    if (result == 10) {
        console.log('error during fetch')
        return
    }
    if (result != null) {
        //@ts-ignore
        people.push(...result)
        var value: number = Math.round(Math.random() * 2)
        // @ts-ignore
        selectBg.value.classList.add(backgroundClass[value])
    }
})
</script>

<template>
    <div class="allPeople" ref="selectBg">
        <div class="helper-layer">
            <div v-for="person in people" :key="person" class="person">
                <div class="forImg">
                    <img :src="person.userImg" :alt="person.userName" />
                </div>
                <div class="text-content">
                    <div v-text="person.userName"></div>
                    <div>
                        <router-link :to="person.userName"
                            >view more
                        </router-link>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped></style>
