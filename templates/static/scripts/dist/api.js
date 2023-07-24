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
exports.register_user = exports.logout_user = exports.login_user = exports.delete_user = exports.update_user = void 0;
const notifications_1 = require("./controllers/notifications");
const root = "http://localhost:8080/api/";
const POST = "POST";
const GET = "GET";
const PUT = "PUT";
const DELETE = "DELETE";
const show_response = (r) => {
    (0, notifications_1.notify_modal)(r.message_type, r.message, "Nice Shirt!", r.code.toString());
};
const update_user = (update) => __awaiter(void 0, void 0, void 0, function* () {
    const opts = {
        method: PUT,
        body: JSON.stringify(update),
        headers: {
            "Content-Type": "application/json"
        }
    };
    const result = yield fetch(root + "users", opts);
    const response = yield result.json();
    show_response(response);
    return response;
});
exports.update_user = update_user;
const delete_user = () => __awaiter(void 0, void 0, void 0, function* () {
    const opts = {
        method: DELETE,
        headers: {
            "Content-Type": "application/json"
        }
    };
    const result = yield fetch(root + "users", opts);
    const response = yield result.json();
    show_response(response);
    return response;
});
exports.delete_user = delete_user;
const login_user = (user) => __awaiter(void 0, void 0, void 0, function* () {
    const opts = {
        method: "POST",
        body: JSON.stringify(user),
        headers: {
            "Content-Type": "application/json"
        }
    };
    const result = yield fetch(root + "users/auth", opts);
    const response = yield result.json();
    show_response(response);
    return response;
});
exports.login_user = login_user;
const logout_user = () => __awaiter(void 0, void 0, void 0, function* () {
    const opts = {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
        }
    };
    const result = yield fetch(root + "users/auth", opts);
    const response = yield result.json();
    show_response(response);
    return response;
});
exports.logout_user = logout_user;
const register_user = (user) => __awaiter(void 0, void 0, void 0, function* () {
    const opts = {
        method: POST,
        body: JSON.stringify(user),
        headers: {
            "Content-Type": "application/json",
        }
    };
    const result = yield fetch(root + "users", opts);
    const response = yield result.json();
    show_response(response);
    return response;
});
exports.register_user = register_user;
