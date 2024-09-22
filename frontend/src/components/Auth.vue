<template>
    <div class="auth-container">
        <router-link class="auth-link" to="/auth/login">Авторизация</router-link>
        <router-link class="auth-link" to="/auth/signup">Регистрация</router-link>
    </div>
</template>

<script>
export default {
    name: 'AuthComponent'
};
</script>

<style scoped>
.auth-container {
    display: flex;
    flex-direction: column; /* Размещаем ссылки вертикально */
    justify-content: center; /* Вертикальное центрирование */
    align-items: center; /* Горизонтальное центрирование */
    height: 100vh; /* Занимаем полную высоту экрана */
    text-align: center; /* Центрируем текст */
}

.auth-link {
    margin: 10px 0; /* Отступы между ссылками */
    padding: 10px 20px; /* Добавляем внутренние отступы */
    border: 2px solid #333; /* Контур вокруг ссылок */
    border-radius: 5px; /* Скругленные углы */
    text-decoration: none; /* Убираем стандартное подчеркивание */
    color: #333; /* Цвет текста */
}

.auth-link:hover {
    background-color: #f0f0f0; /* Фон при наведении */
}
</style>