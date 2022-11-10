<script setup lang="ts">
import { ref } from 'vue'
//@ts-ignore
import Field from '../components/Field.vue'
let deleteButt = ref(false)
let deactivateButt = ref(false)
</script>

<template>
    <div>
        <button @click="deleteButt = true">show Form to Delete</button>
        <Field work="deleteUser" />
    </div>
    <div>
        <button @click="deactivateButt = true">show Form to Deactivate</button>
        <Field work="deactivateUser" />
    </div>
</template>

<style></style>
