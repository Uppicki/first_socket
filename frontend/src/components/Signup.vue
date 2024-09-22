<template>
    <div>
        <h1>Регистрация</h1>
        <form @submit.prevent="signup">
            <input type="text"  
                placeholder="Введите логин"
                v-model="login"
                required/>
            <input type="password"
                placeholder="Введите пароль"
                v-model="password" 
                required/>
            <button type="submit">Зарегистрироваться</button>
        </form>
    </div>
</template>

<script>
export default {
    name: 'SignupComponent',
    data() {
        return {
            login: '',
            password: ''
        }
    },
    methods: {
        signup() {
            alert(`Регистрация ${this.login}`)
        }
    }
}
</script>