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
const update_button_element = document.getElementById("update-button");
const delete_button_element = document.getElementById("delete-button");
const config_1 = require("./config");
const sweetalert2_1 = __importDefault(require("sweetalert2"));
const api_1 = require("../api");
const notifications_1 = require("./notifications");
const update_button_handler = () => {
    const on_click = () => {
        // Show SweetAlert2 modal with input fields
        sweetalert2_1.default.fire({
            title: "Update User",
            html: `
                <div class="container">
                    <div class="row">
                        <input id="username" class="swal2-input" placeholder="New Username">
                        <input id="email" class="swal2-input" placeholder="New Email">
                    </div>
                </div>
            `,
            confirmButtonText: "Update",
            showCancelButton: true,
            preConfirm: () => __awaiter(void 0, void 0, void 0, function* () {
                var _a, _b;
                const new_username_input = sweetalert2_1.default.getPopup().querySelector("#username");
                const new_email_input = sweetalert2_1.default.getPopup().querySelector("#email");
                // Retrieve user input and handle data
                const new_username = new_username_input ? new_username_input.value : (_a = config_1.store.get_user()) === null || _a === void 0 ? void 0 : _a.name;
                const new_email = new_email_input ? new_email_input.value : (_b = config_1.store.get_user()) === null || _b === void 0 ? void 0 : _b.email;
                // Do something with the newUsername and newEmail, e.g., send it to the server
                const user_to_update = {
                    name: new_username,
                    email: new_email,
                };
                const res = yield (0, api_1.update_user)(user_to_update);
                if (res.message_type === "success") {
                    window.location.replace("/");
                }
            }),
        });
    };
    update_button_element.addEventListener("click", on_click);
};
const delete_button_handler = () => {
    console.log("in the delete button handler");
    const on_click = (e) => {
        (0, notifications_1.notify)("you clicked the button...good for you!", "success");
        (0, api_1.delete_user)()
            .then(res => res.message_type === "success" && window.location.replace("/"));
    };
    delete_button_element.addEventListener("click", on_click);
};
const run = () => {
    console.log("looking for update and delete buttons");
    if ((0, config_1.element_exists)("update-button")) {
        console.log("found update button");
        update_button_handler();
    }
    if ((0, config_1.element_exists)("delete-button")) {
        console.log("found the delete button");
        delete_button_handler();
    }
};
exports.default = run;
