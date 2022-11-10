<script setup lang="ts">
import { ref } from 'vue'
let props = defineProps(['modelValue'])
let emit = defineEmits(['update'])
let imgRef = ref(null)
function clicked() {
    //@ts-ignore
    imgRef.value.click()
}
function changed() {
    // @ts-ignore
    let file = imgRef.value.files[0]
    let reader = new FileReader()

    reader.addEventListener('load', () => {
        //@ts-ignore
        emit('update:modelValue', reader.result)
    })
    reader.readAsDataURL(file)
}
</script>
<template>

    <div @click="clicked">click here to upload Img</div>
    <input
        type="file"
        @change="changed"
        ref="imgRef"
        placeholder="testInputs"
    />
<img :src='props.modelValue' />
</template>
<style lang="scss" scoped>
input,img{
display:none;
}
</style>
