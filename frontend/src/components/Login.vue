<template>
    <div class="login-container">
        <div class="login-form">
            <h1>Авторизация</h1>
            <form @submit.prevent="signin">
                <!-- Поле для логина -->
                <div class="input-wrapper">
                    <input type="text"
                        placeholder="Введите логин"
                        v-model="login"
                        @input="handleLoginInput"
                        required />
            
                </div>
           <!-- Поле для пароля -->
                <div class="input-wrapper">
                    <input :type="passwordInputType"
                        placeholder="Введите пароль"
                        v-model="password"
                        @input="handlePasswordInput"
                        required />
                    <span class="input-icon" @click="togglePasswordVisibility">{{ passwordVisibleIcon }}</span>
                    
                </div>
                <p class="password-status">{{ errorMessage }}</p>

                <button type="submit" :disabled="isSubmitting">Войти</button>
            </form>
        </div>
    </div>
</template>

<script>
export default {
    name: 'LoginComponent',
    data() {
        return {
            login: '',
            password: '',
            showPassword: false,
            error: '',
            isSubmitting: false
        };
    },
    computed: {
        passwordVisibleIcon() {
            return this.showPassword ? '👁️' : '🙈';
        },
        errorMessage() {
            return this.error;
        },
        passwordInputType() {
            return this.showPassword ? 'text' : 'password';
        }
    },
    methods: {
        handleLoginInput() {
            this.login = event.target.value.replace(/[^A-Za-z0-9!_]/g, '');
        },
        handlePasswordInput() {
            this.password = event.target.value.replace(/[^A-Za-z0-9!_@#$%^&*()]/g, '');
        },
        async signup() {
            if (this.isFormValid) {
                alert(`Авторизация с логином ${this.login}`);
                // Здесь можете выполнить вызов API для авторизации
            }
        },
        togglePasswordVisibility() {
            this.showPassword = !this.showPassword;
        }
    }
}
</script>

<style scoped>
.login-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
}

.login-form {
    display: flex;
    flex-direction: column;
    padding: 20px;
    border: 2px solid #333;
    border-radius: 10px;
    background-color: #f9f9f9;
    width: 300px;
    box-sizing: border-box;
}

.input-wrapper {
    position: relative;
    margin: 10px 0;
}

.login-form input {
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    width: 100%;
    box-sizing: border-box;
}

.login-form button {
    padding: 10px;
    background-color: #4CAF50;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    font-size: 16px;
    width: 100%;
    box-sizing: border-box;
}

.login-form button:disabled {
    background-color: #ccc;
    cursor: not-allowed;
}

.login-status, .password-status {
    margin-top: 5px;
    font-size: 14px;
    color: #555;
}

.valid-login, .valid-password {
    border-color: green;
}

.invalid-login, .invalid-password {
    border-color: red;
}

.typing-login, .typing-password {
    border-color: blue;
}

.input-icon {
    position: absolute;
    right: 10px;
    top: 50%;
    transform: translateY(-50%);
    font-size: 18px;
    color: #888;
    cursor: pointer;
}


</style>
