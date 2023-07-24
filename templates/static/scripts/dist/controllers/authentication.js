"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
const config_1 = require("./config");
const api_1 = require("../api");
const login_form_element = document.getElementById("login-form");
const logout_button_element = document.getElementById("logout-button");
const register_form_element = document.getElementById("register-form");
const login_form_handler = () => {
    const username = () => {
        return login_form_element.querySelector(`[name="username"]`);
    };
    const password = () => {
        return login_form_element.querySelector(`[name="password"]`);
    };
    const clear_inputs = () => {
        username().value = "";
        password().value = "";
    };
    const onSubmit = (e) => __awaiter(void 0, void 0, void 0, function* () {
        e.preventDefault();
        let data = {
            name: username().value,
            password: password().value
        };
        (0, api_1.login_user)(data)
            .then(res => {
            if (res.message_type === "success") {
                window.location.href = "/profile";
                clear_inputs();
            }
        });
    });
    login_form_element.addEventListener("submit", onSubmit);
};
const logout_button_handler = () => {
    const on_click = (e) => __awaiter(void 0, void 0, void 0, function* () {
        (0, api_1.logout_user)()
            .then(res => {
            if (res.message_type === "success") {
                window.location.replace("/");
            }
        });
    });
    logout_button_element.addEventListener("click", on_click);
};
const register_form_handler = () => {
    const username = () => {
        return login_form_element.querySelector(`[name="username"]`);
    };
    const password = () => {
        return login_form_element.querySelector(`[name="password"]`);
    };
    const email = () => {
        return login_form_element.querySelector(`[name="email"]`);
    };
    const confirm_password = () => {
        return login_form_element.querySelector(`[name="confirm_password"]`);
    };
    const on_submit = (e) => __awaiter(void 0, void 0, void 0, function* () {
        e.preventDefault();
        const data = {
            name: username().value,
            email: email().value,
            password: password().value,
            confirm_password: confirm_password().value
        };
        (0, api_1.register_user)(data);
    });
    register_form_element.addEventListener("submit", on_submit);
};
const run = () => {
    if ((0, config_1.element_exists)("login-form")) {
        login_form_handler();
    }
    if ((0, config_1.element_exists)("register-form")) {
        register_form_handler();
    }
    if ((0, config_1.element_exists)("logout-button")) {
        logout_button_handler();
    }
};
exports.default = run;
