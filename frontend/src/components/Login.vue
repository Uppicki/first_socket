<template>
    <div>
        <h1>Автиоризация</h1>
        <form @submit.prevent="signin">     
            <input type="text"  
                placeholder="Введите логин"
                v-model="login"
                required/>
            <input type="password"
                placeholder="Введите пароль"
                v-model="password" 
                required/>
            <button type="submit">Войти</button>
        </form>
    </div>
</template>



<script>
export default {
    name: 'LoginComponent',
    data() {
    return {
      login: '',
      password: ''
    };
  },
  methods: {
    signin() {
        alert(`Автоизация ${this.login}`)
    }
  }
}
</script>