<script lang="ts" setup>
import { RouterLink } from 'vue-router'
</script>
<template>
    <div>
        toView all users:<router-link to="/">/</router-link><br />
        toView One user:<router-link to="/teej_dv">/:{teej_dv}</router-link><br />
        to Create A user:<router-link to="/create">/create</router-link><br />
        toLogin:<router-link to="/login">Login</router-link><br />
        toAdmin Login:<router-link to="/admin">AdminLogin</router-link><br />
    </div>
</template>
<style lang="scss" scoped></style>
