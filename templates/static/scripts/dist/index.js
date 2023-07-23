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
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const notie_1 = __importDefault(require("notie"));
const path = "http://localhost:8080/api/users/auth";
const login_form_element = document.getElementById("login-form");
const element_exists = (n) => {
    const ele = document.getElementById(n);
    return ele == null ? false : true;
};
const init_login_form = () => {
    const username = () => {
        return login_form_element.namedItem("name");
    };
    const password = () => {
        return login_form_element.elements.namedItem("password");
    };
    const onSubmit = (e) => __awaiter(void 0, void 0, void 0, function* () {
        e.preventDefault();
        let data = {
            name: username().value,
            password: password().value
        };
        const opts = {
            method: "POST",
            body: JSON.stringify(data),
            headers: {
                "Content-Type": "application/json"
            }
        };
        const result = yield fetch(path, opts);
        const response = yield result.json();
        notie_1.default.alert({
            type: response.message_type,
            text: response.message,
        });
    });
    login_form_element.addEventListener("submit", onSubmit);
};
if (element_exists("login-form")) {
    init_login_form();
}
