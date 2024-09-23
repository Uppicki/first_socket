<template>
    <div class="signup-container">
        <div class="signup-form">
            <h1>Регистрация</h1>
            <form @submit.prevent="signup">
                <!-- Поле для логина -->
                <div class="input-wrapper">
                    <input type="text"  
                        placeholder="Введите логин"
                        v-model="login"
                        @input="handleLoginInput"
                        :class="loginStateClass"
                        pattern="[A-Za-z0-9!_]"
                        required/>
                    <span class="input-icon">{{ loginStateIcon }}</span>
                </div>
                <p class="login-status">{{ loginErrorMessage }}</p>

                <!-- Поле для пароля -->
                <div class="input-wrapper">
                    <input :type="passwordInputType"
                        placeholder="Введите пароль"
                        v-model="password" 
                        @input="handlePasswordInput"
                        :class="passwordStateClass"
                        required/>
                    <span class="input-icon visibility-icon" @click="togglePasswordVisibility">{{ passwordVisibleIcon }}</span>
                    <span class="input-icon">{{ passwordStateIcon }}</span>
                </div>
                <p class="password-status">{{ passwordErrorMessage }}</p>

                <!-- Поле для повтора пароля -->
                <div class="input-wrapper">
                    <input :type="passwordRepeatInputType"
                        placeholder="Повторите пароль"
                        v-model="repeatePassword" 
                        @input="handleRepeatPasswordInput"
                        :class="repeatPasswordStateClass"
                        required/>
                    <span class="input-icon visibility-icon" @click="toggleRepeatPasswordVisibility">{{ repeatPasswordVisibleIcon }}</span>
                    <span class="input-icon">{{ repeatPasswordStateIcon }}</span>
                </div>
                <p class="repeat-password-status">{{ repeatPasswordErrorMessage }}</p>

                <button type="submit" :disabled="!isFormValid">Зарегистрироваться</button>
            </form>
        </div>
    </div>
</template>

<script>
import StringConsts from '@/res/string_consts';
import axios from 'axios';



export default {
    name: 'SignupComponent',
    data() {
        return {
            login: '',
            password: '',
            repeatePassword: '',
            loginState: 'default', // default, typing, checking, valid, invalid
            passwordState: 'default', // default, typing, valid, invalid
            repeatPasswordState: 'default', // default, typing, valid, invalid
            showPassword: false,   // Для видимости пароля
            showRepeatPassword: false, // Для видимости повторного пароля
            loginError: '', 
            passwordError: '',
            repeatePasswordError: '',
            loginCheckTimeout: null // Таймер для дебаунсера
        }
    },
    computed: {
        // Состояния логина
        loginStateClass() {
            return {
                'valid-login': this.loginState === 'valid',
                'invalid-login': this.loginState === 'invalid',
                'checking-login': this.loginState === 'checking',
                'typing-login': this.loginState === 'typing',
            };
        },
        loginStateIcon() {
            switch (this.loginState) {
                case 'valid': return '✔️';
                case 'invalid': return '❌';
                case 'checking': return '⏳';
                case 'typing': return '✍️';
                default: return '❌';
            }
        },
        loginErrorMessage() {
            return this.loginError;
        },
        // Состояния для пароля
        passwordStateClass() {
            return {
                'valid-password': this.passwordState === 'valid',
                'invalid-password': this.passwordState === 'invalid',
                'typing-password': this.passwordState === 'typing',
            };
        },
        passwordStateIcon() {
            switch (this.passwordState) {
                case 'valid': return '✔️';
                case 'invalid': return '❌';
                case 'typing': return '✍️';
                default: return '❌';
            }
        },
        passwordVisibleIcon() {
            return this.showPassword ? '👁️' : '🙈';
        },
        passwordErrorMessage() {
            return this.passwordError;
        },
        // Состояния для повторного пароля
        repeatPasswordStateClass() {
            return {
                'valid-repeat-password': this.repeatPasswordState === 'valid',
                'invalid-repeat-password': this.repeatPasswordState === 'invalid',
                'typing-repeat-password': this.repeatPasswordState === 'typing',
            };
        },
        repeatPasswordStateIcon() {
            switch (this.repeatPasswordState) {
                case 'valid': return '✔️';
                case 'invalid': return '❌';
                case 'typing': return '✍️';
                default: return '❌';
            }
        },
        repeatPasswordVisibleIcon() {
            return this.showRepeatPassword ? '👁️' : '🙈';
        },
        repeatPasswordErrorMessage() {
            return this.repeatePasswordError;
        },
        isFormValid() {
            return this.loginState === 'valid' && this.passwordState === 'valid' && this.repeatPasswordState === 'valid';
        },
        passwordInputType() {
            return this.showPassword ? 'text' : 'password';
        },
        passwordRepeatInputType() {
            return this.showRepeatPassword ? 'text' : 'password';
        }
    },
    methods: {
        handleLoginInput() {
            this.loginState = 'typing';
            this.login = event.target.value.replace(/[^A-Za-z0-9!_]/g, '');
            this.loginError = ''; // Сброс ошибки при вводе
            clearTimeout(this.loginCheckTimeout); // Очистка предыдущего таймера

            if (this.login.length < 5 || this.login.length > 32) {
                this.loginState = 'invalid';
                this.loginError = 'Длинна логина должна быть от 5 до 32 символов';
                return;
            }

            // Устанавливаем таймер для дебаунсера
            this.loginCheckTimeout = setTimeout(() => {
                this.loginState = 'checking';
                this.checkLoginAvailability();
            }, 500); // Задержка в 500 мс
        },
        async checkLoginAvailability() {
            console.log('API URL:', StringConsts.VUE_APP_API_URL);
            try {
                // Отправка POST-запроса для проверки логина
                const response = await axios.get(`${StringConsts.VUE_APP_API_URL}/api/v1/auth/availableLogin`, {
                    login: this.login // Передаем логин на сервер
                });

                // Обрабатываем ответ от сервера
                if (response.data.is_available) {
                    this.loginState = 'valid';
                } else {
                    this.loginState = 'invalid';
                }
            } catch (error) {
                console.error('Ошибка при проверке логина:', error);
                this.loginState = 'invalid'; // Если произошла ошибка, устанавливаем логин как недоступный
            }
        },
        handlePasswordInput() {
            this.passwordState = 'typing';
            this.password = event.target.value.replace(/[^A-Za-z0-9!_@#$%^&*()]/g, '');
            this.passwordError = ''; // Сброс ошибки

            if (this.password.length < 8 || this.password.length > 128) {
                this.passwordState = 'invalid';
                this.passwordError = 'Пароль должен быть от 8 до 128 символов';
                return;
            }

            const hasUpperCase = /[A-Z]/.test(this.password);
            const hasLowerCase = /[a-z]/.test(this.password);
            const hasNumbers = /[0-9]/.test(this.password);

            if (!hasUpperCase) {
                this.passwordState = 'invalid';
                this.passwordError = 'Пароль должен содержать заглавные буквы';
                return;
            }
            if (!hasLowerCase) {
                this.passwordState = 'invalid';
                this.passwordError = 'Пароль должен содержать маленькие буквы';
                return;
            }
            if (!hasNumbers) {
                this.passwordState = 'invalid';
                this.passwordError = 'Пароль должен содержать цифры';
                return;
            }
            this.passwordState = 'valid';
        },
        handleRepeatPasswordInput() {
            this.repeatPasswordState = 'typing';
            this.repeatePassword = event.target.value.replace(/[^A-Za-z0-9!_@#$%^&*()]/g, '');
            this.repeatePasswordError = ''; // Сброс ошибки
            if (this.repeatePassword === this.password && this.passwordState==='valid') {
                this.repeatPasswordState = 'valid';
            } else {
                this.repeatPasswordState = 'invalid';
                this.repeatePasswordError = 'Пароли не совпадают или пароль некорректный';
            }
        },
        signup() {
            if (this.isFormValid) {
                alert(`Регистрация с логином ${this.login}`);
            }
        },
        togglePasswordVisibility() {
            this.showPassword = !this.showPassword;
        },
        toggleRepeatPasswordVisibility() {
            this.showRepeatPassword = !this.showRepeatPassword;
        }
    }
}
</script>

<style scoped>
.signup-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
}

.signup-form {
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

.signup-form input {
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    width: 100%;
    box-sizing: border-box;
}

.signup-form button {
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

.signup-form button:disabled {
    background-color: #ccc;
    cursor: not-allowed;
}

.login-status, .password-status, .repeat-password-status {
    margin-top: 5px;
    font-size: 14px;
    color: #555;
}

.valid-login, .valid-password, .valid-repeat-password {
    border-color: green;
}

.invalid-login, .invalid-password, .invalid-repeat-password {
    border-color: red;
}

.typing-login, .typing-password, .typing-repeat-password {
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

.visibility-icon {
    right: 40px; /* Смещение для иконки видимости */
}
</style>
