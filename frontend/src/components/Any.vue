<template>
    <header>
        <h1>Добро пожаловать, {{ userLogin }}</h1>
    </header>
</template>

<script>
import { mapState } from 'vuex';

export default {
    name: 'AnyComponent',
    computed: {
        ...mapState(['userLogin']) // Подключаем состояние пользователя из Vuex
    }
}
</script>

<style scoped>
/* Добавьте ваши стили здесь */
</style>
