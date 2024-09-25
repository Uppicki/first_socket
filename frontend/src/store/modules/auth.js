import axios from 'axios';
import StringConsts from '@/res/string_consts';

const state = {
    userLogin: "",
    accessToken: "",
    refreshToken: ""
};


const mutations = {
    SET_USER_LOGIN(state, {userLogin}){
        state.userLogin = userLogin;
    },
    CLEAR_USER_LOGIN(state) {
        state.userLogin = "";
    },
    SET_TOKENS(state, {accessToken, refreshToken}) {
        state.accessToken = accessToken;
        state.refreshToken = refreshToken;
    },
    CLEAR_TOKENS(state) {
        state.accessToken = "";
        state.refreshToken = "";
    }
};

const actions = {
    loginState({commit}, {userLogin, accessToken, refreshToken}) {
        commit('SET_USER_LOGIN', {userLogin: userLogin});
        commit('SET_TOKENS', {
            accessToken: accessToken,
            refreshToken: refreshToken
        })
        localStorage.setItem('accessToken', accessToken);
        localStorage.setItem('refreshToken', refreshToken);
    },
    logout({commit}) {
        commit('CLEAR_USER_LOGIN')
        commit('CLEAR_TOKENS')
        localStorage.removeItem('accessToken');
        localStorage.removeItem('refreshToken');
    },
    async init({ commit, dispatch, getters }) {
        const accessToken = localStorage.getItem('accessToken');
        const refreshToken = localStorage.getItem('refreshToken');

        if (accessToken && refreshToken) {
            commit('SET_TOKENS', { accessToken, refreshToken });
            let userLoginFromToken = await getters.userLoginFromToken;
            if (userLoginFromToken == "") {
                try {
                    await dispatch('updateToken');

                    userLoginFromToken = await getters.userLoginFromToken();

                    if (userLoginFromToken == "") {
                        commit('CLEAR_TOKENS');     
                    } else {
                        commit('SET_USER_LOGIN', { userLogin: userLoginFromToken });
                    }
                } catch (error) {
                    console.log(error)
                }
            } else {
                commit('SET_USER_LOGIN', { userLogin: userLoginFromToken });
            }
        } else {
            commit('CLEAR_TOKENS');
            commit('CLEAR_USER_LOGIN');
        }
    },
    async checkLogin({commit, dispatch, state, getters}) {
        if (state.accessToken.length > 0 && state.refreshToken.length > 0) {
            try {
                if (!(await getters.isValidToken)) {
                    await dispatch('updateToken');
                }
            } catch (error) {
                commit('CLEAR_TOKENS');
                commit('CLEAR_USER_LOGIN');
            }
        }
    },
    async updateToken({commit, getters}) {
        try {
            const response = await axios.post(`${process.env.VUE_APP_BACKEND_URL}/updateToken`, null, {
                headers: {
                    Authorization: `Bearer ${getters.refreshToken}`
                }
            });

            if (response.status === 200) {
                localStorage.setItem('accessToken', response.data.accessToken);
                localStorage.setItem('refreshToken', response.data.refreshToken);
                commit('SET_TOKENS', {
                    accessToken: response.data.accessToken,
                    refreshToken: response.data.refreshToken,
                })
            } else {
                localStorage.removeItem('accessToken');
                localStorage.removeItem('refreshToken');
                commit('SET_TOKENS', {
                    accessToken: '',
                    refreshToken: ''
                });
            }
        } catch (error) {
            console.log(error);
        }
    },
};

const getters = {
    isAuthenticated: (state) => () => state.accessToken!=="",
    accessToken: (state) => state.accessToken,
    refreshToken: (state) => state.refreshToken,
    isValidToken: async (state) => {
        try {
            const response = await axios.get(
                `${StringConsts.VUE_APP_API_URL}${StringConsts.userInfoUri}`, {
                headers: {
                    Authorization: `Bearer ${state.accessToken}`
                }
            });

            return response.status === 200;
        } catch (error) {
            console.log(error);
            return false;
        }
    },
    userLoginFromToken: async (state) => {
        try {
            const response = await axios.get(
                `${StringConsts.VUE_APP_API_URL}${StringConsts.userInfoUri}`, {
                headers: {
                    Authorization: `Bearer ${state.accessToken}`
                }
            });

            if (response.status === 200) {
                return response.data.login;
            }
            return "";
        } catch (error) {
            return "";
        }
    },
};

export default {
    state,
    mutations,
    actions,
    getters,
};